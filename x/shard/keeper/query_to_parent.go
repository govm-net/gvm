package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/govm-net/gvm/x/shard/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ToParentAll(ctx context.Context, req *types.QueryAllToParentRequest) (*types.QueryAllToParentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var toParents []types.ToParent

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	toParentStore := prefix.NewStore(store, types.KeyPrefix(types.ToParentKey))

	pageRes, err := query.Paginate(toParentStore, req.Pagination, func(key []byte, value []byte) error {
		var toParent types.ToParent
		if err := k.cdc.Unmarshal(value, &toParent); err != nil {
			return err
		}

		toParents = append(toParents, toParent)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllToParentResponse{ToParent: toParents, Pagination: pageRes}, nil
}

func (k Keeper) ToParent(ctx context.Context, req *types.QueryGetToParentRequest) (*types.QueryGetToParentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	toParent, found := k.GetToParent(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetToParentResponse{ToParent: toParent}, nil
}
