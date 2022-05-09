package coinbase

import (
	"time"

	"github.com/alpine-hodler/sdk/pkg/scalar"
)

// * This is a generated file, do not edit

// AccountHoldsOptions are options for API requests.
type AccountHoldsOptions struct {
	After  *string `json:"after" bson:"after"`
	Before *string `json:"before" bson:"before"`
	Limit  *int    `json:"limit" bson:"limit"`
}

// AccountLedgerOptions are options for API requests.
type AccountLedgerOptions struct {
	After     *string `json:"after" bson:"after"`
	Before    *string `json:"before" bson:"before"`
	EndDate   *string `json:"end_date" bson:"end_date"`
	Limit     *int    `json:"limit" bson:"limit"`
	ProfileId *string `json:"profile_id" bson:"profile_id"`
	StartDate *string `json:"start_date" bson:"start_date"`
}

// AccountTransfersOptions are options for API requests.
type AccountTransfersOptions struct {
	After  *string `json:"after" bson:"after"`
	Before *string `json:"before" bson:"before"`
	Limit  *int    `json:"limit" bson:"limit"`
	Type   *string `json:"type" bson:"type"`
}

// BookOptions are options for API requests.
type BookOptions struct {
	// Levels 1 and 2 are aggregated. The size field is the sum of the size of the orders at that price, and num-orders is
	// the count of orders at that price; size should not be multiplied by num-orders. Level 3 is non-aggregated and returns
	// the entire order book. While the book is in an auction, the L1, L2 and L3 book will also contain the most recent
	// indicative quote disseminated during the auction, and auction_mode will be set to true. These indicative quote
	// messages are sent on an interval basis (approximately once a second) during the collection phase of an auction and
	// includes information about the tentative price and size affiliated with the completion. Level 1 and Level 2 are
	// recommended for polling. For the most up-to-date data, consider using the websocket stream. Level 3 is only
	// recommended for users wishing to maintain a full real-time order book using the websocket stream. Abuse of Level 3
	// via polling will cause your access to be limited or blocked.
	Level *int32 `json:"level" bson:"level"`
}

// ConversionsOptions are options for API requests.
type ConversionsOptions struct {
	Amount    float64 `json:"amount" bson:"amount"`
	From      string  `json:"from" bson:"from"`
	Nonce     *string `json:"nonce" bson:"nonce"`
	ProfileId *string `json:"profile_id" bson:"profile_id"`
	To        string  `json:"to" bson:"to"`
}

// ConversionOptions are options for API requests.
type ConversionOptions struct {
	ProfileId *string `json:"profile_id" bson:"profile_id"`
}

// AccountDepositOptions are options for API requests.
type AccountDepositOptions struct {
	Amount            float64 `json:"amount" bson:"amount"`
	CoinbaseAccountId string  `json:"coinbase_account_id" bson:"coinbase_account_id"`
	Currency          string  `json:"currency" bson:"currency"`
	ProfileId         *string `json:"profile_id" bson:"profile_id"`
}

// PaymentMethodDepositOptions are options for API requests.
type PaymentMethodDepositOptions struct {
	Amount          float64 `json:"amount" bson:"amount"`
	Currency        string  `json:"currency" bson:"currency"`
	PaymentMethodId string  `json:"payment_method_id" bson:"payment_method_id"`
	ProfileId       *string `json:"profile_id" bson:"profile_id"`
}

// FillsOptions are options for API requests.
type FillsOptions struct {
	After     *int    `json:"after" bson:"after"`
	Before    *int    `json:"before" bson:"before"`
	Limit     *int    `json:"limit" bson:"limit"`
	OrderId   *string `json:"order_id" bson:"order_id"`
	ProductId *string `json:"product_id" bson:"product_id"`
	ProfileId *string `json:"profile_id" bson:"profile_id"`
}

