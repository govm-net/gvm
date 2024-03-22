package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/govm-net/gvm/testutil/keeper"
	"github.com/govm-net/gvm/x/spc/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.SpcKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
