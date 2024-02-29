package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/govm-net/gvm/x/shard/types"
)

// GetToLeftChildCount get the total number of toLeftChild
func (k Keeper) GetToLeftChildCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ToLeftChildCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetToLeftChildCount set the total number of toLeftChild
func (k Keeper) SetToLeftChildCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ToLeftChildCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendToLeftChild appends a toLeftChild in the store with a new id and update the count
func (k Keeper) AppendToLeftChild(
	ctx context.Context,
	toLeftChild types.ToLeftChild,
) uint64 {
	// Create the toLeftChild
	count := k.GetToLeftChildCount(ctx)

	// Set the ID of the appended value
	toLeftChild.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ToLeftChildKey))
	appendedValue := k.cdc.MustMarshal(&toLeftChild)
	store.Set(GetToLeftChildIDBytes(toLeftChild.Id), appendedValue)

	// Update toLeftChild count
	k.SetToLeftChildCount(ctx, count+1)

	return count
}

// SetToLeftChild set a specific toLeftChild in the store
func (k Keeper) SetToLeftChild(ctx context.Context, toLeftChild types.ToLeftChild) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ToLeftChildKey))
	b := k.cdc.MustMarshal(&toLeftChild)
	store.Set(GetToLeftChildIDBytes(toLeftChild.Id), b)
}

// GetToLeftChild returns a toLeftChild from its id
func (k Keeper) GetToLeftChild(ctx context.Context, id uint64) (val types.ToLeftChild, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ToLeftChildKey))
	b := store.Get(GetToLeftChildIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveToLeftChild removes a toLeftChild from the store
func (k Keeper) RemoveToLeftChild(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ToLeftChildKey))
	store.Delete(GetToLeftChildIDBytes(id))
}

// GetAllToLeftChild returns all toLeftChild
func (k Keeper) GetAllToLeftChild(ctx context.Context) (list []types.ToLeftChild) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ToLeftChildKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ToLeftChild
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetToLeftChildIDBytes returns the byte representation of the ID
func GetToLeftChildIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.ToLeftChildKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
