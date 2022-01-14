package protomodel

import "github.com/alpine-hodler/sdk/serial"

// * This is a generated file, do not edit

// CoinbaseSingleProduct is information on a single product
type CoinbaseSingleProduct struct {
	AuctionMode           bool    `json:"auction_mode"`
	BaseCurrency          string  `json:"base_currency"`
	BaseIncrement         float64 `json:"base_increment"`
	BaseMaxSize           float64 `json:"base_max_size"`
	BaseMinSize           float64 `json:"base_min_size"`
	CancelOnly            bool    `json:"cancel_only"`
	DisplayName           string  `json:"display_name"`
	FxStablecoin          bool    `json:"fx_stablecoin"`
	Id                    string  `json:"id"`
	LimitOnly             bool    `json:"limit_only"`
	MarginEnabled         bool    `json:"margin_enabled"`
	MaxMarketFunds        float64 `json:"max_market_funds"`
	MaxSlippagePercentage float64 `json:"max_slippage_percentage"`
	MinMarketFunds        float64 `json:"min_market_funds"`
	PostOnly              bool    `json:"post_only"`
	QuoteCurrency         string  `json:"quote_currency"`
	QuoteIncrement        float64 `json:"quote_increment"`
	Status                string  `json:"status"`
	StatusMessage         string  `json:"status_message"`
	TradingDisabled       bool    `json:"trading_disabled"`
}

// UnmarshalJSON will deserialize bytes into a CoinbaseSingleProduct model
func (coinbaseSingleProduct *CoinbaseSingleProduct) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag                    = "id"
		baseCurrencyJsonTag          = "base_currency"
		quoteCurrencyJsonTag         = "quote_currency"
		baseMinSizeJsonTag           = "base_min_size"
		baseMaxSizeJsonTag           = "base_max_size"
		quoteIncrementJsonTag        = "quote_increment"
		baseIncrementJsonTag         = "base_increment"
		displayNameJsonTag           = "display_name"
		minMarketFundsJsonTag        = "min_market_funds"
		maxMarketFundsJsonTag        = "max_market_funds"
		marginEnabledJsonTag         = "margin_enabled"
		postOnlyJsonTag              = "post_only"
		limitOnlyJsonTag             = "limit_only"
		cancelOnlyJsonTag            = "cancel_only"
		statusJsonTag                = "status"
		statusMessageJsonTag         = "status_message"
		tradingDisabledJsonTag       = "trading_disabled"
		fxStablecoinJsonTag          = "fx_stablecoin"
		maxSlippagePercentageJsonTag = "max_slippage_percentage"
		auctionModeJsonTag           = "auction_mode"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(auctionModeJsonTag, &coinbaseSingleProduct.AuctionMode)
	data.UnmarshalBool(cancelOnlyJsonTag, &coinbaseSingleProduct.CancelOnly)
	data.UnmarshalBool(fxStablecoinJsonTag, &coinbaseSingleProduct.FxStablecoin)
	data.UnmarshalBool(limitOnlyJsonTag, &coinbaseSingleProduct.LimitOnly)
	data.UnmarshalBool(marginEnabledJsonTag, &coinbaseSingleProduct.MarginEnabled)
	data.UnmarshalBool(postOnlyJsonTag, &coinbaseSingleProduct.PostOnly)
	data.UnmarshalBool(tradingDisabledJsonTag, &coinbaseSingleProduct.TradingDisabled)
	data.UnmarshalFloatFromString(baseIncrementJsonTag, &coinbaseSingleProduct.BaseIncrement)
	data.UnmarshalFloatFromString(baseMaxSizeJsonTag, &coinbaseSingleProduct.BaseMaxSize)
	data.UnmarshalFloatFromString(baseMinSizeJsonTag, &coinbaseSingleProduct.BaseMinSize)
	data.UnmarshalFloatFromString(maxMarketFundsJsonTag, &coinbaseSingleProduct.MaxMarketFunds)
	data.UnmarshalFloatFromString(maxSlippagePercentageJsonTag, &coinbaseSingleProduct.MaxSlippagePercentage)
	data.UnmarshalFloatFromString(minMarketFundsJsonTag, &coinbaseSingleProduct.MinMarketFunds)
	data.UnmarshalFloatFromString(quoteIncrementJsonTag, &coinbaseSingleProduct.QuoteIncrement)
	data.UnmarshalString(baseCurrencyJsonTag, &coinbaseSingleProduct.BaseCurrency)
	data.UnmarshalString(displayNameJsonTag, &coinbaseSingleProduct.DisplayName)
	data.UnmarshalString(idJsonTag, &coinbaseSingleProduct.Id)
	data.UnmarshalString(quoteCurrencyJsonTag, &coinbaseSingleProduct.QuoteCurrency)
	data.UnmarshalString(statusJsonTag, &coinbaseSingleProduct.Status)
	data.UnmarshalString(statusMessageJsonTag, &coinbaseSingleProduct.StatusMessage)
	return nil
}
