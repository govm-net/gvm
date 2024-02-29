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

func createNToRightChild(keeper keeper.Keeper, ctx context.Context, n int) []types.ToRightChild {
	items := make([]types.ToRightChild, n)
	for i := range items {
		items[i].Id = keeper.AppendToRightChild(ctx, items[i])
	}
	return items
}

func TestToRightChildGet(t *testing.T) {
	keeper, ctx := keepertest.ShardKeeper(t)
	items := createNToRightChild(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetToRightChild(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestToRightChildRemove(t *testing.T) {
	keeper, ctx := keepertest.ShardKeeper(t)
	items := createNToRightChild(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveToRightChild(ctx, item.Id)
		_, found := keeper.GetToRightChild(ctx, item.Id)
		require.False(t, found)
	}
}

func TestToRightChildGetAll(t *testing.T) {
	keeper, ctx := keepertest.ShardKeeper(t)
	items := createNToRightChild(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllToRightChild(ctx)),
	)
}

func TestToRightChildCount(t *testing.T) {
	keeper, ctx := keepertest.ShardKeeper(t)
	items := createNToRightChild(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetToRightChildCount(ctx))
}
