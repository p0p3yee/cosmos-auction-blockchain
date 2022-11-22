package keeper

import (
	"context"

	"auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAuction(goCtx context.Context, msg *types.MsgCreateAuction) (*types.MsgCreateAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	auction := types.Auction{
		Creator:      msg.Creator,
		Name:         msg.Name,
		StartPrice:   msg.StartPrice,
		MinStepPrice: msg.MinPriceStep,
	}

	id := k.AppendAuction(ctx, auction)

	return &types.MsgCreateAuctionResponse{Id: id}, nil
}
