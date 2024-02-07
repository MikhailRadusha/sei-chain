package state

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

func (s *DBImpl) SubBalance(evmAddr common.Address, amt *big.Int) {
	if amt.Sign() == 0 {
		return
	}
	if amt.Sign() < 0 {
		s.AddBalance(evmAddr, new(big.Int).Neg(amt))
		return
	}

	usei, _ := SplitUseiWeiAmount(amt) // ignore wei in loadtest
	s.err = s.k.BankKeeper().SubUnlockedCoins(s.ctx, s.getSeiAddress(evmAddr), sdk.NewCoins(sdk.NewCoin("usei", sdk.NewIntFromBigInt(usei))), false)
	// s.send(s.getSeiAddress(evmAddr), s.middleManAddress, amt)
}

func (s *DBImpl) AddBalance(evmAddr common.Address, amt *big.Int) {
	if amt.Sign() == 0 {
		return
	}
	if amt.Sign() < 0 {
		s.SubBalance(evmAddr, new(big.Int).Neg(amt))
		return
	}

	usei, _ := SplitUseiWeiAmount(amt) // ignore wei in loadtest
	s.err = s.k.BankKeeper().AddCoins(s.ctx, s.getSeiAddress(evmAddr), sdk.NewCoins(sdk.NewCoin("usei", sdk.NewIntFromBigInt(usei))), false)
	// s.send(s.middleManAddress, s.getSeiAddress(evmAddr), amt)
}

func (s *DBImpl) GetBalance(evmAddr common.Address) *big.Int {
	usei := s.k.BankKeeper().GetBalance(s.ctx, s.getSeiAddress(evmAddr), s.k.GetBaseDenom(s.ctx)).Amount
	wei := s.k.BankKeeper().GetWeiBalance(s.ctx, s.getSeiAddress(evmAddr))
	return usei.Mul(sdk.NewIntFromBigInt(UseiToSweiMultiplier)).Add(wei).BigInt()
}

// should only be called during simulation
func (s *DBImpl) SetBalance(evmAddr common.Address, amt *big.Int) {
	if !s.simulation {
		panic("should never call SetBalance in a non-simulation setting")
	}
	seiAddr := s.getSeiAddress(evmAddr)
	moduleAddr := s.k.AccountKeeper().GetModuleAddress(types.ModuleName)
	s.send(seiAddr, moduleAddr, s.GetBalance(evmAddr))
	if s.err != nil {
		panic(s.err)
	}
	usei, _ := SplitUseiWeiAmount(amt)
	coinsAmt := sdk.NewCoins(sdk.NewCoin(s.k.GetBaseDenom(s.ctx), sdk.NewIntFromBigInt(usei).Add(sdk.OneInt())))
	if err := s.k.BankKeeper().MintCoins(s.ctx, types.ModuleName, coinsAmt); err != nil {
		panic(err)
	}
	s.send(moduleAddr, seiAddr, amt)
	if s.err != nil {
		panic(s.err)
	}
}

func (s *DBImpl) getSeiAddress(evmAddr common.Address) sdk.AccAddress {
	if feeCollector, _ := s.k.GetFeeCollectorAddress(s.ctx); feeCollector == evmAddr {
		return s.coinbaseAddress
	}
	return s.k.GetSeiAddressOrDefault(s.ctx, evmAddr)
}

func (s *DBImpl) send(from sdk.AccAddress, to sdk.AccAddress, amt *big.Int) {
	usei, wei := SplitUseiWeiAmount(amt)
	s.err = s.k.BankKeeper().SendCoinsAndWei(s.ctx, from, to, nil, s.k.GetBaseDenom(s.ctx), sdk.NewIntFromBigInt(usei), sdk.NewIntFromBigInt(wei))
}
