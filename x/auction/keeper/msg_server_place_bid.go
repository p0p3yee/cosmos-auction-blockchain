package keeper

import (
	"auction/x/auction/types"
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PlaceBid(goCtx context.Context, msg *types.MsgPlaceBid) (*types.MsgPlaceBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	auction, found := k.GetAuction(ctx, msg.AuctionId)
	if !found {
		return nil, types.AuctionNotFound
	}

	if auction.Ended {
		return nil, types.AuctionEnded
	}

	if auction.StartPrice > msg.BidPrice {
		return nil, types.BidPriceTooLow
	}

	if msg.BidPrice-auction.StartPrice < auction.MinStepPrice {
		return nil, types.BidPriceStepTooLow
	}

	bid := types.Bid{
		Creator:   msg.Creator,
		AuctionId: msg.AuctionId,
		BidPrice:  msg.BidPrice,
	}

	if id, err := k.AppendBid(ctx, bid); err != nil {
		return nil, err
	} else {
		return &types.MsgPlaceBidResponse{Id: id}, nil
	}
}
