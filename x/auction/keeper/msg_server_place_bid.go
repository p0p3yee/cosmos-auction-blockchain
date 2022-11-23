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

	if uint64(auction.CreatedAt)+auction.Duration <= uint64(ctx.BlockHeight()) {
		if err := k.EndAuction(ctx, auction.Id); err != nil {
			return nil, err
		}
		return nil, types.AuctionEnded
	}

	if auction.StartPrice > msg.BidPrice {
		return nil, types.BidPriceTooLow
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
