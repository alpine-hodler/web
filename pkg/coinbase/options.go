package coinbase

import (
	"github.com/alpine-hodler/sdk/internal/client"
	"github.com/alpine-hodler/sdk/pkg/scalar"
)

// * This is a generated file, do not edit

// AccountHoldsOptions are options for API requests.
type AccountHoldsOptions struct {
	// After is used for pagination and sets end cursor to `after` date.
	After *string `json:"after" bson:"after"`

	// Before is used for pagination and sets start cursor to `before` date.
	Before *string `json:"before" bson:"before"`

	// Limit puts a limit on number of results to return.
	Limit *int `json:"limit" bson:"limit"`
}

// AccountLedgerOptions are options for API requests.
type AccountLedgerOptions struct {
	// After is used for pagination. Sets end cursor to `after` date.
	After *int `json:"after" bson:"after"`

	// Before is used for pagination. Sets start cursor to `before` date.
	Before *int `json:"before" bson:"before"`

	// EndDate will filter results by maximum posted date.
	EndDate *string `json:"end_date" bson:"end_date"`

	// Limit puts a limit on number of results to return.
	Limit *int `json:"limit" bson:"limit"`

	// StartDate will filter results by minimum posted date.
	StartDate *string `json:"start_date" bson:"start_date"`
	ProfileID *string `json:"profile_id" bson:"profile_id"`
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
	// End is a timestamp for ending range of aggregations.
	End *string `json:"end" bson:"end"`

	// Granularity is one of the following values: {60, 300, 900, 3600, 21600, 86400}. Otherwise, your request will be
	// rejected. These values correspond to timeslices representing one minute, five minutes, fifteen minutes, one hour, six
	// hours, and one day, respectively.
	Granularity *scalar.Granularity `json:"granularity" bson:"granularity"`

	// Start is a timestamp for starting range of aggregations.
	Start *string `json:"start" bson:"start"`
}

// CreateOrderOptions are options for API requests.
type CreateOrderOptions struct {
	CancelAfter *scalar.CancelAfter `json:"cancel_after" bson:"cancel_after"`
	ClientOid   *string             `json:"client_oid" bson:"client_oid"`
	Funds       *float64            `json:"funds" bson:"funds"`
	PostOnly    *bool               `json:"post_only" bson:"post_only"`
	Price       *float64            `json:"price" bson:"price"`
	ProductID   string              `json:"product_id" bson:"product_id"`
	ProfileID   *string             `json:"profile_id" bson:"profile_id"`
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
	AccountID *string `json:"account_id" bson:"account_id"`

	// Email to send generated report to
	Email *string `json:"email" bson:"email"`

	// End date for items to be included in report
	EndDate *string `json:"end_date" bson:"end_date"`

	// Portfolio - Which portfolio to generate the report for
	ProfileID *string `json:"profile_id" bson:"profile_id"`

	// Product - required for fills-type reports
	ProductID *string `json:"product_id" bson:"product_id"`

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
	ProfileID *string `json:"profile_id" bson:"profile_id"`
	To        string  `json:"to" bson:"to"`
}

// CurrencyConversionOptions are options for API requests.
type CurrencyConversionOptions struct {
	ProfileID *string `json:"profile_id" bson:"profile_id"`
}

// CoinbaseAccountDepositOptions are options for API requests.
type CoinbaseAccountDepositOptions struct {
	Amount            float64 `json:"amount" bson:"amount"`
	CoinbaseAccountID string  `json:"coinbase_account_id" bson:"coinbase_account_id"`
	Currency          string  `json:"currency" bson:"currency"`
	ProfileID         *string `json:"profile_id" bson:"profile_id"`
}

// PaymentMethodDepositOptions are options for API requests.
type PaymentMethodDepositOptions struct {
	Amount          float64 `json:"amount" bson:"amount"`
	Currency        string  `json:"currency" bson:"currency"`
	PaymentMethodID string  `json:"payment_method_id" bson:"payment_method_id"`
	ProfileID       *string `json:"profile_id" bson:"profile_id"`
}

// FillsOptions are options for API requests.
type FillsOptions struct {
	After     *int    `json:"after" bson:"after"`
	Before    *int    `json:"before" bson:"before"`
	Limit     *int    `json:"limit" bson:"limit"`
	OrderID   *string `json:"order_id" bson:"order_id"`
	ProductID *string `json:"product_id" bson:"product_id"`
	ProfileID *string `json:"profile_id" bson:"profile_id"`
}

// OrdersOptions are options for API requests.
type OrdersOptions struct {
	After     *string  `json:"after" bson:"after"`
	Before    *string  `json:"before" bson:"before"`
	EndDate   *string  `json:"end_date" bson:"end_date"`
	Limit     int      `json:"limit" bson:"limit"`
	ProductID *string  `json:"product_id" bson:"product_id"`
	ProfileID *string  `json:"profile_id" bson:"profile_id"`
	SortedBy  *string  `json:"sortedBy" bson:"sortedBy"`
	Sorting   *string  `json:"sorting" bson:"sorting"`
	StartDate *string  `json:"start_date" bson:"start_date"`
	Status    []string `json:"status" bson:"status"`
}

