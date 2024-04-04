package types

import (
	"testing"

	"Bounty/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreatePlatform_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreatePlatform
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreatePlatform{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreatePlatform{
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

func TestMsgUpdatePlatform_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdatePlatform
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdatePlatform{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdatePlatform{
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

func TestMsgDeletePlatform_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeletePlatform
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeletePlatform{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeletePlatform{
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