// NewOrderOptions are options for API requests.
type NewOrderOptions struct {
	CancelAfter *scalar.CancelAfter `json:"cancel_after" bson:"cancel_after"`
	ClientOid   *string             `json:"client_oid" bson:"client_oid"`
	Funds       *float64            `json:"funds" bson:"funds"`
	PostOnly    *bool               `json:"post_only" bson:"post_only"`
	Price       *float64            `json:"price" bson:"price"`
	ProductId   string              `json:"product_id" bson:"product_id"`
	ProfileId   *string             `json:"profile_id" bson:"profile_id"`
	Side        scalar.OrderSide    `json:"side" bson:"side"`
	Size        *float64            `json:"size" bson:"size"`
	Stop        *scalar.OrderStop   `json:"stop" bson:"stop"`
	StopPrice   *float64            `json:"stop_price" bson:"stop_price"`
	Stp         *scalar.OrderSTP    `json:"stp" bson:"stp"`
	TimeInForce *scalar.TimeInForce `json:"time_in_force" bson:"time_in_force"`
	Type        *scalar.OrderType   `json:"type" bson:"type"`
}

// OrdersOptions are options for API requests.
type OrdersOptions struct {
	After     *string    `json:"after" bson:"after"`
	Before    *string    `json:"before" bson:"before"`
	EndDate   *time.Time `json:"end_date" bson:"end_date"`
	Limit     int        `json:"limit" bson:"limit"`
	ProductId *string    `json:"product_id" bson:"product_id"`
	ProfileId *string    `json:"profile_id" bson:"profile_id"`
	SortedBy  *string    `json:"sortedBy" bson:"sortedBy"`
	Sorting   *string    `json:"sorting" bson:"sorting"`
	StartDate *time.Time `json:"start_date" bson:"start_date"`
	Status    []string   `json:"status" bson:"status"`
}

// ProductsOptions are options for API requests.
type ProductsOptions struct {
	Type *string `json:"type" bson:"type"`
}

// ProfilesOptions are options for API requests.
type ProfilesOptions struct {
	Active *bool `json:"active" bson:"active"`
}

// AccountWithdrawalOptions are options for API requests.
type AccountWithdrawalOptions struct {
	Amount            float64 `json:"amount" bson:"amount"`
	CoinbaseAccountId string  `json:"coinbase_account_id" bson:"coinbase_account_id"`
	Currency          string  `json:"currency" bson:"currency"`
	ProfileId         *string `json:"profile_id" bson:"profile_id"`
}

// CryptoWithdrawalOptions are options for API requests.
type CryptoWithdrawalOptions struct {
	Amount           float64  `json:"amount" bson:"amount"`
	CryptoAddress    string   `json:"crypto_address" bson:"crypto_address"`
	Currency         string   `json:"currency" bson:"currency"`
	DestinationTag   *string  `json:"destination_tag" bson:"destination_tag"`
	Fee              *float64 `json:"fee" bson:"fee"`
	NoDestinationTag *bool    `json:"no_destination_tag" bson:"no_destination_tag"`
	Nonce            *int     `json:"nonce" bson:"nonce"`
	ProfileId        *string  `json:"profile_id" bson:"profile_id"`
	TwoFactorCode    *string  `json:"two_factor_code" bson:"two_factor_code"`
}

// PaymentMethodWithdrawalOptions are options for API requests.
type PaymentMethodWithdrawalOptions struct {
	Amount          float64 `json:"amount" bson:"amount"`
	Currency        string  `json:"currency" bson:"currency"`
	PaymentMethodId string  `json:"payment_method_id" bson:"payment_method_id"`
	ProfileId       *string `json:"profile_id" bson:"profile_id"`
}

// WithdrawalFeeEstimateOptions are options for API requests.
type WithdrawalFeeEstimateOptions struct {
	CryptoAddress *string `json:"crypto_address" bson:"crypto_address"`
	Currency      *string `json:"currency" bson:"currency"`
}

// SetBefore sets the Before field on AccountHoldsOptions.
func (opts *AccountHoldsOptions) SetBefore(before string) *AccountHoldsOptions {
	opts.Before = &before
	return opts
}

// SetAfter sets the After field on AccountHoldsOptions.
func (opts *AccountHoldsOptions) SetAfter(after string) *AccountHoldsOptions {
	opts.After = &after
	return opts
}

// SetLimit sets the Limit field on AccountHoldsOptions.
func (opts *AccountHoldsOptions) SetLimit(limit int) *AccountHoldsOptions {
	opts.Limit = &limit
	return opts
}

