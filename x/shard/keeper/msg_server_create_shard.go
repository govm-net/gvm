package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/govm-net/gvm/x/shard/types"
)

func (k msgServer) CreateShard(goCtx context.Context, msg *types.MsgCreateShard) (*types.MsgCreateShardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateShardResponse{}, nil
}
