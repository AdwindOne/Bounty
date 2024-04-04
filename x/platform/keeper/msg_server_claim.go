package keeper

import (
	"context"

	"Bounty/x/platform/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateClaim(goCtx context.Context, msg *types.MsgCreateClaim) (*types.MsgCreateClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetClaim(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var claim = types.Claim{
		Creator:  msg.Creator,
		Index:    msg.Index,
		BountyId: msg.BountyId,
		Hacker:   msg.Hacker,
		Status:   msg.Status,
	}

	k.SetClaim(
		ctx,
		claim,
	)
	return &types.MsgCreateClaimResponse{}, nil
}

func (k msgServer) UpdateClaim(goCtx context.Context, msg *types.MsgUpdateClaim) (*types.MsgUpdateClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetClaim(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var claim = types.Claim{
		Creator:  msg.Creator,
		Index:    msg.Index,
		BountyId: msg.BountyId,
		Hacker:   msg.Hacker,
		Status:   msg.Status,
	}

	k.SetClaim(ctx, claim)

	return &types.MsgUpdateClaimResponse{}, nil
}

func (k msgServer) DeleteClaim(goCtx context.Context, msg *types.MsgDeleteClaim) (*types.MsgDeleteClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetClaim(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveClaim(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteClaimResponse{}, nil
}
