package keeper

import (
	"context"

	"Bounty/x/platform/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ClaimBounty(goCtx context.Context, msg *types.MsgClaimBounty) (*types.MsgClaimBountyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgClaimBountyResponse{}, nil
}
