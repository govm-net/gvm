package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/govm-net/gvm/x/shard/types"
)

// SetShard set a specific shard in the store from its index
func (k Keeper) SetShard(ctx context.Context, shard types.Shard) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ShardKeyPrefix))
	b := k.cdc.MustMarshal(&shard)
	store.Set(types.ShardKey(
		shard.Index,
	), b)
}

// GetShard returns a shard from its index
func (k Keeper) GetShard(
	ctx context.Context,
	index uint64,

) (val types.Shard, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ShardKeyPrefix))

	b := store.Get(types.ShardKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveShard removes a shard from the store
func (k Keeper) RemoveShard(
	ctx context.Context,
	index uint64,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ShardKeyPrefix))
	store.Delete(types.ShardKey(
		index,
	))
}

// GetAllShard returns all shard
func (k Keeper) GetAllShard(ctx context.Context) (list []types.Shard) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ShardKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Shard
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
