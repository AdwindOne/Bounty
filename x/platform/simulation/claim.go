package simulation

import (
	"math/rand"
	"strconv"

	"Bounty/x/platform/keeper"
	"Bounty/x/platform/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func SimulateMsgCreateClaim(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		i := r.Int()
		msg := &types.MsgCreateClaim{
			Creator: simAccount.Address.String(),
			Index:   strconv.Itoa(i),
		}

		_, found := k.GetClaim(ctx, msg.Index)
		if found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "Claim already exist"), nil, nil
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           moduletestutil.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgUpdateClaim(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount = simtypes.Account{}
			claim      = types.Claim{}
			msg        = &types.MsgUpdateClaim{}
			allClaim   = k.GetAllClaim(ctx)
			found      = false
		)
		for _, obj := range allClaim {
			simAccount, found = FindAccount(accs, obj.Creator)
			if found {
				claim = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "claim creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.Index = claim.Index

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           moduletestutil.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgDeleteClaim(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount = simtypes.Account{}
			claim      = types.Claim{}
			msg        = &types.MsgUpdateClaim{}
			allClaim   = k.GetAllClaim(ctx)
			found      = false
		)
		for _, obj := range allClaim {
			simAccount, found = FindAccount(accs, obj.Creator)
			if found {
				claim = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "claim creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.Index = claim.Index

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           moduletestutil.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