// CancelOpenOrdersOptions are options for API requests.
type CancelOpenOrdersOptions struct {
	ProductID *string `json:"product_id" bson:"product_id"`
	ProfileID *string `json:"profile_id" bson:"profile_id"`
}

// ProductsOptions are options for API requests.
type ProductsOptions struct {
	Type *string `json:"type" bson:"type"`
}

// DeleteProfileOptions are options for API requests.
type DeleteProfileOptions struct {
	ProfileID *string `json:"profile_id" bson:"profile_id"`
	To        *string `json:"to" bson:"to"`
}

// ProfileOptions are options for API requests.
type ProfileOptions struct {
	Active *bool `json:"active" bson:"active"`
}

// RenameProfileOptions are options for API requests.
type RenameProfileOptions struct {
	Name      *string `json:"name" bson:"name"`
	ProfileID *string `json:"profile_id" bson:"profile_id"`
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
	After *string `json:"after" bson:"after"`

	// Filter results by a specific profile_id
	PortfolioID *string `json:"portfolio_id" bson:"portfolio_id"`

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

// AccountTransfersOptions are options for API requests.
type AccountTransfersOptions struct {
	// After is used for pagination. Sets end cursor to `after` date.
	After *string `json:"after" bson:"after"`

	// Before is used for pagination. Sets start cursor to `before` date.
	Before *string `json:"before" bson:"before"`

	// Limit puts a limit on number of results to return.
	Limit *int                   `json:"limit" bson:"limit"`
	Type  *scalar.TransferMethod `json:"type" bson:"type"`
}

// CoinbaseAccountWithdrawalOptions are options for API requests.
type CoinbaseAccountWithdrawalOptions struct {
	Amount            float64 `json:"amount" bson:"amount"`
	CoinbaseAccountID string  `json:"coinbase_account_id" bson:"coinbase_account_id"`
	Currency          string  `json:"currency" bson:"currency"`
	ProfileID         *string `json:"profile_id" bson:"profile_id"`
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
	ProfileID        *string  `json:"profile_id" bson:"profile_id"`
	TwoFactorCode    *string  `json:"two_factor_code" bson:"two_factor_code"`
}

// PaymentMethodWithdrawalOptions are options for API requests.
type PaymentMethodWithdrawalOptions struct {
	Amount          float64 `json:"amount" bson:"amount"`
	Currency        string  `json:"currency" bson:"currency"`
	PaymentMethodID string  `json:"payment_method_id" bson:"payment_method_id"`
	ProfileID       *string `json:"profile_id" bson:"profile_id"`
}

// WithdrawalFeeEstimateOptions are options for API requests.
type WithdrawalFeeEstimateOptions struct {
	CryptoAddress *string `json:"crypto_address" bson:"crypto_address"`
	Currency      *string `json:"currency" bson:"currency"`
}

// SetBefore sets the Before field on AccountHoldsOptions. Before is used for pagination and sets start cursor to
// `before` date.
func (opts *AccountHoldsOptions) SetBefore(Before string) *AccountHoldsOptions {
	opts.Before = &Before
	return opts
}

// SetAfter sets the After field on AccountHoldsOptions. After is used for pagination and sets end cursor to `after`
// date.
func (opts *AccountHoldsOptions) SetAfter(After string) *AccountHoldsOptions {
	opts.After = &After
	return opts
}

// SetLimit sets the Limit field on AccountHoldsOptions. Limit puts a limit on number of results to return.
func (opts *AccountHoldsOptions) SetLimit(Limit int) *AccountHoldsOptions {
	opts.Limit = &Limit
	return opts
}

// SetStartDate sets the StartDate field on AccountLedgerOptions. StartDate will filter results by minimum posted date.
func (opts *AccountLedgerOptions) SetStartDate(StartDate string) *AccountLedgerOptions {
	opts.StartDate = &StartDate
	return opts
}

// SetEndDate sets the EndDate field on AccountLedgerOptions. EndDate will filter results by maximum posted date.
func (opts *AccountLedgerOptions) SetEndDate(EndDate string) *AccountLedgerOptions {
	opts.EndDate = &EndDate
	return opts
}

// SetBefore sets the Before field on AccountLedgerOptions. Before is used for pagination. Sets start cursor to `before`
// date.
func (opts *AccountLedgerOptions) SetBefore(Before int) *AccountLedgerOptions {
	opts.Before = &Before
	return opts
}

// SetAfter sets the After field on AccountLedgerOptions. After is used for pagination. Sets end cursor to `after` date.
func (opts *AccountLedgerOptions) SetAfter(After int) *AccountLedgerOptions {
	opts.After = &After
	return opts
}

// SetProfileID sets the ProfileID field on AccountLedgerOptions.
func (opts *AccountLedgerOptions) SetProfileID(ProfileID string) *AccountLedgerOptions {
	opts.ProfileID = &ProfileID
	return opts
}

// SetLimit sets the Limit field on AccountLedgerOptions. Limit puts a limit on number of results to return.
func (opts *AccountLedgerOptions) SetLimit(Limit int) *AccountLedgerOptions {
	opts.Limit = &Limit
	return opts
}

// SetBefore sets the Before field on AccountTransfersOptions. Before is used for pagination. Sets start cursor to
// `before` date.
func (opts *AccountTransfersOptions) SetBefore(Before string) *AccountTransfersOptions {
	opts.Before = &Before
	return opts
}

// SetAfter sets the After field on AccountTransfersOptions. After is used for pagination. Sets end cursor to `after`
// date.
func (opts *AccountTransfersOptions) SetAfter(After string) *AccountTransfersOptions {
	opts.After = &After
	return opts
}

// SetLimit sets the Limit field on AccountTransfersOptions. Limit puts a limit on number of results to return.
func (opts *AccountTransfersOptions) SetLimit(Limit int) *AccountTransfersOptions {
	opts.Limit = &Limit
	return opts
}

// SetType sets the Type field on AccountTransfersOptions.
func (opts *AccountTransfersOptions) SetType(Type scalar.TransferMethod) *AccountTransfersOptions {
	opts.Type = &Type
	return opts
}

// SetLevel sets the Level field on BookOptions. Levels 1 and 2 are aggregated. The size field is the sum of the size of
// the orders at that price, and num-orders is the count of orders at that price; size should not be multiplied by
// num-orders. Level 3 is non-aggregated and returns the entire order book. While the book is in an auction, the L1, L2
// and L3 book will also contain the most recent indicative quote disseminated during the auction, and auction_mode will
// be set to true. These indicative quote messages are sent on an interval basis (approximately once a second) during
// the collection phase of an auction and includes information about the tentative price and size affiliated with the
// completion. Level 1 and Level 2 are recommended for polling. For the most up-to-date data, consider using the
// websocket stream. Level 3 is only recommended for users wishing to maintain a full real-time order book using the
// websocket stream. Abuse of Level 3 via polling will cause your access to be limited or blocked.
func (opts *BookOptions) SetLevel(Level int32) *BookOptions {
	opts.Level = &Level
	return opts
}

// SetProfileID sets the ProfileID field on CancelOpenOrdersOptions.
func (opts *CancelOpenOrdersOptions) SetProfileID(ProfileID string) *CancelOpenOrdersOptions {
	opts.ProfileID = &ProfileID
	return opts
}

// SetProductID sets the ProductID field on CancelOpenOrdersOptions.
func (opts *CancelOpenOrdersOptions) SetProductID(ProductID string) *CancelOpenOrdersOptions {
	opts.ProductID = &ProductID
	return opts
}

// SetGranularity sets the Granularity field on CandlesOptions. Granularity is one of the following values: {60, 300,
// 900, 3600, 21600, 86400}. Otherwise, your request will be rejected. These values correspond to timeslices
// representing one minute, five minutes, fifteen minutes, one hour, six hours, and one day, respectively.
func (opts *CandlesOptions) SetGranularity(Granularity scalar.Granularity) *CandlesOptions {
	opts.Granularity = &Granularity
	return opts
}

// SetStart sets the Start field on CandlesOptions. Start is a timestamp for starting range of aggregations.
func (opts *CandlesOptions) SetStart(Start string) *CandlesOptions {
	opts.Start = &Start
	return opts
}

// SetEnd sets the End field on CandlesOptions. End is a timestamp for ending range of aggregations.
func (opts *CandlesOptions) SetEnd(End string) *CandlesOptions {
	opts.End = &End
	return opts
}

// SetProfileID sets the ProfileID field on CoinbaseAccountDepositOptions.
func (opts *CoinbaseAccountDepositOptions) SetProfileID(ProfileID string) *CoinbaseAccountDepositOptions {
	opts.ProfileID = &ProfileID
	return opts
}

// SetAmount sets the Amount field on CoinbaseAccountDepositOptions.
func (opts *CoinbaseAccountDepositOptions) SetAmount(Amount float64) *CoinbaseAccountDepositOptions {
	opts.Amount = Amount
	return opts
}

// SetCoinbaseAccountID sets the CoinbaseAccountID field on CoinbaseAccountDepositOptions.
func (opts *CoinbaseAccountDepositOptions) SetCoinbaseAccountID(CoinbaseAccountID string) *CoinbaseAccountDepositOptions {
	opts.CoinbaseAccountID = CoinbaseAccountID
	return opts
}

// SetCurrency sets the Currency field on CoinbaseAccountDepositOptions.
func (opts *CoinbaseAccountDepositOptions) SetCurrency(Currency string) *CoinbaseAccountDepositOptions {
	opts.Currency = Currency
	return opts
}

// SetProfileID sets the ProfileID field on CoinbaseAccountWithdrawalOptions.
func (opts *CoinbaseAccountWithdrawalOptions) SetProfileID(ProfileID string) *CoinbaseAccountWithdrawalOptions {
	opts.ProfileID = &ProfileID
	return opts
}

// SetAmount sets the Amount field on CoinbaseAccountWithdrawalOptions.
func (opts *CoinbaseAccountWithdrawalOptions) SetAmount(Amount float64) *CoinbaseAccountWithdrawalOptions {
	opts.Amount = Amount
	return opts
}

// SetCoinbaseAccountID sets the CoinbaseAccountID field on CoinbaseAccountWithdrawalOptions.
func (opts *CoinbaseAccountWithdrawalOptions) SetCoinbaseAccountID(CoinbaseAccountID string) *CoinbaseAccountWithdrawalOptions {
	opts.CoinbaseAccountID = CoinbaseAccountID
	return opts
}

// SetCurrency sets the Currency field on CoinbaseAccountWithdrawalOptions.
func (opts *CoinbaseAccountWithdrawalOptions) SetCurrency(Currency string) *CoinbaseAccountWithdrawalOptions {
	opts.Currency = Currency
	return opts
}

// SetProfileID sets the ProfileID field on ConvertCurrencyOptions.
func (opts *ConvertCurrencyOptions) SetProfileID(ProfileID string) *ConvertCurrencyOptions {
	opts.ProfileID = &ProfileID
	return opts
}

// SetFrom sets the From field on ConvertCurrencyOptions.
func (opts *ConvertCurrencyOptions) SetFrom(From string) *ConvertCurrencyOptions {
	opts.From = From
	return opts
}

// SetTo sets the To field on ConvertCurrencyOptions.
func (opts *ConvertCurrencyOptions) SetTo(To string) *ConvertCurrencyOptions {
	opts.To = To
	return opts
}

// SetAmount sets the Amount field on ConvertCurrencyOptions.
func (opts *ConvertCurrencyOptions) SetAmount(Amount float64) *ConvertCurrencyOptions {
	opts.Amount = Amount
	return opts
}

// SetNonce sets the Nonce field on ConvertCurrencyOptions.
func (opts *ConvertCurrencyOptions) SetNonce(Nonce string) *ConvertCurrencyOptions {
	opts.Nonce = &Nonce
	return opts
}

// SetProfileID sets the ProfileID field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetProfileID(ProfileID string) *CreateOrderOptions {
	opts.ProfileID = &ProfileID
	return opts
}

