package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/govm-net/gvm/x/shard/types"
)

// GetToRightChildCount get the total number of toRightChild
func (k Keeper) GetToRightChildCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ToRightChildCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetToRightChildCount set the total number of toRightChild
func (k Keeper) SetToRightChildCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ToRightChildCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendToRightChild appends a toRightChild in the store with a new id and update the count
func (k Keeper) AppendToRightChild(
	ctx context.Context,
	toRightChild types.ToRightChild,
) uint64 {
	// Create the toRightChild
	count := k.GetToRightChildCount(ctx)

	// Set the ID of the appended value
	toRightChild.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ToRightChildKey))
	appendedValue := k.cdc.MustMarshal(&toRightChild)
	store.Set(GetToRightChildIDBytes(toRightChild.Id), appendedValue)

	// Update toRightChild count
	k.SetToRightChildCount(ctx, count+1)

	return count
}

// SetToRightChild set a specific toRightChild in the store
func (k Keeper) SetToRightChild(ctx context.Context, toRightChild types.ToRightChild) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ToRightChildKey))
	b := k.cdc.MustMarshal(&toRightChild)
	store.Set(GetToRightChildIDBytes(toRightChild.Id), b)
}

// GetToRightChild returns a toRightChild from its id
func (k Keeper) GetToRightChild(ctx context.Context, id uint64) (val types.ToRightChild, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ToRightChildKey))
	b := store.Get(GetToRightChildIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveToRightChild removes a toRightChild from the store
func (k Keeper) RemoveToRightChild(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ToRightChildKey))
	store.Delete(GetToRightChildIDBytes(id))
}

// GetAllToRightChild returns all toRightChild
func (k Keeper) GetAllToRightChild(ctx context.Context) (list []types.ToRightChild) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ToRightChildKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ToRightChild
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetToRightChildIDBytes returns the byte representation of the ID
func GetToRightChildIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.ToRightChildKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
