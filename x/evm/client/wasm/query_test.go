package wasm_test

import (
	"testing"

	testkeeper "github.com/sei-protocol/sei-chain/testutil/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/client/wasm"
	"github.com/stretchr/testify/require"
)

func TestERC721TransferPayload(t *testing.T) {
	k, ctx := testkeeper.MockEVMKeeper()
	addr1, _ := testkeeper.MockAddressPair()
	addr2, _ := testkeeper.MockAddressPair()
	h := wasm.NewEVMQueryHandler(k)
	res, err := h.HandleERC721TransferPayload(ctx, addr1.String(), addr2.String(), "1")
	require.Nil(t, err)
	require.NotEmpty(t, res)
}

func TestERC721ApprovePayload(t *testing.T) {
	k, ctx := testkeeper.MockEVMKeeper()
	addr1, _ := testkeeper.MockAddressPair()
	h := wasm.NewEVMQueryHandler(k)
	res, err := h.HandleERC721ApprovePayload(ctx, addr1.String(), "1")
	require.Nil(t, err)
	require.NotEmpty(t, res)
}

func TestERC721ApproveAllPayload(t *testing.T) {
	k, ctx := testkeeper.MockEVMKeeper()
	addr1, _ := testkeeper.MockAddressPair()
	h := wasm.NewEVMQueryHandler(k)
	res, err := h.HandleERC721SetApprovalAllPayload(ctx, addr1.String(), true)
	require.Nil(t, err)
	require.NotEmpty(t, res)
}