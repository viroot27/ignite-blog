package keeper

import (
	"blog/x/blog/types"
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	// Get the current number of posts in the store
	count := k.GetPostCount(ctx)
	// Assign an ID to the post based on teh number of posts in the store
	post.Id = count
	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostKey))
	// Convert the post ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, post.Id)
	// Marshall the post into bytes
	appendedValue := k.cdc.MustMarshal(&post)
	// Insert the post bytes using post ID as a key
	store.Set(byteKey, appendedValue)
	// Update the post count
	k.SetPostCount(ctx, count+1)
	return count
}

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))
	byteKey := []byte(types.PostCountKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))
	byteKey := []byte(types.PostCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}
