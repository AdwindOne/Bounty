package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgClaimBounty{}

func NewMsgClaimBounty(creator string, bountyId string) *MsgClaimBounty {
	return &MsgClaimBounty{
		Creator:  creator,
		BountyId: bountyId,
	}
}

func (msg *MsgClaimBounty) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
