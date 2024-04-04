package bounty_test

import (
	"testing"

	keepertest "Bounty/testutil/keeper"
	"Bounty/testutil/nullify"
	bounty "Bounty/x/bounty/module"
	"Bounty/x/bounty/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PlatformList: []types.Platform{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BountyKeeper(t)
	bounty.InitGenesis(ctx, k, genesisState)
	got := bounty.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PlatformList, got.PlatformList)
	// this line is used by starport scaffolding # genesis/test/assert
}
