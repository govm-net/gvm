package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/govm-net/gvm/x/shard/types"
)

// GetToParentCount get the total number of toParent
func (k Keeper) GetToParentCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ToParentCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetToParentCount set the total number of toParent
func (k Keeper) SetToParentCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ToParentCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendToParent appends a toParent in the store with a new id and update the count
func (k Keeper) AppendToParent(
	ctx context.Context,
	toParent types.ToParent,
) uint64 {
	// Create the toParent
	count := k.GetToParentCount(ctx)

	// Set the ID of the appended value
	toParent.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ToParentKey))
	appendedValue := k.cdc.MustMarshal(&toParent)
	store.Set(GetToParentIDBytes(toParent.Id), appendedValue)

	// Update toParent count
	k.SetToParentCount(ctx, count+1)

	return count
}

// SetToParent set a specific toParent in the store
func (k Keeper) SetToParent(ctx context.Context, toParent types.ToParent) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ToParentKey))
	b := k.cdc.MustMarshal(&toParent)
	store.Set(GetToParentIDBytes(toParent.Id), b)
}

// GetToParent returns a toParent from its id
func (k Keeper) GetToParent(ctx context.Context, id uint64) (val types.ToParent, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ToParentKey))
	b := store.Get(GetToParentIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveToParent removes a toParent from the store
func (k Keeper) RemoveToParent(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ToParentKey))
	store.Delete(GetToParentIDBytes(id))
}

// GetAllToParent returns all toParent
func (k Keeper) GetAllToParent(ctx context.Context) (list []types.ToParent) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ToParentKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ToParent
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetToParentIDBytes returns the byte representation of the ID
func GetToParentIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.ToParentKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
