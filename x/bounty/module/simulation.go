package bounty

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"Bounty/testutil/sample"
	bountysimulation "Bounty/x/bounty/simulation"
	"Bounty/x/bounty/types"
)

// avoid unused import issue
var (
	_ = bountysimulation.FindAccount
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

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	bountyGenesis := types.GenesisState{
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
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&bountyGenesis)
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
		bountysimulation.SimulateMsgCreatePlatform(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePlatform int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePlatform, &weightMsgUpdatePlatform, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePlatform = defaultWeightMsgUpdatePlatform
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePlatform,
		bountysimulation.SimulateMsgUpdatePlatform(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePlatform int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePlatform, &weightMsgDeletePlatform, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePlatform = defaultWeightMsgDeletePlatform
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePlatform,
		bountysimulation.SimulateMsgDeletePlatform(am.accountKeeper, am.bankKeeper, am.keeper),
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
				bountysimulation.SimulateMsgCreatePlatform(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdatePlatform,
			defaultWeightMsgUpdatePlatform,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bountysimulation.SimulateMsgUpdatePlatform(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeletePlatform,
			defaultWeightMsgDeletePlatform,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				bountysimulation.SimulateMsgDeletePlatform(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
