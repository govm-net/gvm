package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendToShard{}

func NewMsgSendToShard(creator string, direction uint64, amount sdk.Coin, info string) *MsgSendToShard {
	return &MsgSendToShard{
		Creator:   creator,
		Direction: direction,
		Amount:    amount,
		Info:      info,
	}
}

func (msg *MsgSendToShard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
