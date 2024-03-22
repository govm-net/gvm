package gvm_test

import (
	"testing"

	keepertest "github.com/govm-net/gvm/testutil/keeper"
	"github.com/govm-net/gvm/testutil/nullify"
	gvm "github.com/govm-net/gvm/x/gvm/module"
	"github.com/govm-net/gvm/x/gvm/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ShardBlockList: []types.ShardBlock{
			{
				Index: 0,
			},
			{
				Index: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GvmKeeper(t)
	gvm.InitGenesis(ctx, k, genesisState)
	got := gvm.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ShardBlockList, got.ShardBlockList)
	// this line is used by starport scaffolding # genesis/test/assert
}
