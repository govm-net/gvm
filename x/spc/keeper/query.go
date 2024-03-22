package keeper

import (
	"github.com/govm-net/gvm/x/spc/types"
)

var _ types.QueryServer = Keeper{}
