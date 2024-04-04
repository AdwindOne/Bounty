package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "Bounty/testutil/keeper"
	"Bounty/testutil/nullify"
	"Bounty/x/bounty/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPlatformQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.BountyKeeper(t)
	msgs := createNPlatform(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetPlatformRequest
		response *types.QueryGetPlatformResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetPlatformRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetPlatformResponse{Platform: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetPlatformRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetPlatformResponse{Platform: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetPlatformRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Platform(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestPlatformQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.BountyKeeper(t)
	msgs := createNPlatform(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPlatformRequest {
		return &types.QueryAllPlatformRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PlatformAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Platform), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Platform),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.PlatformAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Platform), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Platform),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.PlatformAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Platform),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.PlatformAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