// SetType sets the Type field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetType(Type scalar.OrderType) *CreateOrderOptions {
	opts.Type = &Type
	return opts
}

// SetSide sets the Side field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetSide(Side scalar.OrderSide) *CreateOrderOptions {
	opts.Side = Side
	return opts
}

// SetStp sets the Stp field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetStp(Stp scalar.OrderSTP) *CreateOrderOptions {
	opts.Stp = &Stp
	return opts
}

// SetStop sets the Stop field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetStop(Stop scalar.OrderStop) *CreateOrderOptions {
	opts.Stop = &Stop
	return opts
}

// SetStopPrice sets the StopPrice field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetStopPrice(StopPrice float64) *CreateOrderOptions {
	opts.StopPrice = &StopPrice
	return opts
}

// SetPrice sets the Price field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetPrice(Price float64) *CreateOrderOptions {
	opts.Price = &Price
	return opts
}

// SetSize sets the Size field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetSize(Size float64) *CreateOrderOptions {
	opts.Size = &Size
	return opts
}

// SetFunds sets the Funds field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetFunds(Funds float64) *CreateOrderOptions {
	opts.Funds = &Funds
	return opts
}

// SetProductID sets the ProductID field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetProductID(ProductID string) *CreateOrderOptions {
	opts.ProductID = ProductID
	return opts
}

