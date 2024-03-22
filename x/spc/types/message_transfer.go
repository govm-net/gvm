package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTransfer{}

func NewMsgTransfer(creator string, appId string, authenticatorData string, clientData string, credential string, signature string) *MsgTransfer {
	return &MsgTransfer{
		Creator:           creator,
		AppId:             appId,
		AuthenticatorData: authenticatorData,
		ClientData:        clientData,
		Credential:        credential,
		Signature:         signature,
	}
}

func (msg *MsgTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
