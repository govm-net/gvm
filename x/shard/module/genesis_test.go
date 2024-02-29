package shard_test

import (
	"testing"

	keepertest "github.com/govm-net/gvm/testutil/keeper"
	"github.com/govm-net/gvm/testutil/nullify"
	shard "github.com/govm-net/gvm/x/shard/module"
	"github.com/govm-net/gvm/x/shard/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ShardKeeper(t)
	shard.InitGenesis(ctx, k, genesisState)
	got := shard.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
