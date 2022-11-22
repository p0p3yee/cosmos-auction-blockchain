package keeper

import (
	"context"

	"auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PlaceBid(goCtx context.Context, msg *types.MsgPlaceBid) (*types.MsgPlaceBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	bid := types.Bid{
		Creator:   msg.Creator,
		AuctionId: msg.AuctionId,
		BidPrice:  msg.BidPrice,
	}

	id := k.AppendBid(ctx, bid)

	return &types.MsgPlaceBidResponse{Id: id}, nil
}
