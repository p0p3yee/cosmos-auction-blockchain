package keeper

import (
	"context"

	"auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) FinalizeAuction(goCtx context.Context, msg *types.MsgFinalizeAuction) (*types.MsgFinalizeAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	auction, found := k.GetAuction(ctx, msg.AuctionId)
	if !found {
		return nil, types.AuctionNotFound
	}

	if auction.Ended {
		return nil, types.AuctionEnded
	}

	if auction.Creator != msg.Creator {
		return nil, types.NotAuctionOwner
	}

	bid, found := k.GetBid(ctx, msg.BidId)
	if !found {
		return nil, types.BidNotFound
	}

	finalize := types.FinalizeAuction{
		Creator:    msg.Creator,
		AuctionId:  msg.AuctionId,
		BidId:      msg.BidId,
		FinalPrice: bid.BidPrice,
	}

	if id, finalPrice, err := k.AppendFinalizeAuction(ctx, finalize); err != nil {
		return nil, err
	} else {
		if err = k.EndAuction(ctx, msg.AuctionId); err != nil {
			return nil, err
		}

		return &types.MsgFinalizeAuctionResponse{
			Id:         id,
			FinalPrice: finalPrice,
		}, nil
	}
}
