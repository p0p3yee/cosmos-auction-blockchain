package keeper

import (
	"encoding/binary"
	"errors"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"auction/x/auction/types"
)

func (k Keeper) GetFinalizeAuctionCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.FinalizeAuctionCountKey))

	byteKey := []byte(types.FinalizeAuctionCountKey)

	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetFinalizeAuctionCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.FinalizeAuctionCountKey))

	byteKey := []byte(types.FinalizeAuctionCountKey)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	store.Set(byteKey, bz)
}

func (k Keeper) AppendFinalizeAuction(ctx sdk.Context, finalizeAuction types.FinalizeAuction) (uint64, uint64, error) {

	auctionStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AuctionKey))

	byteAuctionId := make([]byte, 8)
	binary.BigEndian.PutUint64(byteAuctionId, finalizeAuction.AuctionId)

	targetAuctionByte := auctionStore.Get(byteAuctionId)
	var targetAuction types.Auction

	if err := k.cdc.Unmarshal(targetAuctionByte, &targetAuction); err != nil {
		return 0, 0, err
	}

	if targetAuction.Creator != finalizeAuction.Creator {
		return 0, 0, errors.New("not target auction creator")
	}

	if targetAuction.Ended {
		return 0, 0, errors.New("target auction already ended")
	}

	bidStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BidKey))

	byteBidId := make([]byte, 8)
	binary.BigEndian.PutUint64(byteBidId, finalizeAuction.BidId)

	targetBidByte := bidStore.Get(byteBidId)
	var targetBid types.Bid

	if err := k.cdc.Unmarshal(targetBidByte, &targetBid); err != nil {
		return 0, 0, err
	}

	count := k.GetFinalizeAuctionCount(ctx)

	finalizeAuction.Id = count
	finalizeAuction.FinalPrice = targetBid.BidPrice

	// Update target auction status to ended
	targetAuction.Ended = true
	updatedTargetAuction := k.cdc.MustMarshal(&targetAuction)
	auctionStore.Set(byteAuctionId, updatedTargetAuction)

	// Save the finalized auction
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.FinalizeAuctionKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, finalizeAuction.Id)

	appendedValue := k.cdc.MustMarshal(&finalizeAuction)

	store.Set(byteKey, appendedValue)

	k.SetFinalizeAuctionCount(ctx, count+1)
	return count, finalizeAuction.FinalPrice, nil
}
