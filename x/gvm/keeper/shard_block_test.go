package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "github.com/govm-net/gvm/testutil/keeper"
	"github.com/govm-net/gvm/testutil/nullify"
	"github.com/govm-net/gvm/x/gvm/keeper"
	"github.com/govm-net/gvm/x/gvm/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNShardBlock(keeper keeper.Keeper, ctx context.Context, n int) []types.ShardBlock {
	items := make([]types.ShardBlock, n)
	for i := range items {
		items[i].Index = uint64(i)

		keeper.SetShardBlock(ctx, items[i])
	}
	return items
}

func TestShardBlockGet(t *testing.T) {
	keeper, ctx := keepertest.GvmKeeper(t)
	items := createNShardBlock(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetShardBlock(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestShardBlockRemove(t *testing.T) {
	keeper, ctx := keepertest.GvmKeeper(t)
	items := createNShardBlock(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveShardBlock(ctx,
			item.Index,
		)
		_, found := keeper.GetShardBlock(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestShardBlockGetAll(t *testing.T) {
	keeper, ctx := keepertest.GvmKeeper(t)
	items := createNShardBlock(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllShardBlock(ctx)),
	)
}
