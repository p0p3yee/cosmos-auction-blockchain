package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/auction module sentinel errors
var (
	AuctionEnded           = sdkerrors.Register(ModuleName, 1100, "target auction already ended")
	AuctionNotFound        = sdkerrors.Register(ModuleName, 1200, "target auction not found")
	AuctionPriceInvalid    = sdkerrors.Register(ModuleName, 1300, "auction start price must larger than 0")
	AuctionDurationInvalid = sdkerrors.Register(ModuleName, 1400, "auction duration must be at least 100")

	BidNotFound    = sdkerrors.Register(ModuleName, 1500, "target bid not found")
	BidPriceTooLow = sdkerrors.Register(ModuleName, 1600, "bid price is lower than target auction start price")

	NotAuctionOwner = sdkerrors.Register(ModuleName, 1700, "you are not the owner of this auction")
)
