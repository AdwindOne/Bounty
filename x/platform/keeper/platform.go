package keeper

import (
	"context"

	"Bounty/x/platform/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetPlatform set a specific platform in the store from its index
func (k Keeper) SetPlatform(ctx context.Context, platform types.Platform) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PlatformKeyPrefix))
	b := k.cdc.MustMarshal(&platform)
	store.Set(types.PlatformKey(
		platform.Index,
	), b)
}

// GetPlatform returns a platform from its index
func (k Keeper) GetPlatform(
	ctx context.Context,
	index string,

) (val types.Platform, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PlatformKeyPrefix))

	b := store.Get(types.PlatformKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePlatform removes a platform from the store
func (k Keeper) RemovePlatform(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PlatformKeyPrefix))
	store.Delete(types.PlatformKey(
		index,
	))
}

// GetAllPlatform returns all platform
func (k Keeper) GetAllPlatform(ctx context.Context) (list []types.Platform) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PlatformKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Platform
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
