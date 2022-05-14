package coinbase

import (
	"time"

	"github.com/alpine-hodler/sdk/internal/client"
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

// CandlesOptions are options for API requests.
type CandlesOptions struct {
	End         *time.Time          `json:"end" bson:"end"`
	Granularity *scalar.Granularity `json:"granularity" bson:"granularity"`
	Start       *time.Time          `json:"start" bson:"start"`
}

// CreateOrderOptions are options for API requests.
type CreateOrderOptions struct {
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

// CreateReportOptions are options for API requests.
type CreateReportOptions struct {
	// Account - required for account-type reports
	AccountId *string `json:"account_id" bson:"account_id"`

	// Email to send generated report to
	Email *string `json:"email" bson:"email"`

	// End date for items to be included in report
	EndDate *string `json:"end_date" bson:"end_date"`

	// Portfolio - Which portfolio to generate the report for
	ProfileId *string `json:"profile_id" bson:"profile_id"`

	// Product - required for fills-type reports
	ProductId *string `json:"product_id" bson:"product_id"`

	// Start date for items to be included in report.
	StartDate *string `json:"start_date" bson:"start_date"`

	// required for 1099k-transaction-history-type reports
	Year   *string           `json:"year" bson:"year"`
	Format *scalar.Format    `json:"format" bson:"format"`
	Type   scalar.ReportType `json:"type" bson:"type"`
}

// ConvertCurrencyOptions are options for API requests.
type ConvertCurrencyOptions struct {
	Amount    float64 `json:"amount" bson:"amount"`
	From      string  `json:"from" bson:"from"`
	Nonce     *string `json:"nonce" bson:"nonce"`
	ProfileId *string `json:"profile_id" bson:"profile_id"`
	To        string  `json:"to" bson:"to"`
}

// CurrencyConversionOptions are options for API requests.
type CurrencyConversionOptions struct {
	ProfileId *string `json:"profile_id" bson:"profile_id"`
}

// CoinbaseAccountDepositOptions are options for API requests.
type CoinbaseAccountDepositOptions struct {
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

// CancelOpenOrdersOptions are options for API requests.
type CancelOpenOrdersOptions struct {
	ProductId *string `json:"product_id" bson:"product_id"`
	ProfileId *string `json:"profile_id" bson:"profile_id"`
}

// ProductsOptions are options for API requests.
type ProductsOptions struct {
	Type *string `json:"type" bson:"type"`
}

// DeleteProfileOptions are options for API requests.
type DeleteProfileOptions struct {
	ProfileId *string `json:"profile_id" bson:"profile_id"`
	To        *string `json:"to" bson:"to"`
}

// ProfileOptions are options for API requests.
type ProfileOptions struct {
	Active *bool `json:"active" bson:"active"`
}

// RenameProfileOptions are options for API requests.
type RenameProfileOptions struct {
	Name      *string `json:"name" bson:"name"`
	ProfileId *string `json:"profile_id" bson:"profile_id"`
}

// ProfilesOptions are options for API requests.
type ProfilesOptions struct {
	Active *bool `json:"active" bson:"active"`
}

// CreateProfileOptions are options for API requests.
type CreateProfileOptions struct {
	Name *string `json:"name" bson:"name"`
}

// CreateProfileTransferOptions are options for API requests.
type CreateProfileTransferOptions struct {
	Amount   *string `json:"amount" bson:"amount"`
	Currency *string `json:"currency" bson:"currency"`
	From     *string `json:"from" bson:"from"`
	To       *string `json:"to" bson:"to"`
}

// ReportsOptions are options for API requests.
type ReportsOptions struct {
	// Filter results after a specific date
	After *time.Time `json:"after" bson:"after"`

	// Filter results by a specific profile_id
	PortfolioId *string `json:"portfolio_id" bson:"portfolio_id"`

	// Filter results by type of report (fills or account) - otc_fills: real string is otc-fills -
	// type_1099k_transaction_history: real string is 1099-transaction-history - tax_invoice: real string is tax-invoice
	Type *scalar.ReportType `json:"type" bson:"type"`

	// Ignore expired results
	IgnoredExpired *bool `json:"ignored_expired" bson:"ignored_expired"`

	// Limit results to a specific number
	Limit *int `json:"limit" bson:"limit"`
}

// TradesOptions are options for API requests.
type TradesOptions struct {
	After  *int32 `json:"after" bson:"after"`
	Before *int32 `json:"before" bson:"before"`
	Limit  *int32 `json:"limit" bson:"limit"`
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

// SetLevel sets the Level field on BookOptions.
func (opts *BookOptions) SetLevel(level int32) *BookOptions {
	opts.Level = &level
	return opts
}

// SetProfileId sets the ProfileId field on CancelOpenOrdersOptions.
func (opts *CancelOpenOrdersOptions) SetProfileId(profileId string) *CancelOpenOrdersOptions {
	opts.ProfileId = &profileId
	return opts
}

// SetProductId sets the ProductId field on CancelOpenOrdersOptions.
func (opts *CancelOpenOrdersOptions) SetProductId(productId string) *CancelOpenOrdersOptions {
	opts.ProductId = &productId
	return opts
}

// SetGranularity sets the Granularity field on CandlesOptions.
func (opts *CandlesOptions) SetGranularity(granularity scalar.Granularity) *CandlesOptions {
	opts.Granularity = &granularity
	return opts
}

// SetStart sets the Start field on CandlesOptions.
func (opts *CandlesOptions) SetStart(start time.Time) *CandlesOptions {
	opts.Start = &start
	return opts
}

// SetEnd sets the End field on CandlesOptions.
func (opts *CandlesOptions) SetEnd(end time.Time) *CandlesOptions {
	opts.End = &end
	return opts
}

// SetProfileId sets the ProfileId field on CoinbaseAccountDepositOptions.
func (opts *CoinbaseAccountDepositOptions) SetProfileId(profileId string) *CoinbaseAccountDepositOptions {
	opts.ProfileId = &profileId
	return opts
}

// SetAmount sets the Amount field on CoinbaseAccountDepositOptions.
func (opts *CoinbaseAccountDepositOptions) SetAmount(amount float64) *CoinbaseAccountDepositOptions {
	opts.Amount = amount
	return opts
}

// SetCoinbaseAccountId sets the CoinbaseAccountId field on CoinbaseAccountDepositOptions.
func (opts *CoinbaseAccountDepositOptions) SetCoinbaseAccountId(coinbaseAccountId string) *CoinbaseAccountDepositOptions {
	opts.CoinbaseAccountId = coinbaseAccountId
	return opts
}

// SetCurrency sets the Currency field on CoinbaseAccountDepositOptions.
func (opts *CoinbaseAccountDepositOptions) SetCurrency(currency string) *CoinbaseAccountDepositOptions {
	opts.Currency = currency
	return opts
}

// SetProfileId sets the ProfileId field on ConvertCurrencyOptions.
func (opts *ConvertCurrencyOptions) SetProfileId(profileId string) *ConvertCurrencyOptions {
	opts.ProfileId = &profileId
	return opts
}

// SetFrom sets the From field on ConvertCurrencyOptions.
func (opts *ConvertCurrencyOptions) SetFrom(from string) *ConvertCurrencyOptions {
	opts.From = from
	return opts
}

// SetTo sets the To field on ConvertCurrencyOptions.
func (opts *ConvertCurrencyOptions) SetTo(to string) *ConvertCurrencyOptions {
	opts.To = to
	return opts
}

// SetAmount sets the Amount field on ConvertCurrencyOptions.
func (opts *ConvertCurrencyOptions) SetAmount(amount float64) *ConvertCurrencyOptions {
	opts.Amount = amount
	return opts
}

// SetNonce sets the Nonce field on ConvertCurrencyOptions.
func (opts *ConvertCurrencyOptions) SetNonce(nonce string) *ConvertCurrencyOptions {
	opts.Nonce = &nonce
	return opts
}

// SetProfileId sets the ProfileId field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetProfileId(profileId string) *CreateOrderOptions {
	opts.ProfileId = &profileId
	return opts
}

// SetType sets the Type field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetType(typ scalar.OrderType) *CreateOrderOptions {
	opts.Type = &typ
	return opts
}

// SetSide sets the Side field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetSide(side scalar.OrderSide) *CreateOrderOptions {
	opts.Side = side
	return opts
}

// SetStp sets the Stp field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetStp(stp scalar.OrderSTP) *CreateOrderOptions {
	opts.Stp = &stp
	return opts
}

