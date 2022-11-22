package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"auction/x/auction/types"
)

func (k Keeper) GetAuctionCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AuctionCountKey))

	byteKey := []byte(types.AuctionCountKey)

	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetAuctionCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AuctionCountKey))

	byteKey := []byte(types.AuctionCountKey)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	store.Set(byteKey, bz)
}

func (k Keeper) AppendAuction(ctx sdk.Context, auction types.Auction) uint64 {
	count := k.GetAuctionCount(ctx)

	auction.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AuctionKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, auction.Id)

	appendedValue := k.cdc.MustMarshal(&auction)

	store.Set(byteKey, appendedValue)

	k.SetAuctionCount(ctx, count+1)
	return count
}
