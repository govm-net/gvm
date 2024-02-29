package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "github.com/govm-net/gvm/testutil/keeper"
	"github.com/govm-net/gvm/testutil/nullify"
	"github.com/govm-net/gvm/x/shard/keeper"
	"github.com/govm-net/gvm/x/shard/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNShard(keeper keeper.Keeper, ctx context.Context, n int) []types.Shard {
	items := make([]types.Shard, n)
	for i := range items {
		items[i].Index = uint64(i)

		keeper.SetShard(ctx, items[i])
	}
	return items
}

func TestShardGet(t *testing.T) {
	keeper, ctx := keepertest.ShardKeeper(t)
	items := createNShard(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetShard(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestShardRemove(t *testing.T) {
	keeper, ctx := keepertest.ShardKeeper(t)
	items := createNShard(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveShard(ctx,
			item.Index,
		)
		_, found := keeper.GetShard(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestShardGetAll(t *testing.T) {
	keeper, ctx := keepertest.ShardKeeper(t)
	items := createNShard(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllShard(ctx)),
	)
}