// SetTimeInForce sets the TimeInForce field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetTimeInForce(TimeInForce scalar.TimeInForce) *CreateOrderOptions {
	opts.TimeInForce = &TimeInForce
	return opts
}

// SetCancelAfter sets the CancelAfter field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetCancelAfter(CancelAfter scalar.CancelAfter) *CreateOrderOptions {
	opts.CancelAfter = &CancelAfter
	return opts
}

// SetPostOnly sets the PostOnly field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetPostOnly(PostOnly bool) *CreateOrderOptions {
	opts.PostOnly = &PostOnly
	return opts
}

// SetClientOid sets the ClientOid field on CreateOrderOptions.
func (opts *CreateOrderOptions) SetClientOid(ClientOid string) *CreateOrderOptions {
	opts.ClientOid = &ClientOid
	return opts
}

// SetName sets the Name field on CreateProfileOptions.
func (opts *CreateProfileOptions) SetName(Name string) *CreateProfileOptions {
	opts.Name = &Name
	return opts
}

// SetFrom sets the From field on CreateProfileTransferOptions.
func (opts *CreateProfileTransferOptions) SetFrom(From string) *CreateProfileTransferOptions {
	opts.From = &From
	return opts
}

// SetTo sets the To field on CreateProfileTransferOptions.
func (opts *CreateProfileTransferOptions) SetTo(To string) *CreateProfileTransferOptions {
	opts.To = &To
	return opts
}

