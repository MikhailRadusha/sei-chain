package keeper

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	lru "github.com/hashicorp/golang-lru/v2/simplelru"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/sei-protocol/sei-chain/x/evm/types"
)

type Keeper struct {
	logger      log.Logger
	storeKey    sdk.StoreKey
	memStoreKey sdk.StoreKey
	Paramstore  paramtypes.Subspace

	bankKeeper    bankkeeper.Keeper
	accountKeeper *authkeeper.AccountKeeper
	stakingKeeper *stakingkeeper.Keeper

	cachedFeeCollectorAddressMtx *sync.RWMutex
	cachedFeeCollectorAddress    *common.Address
	nonceMx                      *sync.RWMutex
	pendingNonces                map[string][]*addressNoncePair
	completedNonces              *lru.LRU[string, bool]
	keyToNonce                   map[tmtypes.TxKey]*addressNoncePair
}

// NonceExpiration is able to be overridden for a test
var NonceExpiration = 1 * time.Minute

type EvmTxDeferredInfo struct {
	TxIndx  int
	TxHash  common.Hash
	TxBloom ethtypes.Bloom
}

type addressNoncePair struct {
	key       tmtypes.TxKey
	address   common.Address
	nonce     uint64
	timestamp time.Time
}

func (p addressNoncePair) IsExpired(now time.Time) bool {
	return now.Sub(p.timestamp) > NonceExpiration
}

func NewKeeper(
	logger log.Logger,
	storeKey sdk.StoreKey, memStoreKey sdk.StoreKey, paramstore paramtypes.Subspace,
	bankKeeper bankkeeper.Keeper, accountKeeper *authkeeper.AccountKeeper, stakingKeeper *stakingkeeper.Keeper) *Keeper {
	if !paramstore.HasKeyTable() {
		paramstore = paramstore.WithKeyTable(types.ParamKeyTable())
	}
	// needs to be bounded to avoid leaking forever
	cn, err := lru.NewLRU[string, bool](100000, nil)
	if err != nil {
		panic(fmt.Sprintf("could not create lru: %v", err))
	}
	k := &Keeper{
		logger:                       logger,
		storeKey:                     storeKey,
		memStoreKey:                  memStoreKey,
		Paramstore:                   paramstore,
		bankKeeper:                   bankKeeper,
		accountKeeper:                accountKeeper,
		stakingKeeper:                stakingKeeper,
		pendingNonces:                make(map[string][]*addressNoncePair),
		completedNonces:              cn,
		nonceMx:                      &sync.RWMutex{},
		cachedFeeCollectorAddressMtx: &sync.RWMutex{},
		keyToNonce:                   make(map[tmtypes.TxKey]*addressNoncePair),
	}
	go k.startNonceReaper()
	return k
}

func (k *Keeper) AccountKeeper() *authkeeper.AccountKeeper {
	return k.accountKeeper
}

func (k *Keeper) BankKeeper() bankkeeper.Keeper {
	return k.bankKeeper
}

func (k *Keeper) GetStoreKey() sdk.StoreKey {
	return k.storeKey
}

func (k *Keeper) PrefixStore(ctx sdk.Context, pref []byte) sdk.KVStore {
	store := ctx.KVStore(k.GetStoreKey())
	return prefix.NewStore(store, pref)
}

func (k *Keeper) PurgePrefix(ctx sdk.Context, pref []byte) {
	store := k.PrefixStore(ctx, pref)
	if err := store.DeleteAll(nil, nil); err != nil {
		panic(err)
	}
}