// SetStartDate sets the StartDate field on AccountLedgerOptions.
func (opts *AccountLedgerOptions) SetStartDate(startDate string) *AccountLedgerOptions {
	opts.StartDate = &startDate
	return opts
}

// SetEndDate sets the EndDate field on AccountLedgerOptions.
func (opts *AccountLedgerOptions) SetEndDate(endDate string) *AccountLedgerOptions {
	opts.EndDate = &endDate
	return opts
}

// SetBefore sets the Before field on AccountLedgerOptions.
func (opts *AccountLedgerOptions) SetBefore(before string) *AccountLedgerOptions {
	opts.Before = &before
	return opts
}

// SetAfter sets the After field on AccountLedgerOptions.
func (opts *AccountLedgerOptions) SetAfter(after string) *AccountLedgerOptions {
	opts.After = &after
	return opts
}

// SetProfileId sets the ProfileId field on AccountLedgerOptions.
func (opts *AccountLedgerOptions) SetProfileId(profileId string) *AccountLedgerOptions {
	opts.ProfileId = &profileId
	return opts
}

// SetLimit sets the Limit field on AccountLedgerOptions.
func (opts *AccountLedgerOptions) SetLimit(limit int) *AccountLedgerOptions {
	opts.Limit = &limit
	return opts
}

// SetBefore sets the Before field on AccountTransfersOptions.
func (opts *AccountTransfersOptions) SetBefore(before string) *AccountTransfersOptions {
	opts.Before = &before
	return opts
}

// SetAfter sets the After field on AccountTransfersOptions.
func (opts *AccountTransfersOptions) SetAfter(after string) *AccountTransfersOptions {
	opts.After = &after
	return opts
}

// SetLimit sets the Limit field on AccountTransfersOptions.
func (opts *AccountTransfersOptions) SetLimit(limit int) *AccountTransfersOptions {
	opts.Limit = &limit
	return opts
}

// SetType sets the Type field on AccountTransfersOptions.
func (opts *AccountTransfersOptions) SetType(typ string) *AccountTransfersOptions {
	opts.Type = &typ
	return opts
}

// SetLevel sets the Level field on BookOptions.
func (opts *BookOptions) SetLevel(level int32) *BookOptions {
	opts.Level = &level
	return opts
}

// SetProfileId sets the ProfileId field on ConversionsOptions.
func (opts *ConversionsOptions) SetProfileId(profileId string) *ConversionsOptions {
	opts.ProfileId = &profileId
	return opts
}

// SetFrom sets the From field on ConversionsOptions.
func (opts *ConversionsOptions) SetFrom(from string) *ConversionsOptions {
	opts.From = from
	return opts
}

// SetTo sets the To field on ConversionsOptions.
func (opts *ConversionsOptions) SetTo(to string) *ConversionsOptions {
	opts.To = to
	return opts
}

// SetAmount sets the Amount field on ConversionsOptions.
func (opts *ConversionsOptions) SetAmount(amount float64) *ConversionsOptions {
	opts.Amount = amount
	return opts
}

// SetNonce sets the Nonce field on ConversionsOptions.
func (opts *ConversionsOptions) SetNonce(nonce string) *ConversionsOptions {
	opts.Nonce = &nonce
	return opts
}

// SetProfileId sets the ProfileId field on ConversionOptions.
func (opts *ConversionOptions) SetProfileId(profileId string) *ConversionOptions {
	opts.ProfileId = &profileId
	return opts
}

// SetProfileId sets the ProfileId field on AccountDepositOptions.
func (opts *AccountDepositOptions) SetProfileId(profileId string) *AccountDepositOptions {
	opts.ProfileId = &profileId
	return opts
}

// SetAmount sets the Amount field on AccountDepositOptions.
func (opts *AccountDepositOptions) SetAmount(amount float64) *AccountDepositOptions {
	opts.Amount = amount
	return opts
}

// SetCoinbaseAccountId sets the CoinbaseAccountId field on AccountDepositOptions.
func (opts *AccountDepositOptions) SetCoinbaseAccountId(coinbaseAccountId string) *AccountDepositOptions {
	opts.CoinbaseAccountId = coinbaseAccountId
	return opts
}

// SetCurrency sets the Currency field on AccountDepositOptions.
func (opts *AccountDepositOptions) SetCurrency(currency string) *AccountDepositOptions {
	opts.Currency = currency
	return opts
}

