package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "Bounty/testutil/keeper"
	"Bounty/x/bounty/keeper"
	"Bounty/x/bounty/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPlatformMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.BountyKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreatePlatform{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreatePlatform(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetPlatform(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestPlatformMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdatePlatform
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdatePlatform{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdatePlatform{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdatePlatform{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BountyKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreatePlatform{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreatePlatform(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdatePlatform(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetPlatform(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestPlatformMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeletePlatform
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeletePlatform{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeletePlatform{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeletePlatform{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BountyKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreatePlatform(ctx, &types.MsgCreatePlatform{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeletePlatform(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetPlatform(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
