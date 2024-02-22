package bank_test

import (
	"math/big"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/sei-protocol/sei-chain/precompiles/bank"
	testkeeper "github.com/sei-protocol/sei-chain/testutil/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/state"
	"github.com/sei-protocol/sei-chain/x/evm/types"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	k, ctx := testkeeper.MockEVMKeeper()

	senderAddr, senderEVMAddr := testkeeper.MockAddressPair()
	k.SetAddressMapping(ctx, senderAddr, senderEVMAddr)
	err := k.BankKeeper().MintCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin("usei", sdk.NewInt(100))))
	require.Nil(t, err)
	err = k.BankKeeper().SendCoinsFromModuleToAccount(ctx, types.ModuleName, senderAddr, sdk.NewCoins(sdk.NewCoin("usei", sdk.NewInt(100))))
	require.Nil(t, err)

	seiAddr, evmAddr := testkeeper.MockAddressPair()
	k.SetAddressMapping(ctx, seiAddr, evmAddr)
	p, err := bank.NewPrecompile(k.BankKeeper(), k)
	require.Nil(t, err)
	statedb := state.NewDBImpl(ctx, k, true)
	evm := vm.EVM{
		StateDB:   statedb,
		TxContext: vm.TxContext{Origin: senderEVMAddr},
	}

	send, err := p.ABI.MethodById(p.SendID)
	require.Nil(t, err)
	args, err := send.Inputs.Pack(senderEVMAddr, evmAddr, "usei", big.NewInt(25))
	require.Nil(t, err)
	_, err = p.Run(&evm, senderEVMAddr, append(p.SendID, args...)) // should error because address is not whitelisted
	require.NotNil(t, err)

	sendNative, err := p.ABI.MethodById(p.SendNativeID)
	require.Nil(t, err)
	seiAddrString := seiAddr.String()
	argsNativeZero, err := sendNative.Inputs.Pack(seiAddrString, big.NewInt(0)) // no error and return early with 0 amount
	require.Nil(t, err)
	_, err = p.Run(&evm, senderEVMAddr, append(p.SendNativeID, argsNativeZero...))
	require.Nil(t, err)

	// Send native 10_000_000_000_000, split into 10 usei
	argsNative, err := sendNative.Inputs.Pack(seiAddrString, big.NewInt(10_000_000_000_000))
	require.Nil(t, err)
	_, err = p.Run(&evm, senderEVMAddr, append(p.SendNativeID, argsNative...))
	require.Nil(t, err)

	balance, err := p.ABI.MethodById(p.BalanceID)
	require.Nil(t, err)
	args, err = balance.Inputs.Pack(evmAddr, "usei")
	require.Nil(t, err)
	res, err := p.Run(&evm, common.Address{}, append(p.BalanceID, args...))
	require.Nil(t, err)
	is, err := balance.Outputs.Unpack(res)
	require.Nil(t, err)
	require.Equal(t, 1, len(is))
	require.Equal(t, big.NewInt(10000000000000), is[0].(*big.Int))
	res, err = p.Run(&evm, common.Address{}, append(p.BalanceID, args[:1]...))
	require.NotNil(t, err)
	args, err = balance.Inputs.Pack(evmAddr, "")
	require.Nil(t, err)
	res, err = p.Run(&evm, common.Address{}, append(p.BalanceID, args...))
	require.NotNil(t, err)

	// invalid input
	_, err = p.Run(&evm, common.Address{}, []byte{1, 2, 3, 4})
	require.NotNil(t, err)
}

func TestMetadata(t *testing.T) {
	k, ctx := testkeeper.MockEVMKeeper()
	k.BankKeeper().SetDenomMetaData(ctx, banktypes.Metadata{Name: "SEI", Symbol: "usei", Base: "usei"})
	p, err := bank.NewPrecompile(k.BankKeeper(), k)
	require.Nil(t, err)
	statedb := state.NewDBImpl(ctx, k, true)
	evm := vm.EVM{
		StateDB: statedb,
	}
	name, err := p.ABI.MethodById(p.NameID)
	require.Nil(t, err)
	args, err := name.Inputs.Pack("usei")
	require.Nil(t, err)
	res, err := p.Run(&evm, common.Address{}, append(p.NameID, args...))
	require.Nil(t, err)
	outputs, err := name.Outputs.Unpack(res)
	require.Nil(t, err)
	require.Equal(t, "SEI", outputs[0])

	symbol, err := p.ABI.MethodById(p.SymbolID)
	require.Nil(t, err)
	args, err = symbol.Inputs.Pack("usei")
	require.Nil(t, err)
	res, err = p.Run(&evm, common.Address{}, append(p.SymbolID, args...))
	require.Nil(t, err)
	outputs, err = symbol.Outputs.Unpack(res)
	require.Nil(t, err)
	require.Equal(t, "usei", outputs[0])

	decimal, err := p.ABI.MethodById(p.DecimalsID)
	require.Nil(t, err)
	args, err = decimal.Inputs.Pack("usei")
	require.Nil(t, err)
	res, err = p.Run(&evm, common.Address{}, append(p.DecimalsID, args...))
	require.Nil(t, err)
	outputs, err = decimal.Outputs.Unpack(res)
	require.Nil(t, err)
	require.Equal(t, uint8(0), outputs[0])

	supply, err := p.ABI.MethodById(p.SupplyID)
	require.Nil(t, err)
	args, err = supply.Inputs.Pack("usei")
	require.Nil(t, err)
	res, err = p.Run(&evm, common.Address{}, append(p.SupplyID, args...))
	require.Nil(t, err)
	outputs, err = supply.Outputs.Unpack(res)
	require.Nil(t, err)
	require.Equal(t, big.NewInt(10), outputs[0])
}

func TestRequiredGas(t *testing.T) {
	k, _ := testkeeper.MockEVMKeeper()
	p, err := bank.NewPrecompile(k.BankKeeper(), k)
	require.Nil(t, err)
	balanceRequiredGas := p.RequiredGas(p.BalanceID)
	require.Equal(t, uint64(1000), balanceRequiredGas)
	// invalid method
	require.Equal(t, uint64(0), p.RequiredGas([]byte{1, 1, 1, 1}))
}

func TestAddress(t *testing.T) {
	k, _ := testkeeper.MockEVMKeeper()
	p, err := bank.NewPrecompile(k.BankKeeper(), k)
	require.Nil(t, err)
	require.Equal(t, common.HexToAddress(bank.BankAddress), p.Address())
}
