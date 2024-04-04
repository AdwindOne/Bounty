package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateBounty{}

func NewMsgCreateBounty(creator string, title string, description string, reward string) *MsgCreateBounty {
	return &MsgCreateBounty{
		Creator:     creator,
		Title:       title,
		Description: description,
		Reward:      reward,
	}
}

func (msg *MsgCreateBounty) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
