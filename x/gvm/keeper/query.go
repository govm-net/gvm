package keeper

import (
	"github.com/govm-net/gvm/x/gvm/types"
)

var _ types.QueryServer = Keeper{}
