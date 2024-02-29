package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/govm-net/gvm/x/shard/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShardAll(ctx context.Context, req *types.QueryAllShardRequest) (*types.QueryAllShardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var shards []types.Shard

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	shardStore := prefix.NewStore(store, types.KeyPrefix(types.ShardKeyPrefix))

	pageRes, err := query.Paginate(shardStore, req.Pagination, func(key []byte, value []byte) error {
		var shard types.Shard
		if err := k.cdc.Unmarshal(value, &shard); err != nil {
			return err
		}

		shards = append(shards, shard)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllShardResponse{Shard: shards, Pagination: pageRes}, nil
}

func (k Keeper) Shard(ctx context.Context, req *types.QueryGetShardRequest) (*types.QueryGetShardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetShard(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetShardResponse{Shard: val}, nil
}