// SetProfileId sets the ProfileId field on PaymentMethodDepositOptions.
func (opts *PaymentMethodDepositOptions) SetProfileId(profileId string) *PaymentMethodDepositOptions {
	opts.ProfileId = &profileId
	return opts
}

// SetAmount sets the Amount field on PaymentMethodDepositOptions.
func (opts *PaymentMethodDepositOptions) SetAmount(amount float64) *PaymentMethodDepositOptions {
	opts.Amount = amount
	return opts
}

// SetPaymentMethodId sets the PaymentMethodId field on PaymentMethodDepositOptions.
func (opts *PaymentMethodDepositOptions) SetPaymentMethodId(paymentMethodId string) *PaymentMethodDepositOptions {
	opts.PaymentMethodId = paymentMethodId
	return opts
}

// SetCurrency sets the Currency field on PaymentMethodDepositOptions.
func (opts *PaymentMethodDepositOptions) SetCurrency(currency string) *PaymentMethodDepositOptions {
	opts.Currency = currency
	return opts
}

// SetOrderId sets the OrderId field on FillsOptions.
func (opts *FillsOptions) SetOrderId(orderId string) *FillsOptions {
	opts.OrderId = &orderId
	return opts
}

// SetProductId sets the ProductId field on FillsOptions.
func (opts *FillsOptions) SetProductId(productId string) *FillsOptions {
	opts.ProductId = &productId
	return opts
}

// SetProfileId sets the ProfileId field on FillsOptions.
func (opts *FillsOptions) SetProfileId(profileId string) *FillsOptions {
	opts.ProfileId = &profileId
	return opts
}

// SetLimit sets the Limit field on FillsOptions.
func (opts *FillsOptions) SetLimit(limit int) *FillsOptions {
	opts.Limit = &limit
	return opts
}

// SetBefore sets the Before field on FillsOptions.
func (opts *FillsOptions) SetBefore(before int) *FillsOptions {
	opts.Before = &before
	return opts
}

// SetAfter sets the After field on FillsOptions.
func (opts *FillsOptions) SetAfter(after int) *FillsOptions {
	opts.After = &after
	return opts
}

// SetProfileId sets the ProfileId field on NewOrderOptions.
func (opts *NewOrderOptions) SetProfileId(profileId string) *NewOrderOptions {
	opts.ProfileId = &profileId
	return opts
}

// SetType sets the Type field on NewOrderOptions.
func (opts *NewOrderOptions) SetType(typ scalar.OrderType) *NewOrderOptions {
	opts.Type = &typ
	return opts
}

// SetSide sets the Side field on NewOrderOptions.
func (opts *NewOrderOptions) SetSide(side scalar.OrderSide) *NewOrderOptions {
	opts.Side = side
	return opts
}

// SetStp sets the Stp field on NewOrderOptions.
func (opts *NewOrderOptions) SetStp(stp scalar.OrderSTP) *NewOrderOptions {
	opts.Stp = &stp
	return opts
}

// SetStop sets the Stop field on NewOrderOptions.
func (opts *NewOrderOptions) SetStop(stop scalar.OrderStop) *NewOrderOptions {
	opts.Stop = &stop
	return opts
}

// SetStopPrice sets the StopPrice field on NewOrderOptions.
func (opts *NewOrderOptions) SetStopPrice(stopPrice float64) *NewOrderOptions {
	opts.StopPrice = &stopPrice
	return opts
}

// SetPrice sets the Price field on NewOrderOptions.
func (opts *NewOrderOptions) SetPrice(price float64) *NewOrderOptions {
	opts.Price = &price
	return opts
}

// SetSize sets the Size field on NewOrderOptions.
func (opts *NewOrderOptions) SetSize(size float64) *NewOrderOptions {
	opts.Size = &size
	return opts
}

// SetFunds sets the Funds field on NewOrderOptions.
func (opts *NewOrderOptions) SetFunds(funds float64) *NewOrderOptions {
	opts.Funds = &funds
	return opts
}

// SetProductId sets the ProductId field on NewOrderOptions.
func (opts *NewOrderOptions) SetProductId(productId string) *NewOrderOptions {
	opts.ProductId = productId
	return opts
}