// SetStop sets the Stop field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetStop(stop scalar.OrderStop) *CreateOrderOptions {
	opts.Stop = &stop
	return opts
}

// SetStopPrice sets the StopPrice field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetStopPrice(stopPrice float64) *CreateOrderOptions {
	opts.StopPrice = &stopPrice
	return opts
}

// SetPrice sets the Price field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetPrice(price float64) *CreateOrderOptions {
	opts.Price = &price
	return opts
}

// SetSize sets the Size field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetSize(size float64) *CreateOrderOptions {
	opts.Size = &size
	return opts
}

// SetFunds sets the Funds field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetFunds(funds float64) *CreateOrderOptions {
	opts.Funds = &funds
	return opts
}

// SetProductId sets the ProductId field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetProductId(productId string) *CreateOrderOptions {
	opts.ProductId = productId
	return opts
}

// SetTimeInForce sets the TimeInForce field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetTimeInForce(timeInForce scalar.TimeInForce) *CreateOrderOptions {
	opts.TimeInForce = &timeInForce
	return opts
}

// SetCancelAfter sets the CancelAfter field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetCancelAfter(cancelAfter scalar.CancelAfter) *CreateOrderOptions {
	opts.CancelAfter = &cancelAfter
	return opts
}

// SetPostOnly sets the PostOnly field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetPostOnly(postOnly bool) *CreateOrderOptions {
	opts.PostOnly = &postOnly
	return opts
}

