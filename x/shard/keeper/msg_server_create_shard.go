package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/govm-net/gvm/x/shard/types"
)

func (k msgServer) CreateShard(goCtx context.Context, msg *types.MsgCreateShard) (*types.MsgCreateShardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	if params.ShardId != 1 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	_, isFound := k.GetShard(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var shards = types.Shard{
		Creator:   msg.Creator,
		Index:     msg.Index,
		StartTime: 1,
	}

	k.SetShard(
		ctx,
		shards,
	)

	return &types.MsgCreateShardResponse{}, nil
}
