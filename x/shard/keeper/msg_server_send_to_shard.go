package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/govm-net/gvm/x/shard/types"
)

func (k msgServer) SendToShard(goCtx context.Context, msg *types.MsgSendToShard) (*types.MsgSendToShardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSendToShardResponse{}, nil
}
