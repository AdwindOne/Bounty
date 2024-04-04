package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "Bounty/testutil/keeper"
	"Bounty/x/platform/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.PlatformKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
