package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/govm-net/gvm/x/shard/types"
)

func (k msgServer) SendToShard(goCtx context.Context, msg *types.MsgSendToShard) (*types.MsgSendToShardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sdk.AccAddress(msg.Creator), types.ModuleName, sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	var toLeftChild = types.ToLeftChild{
		Sender: msg.Creator,
		Amount: msg.Amount,
		Info:   msg.Info,
	}
	id := k.AppendToLeftChild(
		ctx,
		toLeftChild,
	)

	return &types.MsgSendToShardResponse{Id: id}, nil
}
