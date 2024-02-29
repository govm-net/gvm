package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyShardId = []byte("ShardId")
	// TODO: Determine the default value
	DefaultShardId uint64 = 0
)

var (
	KeyProxy = []byte("Proxy")
	// TODO: Determine the default value
	DefaultProxy string = "proxy"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	shardId uint64,
	proxy string,
) Params {
	return Params{
		ShardId: shardId,
		Proxy:   proxy,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultShardId,
		DefaultProxy,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyShardId, &p.ShardId, validateShardId),
		paramtypes.NewParamSetPair(KeyProxy, &p.Proxy, validateProxy),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateShardId(p.ShardId); err != nil {
		return err
	}

	if err := validateProxy(p.Proxy); err != nil {
		return err
	}

	return nil
}

// validateShardId validates the ShardId param
func validateShardId(v interface{}) error {
	shardId, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = shardId

	return nil
}

// validateProxy validates the Proxy param
func validateProxy(v interface{}) error {
	proxy, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = proxy

	return nil
}
