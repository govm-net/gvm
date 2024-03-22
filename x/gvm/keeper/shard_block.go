package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/govm-net/gvm/x/gvm/types"
)

// SetShardBlock set a specific shardBlock in the store from its index
func (k Keeper) SetShardBlock(ctx context.Context, shardBlock types.ShardBlock) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ShardBlockKeyPrefix))
	b := k.cdc.MustMarshal(&shardBlock)
	store.Set(types.ShardBlockKey(
		shardBlock.Index,
	), b)
}

// GetShardBlock returns a shardBlock from its index
func (k Keeper) GetShardBlock(
	ctx context.Context,
	index uint64,

) (val types.ShardBlock, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ShardBlockKeyPrefix))

	b := store.Get(types.ShardBlockKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveShardBlock removes a shardBlock from the store
func (k Keeper) RemoveShardBlock(
	ctx context.Context,
	index uint64,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ShardBlockKeyPrefix))
	store.Delete(types.ShardBlockKey(
		index,
	))
}

// GetAllShardBlock returns all shardBlock
func (k Keeper) GetAllShardBlock(ctx context.Context) (list []types.ShardBlock) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ShardBlockKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ShardBlock
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