// SetCurrency sets the Currency field on CreateProfileTransferOptions.
func (opts *CreateProfileTransferOptions) SetCurrency(Currency string) *CreateProfileTransferOptions {
	opts.Currency = &Currency
	return opts
}

// SetAmount sets the Amount field on CreateProfileTransferOptions.
func (opts *CreateProfileTransferOptions) SetAmount(Amount string) *CreateProfileTransferOptions {
	opts.Amount = &Amount
	return opts
}

// SetStartDate sets the StartDate field on CreateReportOptions. Start date for items to be included in report.
func (opts *CreateReportOptions) SetStartDate(StartDate string) *CreateReportOptions {
	opts.StartDate = &StartDate
	return opts
}

// SetEndDate sets the EndDate field on CreateReportOptions. End date for items to be included in report
func (opts *CreateReportOptions) SetEndDate(EndDate string) *CreateReportOptions {
	opts.EndDate = &EndDate
	return opts
}

// SetType sets the Type field on CreateReportOptions.
func (opts *CreateReportOptions) SetType(Type scalar.ReportType) *CreateReportOptions {
	opts.Type = Type
	return opts
}

// SetYear sets the Year field on CreateReportOptions. required for 1099k-transaction-history-type reports
func (opts *CreateReportOptions) SetYear(Year string) *CreateReportOptions {
	opts.Year = &Year
	return opts
}

// SetFormat sets the Format field on CreateReportOptions.
func (opts *CreateReportOptions) SetFormat(Format scalar.Format) *CreateReportOptions {
	opts.Format = &Format
	return opts
}

// SetProductID sets the ProductID field on CreateReportOptions. Product - required for fills-type reports
func (opts *CreateReportOptions) SetProductID(ProductID string) *CreateReportOptions {
	opts.ProductID = &ProductID
	return opts
}

// SetAccountID sets the AccountID field on CreateReportOptions. Account - required for account-type reports
func (opts *CreateReportOptions) SetAccountID(AccountID string) *CreateReportOptions {
	opts.AccountID = &AccountID
	return opts
}

// SetEmail sets the Email field on CreateReportOptions. Email to send generated report to
func (opts *CreateReportOptions) SetEmail(Email string) *CreateReportOptions {
	opts.Email = &Email
	return opts
}

// SetProfileID sets the ProfileID field on CreateReportOptions. Portfolio - Which portfolio to generate the report for
func (opts *CreateReportOptions) SetProfileID(ProfileID string) *CreateReportOptions {
	opts.ProfileID = &ProfileID
	return opts
}

// SetProfileID sets the ProfileID field on CryptoWithdrawalOptions.
func (opts *CryptoWithdrawalOptions) SetProfileID(ProfileID string) *CryptoWithdrawalOptions {
	opts.ProfileID = &ProfileID
	return opts
}

// SetAmount sets the Amount field on CryptoWithdrawalOptions.
func (opts *CryptoWithdrawalOptions) SetAmount(Amount float64) *CryptoWithdrawalOptions {
	opts.Amount = Amount
	return opts
}

// SetCryptoAddress sets the CryptoAddress field on CryptoWithdrawalOptions.
func (opts *CryptoWithdrawalOptions) SetCryptoAddress(CryptoAddress string) *CryptoWithdrawalOptions {
	opts.CryptoAddress = CryptoAddress
	return opts
}

// SetCurrency sets the Currency field on CryptoWithdrawalOptions.
func (opts *CryptoWithdrawalOptions) SetCurrency(Currency string) *CryptoWithdrawalOptions {
	opts.Currency = Currency
	return opts
}

// SetDestinationTag sets the DestinationTag field on CryptoWithdrawalOptions.
func (opts *CryptoWithdrawalOptions) SetDestinationTag(DestinationTag string) *CryptoWithdrawalOptions {
	opts.DestinationTag = &DestinationTag
	return opts
}

// SetNoDestinationTag sets the NoDestinationTag field on CryptoWithdrawalOptions.
func (opts *CryptoWithdrawalOptions) SetNoDestinationTag(NoDestinationTag bool) *CryptoWithdrawalOptions {
	opts.NoDestinationTag = &NoDestinationTag
	return opts
}

