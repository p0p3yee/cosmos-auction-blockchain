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
		return nil, types.AuctionDurationPassed
	}

	msgBidPrice, _ := sdk.ParseCoinsNormalized(msg.BidPrice)
	auctionStartPrice, _ := sdk.ParseCoinsNormalized(auction.StartPrice)

	if auction.HighestBidExists {
		currentHighestBid, found := k.GetBid(ctx, auction.CurrentHighestBidId)
		if !found {
			return nil, types.InternalError
		}

		receiver, err := sdk.AccAddressFromBech32(currentHighestBid.Creator)
		if err != nil {
			return nil, err
		}

		bidPriceCoins, err := sdk.ParseCoinsNormalized(currentHighestBid.BidPrice)
		if err != nil {
			return nil, err
		}

		if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, bidPriceCoins); err != nil {
			return nil, err
		}

		auctionStartPrice = bidPriceCoins
	}

	if auctionStartPrice.IsAllGTE(msgBidPrice) {
		return nil, types.BidPriceLow
	}

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	senderBalance := k.bankKeeper.SpendableCoins(ctx, sender)

	bidPriceInCoin, err := sdk.ParseCoinsNormalized(string(msg.BidPrice))
	if err != nil {
		return nil, err
	}

	if bidPriceInCoin.IsAllGT(senderBalance) {
		return nil, types.InsufficientBalance
	}

	bid := types.Bid{
		Creator:   msg.Creator,
		AuctionId: msg.AuctionId,
		BidPrice:  msg.BidPrice,
	}

	if id, err := k.AppendBid(ctx, bid); err != nil {
		return nil, err
	} else {
		if err := k.UpdateAuctionHighestBidId(ctx, msg.AuctionId, id); err != nil {
			return nil, err
		}
		if err = k.bankKeeper.SendCoinsFromAccountToModule(
			ctx,
			sender,
			types.ModuleName,
			bidPriceInCoin,
		); err != nil {
			return nil, err
		}
		return &types.MsgPlaceBidResponse{Id: id}, nil
	}
}
