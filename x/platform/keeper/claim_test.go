package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "Bounty/testutil/keeper"
	"Bounty/testutil/nullify"
	"Bounty/x/platform/keeper"
	"Bounty/x/platform/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNClaim(keeper keeper.Keeper, ctx context.Context, n int) []types.Claim {
	items := make([]types.Claim, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetClaim(ctx, items[i])
	}
	return items
}

func TestClaimGet(t *testing.T) {
	keeper, ctx := keepertest.PlatformKeeper(t)
	items := createNClaim(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetClaim(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestClaimRemove(t *testing.T) {
	keeper, ctx := keepertest.PlatformKeeper(t)
	items := createNClaim(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveClaim(ctx,
			item.Index,
		)
		_, found := keeper.GetClaim(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestClaimGetAll(t *testing.T) {
	keeper, ctx := keepertest.PlatformKeeper(t)
	items := createNClaim(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllClaim(ctx)),
	)
}
