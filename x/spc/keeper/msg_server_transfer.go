package keeper

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	"github.com/govm-net/gvm/x/spc/types"
)

type PaymentCurrencyAmount struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

type PaymentCredentialInstrument struct {
	DisplayName     string `json:"displayName"`
	Icon            string `json:"icon"`
	IconMustBeShown bool   `json:"iconMustBeShown,omitempty"`
}

type CollectedClientAdditionalPaymentData struct {
	RpId        string                      `json:"rpId"`
	TopOrigin   string                      `json:"topOrigin"`
	PayeeName   string                      `json:"payeeName,omitempty"`
	PayeeOrigin string                      `json:"payeeOrigin,omitempty"`
	Total       PaymentCurrencyAmount       `json:"total"`
	Instrument  PaymentCredentialInstrument `json:"instrument"`
}

type CollectedClientData struct {
	Type      string                               `json:"type"`
	Challenge string                               `json:"challenge"`
	Origin    string                               `json:"origin"`
	Payment   CollectedClientAdditionalPaymentData `json:"payment"`
}

func base64decode(in string) []byte {
	data := bytes.Trim([]byte(in), "\"")

	// Trim the trailing equal characters.
	data = bytes.TrimRight(data, "=")
	out := make([]byte, len(data))

	n, err := base64.RawURLEncoding.Decode(out, data)
	if err != nil {
		return nil
	}
	return out[:n]
}

func (k msgServer) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	clientData := base64decode(msg.ClientData)
	authenticatorData := base64decode(msg.AuthenticatorData)
	credential := base64decode(msg.Credential)
	sign := base64decode(msg.Signature)

	clientDataHash := sha256.Sum256(clientData)

	sigData := append(authenticatorData, clientDataHash[:]...)

	var (
		key        interface{}
		clientInfo CollectedClientData
		err        error
	)
	err = json.Unmarshal(clientData, &clientInfo)
	if err != nil {
		return nil, err
	}

	if msg.AppId == "" {
		key, err = webauthncose.ParsePublicKey(credential)
	} else {
		key, err = webauthncose.ParseFIDOPublicKey(credential)
	}

	if err != nil {
		return nil, fmt.Errorf("error parsing the assertion public key: %+v", err)
	}

	valid, err := webauthncose.VerifySignature(key, sigData, sign)
	if !valid || err != nil {
		return nil, fmt.Errorf("error validating the assertion signature: %+v", err)
	}
	from := sdk.AccAddress(credential)
	account := k.accountKeeper.GetAccount(ctx, from)

	to, err := sdk.AccAddressFromBech32(clientInfo.Payment.PayeeName)
	if err != nil {
		return nil, fmt.Errorf("fail to parse payee address: %+v", err)
	}
	coinStr := clientInfo.Payment.Total.Value + clientInfo.Payment.Total.Currency
	amount, err := sdk.ParseCoinNormalized(coinStr)
	if err != nil {
		return nil, fmt.Errorf("fail to parse coin: %+v", err)
	}
	if amount.Amount.LT(math.ZeroInt()) {
		return nil, fmt.Errorf("error amount: %s", coinStr)
	}

	bi := new(big.Int)
	bi.SetUint64(account.GetSequence())
	hopeChallenge := make([]byte, protocol.ChallengeLength)
	hopeChallenge = bi.FillBytes(hopeChallenge)
	challenge := base64decode(clientInfo.Challenge)
	if !bytes.Equal(challenge, hopeChallenge) {
		return nil, fmt.Errorf("error challenge")
	}
	account.SetSequence(account.GetSequence() + 1)
	err = k.bankKeeper.SendCoins(ctx, from, to, sdk.NewCoins(amount))
	if err != nil {
		return nil, fmt.Errorf("fail to send coins: %+v", err)
	}

	return &types.MsgTransferResponse{}, nil
}
