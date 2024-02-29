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

func (k Keeper) ToLeftChildAll(ctx context.Context, req *types.QueryAllToLeftChildRequest) (*types.QueryAllToLeftChildResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var toLeftChilds []types.ToLeftChild

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	toLeftChildStore := prefix.NewStore(store, types.KeyPrefix(types.ToLeftChildKey))

	pageRes, err := query.Paginate(toLeftChildStore, req.Pagination, func(key []byte, value []byte) error {
		var toLeftChild types.ToLeftChild
		if err := k.cdc.Unmarshal(value, &toLeftChild); err != nil {
			return err
		}

		toLeftChilds = append(toLeftChilds, toLeftChild)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllToLeftChildResponse{ToLeftChild: toLeftChilds, Pagination: pageRes}, nil
}

func (k Keeper) ToLeftChild(ctx context.Context, req *types.QueryGetToLeftChildRequest) (*types.QueryGetToLeftChildResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	toLeftChild, found := k.GetToLeftChild(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetToLeftChildResponse{ToLeftChild: toLeftChild}, nil
}
