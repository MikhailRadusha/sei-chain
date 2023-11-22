package state_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	testkeeper "github.com/sei-protocol/sei-chain/testutil/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/state"
	"github.com/stretchr/testify/require"
)

func TestAddLog(t *testing.T) {
	k, ctx := testkeeper.MockEVMKeeper()
	statedb := state.NewDBImpl(ctx, k, false)

	logs, err := statedb.GetAllLogs()
	require.Nil(t, err)
	require.Empty(t, logs)

	log1 := ethtypes.Log{Address: common.BytesToAddress([]byte{1}), Topics: []common.Hash{}, Data: []byte{}}
	statedb.AddLog(&log1)
	require.Nil(t, statedb.Err())
	logs, err = statedb.GetAllLogs()
	require.Nil(t, err)
	require.Equal(t, 1, len(logs))
	require.Equal(t, log1, *logs[0])

	log2 := ethtypes.Log{Address: common.BytesToAddress([]byte{2}), Topics: []common.Hash{}, Data: []byte{3}}
	statedb.AddLog(&log2)
	require.Nil(t, statedb.Err())
	logs, err = statedb.GetAllLogs()
	require.Nil(t, err)
	require.Equal(t, 2, len(logs))
	require.Equal(t, log1, *logs[0])
	require.Equal(t, log2, *logs[1])

	logs = statedb.GetLogs(common.Hash{}, 0, common.Hash{})
	require.Equal(t, 2, len(logs))
}