// SetTwoFactorCode sets the TwoFactorCode field on CryptoWithdrawalOptions.
func (opts *CryptoWithdrawalOptions) SetTwoFactorCode(TwoFactorCode string) *CryptoWithdrawalOptions {
	opts.TwoFactorCode = &TwoFactorCode
	return opts
}

// SetNonce sets the Nonce field on CryptoWithdrawalOptions.
func (opts *CryptoWithdrawalOptions) SetNonce(Nonce int) *CryptoWithdrawalOptions {
	opts.Nonce = &Nonce
	return opts
}

// SetFee sets the Fee field on CryptoWithdrawalOptions.
func (opts *CryptoWithdrawalOptions) SetFee(Fee float64) *CryptoWithdrawalOptions {
	opts.Fee = &Fee
	return opts
}

// SetProfileID sets the ProfileID field on CurrencyConversionOptions.
func (opts *CurrencyConversionOptions) SetProfileID(ProfileID string) *CurrencyConversionOptions {
	opts.ProfileID = &ProfileID
	return opts
}

// SetProfileID sets the ProfileID field on DeleteProfileOptions.
func (opts *DeleteProfileOptions) SetProfileID(ProfileID string) *DeleteProfileOptions {
	opts.ProfileID = &ProfileID
	return opts
}

// SetTo sets the To field on DeleteProfileOptions.
func (opts *DeleteProfileOptions) SetTo(To string) *DeleteProfileOptions {
	opts.To = &To
	return opts
}

// SetOrderID sets the OrderID field on FillsOptions.
func (opts *FillsOptions) SetOrderID(OrderID string) *FillsOptions {
	opts.OrderID = &OrderID
	return opts
}

// SetProductID sets the ProductID field on FillsOptions.
func (opts *FillsOptions) SetProductID(ProductID string) *FillsOptions {
	opts.ProductID = &ProductID
	return opts
}

// SetProfileID sets the ProfileID field on FillsOptions.
func (opts *FillsOptions) SetProfileID(ProfileID string) *FillsOptions {
	opts.ProfileID = &ProfileID
	return opts
}

// SetLimit sets the Limit field on FillsOptions.
func (opts *FillsOptions) SetLimit(Limit int) *FillsOptions {
	opts.Limit = &Limit
	return opts
}

// SetBefore sets the Before field on FillsOptions.
func (opts *FillsOptions) SetBefore(Before int) *FillsOptions {
	opts.Before = &Before
	return opts
}

// SetAfter sets the After field on FillsOptions.
func (opts *FillsOptions) SetAfter(After int) *FillsOptions {
	opts.After = &After
	return opts
}

// SetProfileID sets the ProfileID field on OrdersOptions.
func (opts *OrdersOptions) SetProfileID(ProfileID string) *OrdersOptions {
	opts.ProfileID = &ProfileID
	return opts
}

// SetProductID sets the ProductID field on OrdersOptions.
func (opts *OrdersOptions) SetProductID(ProductID string) *OrdersOptions {
	opts.ProductID = &ProductID
	return opts
}

// SetSortedBy sets the SortedBy field on OrdersOptions.
func (opts *OrdersOptions) SetSortedBy(SortedBy string) *OrdersOptions {
	opts.SortedBy = &SortedBy
	return opts
}

// SetSorting sets the Sorting field on OrdersOptions.
func (opts *OrdersOptions) SetSorting(Sorting string) *OrdersOptions {
	opts.Sorting = &Sorting
	return opts
}

// SetStartDate sets the StartDate field on OrdersOptions.
func (opts *OrdersOptions) SetStartDate(StartDate string) *OrdersOptions {
	opts.StartDate = &StartDate
	return opts
}

// SetEndDate sets the EndDate field on OrdersOptions.
func (opts *OrdersOptions) SetEndDate(EndDate string) *OrdersOptions {
	opts.EndDate = &EndDate
	return opts
}

// SetBefore sets the Before field on OrdersOptions.
func (opts *OrdersOptions) SetBefore(Before string) *OrdersOptions {
	opts.Before = &Before
	return opts
}

// SetAfter sets the After field on OrdersOptions.
func (opts *OrdersOptions) SetAfter(After string) *OrdersOptions {
	opts.After = &After
	return opts
}

// SetLimit sets the Limit field on OrdersOptions.
func (opts *OrdersOptions) SetLimit(Limit int) *OrdersOptions {
	opts.Limit = Limit
	return opts
}

// SetStatus sets the Status field on OrdersOptions.
func (opts *OrdersOptions) SetStatus(Status []string) *OrdersOptions {
	opts.Status = Status
	return opts
}

// SetProfileID sets the ProfileID field on PaymentMethodDepositOptions.
func (opts *PaymentMethodDepositOptions) SetProfileID(ProfileID string) *PaymentMethodDepositOptions {
	opts.ProfileID = &ProfileID
	return opts
}