func (k *Keeper) GetVMBlockContext(ctx sdk.Context, gp core.GasPool) (*vm.BlockContext, error) {
	coinbase, err := k.GetFeeCollectorAddress(ctx)
	if err != nil {
		return nil, err
	}
	r, err := ctx.BlockHeader().Time.MarshalBinary()
	if err != nil {
		return nil, err
	}
	rh := common.BytesToHash(r)
	return &vm.BlockContext{
		CanTransfer: core.CanTransfer,
		Transfer:    core.Transfer,
		GetHash:     k.GetHashFn(ctx),
		Coinbase:    coinbase,
		GasLimit:    gp.Gas(),
		BlockNumber: big.NewInt(ctx.BlockHeight()),
		Time:        uint64(ctx.BlockHeader().Time.Unix()),
		Difficulty:  big.NewInt(0),                               // only needed for PoW
		BaseFee:     k.GetBaseFeePerGas(ctx).RoundInt().BigInt(), // feemarket not enabled
		Random:      &rh,
	}, nil
}

// returns a function that provides block header hash based on block number
func (k *Keeper) GetHashFn(ctx sdk.Context) vm.GetHashFunc {
	return func(height uint64) common.Hash {
		if height > math.MaxInt64 {
			ctx.Logger().Error("Sei block height is bounded by int64 range")
			return common.Hash{}
		}
		h := int64(height)
		if ctx.BlockHeight() == h {
			// current header hash is in the context already
			return common.BytesToHash(ctx.HeaderHash())
		}
		if ctx.BlockHeight() < h {
			// future block doesn't have a hash yet
			return common.Hash{}
		}
		// fetch historical hash from historical info
		return k.getHistoricalHash(ctx, h)
	}
}

func (k *Keeper) GetEVMTxDeferredInfo(ctx sdk.Context) (res []EvmTxDeferredInfo) {
	hashMap, bloomMap := map[int]common.Hash{}, map[int]ethtypes.Bloom{}
	hashIter := prefix.NewStore(ctx.KVStore(k.memStoreKey), types.TxHashPrefix).Iterator(nil, nil)
	for ; hashIter.Valid(); hashIter.Next() {
		h := common.Hash{}
		h.SetBytes(hashIter.Value())
		hashMap[int(binary.BigEndian.Uint32(hashIter.Key()))] = h
	}
	hashIter.Close()
	bloomIter := prefix.NewStore(ctx.KVStore(k.memStoreKey), types.TxBloomPrefix).Iterator(nil, nil)
	for ; bloomIter.Valid(); bloomIter.Next() {
		b := ethtypes.Bloom{}
		b.SetBytes(bloomIter.Value())
		bloomMap[int(binary.BigEndian.Uint32(bloomIter.Key()))] = b
	}
	bloomIter.Close()
	for idx, h := range hashMap {
		i := EvmTxDeferredInfo{TxIndx: idx, TxHash: h}
		if b, ok := bloomMap[idx]; ok {
			i.TxBloom = b
			delete(bloomMap, idx)
		}
		res = append(res, i)
	}
	for idx, b := range bloomMap {
		res = append(res, EvmTxDeferredInfo{TxIndx: idx, TxBloom: b})
	}
	sort.SliceStable(res, func(i, j int) bool { return res[i].TxIndx < res[j].TxIndx })
	return
}

func (k *Keeper) AppendToEvmTxDeferredInfo(ctx sdk.Context, bloom ethtypes.Bloom, txHash common.Hash) {
	key := make([]byte, 8)
	binary.BigEndian.PutUint32(key, uint32(ctx.TxIndex()))
	prefix.NewStore(ctx.KVStore(k.memStoreKey), types.TxHashPrefix).Set(key, txHash[:])
	prefix.NewStore(ctx.KVStore(k.memStoreKey), types.TxBloomPrefix).Set(key, bloom[:])
}

