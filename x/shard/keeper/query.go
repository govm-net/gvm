package keeper

import (
	"github.com/govm-net/gvm/x/shard/types"
)

var _ types.QueryServer = Keeper{}
