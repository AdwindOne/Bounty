package keeper

import (
	"context"

	"Bounty/x/platform/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateBounty(goCtx context.Context, msg *types.MsgCreateBounty) (*types.MsgCreateBountyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateBountyResponse{}, nil
}
