package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "Bounty/testutil/keeper"
	"Bounty/x/platform/keeper"
	"Bounty/x/platform/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestClaimMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.PlatformKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateClaim{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateClaim(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetClaim(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestClaimMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateClaim
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateClaim{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateClaim{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateClaim{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.PlatformKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateClaim{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateClaim(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateClaim(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetClaim(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestClaimMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteClaim
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteClaim{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteClaim{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteClaim{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.PlatformKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateClaim(ctx, &types.MsgCreateClaim{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteClaim(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetClaim(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
