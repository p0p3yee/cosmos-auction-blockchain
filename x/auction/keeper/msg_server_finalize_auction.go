package keeper

import (
	"context"

	"auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) FinalizeAuction(goCtx context.Context, msg *types.MsgFinalizeAuction) (*types.MsgFinalizeAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	finalize := types.FinalizeAuction{
		Creator:   msg.Creator,
		AuctionId: msg.AuctionId,
		BidId:     msg.BidId,
	}

	if id, finalPrice, err := k.AppendFinalizeAuction(ctx, finalize); err != nil {
		return nil, err
	} else {
		return &types.MsgFinalizeAuctionResponse{
			Id:         id,
			FinalPrice: finalPrice,
		}, nil
	}
}