func (k *Keeper) ClearEVMTxDeferredInfo(ctx sdk.Context) {
	hashStore := prefix.NewStore(ctx.KVStore(k.memStoreKey), types.TxHashPrefix)
	hashIterator := hashStore.Iterator(nil, nil)
	defer hashIterator.Close()
	hashKeysToDelete := [][]byte{}
	for ; hashIterator.Valid(); hashIterator.Next() {
		hashKeysToDelete = append(hashKeysToDelete, hashIterator.Key())
	}
	// close the first iterator for safety
	hashIterator.Close()
	for _, key := range hashKeysToDelete {
		hashStore.Delete(key)
	}

	bloomStore := prefix.NewStore(ctx.KVStore(k.memStoreKey), types.TxBloomPrefix)
	bloomIterator := bloomStore.Iterator(nil, nil)
	bloomKeysToDelete := [][]byte{}
	defer bloomIterator.Close()
	for ; bloomIterator.Valid(); bloomIterator.Next() {
		bloomKeysToDelete = append(bloomKeysToDelete, bloomIterator.Key())
	}
	// close the second iterator for safety
	bloomIterator.Close()
	for _, key := range bloomKeysToDelete {
		bloomStore.Delete(key)
	}
}

func (k *Keeper) getHistoricalHash(ctx sdk.Context, h int64) common.Hash {
	histInfo, found := k.stakingKeeper.GetHistoricalInfo(ctx, h)
	if !found {
		// too old, already pruned
		return common.Hash{}
	}
	header, _ := tmtypes.HeaderFromProto(&histInfo.Header)

	return common.BytesToHash(header.Hash())
}

// nonceCacheKey is a helper function to create a key for the completed nonces cache
func nonceCacheKey(addr common.Address, nonce uint64) string {
	return fmt.Sprintf("%s|%d", addr.Hex(), nonce)
}

func toNonceList(pair []*addressNoncePair) []uint64 {
	out := make([]uint64, len(pair))
	for i, p := range pair {
		out[i] = p.nonce
	}
	return out
}

// CalculateNextNonce calculates the next nonce for an address
// If includePending is true, it will consider pending nonces
// If includePending is false, it will only return the next nonce from GetNonce
func (k *Keeper) CalculateNextNonce(ctx sdk.Context, addr common.Address, includePending bool) uint64 {
	k.nonceMx.Lock()
	defer k.nonceMx.Unlock()

	nextNonce := k.GetNonce(ctx, addr)

	// we only want the latest nonce if we're not including pending
	if !includePending {
		return nextNonce
	}

	// get the pending nonces (nil is fine)
	pending := toNonceList(k.pendingNonces[addr.Hex()])

	// Check each nonce starting from latest until we find a gap
	// That gap is the next nonce we should use.
	// The completed nonces are limited to 100k entries
	for {
		// if it's not in pending and not completed, then it's the next nonce
		if !sortedListContains(pending, nextNonce) && !k.completedNonces.Contains(nonceCacheKey(addr, nextNonce)) {
			return nextNonce
		}
		nextNonce++
	}
}

// sortedListContains is a helper function to check if a sorted slice contains a specific element
func sortedListContains(slice []uint64, item uint64) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
		// because it's sorted, we can bail if it's higher
		if v > item {
			return false
		}
	}
	return false
}

// AddPendingNonce adds a pending nonce to the keeper
func (k *Keeper) AddPendingNonce(key tmtypes.TxKey, addr common.Address, nonce uint64) {
	k.nonceMx.Lock()
	defer k.nonceMx.Unlock()

	addrStr := addr.Hex()

	pair := &addressNoncePair{
		address:   addr,
		nonce:     nonce,
		key:       key,
		timestamp: time.Now().UTC(),
	}

	k.keyToNonce[key] = pair
	k.pendingNonces[addrStr] = append(k.pendingNonces[addrStr], pair)

	sort.Slice(k.pendingNonces[addrStr], func(i, j int) bool {
		return k.pendingNonces[addrStr][i].nonce < k.pendingNonces[addrStr][j].nonce
	})
}

// ExpirePendingNonce removes a pending nonce from the keeper but leaves a hole
// so that a future transaction must use this nonce
func (k *Keeper) ExpirePendingNonce(key tmtypes.TxKey) {
	k.nonceMx.Lock()
	defer k.nonceMx.Unlock()
	k.expirePendingNonceUnsafe(key)
}

