package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/govm-net/gvm/x/gvm/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShardBlockAll(ctx context.Context, req *types.QueryAllShardBlockRequest) (*types.QueryAllShardBlockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var shardBlocks []types.ShardBlock

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	shardBlockStore := prefix.NewStore(store, types.KeyPrefix(types.ShardBlockKeyPrefix))

	pageRes, err := query.Paginate(shardBlockStore, req.Pagination, func(key []byte, value []byte) error {
		var shardBlock types.ShardBlock
		if err := k.cdc.Unmarshal(value, &shardBlock); err != nil {
			return err
		}

		shardBlocks = append(shardBlocks, shardBlock)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllShardBlockResponse{ShardBlock: shardBlocks, Pagination: pageRes}, nil
}

func (k Keeper) ShardBlock(ctx context.Context, req *types.QueryGetShardBlockRequest) (*types.QueryGetShardBlockResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetShardBlock(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetShardBlockResponse{ShardBlock: val}, nil
}
