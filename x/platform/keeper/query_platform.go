package keeper

import (
	"context"

	"Bounty/x/platform/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PlatformAll(ctx context.Context, req *types.QueryAllPlatformRequest) (*types.QueryAllPlatformResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var platforms []types.Platform

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	platformStore := prefix.NewStore(store, types.KeyPrefix(types.PlatformKeyPrefix))

	pageRes, err := query.Paginate(platformStore, req.Pagination, func(key []byte, value []byte) error {
		var platform types.Platform
		if err := k.cdc.Unmarshal(value, &platform); err != nil {
			return err
		}

		platforms = append(platforms, platform)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPlatformResponse{Platform: platforms, Pagination: pageRes}, nil
}

func (k Keeper) Platform(ctx context.Context, req *types.QueryGetPlatformRequest) (*types.QueryGetPlatformResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetPlatform(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPlatformResponse{Platform: val}, nil
}