// SetClientOid sets the ClientOid field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetClientOid(clientOid string) *CreateOrderOptions {
	opts.ClientOid = &clientOid
	return opts
}

// SetName sets the Name field on CreateProfileOptions.
func (opts *CreateProfileOptions) SetName(name string) *CreateProfileOptions {
	opts.Name = &name
	return opts
}

// SetFrom sets the From field on CreateProfileTransferOptions.
func (opts *CreateProfileTransferOptions) SetFrom(from string) *CreateProfileTransferOptions {
	opts.From = &from
	return opts
}

// SetTo sets the To field on CreateProfileTransferOptions.
func (opts *CreateProfileTransferOptions) SetTo(to string) *CreateProfileTransferOptions {
	opts.To = &to
	return opts
}

// SetCurrency sets the Currency field on CreateProfileTransferOptions.
func (opts *CreateProfileTransferOptions) SetCurrency(currency string) *CreateProfileTransferOptions {
	opts.Currency = &currency
	return opts
}

// SetAmount sets the Amount field on CreateProfileTransferOptions.
func (opts *CreateProfileTransferOptions) SetAmount(amount string) *CreateProfileTransferOptions {
	opts.Amount = &amount
	return opts
}

// SetStartDate sets the StartDate field on CreateReportOptions.
func (opts *CreateReportOptions) SetStartDate(startDate string) *CreateReportOptions {
	opts.StartDate = &startDate
	return opts
}

// SetEndDate sets the EndDate field on CreateReportOptions.
func (opts *CreateReportOptions) SetEndDate(endDate string) *CreateReportOptions {
	opts.EndDate = &endDate
	return opts
}

// SetType sets the Type field on CreateReportOptions.
func (opts *CreateReportOptions) SetType(typ scalar.ReportType) *CreateReportOptions {
	opts.Type = typ
	return opts
}

// SetYear sets the Year field on CreateReportOptions.
func (opts *CreateReportOptions) SetYear(year string) *CreateReportOptions {
	opts.Year = &year
	return opts
}

// SetFormat sets the Format field on CreateReportOptions.
func (opts *CreateReportOptions) SetFormat(format scalar.Format) *CreateReportOptions {
	opts.Format = &format
	return opts
}

// SetProductId sets the ProductId field on CreateReportOptions.
func (opts *CreateReportOptions) SetProductId(productId string) *CreateReportOptions {
	opts.ProductId = &productId
	return opts
}

// SetAccountId sets the AccountId field on CreateReportOptions.
func (opts *CreateReportOptions) SetAccountId(accountId string) *CreateReportOptions {
	opts.AccountId = &accountId
	return opts
}

// SetEmail sets the Email field on CreateReportOptions.
func (opts *CreateReportOptions) SetEmail(email string) *CreateReportOptions {
	opts.Email = &email
	return opts
}

