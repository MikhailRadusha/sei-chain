package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	// module name
	ModuleName = "evm"

	RouterKey = ModuleName

	// StoreKey is string representation of the store key for auth
	StoreKey = "evm"

	// QuerierRoute is the querier route for auth
	QuerierRoute = ModuleName
)

var (
	BalanceKeyPrefix                = []byte{0x01}
	EVMAddressToSeiAddressKeyPrefix = []byte{0x02}
	SeiAddressToEVMAddressKeyPrefix = []byte{0x03}
	StateKeyPrefix                  = []byte{0x04}
	TransientStateKeyPrefix         = []byte{0x05}
	AccountTransientStateKeyPrefix  = []byte{0x06}
	TransientModuleStateKeyPrefix   = []byte{0x07}
	CodeKeyPrefix                   = []byte{0x08}
	CodeHashKeyPrefix               = []byte{0x09}
	CodeSizeKeyPrefix               = []byte{0x0a}
	NonceKeyPrefix                  = []byte{0x0b}
	ReceiptKeyPrefix                = []byte{0x0c}
)

/*
*
Transient Module State Keys
*/
var (
	// Represents the sum of all unassociated evm account balances
	// If evm module balance is higher than this value at the end of
	// the transaction, we need to burn from module balance in order
	// for this number to align.
	TotalUnassociatedBalanceKey = []byte{0x01}
	GasRefundKey                = []byte{0x02}
	LogsKey                     = []byte{0x03}
	AccessListKey               = []byte{0x04}
)

/*
*
Transient Account State Keys
*/
var (
	AccountCreated = []byte{0x01}
	AccountDeleted = []byte{0x02}
)

func BalanceKey(addr common.Address) []byte {
	return append(BalanceKeyPrefix, addr[:]...)
}

func EVMAddressToSeiAddressKey(evmAddress common.Address) []byte {
	return append(EVMAddressToSeiAddressKeyPrefix, evmAddress[:]...)
}

func SeiAddressToEVMAddressKey(seiAddress sdk.AccAddress) []byte {
	return append(SeiAddressToEVMAddressKeyPrefix, seiAddress...)
}

func StateKey(evmAddress common.Address) []byte {
	return append(StateKeyPrefix, evmAddress[:]...)
}

func TransientStateKey(evmAddress common.Address) []byte {
	return append(TransientStateKeyPrefix, evmAddress[:]...)
}

func ReceiptKey(txHash common.Hash) []byte {
	return append(ReceiptKeyPrefix, txHash[:]...)
}