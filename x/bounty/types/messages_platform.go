package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePlatform{}

func NewMsgCreatePlatform(
	creator string,
	index string,

) *MsgCreatePlatform {
	return &MsgCreatePlatform{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgCreatePlatform) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePlatform{}

func NewMsgUpdatePlatform(
	creator string,
	index string,

) *MsgUpdatePlatform {
	return &MsgUpdatePlatform{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgUpdatePlatform) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePlatform{}

func NewMsgDeletePlatform(
	creator string,
	index string,

) *MsgDeletePlatform {
	return &MsgDeletePlatform{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeletePlatform) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