// SetTimeInForce sets the TimeInForce field on NewOrderOptions.
func (opts *NewOrderOptions) SetTimeInForce(timeInForce scalar.TimeInForce) *NewOrderOptions {
	opts.TimeInForce = &timeInForce
	return opts
}

// SetCancelAfter sets the CancelAfter field on NewOrderOptions.
func (opts *NewOrderOptions) SetCancelAfter(cancelAfter scalar.CancelAfter) *NewOrderOptions {
	opts.CancelAfter = &cancelAfter
	return opts
}

// SetPostOnly sets the PostOnly field on NewOrderOptions.
func (opts *NewOrderOptions) SetPostOnly(postOnly bool) *NewOrderOptions {
	opts.PostOnly = &postOnly
	return opts
}

// SetClientOid sets the ClientOid field on NewOrderOptions.
func (opts *NewOrderOptions) SetClientOid(clientOid string) *NewOrderOptions {
	opts.ClientOid = &clientOid
	return opts
}

// SetProfileId sets the ProfileId field on OrdersOptions.
func (opts *OrdersOptions) SetProfileId(profileId string) *OrdersOptions {
	opts.ProfileId = &profileId
	return opts
}

// SetProductId sets the ProductId field on OrdersOptions.
func (opts *OrdersOptions) SetProductId(productId string) *OrdersOptions {
	opts.ProductId = &productId
	return opts
}

// SetSortedBy sets the SortedBy field on OrdersOptions.
func (opts *OrdersOptions) SetSortedBy(sortedBy string) *OrdersOptions {
	opts.SortedBy = &sortedBy
	return opts
}

// SetSorting sets the Sorting field on OrdersOptions.
func (opts *OrdersOptions) SetSorting(sorting string) *OrdersOptions {
	opts.Sorting = &sorting
	return opts
}

// SetStartDate sets the StartDate field on OrdersOptions.
func (opts *OrdersOptions) SetStartDate(startDate time.Time) *OrdersOptions {
	opts.StartDate = &startDate
	return opts
}

// SetEndDate sets the EndDate field on OrdersOptions.
func (opts *OrdersOptions) SetEndDate(endDate time.Time) *OrdersOptions {
	opts.EndDate = &endDate
	return opts
}

// SetBefore sets the Before field on OrdersOptions.
func (opts *OrdersOptions) SetBefore(before string) *OrdersOptions {
	opts.Before = &before
	return opts
}

// SetAfter sets the After field on OrdersOptions.
func (opts *OrdersOptions) SetAfter(after string) *OrdersOptions {
	opts.After = &after
	return opts
}

// SetLimit sets the Limit field on OrdersOptions.
func (opts *OrdersOptions) SetLimit(limit int) *OrdersOptions {
	opts.Limit = limit
	return opts
}

// SetStatus sets the Status field on OrdersOptions.
func (opts *OrdersOptions) SetStatus(status []string) *OrdersOptions {
	opts.Status = status
	return opts
}

// SetType sets the Type field on ProductsOptions.
func (opts *ProductsOptions) SetType(typ string) *ProductsOptions {
	opts.Type = &typ
	return opts
}

// SetActive sets the Active field on ProfilesOptions.
func (opts *ProfilesOptions) SetActive(active bool) *ProfilesOptions {
	opts.Active = &active
	return opts
}

// SetProfileId sets the ProfileId field on AccountWithdrawalOptions.
func (opts *AccountWithdrawalOptions) SetProfileId(profileId string) *AccountWithdrawalOptions {
	opts.ProfileId = &profileId
	return opts
}

// SetAmount sets the Amount field on AccountWithdrawalOptions.
func (opts *AccountWithdrawalOptions) SetAmount(amount float64) *AccountWithdrawalOptions {
	opts.Amount = amount
	return opts
}

// SetCoinbaseAccountId sets the CoinbaseAccountId field on AccountWithdrawalOptions.
func (opts *AccountWithdrawalOptions) SetCoinbaseAccountId(coinbaseAccountId string) *AccountWithdrawalOptions {
	opts.CoinbaseAccountId = coinbaseAccountId
	return opts
}

