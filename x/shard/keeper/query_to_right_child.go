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

func (k Keeper) ToRightChildAll(ctx context.Context, req *types.QueryAllToRightChildRequest) (*types.QueryAllToRightChildResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var toRightChilds []types.ToRightChild

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	toRightChildStore := prefix.NewStore(store, types.KeyPrefix(types.ToRightChildKey))

	pageRes, err := query.Paginate(toRightChildStore, req.Pagination, func(key []byte, value []byte) error {
		var toRightChild types.ToRightChild
		if err := k.cdc.Unmarshal(value, &toRightChild); err != nil {
			return err
		}

		toRightChilds = append(toRightChilds, toRightChild)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllToRightChildResponse{ToRightChild: toRightChilds, Pagination: pageRes}, nil
}

func (k Keeper) ToRightChild(ctx context.Context, req *types.QueryGetToRightChildRequest) (*types.QueryGetToRightChildResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	toRightChild, found := k.GetToRightChild(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetToRightChildResponse{ToRightChild: toRightChild}, nil
}
