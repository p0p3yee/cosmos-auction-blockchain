package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/auction module sentinel errors
var (
	AuctionEnded    = sdkerrors.Register(ModuleName, 1100, "target auction already ended")
	AuctionNotFound = sdkerrors.Register(ModuleName, 1200, "target auction not found")

	BidNotFound        = sdkerrors.Register(ModuleName, 1300, "target bid not found")
	BidPriceTooLow     = sdkerrors.Register(ModuleName, 1400, "bid price is lower than target auction start price")
	BidPriceStepTooLow = sdkerrors.Register(ModuleName, 1500, "bid price increment is lower than target auction min step price")

	NotAuctionOwner = sdkerrors.Register(ModuleName, 1600, "you are not the owner of this auction")
)