// SetProfileId sets the ProfileId field on CreateReportOptions.
func (opts *CreateReportOptions) SetProfileId(profileId string) *CreateReportOptions {
	opts.ProfileId = &profileId
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

// SetProfileId sets the ProfileId field on CurrencyConversionOptions.
func (opts *CurrencyConversionOptions) SetProfileId(profileId string) *CurrencyConversionOptions {
	opts.ProfileId = &profileId
	return opts
}

// SetProfileId sets the ProfileId field on DeleteProfileOptions.
func (opts *DeleteProfileOptions) SetProfileId(profileId string) *DeleteProfileOptions {
	opts.ProfileId = &profileId
	return opts
}

// SetTo sets the To field on DeleteProfileOptions.
func (opts *DeleteProfileOptions) SetTo(to string) *DeleteProfileOptions {
	opts.To = &to
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

// SetType sets the Type field on ProductsOptions.
func (opts *ProductsOptions) SetType(typ string) *ProductsOptions {
	opts.Type = &typ
	return opts
}

// SetActive sets the Active field on ProfileOptions.
func (opts *ProfileOptions) SetActive(active bool) *ProfileOptions {
	opts.Active = &active
	return opts
}

// SetActive sets the Active field on ProfilesOptions.
func (opts *ProfilesOptions) SetActive(active bool) *ProfilesOptions {
	opts.Active = &active
	return opts
}

// SetProfileId sets the ProfileId field on RenameProfileOptions.
func (opts *RenameProfileOptions) SetProfileId(profileId string) *RenameProfileOptions {
	opts.ProfileId = &profileId
	return opts
}

// SetName sets the Name field on RenameProfileOptions.
func (opts *RenameProfileOptions) SetName(name string) *RenameProfileOptions {
	opts.Name = &name
	return opts
}

// SetPortfolioId sets the PortfolioId field on ReportsOptions.
func (opts *ReportsOptions) SetPortfolioId(portfolioId string) *ReportsOptions {
	opts.PortfolioId = &portfolioId
	return opts
}

// SetAfter sets the After field on ReportsOptions.
func (opts *ReportsOptions) SetAfter(after time.Time) *ReportsOptions {
	opts.After = &after
	return opts
}

// SetLimit sets the Limit field on ReportsOptions.
func (opts *ReportsOptions) SetLimit(limit int) *ReportsOptions {
	opts.Limit = &limit
	return opts
}

// SetType sets the Type field on ReportsOptions.
func (opts *ReportsOptions) SetType(typ scalar.ReportType) *ReportsOptions {
	opts.Type = &typ
	return opts
}

// SetIgnoredExpired sets the IgnoredExpired field on ReportsOptions.
func (opts *ReportsOptions) SetIgnoredExpired(ignoredExpired bool) *ReportsOptions {
	opts.IgnoredExpired = &ignoredExpired
	return opts
}

// SetLimit sets the Limit field on TradesOptions.
func (opts *TradesOptions) SetLimit(limit int32) *TradesOptions {
	opts.Limit = &limit
	return opts
}

// SetBefore sets the Before field on TradesOptions.
func (opts *TradesOptions) SetBefore(before int32) *TradesOptions {
	opts.Before = &before
	return opts
}

// SetAfter sets the After field on TradesOptions.
func (opts *TradesOptions) SetAfter(after int32) *TradesOptions {
	opts.After = &after
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

func (opts *AccountWithdrawalOptions) SetBody(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetBodyString("profile_id", opts.ProfileId).
		SetBodyFloat("amount", &opts.Amount).
		SetBodyString("coinbase_account_id", &opts.CoinbaseAccountId).
		SetBodyString("currency", &opts.Currency)
}

func (opts *CoinbaseAccountDepositOptions) SetBody(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetBodyString("profile_id", opts.ProfileId).
		SetBodyFloat("amount", &opts.Amount).
		SetBodyString("coinbase_account_id", &opts.CoinbaseAccountId).
		SetBodyString("currency", &opts.Currency)
}

func (opts *ConvertCurrencyOptions) SetBody(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetBodyString("profile_id", opts.ProfileId).
		SetBodyString("from", &opts.From).
		SetBodyString("to", &opts.To).
		SetBodyFloat("amount", &opts.Amount).
		SetBodyString("nonce", opts.Nonce)
}

func (opts *CreateOrderOptions) SetBody(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetBodyString("profile_id", opts.ProfileId).
		SetStringer("type", opts.Type).
		SetStringer("side", &opts.Side).
		SetStringer("stp", opts.Stp).
		SetStringer("stop", opts.Stop).
		SetBodyFloat("stop_price", opts.StopPrice).
		SetBodyFloat("price", opts.Price).
		SetBodyFloat("size", opts.Size).
		SetBodyFloat("funds", opts.Funds).
		SetBodyString("product_id", &opts.ProductId).
		SetStringer("time_in_force", opts.TimeInForce).
		SetStringer("cancel_after", opts.CancelAfter).
		SetBodyBool("post_only", opts.PostOnly).
		SetBodyString("client_oid", opts.ClientOid)
}

func (opts *CreateProfileOptions) SetBody(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetBodyString("name", opts.Name)
}

func (opts *CreateProfileTransferOptions) SetBody(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetBodyString("from", opts.From).
		SetBodyString("to", opts.To).
		SetBodyString("currency", opts.Currency).
		SetBodyString("amount", opts.Amount)
}

func (opts *CreateReportOptions) SetBody(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetBodyString("start_date", opts.StartDate).
		SetBodyString("end_date", opts.EndDate).
		SetStringer("type", &opts.Type).
		SetBodyString("year", opts.Year).
		SetStringer("format", opts.Format).
		SetBodyString("product_id", opts.ProductId).
		SetBodyString("account_id", opts.AccountId).
		SetBodyString("email", opts.Email).
		SetBodyString("profile_id", opts.ProfileId)
}

func (opts *CryptoWithdrawalOptions) SetBody(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetBodyString("profile_id", opts.ProfileId).
		SetBodyFloat("amount", &opts.Amount).
		SetBodyString("crypto_address", &opts.CryptoAddress).
		SetBodyString("currency", &opts.Currency).
		SetBodyString("destination_tag", opts.DestinationTag).
		SetBodyBool("no_destination_tag", opts.NoDestinationTag).
		SetBodyString("two_factor_code", opts.TwoFactorCode).
		SetBodyInt("nonce", opts.Nonce).
		SetBodyFloat("fee", opts.Fee)
}

func (opts *PaymentMethodDepositOptions) SetBody(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetBodyString("profile_id", opts.ProfileId).
		SetBodyFloat("amount", &opts.Amount).
		SetBodyString("payment_method_id", &opts.PaymentMethodId).
		SetBodyString("currency", &opts.Currency)
}

func (opts *PaymentMethodWithdrawalOptions) SetBody(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetBodyString("profile_id", opts.ProfileId).
		SetBodyFloat("amount", &opts.Amount).
		SetBodyString("payment_method_id", &opts.PaymentMethodId).
		SetBodyString("currency", &opts.Currency)
}

func (opts *AccountHoldsOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("before", opts.Before).
		SetQueryParamString("after", opts.After).
		SetQueryParamInt("limit", opts.Limit)
}

func (opts *AccountLedgerOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("start_date", opts.StartDate).
		SetQueryParamString("end_date", opts.EndDate).
		SetQueryParamString("before", opts.Before).
		SetQueryParamString("after", opts.After).
		SetQueryParamString("profile_id", opts.ProfileId).
		SetQueryParamInt("limit", opts.Limit)
}

func (opts *AccountTransfersOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("before", opts.Before).
		SetQueryParamString("after", opts.After).
		SetQueryParamInt("limit", opts.Limit).
		SetQueryParamString("type", opts.Type)
}

func (opts *BookOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamInt32("level", opts.Level)
}

func (opts *CancelOpenOrdersOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("profile_id", opts.ProfileId).
		SetQueryParamString("product_id", opts.ProductId)
}

func (opts *CandlesOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamTime("start", opts.Start).
		SetQueryParamTime("end", opts.End)
}

func (opts *CurrencyConversionOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("profile_id", opts.ProfileId)
}

func (opts *DeleteProfileOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("profile_id", opts.ProfileId).
		SetQueryParamString("to", opts.To)
}

func (opts *FillsOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("order_id", opts.OrderId).
		SetQueryParamString("product_id", opts.ProductId).
		SetQueryParamString("profile_id", opts.ProfileId).
		SetQueryParamInt("limit", opts.Limit).
		SetQueryParamInt("before", opts.Before).
		SetQueryParamInt("after", opts.After)
}

func (opts *OrdersOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("profile_id", opts.ProfileId).
		SetQueryParamString("product_id", opts.ProductId).
		SetQueryParamString("sortedBy", opts.SortedBy).
		SetQueryParamString("sorting", opts.Sorting).
		SetQueryParamTime("start_date", opts.StartDate).
		SetQueryParamTime("end_date", opts.EndDate).
		SetQueryParamString("before", opts.Before).
		SetQueryParamString("after", opts.After).
		SetQueryParamInt("limit", &opts.Limit).
		SetQueryParamStrings("status", opts.Status)
}

func (opts *ProductsOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("type", opts.Type)
}

func (opts *ProfileOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamBool("active", opts.Active)
}

func (opts *ProfilesOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamBool("active", opts.Active)
}

func (opts *RenameProfileOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("profile_id", opts.ProfileId).
		SetQueryParamString("name", opts.Name)
}

func (opts *ReportsOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("portfolio_id", opts.PortfolioId).
		SetQueryParamTime("after", opts.After).
		SetQueryParamInt("limit", opts.Limit).
		SetQueryParamBool("ignored_expired", opts.IgnoredExpired)
}

func (opts *TradesOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamInt32("limit", opts.Limit).
		SetQueryParamInt32("before", opts.Before).
		SetQueryParamInt32("after", opts.After)
}

func (opts *WithdrawalFeeEstimateOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("currency", opts.Currency).
		SetQueryParamString("crypto_address", opts.CryptoAddress)
}
