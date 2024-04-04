package types

import (
	"testing"

	"Bounty/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateClaim_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateClaim
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateClaim{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateClaim{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateClaim_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateClaim
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateClaim{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateClaim{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteClaim_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteClaim
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteClaim{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteClaim{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