func (k *Keeper) expirePendingNonceUnsafe(key tmtypes.TxKey) {
	tx, ok := k.keyToNonce[key]

	if !ok {
		return
	}

	delete(k.keyToNonce, key)
	addr := tx.address.Hex()
	for i, pair := range k.pendingNonces[addr] {
		if pair.nonce == tx.nonce {
			// remove nonce but keep prior nonces in the slice (unlike the completion scenario)
			k.pendingNonces[addr] = append(k.pendingNonces[addr][:i], k.pendingNonces[addr][i+1:]...)
			// If the slice is empty, delete the key from the map
			if len(k.pendingNonces[addr]) == 0 {
				delete(k.pendingNonces, addr)
			}
			return
		}
	}
}

func (k *Keeper) ReapExpiredNonces() {
	k.nonceMx.Lock()
	defer k.nonceMx.Unlock()
	k.print()
	now := time.Now().UTC()
	for addr, nonces := range k.pendingNonces {
		var remaining []*addressNoncePair
		for _, nonce := range nonces {
			if !nonce.IsExpired(now) {
				remaining = append(remaining, nonce)
				continue
			}
			k.logger.Error("reaper expiring nonce",
				"nonce", nonce.nonce,
				"address", addr,
				"age_ms", now.Sub(nonce.timestamp).Milliseconds())
			delete(k.keyToNonce, nonce.key)
		}
		if len(remaining) == 0 {
			delete(k.pendingNonces, addr)
			continue
		}
		k.pendingNonces[addr] = remaining
	}
	for key, v := range k.keyToNonce {
		if v.IsExpired(now) {
			panic(fmt.Sprintf("keyToNonce has a key that is not in pendingNonces: %X, %v", key, v))
		}
	}
}

// startNonceReaper is a background process that periodically checks for expired nonces
// this exists for safety in case a bug is introduced that does not clear a nonce
func (k *Keeper) startNonceReaper() {
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		k.logger.Info("DEBUG: nonce reaper start", "keyToNonceLen", len(k.keyToNonce), "pendingNonceLen", len(k.pendingNonces))
		k.ReapExpiredNonces()
	}
}

func (k *Keeper) print() {
	for addr, nonces := range k.pendingNonces {
		var nonceStrs []string
		for _, n := range nonces {
			nonceStrs = append(nonceStrs, fmt.Sprintf("%d(%dms)", n.nonce, time.Since(n.timestamp).Milliseconds()))
		}
		k.logger.Info("DEBUG: nonce reaper print()", "address", addr, "nonces", strings.Join(nonceStrs, ","))
	}
}

// CompletePendingNonce removes a pending nonce from the keeper
// success means this transaction was processed and this nonce is used
func (k *Keeper) CompletePendingNonce(key tmtypes.TxKey) {
	k.nonceMx.Lock()
	defer k.nonceMx.Unlock()

	acctNonce, ok := k.keyToNonce[key]
	if !ok {
		return
	}
	address := acctNonce.address
	nonce := acctNonce.nonce

	delete(k.keyToNonce, key)
	k.completedNonces.Add(nonceCacheKey(address, nonce), true)

	addrStr := address.Hex()
	if _, ok := k.pendingNonces[addrStr]; !ok {
		return
	}

	for i, pair := range k.pendingNonces[addrStr] {
		if pair.nonce >= nonce {
			// remove the nonce and all prior nonces from the slice
			copy(k.pendingNonces[addrStr], k.pendingNonces[addrStr][i+1:])
			k.pendingNonces[addrStr] = k.pendingNonces[addrStr][:len(k.pendingNonces[addrStr])-i-1]

			// If the slice is empty, delete the key from the map
			if len(k.pendingNonces[addrStr]) == 0 {
				delete(k.pendingNonces, addrStr)
			}
			return
		}
	}
}
