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

func createNToParent(keeper keeper.Keeper, ctx context.Context, n int) []types.ToParent {
	items := make([]types.ToParent, n)
	for i := range items {
		items[i].Id = keeper.AppendToParent(ctx, items[i])
	}
	return items
}

func TestToParentGet(t *testing.T) {
	keeper, ctx := keepertest.ShardKeeper(t)
	items := createNToParent(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetToParent(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestToParentRemove(t *testing.T) {
	keeper, ctx := keepertest.ShardKeeper(t)
	items := createNToParent(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveToParent(ctx, item.Id)
		_, found := keeper.GetToParent(ctx, item.Id)
		require.False(t, found)
	}
}

func TestToParentGetAll(t *testing.T) {
	keeper, ctx := keepertest.ShardKeeper(t)
	items := createNToParent(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllToParent(ctx)),
	)
}

func TestToParentCount(t *testing.T) {
	keeper, ctx := keepertest.ShardKeeper(t)
	items := createNToParent(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetToParentCount(ctx))
}
