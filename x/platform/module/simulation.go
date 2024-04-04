package platform

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"Bounty/testutil/sample"
	platformsimulation "Bounty/x/platform/simulation"
	"Bounty/x/platform/types"
)

// avoid unused import issue
var (
	_ = platformsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreatePlatform = "op_weight_msg_platform"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePlatform int = 100

	opWeightMsgUpdatePlatform = "op_weight_msg_platform"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePlatform int = 100

	opWeightMsgDeletePlatform = "op_weight_msg_platform"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePlatform int = 100

	opWeightMsgCreateClaim = "op_weight_msg_claim"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateClaim int = 100

	opWeightMsgUpdateClaim = "op_weight_msg_claim"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateClaim int = 100

	opWeightMsgDeleteClaim = "op_weight_msg_claim"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteClaim int = 100

	opWeightMsgCreateBounty = "op_weight_msg_create_bounty"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateBounty int = 100

	opWeightMsgClaimBounty = "op_weight_msg_claim_bounty"
	// TODO: Determine the simulation weight value
	defaultWeightMsgClaimBounty int = 100

	opWeightMsgCompleteBounty = "op_weight_msg_complete_bounty"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCompleteBounty int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	platformGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PlatformList: []types.Platform{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		ClaimList: []types.Claim{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&platformGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePlatform int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePlatform, &weightMsgCreatePlatform, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePlatform = defaultWeightMsgCreatePlatform
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePlatform,
		platformsimulation.SimulateMsgCreatePlatform(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePlatform int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePlatform, &weightMsgUpdatePlatform, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePlatform = defaultWeightMsgUpdatePlatform
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePlatform,
		platformsimulation.SimulateMsgUpdatePlatform(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePlatform int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePlatform, &weightMsgDeletePlatform, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePlatform = defaultWeightMsgDeletePlatform
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePlatform,
		platformsimulation.SimulateMsgDeletePlatform(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateClaim int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateClaim, &weightMsgCreateClaim, nil,
		func(_ *rand.Rand) {
			weightMsgCreateClaim = defaultWeightMsgCreateClaim
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateClaim,
		platformsimulation.SimulateMsgCreateClaim(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateClaim int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateClaim, &weightMsgUpdateClaim, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClaim = defaultWeightMsgUpdateClaim
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateClaim,
		platformsimulation.SimulateMsgUpdateClaim(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteClaim int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteClaim, &weightMsgDeleteClaim, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteClaim = defaultWeightMsgDeleteClaim
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteClaim,
		platformsimulation.SimulateMsgDeleteClaim(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateBounty int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateBounty, &weightMsgCreateBounty, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBounty = defaultWeightMsgCreateBounty
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateBounty,
		platformsimulation.SimulateMsgCreateBounty(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgClaimBounty int
	simState.AppParams.GetOrGenerate(opWeightMsgClaimBounty, &weightMsgClaimBounty, nil,
		func(_ *rand.Rand) {
			weightMsgClaimBounty = defaultWeightMsgClaimBounty
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaimBounty,
		platformsimulation.SimulateMsgClaimBounty(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCompleteBounty int
	simState.AppParams.GetOrGenerate(opWeightMsgCompleteBounty, &weightMsgCompleteBounty, nil,
		func(_ *rand.Rand) {
			weightMsgCompleteBounty = defaultWeightMsgCompleteBounty
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCompleteBounty,
		platformsimulation.SimulateMsgCompleteBounty(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreatePlatform,
			defaultWeightMsgCreatePlatform,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				platformsimulation.SimulateMsgCreatePlatform(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdatePlatform,
			defaultWeightMsgUpdatePlatform,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				platformsimulation.SimulateMsgUpdatePlatform(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeletePlatform,
			defaultWeightMsgDeletePlatform,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				platformsimulation.SimulateMsgDeletePlatform(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateClaim,
			defaultWeightMsgCreateClaim,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				platformsimulation.SimulateMsgCreateClaim(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateClaim,
			defaultWeightMsgUpdateClaim,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				platformsimulation.SimulateMsgUpdateClaim(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteClaim,
			defaultWeightMsgDeleteClaim,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				platformsimulation.SimulateMsgDeleteClaim(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateBounty,
			defaultWeightMsgCreateBounty,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				platformsimulation.SimulateMsgCreateBounty(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgClaimBounty,
			defaultWeightMsgClaimBounty,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				platformsimulation.SimulateMsgClaimBounty(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCompleteBounty,
			defaultWeightMsgCompleteBounty,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				platformsimulation.SimulateMsgCompleteBounty(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
