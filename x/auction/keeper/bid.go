package keeper

import (
	"encoding/binary"
	"errors"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"auction/x/auction/types"
)

func (k Keeper) GetBidCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BidCountKey))

	byteKey := []byte(types.BidCountKey)

	bz := store.Get(byteKey)

	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetBidCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BidCountKey))

	byteKey := []byte(types.BidCountKey)

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	store.Set(byteKey, bz)
}

func (k Keeper) AppendBid(ctx sdk.Context, bid types.Bid) (uint64, error) {

	auctionStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AuctionKey))

	byteAuctionId := make([]byte, 8)
	binary.BigEndian.PutUint64(byteAuctionId, bid.AuctionId)

	targetAuctionByte := auctionStore.Get(byteAuctionId)
	var targetAuction types.Auction

	if err := k.cdc.Unmarshal(targetAuctionByte, &targetAuction); err != nil {
		return 0, err
	}

	if targetAuction.Ended {
		return 0, errors.New("target auction already ended")
	}

	if targetAuction.StartPrice > bid.BidPrice {
		return 0, errors.New("bid price is lower than target auction start price")
	}

	if bid.BidPrice-targetAuction.StartPrice < targetAuction.MinStepPrice {
		return 0, errors.New("bid price increment is lower than target auction min step price")
	}

	count := k.GetBidCount(ctx)

	bid.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BidKey))

	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, bid.Id)

	appendedValue := k.cdc.MustMarshal(&bid)

	store.Set(byteKey, appendedValue)

	k.SetBidCount(ctx, count+1)
	return count, nil
}