// SetAmount sets the Amount field on PaymentMethodDepositOptions.
func (opts *PaymentMethodDepositOptions) SetAmount(Amount float64) *PaymentMethodDepositOptions {
	opts.Amount = Amount
	return opts
}

// SetPaymentMethodID sets the PaymentMethodID field on PaymentMethodDepositOptions.
func (opts *PaymentMethodDepositOptions) SetPaymentMethodID(PaymentMethodID string) *PaymentMethodDepositOptions {
	opts.PaymentMethodID = PaymentMethodID
	return opts
}

// SetCurrency sets the Currency field on PaymentMethodDepositOptions.
func (opts *PaymentMethodDepositOptions) SetCurrency(Currency string) *PaymentMethodDepositOptions {
	opts.Currency = Currency
	return opts
}

// SetProfileID sets the ProfileID field on PaymentMethodWithdrawalOptions.
func (opts *PaymentMethodWithdrawalOptions) SetProfileID(ProfileID string) *PaymentMethodWithdrawalOptions {
	opts.ProfileID = &ProfileID
	return opts
}

// SetAmount sets the Amount field on PaymentMethodWithdrawalOptions.
func (opts *PaymentMethodWithdrawalOptions) SetAmount(Amount float64) *PaymentMethodWithdrawalOptions {
	opts.Amount = Amount
	return opts
}

// SetPaymentMethodID sets the PaymentMethodID field on PaymentMethodWithdrawalOptions.
func (opts *PaymentMethodWithdrawalOptions) SetPaymentMethodID(PaymentMethodID string) *PaymentMethodWithdrawalOptions {
	opts.PaymentMethodID = PaymentMethodID
	return opts
}

// SetCurrency sets the Currency field on PaymentMethodWithdrawalOptions.
func (opts *PaymentMethodWithdrawalOptions) SetCurrency(Currency string) *PaymentMethodWithdrawalOptions {
	opts.Currency = Currency
	return opts
}

// SetType sets the Type field on ProductsOptions.
func (opts *ProductsOptions) SetType(Type string) *ProductsOptions {
	opts.Type = &Type
	return opts
}

// SetActive sets the Active field on ProfileOptions.
func (opts *ProfileOptions) SetActive(Active bool) *ProfileOptions {
	opts.Active = &Active
	return opts
}

// SetActive sets the Active field on ProfilesOptions.
func (opts *ProfilesOptions) SetActive(Active bool) *ProfilesOptions {
	opts.Active = &Active
	return opts
}

// SetProfileID sets the ProfileID field on RenameProfileOptions.
func (opts *RenameProfileOptions) SetProfileID(ProfileID string) *RenameProfileOptions {
	opts.ProfileID = &ProfileID
	return opts
}

// SetName sets the Name field on RenameProfileOptions.
func (opts *RenameProfileOptions) SetName(Name string) *RenameProfileOptions {
	opts.Name = &Name
	return opts
}

// SetPortfolioID sets the PortfolioID field on ReportsOptions. Filter results by a specific profile_id
func (opts *ReportsOptions) SetPortfolioID(PortfolioID string) *ReportsOptions {
	opts.PortfolioID = &PortfolioID
	return opts
}

// SetAfter sets the After field on ReportsOptions. Filter results after a specific date
func (opts *ReportsOptions) SetAfter(After string) *ReportsOptions {
	opts.After = &After
	return opts
}

// SetLimit sets the Limit field on ReportsOptions. Limit results to a specific number
func (opts *ReportsOptions) SetLimit(Limit int) *ReportsOptions {
	opts.Limit = &Limit
	return opts
}

// SetType sets the Type field on ReportsOptions. Filter results by type of report (fills or account) - otc_fills: real
// string is otc-fills - type_1099k_transaction_history: real string is 1099-transaction-history - tax_invoice: real
// string is tax-invoice
func (opts *ReportsOptions) SetType(Type scalar.ReportType) *ReportsOptions {
	opts.Type = &Type
	return opts
}

// SetIgnoredExpired sets the IgnoredExpired field on ReportsOptions. Ignore expired results
func (opts *ReportsOptions) SetIgnoredExpired(IgnoredExpired bool) *ReportsOptions {
	opts.IgnoredExpired = &IgnoredExpired
	return opts
}

// SetLimit sets the Limit field on TradesOptions.
func (opts *TradesOptions) SetLimit(Limit int32) *TradesOptions {
	opts.Limit = &Limit
	return opts
}

// SetBefore sets the Before field on TradesOptions.
func (opts *TradesOptions) SetBefore(Before int32) *TradesOptions {
	opts.Before = &Before
	return opts
}

// SetAfter sets the After field on TradesOptions.
func (opts *TradesOptions) SetAfter(After int32) *TradesOptions {
	opts.After = &After
	return opts
}

