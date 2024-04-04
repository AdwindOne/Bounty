package platform_test

import (
	"testing"

	keepertest "Bounty/testutil/keeper"
	"Bounty/testutil/nullify"
	platform "Bounty/x/platform/module"
	"Bounty/x/platform/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PlatformKeeper(t)
	platform.InitGenesis(ctx, k, genesisState)
	got := platform.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
