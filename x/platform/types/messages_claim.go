package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateClaim{}

func NewMsgCreateClaim(
	creator string,
	index string,
	bountyId string,
	hacker string,
	status string,

) *MsgCreateClaim {
	return &MsgCreateClaim{
		Creator:  creator,
		Index:    index,
		BountyId: bountyId,
		Hacker:   hacker,
		Status:   status,
	}
}

func (msg *MsgCreateClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateClaim{}

func NewMsgUpdateClaim(
	creator string,
	index string,
	bountyId string,
	hacker string,
	status string,

) *MsgUpdateClaim {
	return &MsgUpdateClaim{
		Creator:  creator,
		Index:    index,
		BountyId: bountyId,
		Hacker:   hacker,
		Status:   status,
	}
}

func (msg *MsgUpdateClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteClaim{}

func NewMsgDeleteClaim(
	creator string,
	index string,

) *MsgDeleteClaim {
	return &MsgDeleteClaim{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
