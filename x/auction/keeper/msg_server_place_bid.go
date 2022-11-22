package keeper

import (
	"context"

	"auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PlaceBid(goCtx context.Context, msg *types.MsgPlaceBid) (*types.MsgPlaceBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgPlaceBidResponse{}, nil
}
