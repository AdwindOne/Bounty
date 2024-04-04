package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "Bounty/testutil/keeper"
	"Bounty/testutil/nullify"
	"Bounty/x/bounty/keeper"
	"Bounty/x/bounty/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPlatform(keeper keeper.Keeper, ctx context.Context, n int) []types.Platform {
	items := make([]types.Platform, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetPlatform(ctx, items[i])
	}
	return items
}

func TestPlatformGet(t *testing.T) {
	keeper, ctx := keepertest.BountyKeeper(t)
	items := createNPlatform(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPlatform(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPlatformRemove(t *testing.T) {
	keeper, ctx := keepertest.BountyKeeper(t)
	items := createNPlatform(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePlatform(ctx,
			item.Index,
		)
		_, found := keeper.GetPlatform(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestPlatformGetAll(t *testing.T) {
	keeper, ctx := keepertest.BountyKeeper(t)
	items := createNPlatform(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPlatform(ctx)),
	)
}
