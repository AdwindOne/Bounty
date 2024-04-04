package keeper

import (
	"context"

	"Bounty/x/platform/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CompleteBounty(goCtx context.Context, msg *types.MsgCompleteBounty) (*types.MsgCompleteBountyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCompleteBountyResponse{}, nil
}