// SetCurrency sets the Currency field on AccountWithdrawalOptions.
func (opts *AccountWithdrawalOptions) SetCurrency(currency string) *AccountWithdrawalOptions {
	opts.Currency = currency
	return opts
}

// SetProfileId sets the ProfileId field on CryptoWithdrawalOptions.
func (opts *CryptoWithdrawalOptions) SetProfileId(profileId string) *CryptoWithdrawalOptions {
	opts.ProfileId = &profileId
	return opts
}

// SetAmount sets the Amount field on CryptoWithdrawalOptions.
func (opts *CryptoWithdrawalOptions) SetAmount(amount float64) *CryptoWithdrawalOptions {
	opts.Amount = amount
	return opts
}

// SetCryptoAddress sets the CryptoAddress field on CryptoWithdrawalOptions.
func (opts *CryptoWithdrawalOptions) SetCryptoAddress(cryptoAddress string) *CryptoWithdrawalOptions {
	opts.CryptoAddress = cryptoAddress
	return opts
}

// SetCurrency sets the Currency field on CryptoWithdrawalOptions.
func (opts *CryptoWithdrawalOptions) SetCurrency(currency string) *CryptoWithdrawalOptions {
	opts.Currency = currency
	return opts
}

// SetDestinationTag sets the DestinationTag field on CryptoWithdrawalOptions.
func (opts *CryptoWithdrawalOptions) SetDestinationTag(destinationTag string) *CryptoWithdrawalOptions {
	opts.DestinationTag = &destinationTag
	return opts
}

// SetNoDestinationTag sets the NoDestinationTag field on CryptoWithdrawalOptions.
func (opts *CryptoWithdrawalOptions) SetNoDestinationTag(noDestinationTag bool) *CryptoWithdrawalOptions {
	opts.NoDestinationTag = &noDestinationTag
	return opts
}

// SetTwoFactorCode sets the TwoFactorCode field on CryptoWithdrawalOptions.
func (opts *CryptoWithdrawalOptions) SetTwoFactorCode(twoFactorCode string) *CryptoWithdrawalOptions {
	opts.TwoFactorCode = &twoFactorCode
	return opts
}

// SetNonce sets the Nonce field on CryptoWithdrawalOptions.
func (opts *CryptoWithdrawalOptions) SetNonce(nonce int) *CryptoWithdrawalOptions {
	opts.Nonce = &nonce
	return opts
}

// SetFee sets the Fee field on CryptoWithdrawalOptions.
func (opts *CryptoWithdrawalOptions) SetFee(fee float64) *CryptoWithdrawalOptions {
	opts.Fee = &fee
	return opts
}

// SetProfileId sets the ProfileId field on PaymentMethodWithdrawalOptions.
func (opts *PaymentMethodWithdrawalOptions) SetProfileId(profileId string) *PaymentMethodWithdrawalOptions {
	opts.ProfileId = &profileId
	return opts
}

// SetAmount sets the Amount field on PaymentMethodWithdrawalOptions.
func (opts *PaymentMethodWithdrawalOptions) SetAmount(amount float64) *PaymentMethodWithdrawalOptions {
	opts.Amount = amount
	return opts
}

// SetPaymentMethodId sets the PaymentMethodId field on PaymentMethodWithdrawalOptions.
func (opts *PaymentMethodWithdrawalOptions) SetPaymentMethodId(paymentMethodId string) *PaymentMethodWithdrawalOptions {
	opts.PaymentMethodId = paymentMethodId
	return opts
}

// SetCurrency sets the Currency field on PaymentMethodWithdrawalOptions.
func (opts *PaymentMethodWithdrawalOptions) SetCurrency(currency string) *PaymentMethodWithdrawalOptions {
	opts.Currency = currency
	return opts
}

// SetCurrency sets the Currency field on WithdrawalFeeEstimateOptions.
func (opts *WithdrawalFeeEstimateOptions) SetCurrency(currency string) *WithdrawalFeeEstimateOptions {
	opts.Currency = &currency
	return opts
}

// SetCryptoAddress sets the CryptoAddress field on WithdrawalFeeEstimateOptions.
func (opts *WithdrawalFeeEstimateOptions) SetCryptoAddress(cryptoAddress string) *WithdrawalFeeEstimateOptions {
	opts.CryptoAddress = &cryptoAddress
	return opts
}
