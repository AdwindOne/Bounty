package simulation

import (
	"math/rand"

	"Bounty/x/platform/keeper"
	"Bounty/x/platform/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateBounty(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateBounty{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateBounty simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "CreateBounty simulation not implemented"), nil, nil
	}
}
