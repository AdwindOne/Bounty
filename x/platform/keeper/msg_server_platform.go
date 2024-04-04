package keeper

import (
	"context"

	"Bounty/x/platform/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePlatform(goCtx context.Context, msg *types.MsgCreatePlatform) (*types.MsgCreatePlatformResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetPlatform(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var platform = types.Platform{
		Creator:     msg.Creator,
		Index:       msg.Index,
		Title:       msg.Title,
		Description: msg.Description,
		Reward:      msg.Reward,
		Creator1:    msg.Creator1,
	}

	k.SetPlatform(
		ctx,
		platform,
	)
	return &types.MsgCreatePlatformResponse{}, nil
}

func (k msgServer) UpdatePlatform(goCtx context.Context, msg *types.MsgUpdatePlatform) (*types.MsgUpdatePlatformResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetPlatform(
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

	var platform = types.Platform{
		Creator:     msg.Creator,
		Index:       msg.Index,
		Title:       msg.Title,
		Description: msg.Description,
		Reward:      msg.Reward,
		Creator1:    msg.Creator1,
	}

	k.SetPlatform(ctx, platform)

	return &types.MsgUpdatePlatformResponse{}, nil
}

func (k msgServer) DeletePlatform(goCtx context.Context, msg *types.MsgDeletePlatform) (*types.MsgDeletePlatformResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetPlatform(
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

	k.RemovePlatform(
		ctx,
		msg.Index,
	)

	return &types.MsgDeletePlatformResponse{}, nil
}
