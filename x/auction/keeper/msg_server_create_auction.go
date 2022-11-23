package keeper

import (
	"context"

	"auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAuction(goCtx context.Context, msg *types.MsgCreateAuction) (*types.MsgCreateAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.StartPrice <= 0 {
		return nil, types.AuctionPriceInvalid
	}

	if msg.Duration < 100 {
		return nil, types.AuctionDurationInvalid
	}

	auction := types.Auction{
		Creator:    msg.Creator,
		Name:       msg.Name,
		StartPrice: msg.StartPrice,
		Duration:   msg.Duration,
		Ended:      false,
		CreatedAt:  ctx.BlockHeight(),
	}

	id := k.AppendAuction(ctx, auction)

	return &types.MsgCreateAuctionResponse{Id: id}, nil
}