// SetCurrency sets the Currency field on WithdrawalFeeEstimateOptions.
func (opts *WithdrawalFeeEstimateOptions) SetCurrency(Currency string) *WithdrawalFeeEstimateOptions {
	opts.Currency = &Currency
	return opts
}

// SetCryptoAddress sets the CryptoAddress field on WithdrawalFeeEstimateOptions.
func (opts *WithdrawalFeeEstimateOptions) SetCryptoAddress(CryptoAddress string) *WithdrawalFeeEstimateOptions {
	opts.CryptoAddress = &CryptoAddress
	return opts
}

func (opts *CoinbaseAccountDepositOptions) SetBody(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetBodyString("profile_id", opts.ProfileID).
		SetBodyFloat("amount", &opts.Amount).
		SetBodyString("coinbase_account_id", &opts.CoinbaseAccountID).
		SetBodyString("currency", &opts.Currency)
}

func (opts *CoinbaseAccountWithdrawalOptions) SetBody(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetBodyString("profile_id", opts.ProfileID).
		SetBodyFloat("amount", &opts.Amount).
		SetBodyString("coinbase_account_id", &opts.CoinbaseAccountID).
		SetBodyString("currency", &opts.Currency)
}

func (opts *ConvertCurrencyOptions) SetBody(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetBodyString("profile_id", opts.ProfileID).
		SetBodyString("from", &opts.From).
		SetBodyString("to", &opts.To).
		SetBodyFloat("amount", &opts.Amount).
		SetBodyString("nonce", opts.Nonce)
}

func (opts *CreateOrderOptions) SetBody(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetBodyString("profile_id", opts.ProfileID).
		SetStringer("type", opts.Type).
		SetStringer("side", &opts.Side).
		SetStringer("stp", opts.Stp).
		SetStringer("stop", opts.Stop).
		SetBodyFloat("stop_price", opts.StopPrice).
		SetBodyFloat("price", opts.Price).
		SetBodyFloat("size", opts.Size).
		SetBodyFloat("funds", opts.Funds).
		SetBodyString("product_id", &opts.ProductID).
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
		SetBodyString("product_id", opts.ProductID).
		SetBodyString("account_id", opts.AccountID).
		SetBodyString("email", opts.Email).
		SetBodyString("profile_id", opts.ProfileID)
}

func (opts *CryptoWithdrawalOptions) SetBody(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetBodyString("profile_id", opts.ProfileID).
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
	req.SetBodyString("profile_id", opts.ProfileID).
		SetBodyFloat("amount", &opts.Amount).
		SetBodyString("payment_method_id", &opts.PaymentMethodID).
		SetBodyString("currency", &opts.Currency)
}

func (opts *PaymentMethodWithdrawalOptions) SetBody(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetBodyString("profile_id", opts.ProfileID).
		SetBodyFloat("amount", &opts.Amount).
		SetBodyString("payment_method_id", &opts.PaymentMethodID).
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
		SetQueryParamInt("before", opts.Before).
		SetQueryParamInt("after", opts.After).
		SetQueryParamString("profile_id", opts.ProfileID).
		SetQueryParamInt("limit", opts.Limit)
}

func (opts *AccountTransfersOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("before", opts.Before).
		SetQueryParamString("after", opts.After).
		SetQueryParamInt("limit", opts.Limit)
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
	req.SetQueryParamString("profile_id", opts.ProfileID).
		SetQueryParamString("product_id", opts.ProductID)
}

func (opts *CandlesOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("start", opts.Start).
		SetQueryParamString("end", opts.End)
}

func (opts *CurrencyConversionOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("profile_id", opts.ProfileID)
}

func (opts *DeleteProfileOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("profile_id", opts.ProfileID).
		SetQueryParamString("to", opts.To)
}

func (opts *FillsOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("order_id", opts.OrderID).
		SetQueryParamString("product_id", opts.ProductID).
		SetQueryParamString("profile_id", opts.ProfileID).
		SetQueryParamInt("limit", opts.Limit).
		SetQueryParamInt("before", opts.Before).
		SetQueryParamInt("after", opts.After)
}

func (opts *OrdersOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("profile_id", opts.ProfileID).
		SetQueryParamString("product_id", opts.ProductID).
		SetQueryParamString("sortedBy", opts.SortedBy).
		SetQueryParamString("sorting", opts.Sorting).
		SetQueryParamString("start_date", opts.StartDate).
		SetQueryParamString("end_date", opts.EndDate).
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
	req.SetQueryParamString("profile_id", opts.ProfileID).
		SetQueryParamString("name", opts.Name)
}

func (opts *ReportsOptions) SetQueryParams(req *client.Request) {
	if opts == nil {
		return
	}
	req.SetQueryParamString("portfolio_id", opts.PortfolioID).
		SetQueryParamString("after", opts.After).
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
