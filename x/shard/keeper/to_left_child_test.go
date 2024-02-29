package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/govm-net/gvm/testutil/keeper"
	"github.com/govm-net/gvm/testutil/nullify"
	"github.com/govm-net/gvm/x/shard/keeper"
	"github.com/govm-net/gvm/x/shard/types"
	"github.com/stretchr/testify/require"
)

func createNToLeftChild(keeper keeper.Keeper, ctx context.Context, n int) []types.ToLeftChild {
	items := make([]types.ToLeftChild, n)
	for i := range items {
		items[i].Id = keeper.AppendToLeftChild(ctx, items[i])
	}
	return items
}

func TestToLeftChildGet(t *testing.T) {
	keeper, ctx := keepertest.ShardKeeper(t)
	items := createNToLeftChild(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetToLeftChild(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestToLeftChildRemove(t *testing.T) {
	keeper, ctx := keepertest.ShardKeeper(t)
	items := createNToLeftChild(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveToLeftChild(ctx, item.Id)
		_, found := keeper.GetToLeftChild(ctx, item.Id)
		require.False(t, found)
	}
}

func TestToLeftChildGetAll(t *testing.T) {
	keeper, ctx := keepertest.ShardKeeper(t)
	items := createNToLeftChild(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllToLeftChild(ctx)),
	)
}

func TestToLeftChildCount(t *testing.T) {
	keeper, ctx := keepertest.ShardKeeper(t)
	items := createNToLeftChild(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetToLeftChildCount(ctx))
}
