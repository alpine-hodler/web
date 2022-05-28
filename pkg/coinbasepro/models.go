package coinbasepro

import (
	"time"

	"github.com/alpine-hodler/web/internal/serial"
	"github.com/alpine-hodler/web/pkg/scalar"
)

// * This is a generated file, do not edit

// Account holds data for trading account from the profile of the API key
type Account struct {
	Available      string `json:"available" bson:"available"`
	Balance        string `json:"balance" bson:"balance"`
	Currency       string `json:"currency" bson:"currency"`
	Hold           string `json:"hold" bson:"hold"`
	ID             string `json:"id" bson:"id"`
	ProfileID      string `json:"profile_id" bson:"profile_id"`
	TradingEnabled bool   `json:"trading_enabled" bson:"trading_enabled"`
}

// AccountHold represents the hold on an account that belong to the same profile as the API key. Holds are placed on an
// account for any active orders or pending withdraw requests. As an order is filled, the hold amount is updated. If an
// order is canceled, any remaining hold is removed. For withdrawals, the hold is removed after it is completed.
type AccountHold struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	ID        string    `json:"id" bson:"id"`
	Ref       string    `json:"ref" bson:"ref"`
	Type      string    `json:"type" bson:"type"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

// AccountLedger lists ledger activity for an account. This includes anything that would affect the accounts balance -
// transfers, trades, fees, etc.
type AccountLedger struct {
	Amount    string               `json:"amount" bson:"amount"`
	Balance   string               `json:"balance" bson:"balance"`
	CreatedAt time.Time            `json:"created_at" bson:"created_at"`
	Details   AccountLedgerDetails `json:"details" bson:"details"`
	ID        string               `json:"id" bson:"id"`
	Type      scalar.EntryType     `json:"type" bson:"type"`
}

// AccountLedgerDetails are the details for account history.
type AccountLedgerDetails struct {
	OrderID   string `json:"order_id" bson:"order_id"`
	ProductID string `json:"product_id" bson:"product_id"`
	TradeID   string `json:"trade_id" bson:"trade_id"`
}

// AccountTransferDetails are the details for an account transfer.
type AccountTransferDetails struct {
	CoinbaseAccountID       string `json:"coinbase_account_id" bson:"coinbase_account_id"`
	CoinbasePaymentMethodID string `json:"coinbase_payment_method_id" bson:"coinbase_payment_method_id"`
	CoinbaseTransactionID   string `json:"coinbase_transaction_id" bson:"coinbase_transaction_id"`
}

// Auction is an object of data concerning a book request.
type Auction struct {
	AuctionState string    `json:"auction_state" bson:"auction_state"`
	BestAskPrice string    `json:"best_ask_price" bson:"best_ask_price"`
	BestAskSize  string    `json:"best_ask_size" bson:"best_ask_size"`
	BestBidPrice string    `json:"best_bid_price" bson:"best_bid_price"`
	BestBidSize  string    `json:"best_bid_size" bson:"best_bid_size"`
	CanOpen      string    `json:"can_open" bson:"can_open"`
	OpenPrice    string    `json:"open_price" bson:"open_price"`
	OpenSize     string    `json:"open_size" bson:"open_size"`
	Time         time.Time `json:"time" bson:"time"`
}

// AvailableBalance is the available balance on the coinbase account
type AvailableBalance struct {
	Amount   string `json:"amount" bson:"amount"`
	Currency string `json:"currency" bson:"currency"`
	Scale    string `json:"scale" bson:"scale"`
}

// Balance is the balance for picker data
type Balance struct {
	Amount   string `json:"amount" bson:"amount"`
	Currency string `json:"currency" bson:"currency"`
}

// BankCountry are the name and code for the bank's country associated with a wallet
type BankCountry struct {
	Code string `json:"code" bson:"code"`
	Name string `json:"name" bson:"name"`
}

// Book is a list of open orders for a product. The amount of detail shown can be customized with the level parameter.
type Book struct {
	Asks        scalar.BidAsk `json:"asks" bson:"asks"`
	Auction     Auction       `json:"auction" bson:"auction"`
	AuctionMode bool          `json:"auction_mode" bson:"auction_mode"`
	Bids        scalar.BidAsk `json:"bids" bson:"bids"`
	Sequence    float64       `json:"sequence" bson:"sequence"`
}

// Candles are the historic rates for a product. Rates are returned in grouped buckets. Candle schema is of the form
// `[timestamp, price_low, price_high, price_open, price_close]`
type Candles [][]float64

// CreateOrder is the server's response for placing a new order.
type CreateOrder struct {
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	DoneAt        time.Time          `json:"done_at" bson:"done_at"`
	DoneReason    string             `json:"done_reason" bson:"done_reason"`
	ExpireTime    time.Time          `json:"expire_time" bson:"expire_time"`
	FillFees      string             `json:"fill_fees" bson:"fill_fees"`
	FilledSize    string             `json:"filled_size" bson:"filled_size"`
	FundingAmount string             `json:"funding_amount" bson:"funding_amount"`
	Funds         string             `json:"funds" bson:"funds"`
	ID            string             `json:"id" bson:"id"`
	PostOnly      bool               `json:"post_only" bson:"post_only"`
	Price         string             `json:"price" bson:"price"`
	ProductID     string             `json:"product_id" bson:"product_id"`
	ProfileID     string             `json:"profile_id" bson:"profile_id"`
	RejectReason  string             `json:"reject_reason" bson:"reject_reason"`
	Settled       bool               `json:"settled" bson:"settled"`
	Side          scalar.OrderSide   `json:"side" bson:"side"`
	Size          string             `json:"size" bson:"size"`
	SpecificFunds string             `json:"specific_funds" bson:"specific_funds"`
	Status        string             `json:"status" bson:"status"`
	Stop          scalar.OrderStop   `json:"stop" bson:"stop"`
	StopPrice     string             `json:"stop_price" bson:"stop_price"`
	TimeInForce   scalar.TimeInForce `json:"time_in_force" bson:"time_in_force"`
	Type          scalar.OrderType   `json:"type" bson:"type"`
}

// CreateReport represents information for a report created through the client.
type CreateReport struct {
	ID     string            `json:"id" bson:"id"`
	Status scalar.Status     `json:"status" bson:"status"`
	Type   scalar.ReportType `json:"type" bson:"type"`
}

// CryptoAccount references a crypto account that a CoinbasePaymentMethod belongs to
type CryptoAccount struct {
	ID           string `json:"id" bson:"id"`
	Resource     string `json:"resource" bson:"resource"`
	ResourcePath string `json:"resource_path" bson:"resource_path"`
}

// CryptoAddress is used for a one-time crypto address for depositing crypto.
type CryptoAddress struct {
	Address        string                  `json:"address" bson:"address"`
	AddressInfo    CryptoAddressInfo       `json:"address_info" bson:"address_info"`
	CallbackURL    string                  `json:"callback_url" bson:"callback_url"`
	CreateAt       time.Time               `json:"create_at" bson:"create_at"`
	DepositUri     string                  `json:"deposit_uri" bson:"deposit_uri"`
	DestinationTag string                  `json:"destination_tag" bson:"destination_tag"`
	ID             string                  `json:"id" bson:"id"`
	LegacyAddress  string                  `json:"legacy_address" bson:"legacy_address"`
	Name           string                  `json:"name" bson:"name"`
	Network        string                  `json:"network" bson:"network"`
	Resource       string                  `json:"resource" bson:"resource"`
	ResourcePath   string                  `json:"resource_path" bson:"resource_path"`
	UpdatedAt      time.Time               `json:"updated_at" bson:"updated_at"`
	UriScheme      string                  `json:"uri_scheme" bson:"uri_scheme"`
	Warnings       []*CryptoAddressWarning `json:"warnings" bson:"warnings"`
}

// CryptoAddressInfo holds info for a crypto address
type CryptoAddressInfo struct {
	Address        string `json:"address" bson:"address"`
	DestinationTag string `json:"destination_tag" bson:"destination_tag"`
}

// CryptoAddressWarning is a warning for generating a crypting address
type CryptoAddressWarning struct {
	Details  string `json:"details" bson:"details"`
	ImageURL string `json:"image_url" bson:"image_url"`
	Title    string `json:"title" bson:"title"`
}

// Currency is a currency that coinbase knows about. Not al currencies may be currently in use for trading.
type Currency struct {
	ConvertibleTo []string        `json:"convertible_to" bson:"convertible_to"`
	Details       CurrencyDetails `json:"details" bson:"details"`
	ID            string          `json:"id" bson:"id"`
	MaxPrecision  string          `json:"max_precision" bson:"max_precision"`
	Message       string          `json:"message" bson:"message"`
	MinSize       string          `json:"min_size" bson:"min_size"`
	Name          string          `json:"name" bson:"name"`
	Status        string          `json:"status" bson:"status"`
}

// CurrencyConversion is the response that converts funds from from currency to to currency. Funds are converted on the
// from account in the profile_id profile.
type CurrencyConversion struct {
	Amount        string `json:"amount" bson:"amount"`
	From          string `json:"from" bson:"from"`
	FromAccountID string `json:"from_account_id" bson:"from_account_id"`
	ID            string `json:"id" bson:"id"`
	To            string `json:"to" bson:"to"`
	ToAccountID   string `json:"to_account_id" bson:"to_account_id"`
}

// CurrencyDetails are the details for a currency that coinbase knows about
type CurrencyDetails struct {
	CryptoAddressLink     string   `json:"crypto_address_link" bson:"crypto_address_link"`
	CryptoTransactionLink string   `json:"crypto_transaction_link" bson:"crypto_transaction_link"`
	DisplayName           string   `json:"display_name" bson:"display_name"`
	GroupTypes            []string `json:"group_types" bson:"group_types"`
	MaxWithdrawalAmount   float64  `json:"max_withdrawal_amount" bson:"max_withdrawal_amount"`
	MinWithdrawalAmount   float64  `json:"min_withdrawal_amount" bson:"min_withdrawal_amount"`
	NetworkConfirmations  int      `json:"network_confirmations" bson:"network_confirmations"`
	ProcessingTimeSeconds float64  `json:"processing_time_seconds" bson:"processing_time_seconds"`
	PushPaymentMethods    []string `json:"push_payment_methods" bson:"push_payment_methods"`
	SortOrder             int      `json:"sort_order" bson:"sort_order"`
	Symbol                string   `json:"symbol" bson:"symbol"`
	Type                  string   `json:"type" bson:"type"`
}

// CurrencyTransferLimit encapsulates ACH data for a currency via Max/Remaining amounts.
type CurrencyTransferLimit struct {
	Max       float64 `json:"max" bson:"max"`
	Remaining float64 `json:"remaining" bson:"remaining"`
}

// CurrencyTransferLimits encapsulates ACH data for many currencies.
type CurrencyTransferLimits map[string]CurrencyTransferLimit

// Deposit is the response for deposited funds from a www.coinbase.com wallet to the specified profile_id.
type Deposit struct {
	Amount   string `json:"amount" bson:"amount"`
	Currency string `json:"currency" bson:"currency"`
	Fee      string `json:"fee" bson:"fee"`
	ID       string `json:"id" bson:"id"`
	PayoutAt string `json:"payout_at" bson:"payout_at"`
	Subtotal string `json:"subtotal" bson:"subtotal"`
}

// ExchangeLimits represents exchange limit information for a single user.
type ExchangeLimits struct {
	LimitCurrency  string         `json:"limit_currency" bson:"limit_currency"`
	TransferLimits TransferLimits `json:"transfer_limits" bson:"transfer_limits"`
}

// Fees are fees rates and 30 days trailing volume.
type Fees struct {
	MakerFeeRate string `json:"maker_fee_rate" bson:"maker_fee_rate"`
	TakerFeeRate string `json:"taker_fee_rate" bson:"taker_fee_rate"`
	UsdVolume    string `json:"usd_volume" bson:"usd_volume"`
}

// FIATAccount references a FIAT account thata CoinbasePaymentMethod belongs to
type FIATAccount struct {
	ID           string `json:"id" bson:"id"`
	Resource     string `json:"resource" bson:"resource"`
	ResourcePath string `json:"resource_path" bson:"resource_path"`
}

// TODO: Get fill description
type Fill struct {
	Fee       string `json:"fee" bson:"fee"`
	Liquidity string `json:"liquidity" bson:"liquidity"`
	OrderID   string `json:"order_id" bson:"order_id"`
	Price     string `json:"price" bson:"price"`
	ProductID string `json:"product_id" bson:"product_id"`
	ProfileID string `json:"profile_id" bson:"profile_id"`
	Settled   bool   `json:"settled" bson:"settled"`
	Side      string `json:"side" bson:"side"`
	Size      string `json:"size" bson:"size"`
	TradeID   int    `json:"trade_id" bson:"trade_id"`
	UsdVolume string `json:"usd_volume" bson:"usd_volume"`
	UserID    string `json:"user_id" bson:"user_id"`
}

// TODO
type Flags struct{}

// Limits defines limits for a payment method
type Limits struct {
	Name string `json:"name" bson:"name"`
	Type string `json:"type" bson:"type"`
}

// Oracle is cryptographically signed price-info ready to be posted on-chain using Compound's Open Oracle smart
// contract.
type Oracle struct {
	// Messages are an array contains abi-encoded values [kind, timestamp, key, value], where kind always equals to
	// 'prices', timestamp is the time when the price was obtained, key is asset ticker (e.g. 'eth') and value is asset
	// price
	Messages []string `json:"messages" bson:"messages"`

	// Prices contains human-readable asset prices
	Prices OraclePrices `json:"prices" bson:"prices"`

	// Signatures are an array of Ethereum-compatible ECDSA signatures for each message
	Signatures []string `json:"signatures" bson:"signatures"`

	// Timestamp indicates when the latest datapoint was obtained
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

// OraclePrices contain human-readable asset prices.
type OraclePrices struct {
	AdditionalProp string `json:"additionalProp" bson:"additionalProp"`
}

// Order is an open order.
type Order struct {
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	DoneAt         time.Time          `json:"done_at" bson:"done_at"`
	DoneReason     string             `json:"done_reason" bson:"done_reason"`
	ExecutedValue  string             `json:"executed_value" bson:"executed_value"`
	ExpireTime     time.Time          `json:"expire_time" bson:"expire_time"`
	FillFees       string             `json:"fill_fees" bson:"fill_fees"`
	FilledSize     string             `json:"filled_size" bson:"filled_size"`
	FundingAmount  string             `json:"funding_amount" bson:"funding_amount"`
	Funds          string             `json:"funds" bson:"funds"`
	ID             string             `json:"id" bson:"id"`
	PostOnly       bool               `json:"post_only" bson:"post_only"`
	Price          string             `json:"price" bson:"price"`
	ProductID      string             `json:"product_id" bson:"product_id"`
	RejectReason   string             `json:"reject_reason" bson:"reject_reason"`
	Settled        bool               `json:"settled" bson:"settled"`
	Side           scalar.OrderSide   `json:"side" bson:"side"`
	Size           string             `json:"size" bson:"size"`
	SpecifiedFunds string             `json:"specified_funds" bson:"specified_funds"`
	Status         string             `json:"status" bson:"status"`
	Stop           string             `json:"stop" bson:"stop"`
	StopPrice      string             `json:"stop_price" bson:"stop_price"`
	TimeInForce    scalar.TimeInForce `json:"time_in_force" bson:"time_in_force"`
	Type           scalar.OrderType   `json:"type" bson:"type"`
}

// PaymentMethod is a payment method used on coinbase
type PaymentMethod struct {
	AllowBuy           bool                `json:"allow_buy" bson:"allow_buy"`
	AllowDeposit       bool                `json:"allow_deposit" bson:"allow_deposit"`
	AllowSell          bool                `json:"allow_sell" bson:"allow_sell"`
	AllowWithdraw      bool                `json:"allow_withdraw" bson:"allow_withdraw"`
	AvailableBalance   AvailableBalance    `json:"available_balance" bson:"available_balance"`
	CdvStatus          string              `json:"cdv_status" bson:"cdv_status"`
	CreateAt           time.Time           `json:"create_at" bson:"create_at"`
	CryptoAccount      CryptoAccount       `json:"crypto_account" bson:"crypto_account"`
	Currency           string              `json:"currency" bson:"currency"`
	FIATAccount        FIATAccount         `json:"fiat_account" bson:"fiat_account"`
	HoldBusinessDays   int                 `json:"hold_business_days" bson:"hold_business_days"`
	HoldDays           int                 `json:"hold_days" bson:"hold_days"`
	ID                 string              `json:"id" bson:"id"`
	InstantBuy         bool                `json:"instant_buy" bson:"instant_buy"`
	InstantSale        bool                `json:"instant_sale" bson:"instant_sale"`
	Limits             Limits              `json:"limits" bson:"limits"`
	Name               string              `json:"name" bson:"name"`
	PickerData         PickerData          `json:"picker_data" bson:"picker_data"`
	PrimaryBuy         bool                `json:"primary_buy" bson:"primary_buy"`
	PrimarySell        bool                `json:"primary_sell" bson:"primary_sell"`
	RecurringOptions   []*RecurringOptions `json:"recurring_options" bson:"recurring_options"`
	Resource           string              `json:"resource" bson:"resource"`
	ResourcePath       string              `json:"resource_path" bson:"resource_path"`
	Type               string              `json:"type" bson:"type"`
	UpdatedAt          time.Time           `json:"updated_at" bson:"updated_at"`
	VerificationMethod string              `json:"verification_method" bson:"verification_method"`
	Verified           bool                `json:"verified" bson:"verified"`
}

// PickerData ??
type PickerData struct {
	AccountName           string  `json:"account_name" bson:"account_name"`
	AccountNumber         string  `json:"account_number" bson:"account_number"`
	AccountType           string  `json:"account_type" bson:"account_type"`
	Balance               Balance `json:"balance" bson:"balance"`
	BankName              string  `json:"bank_name" bson:"bank_name"`
	BranchName            string  `json:"branch_name" bson:"branch_name"`
	CustomerName          string  `json:"customer_name" bson:"customer_name"`
	Iban                  string  `json:"iban" bson:"iban"`
	IconURL               string  `json:"icon_url" bson:"icon_url"`
	InstitutionCode       string  `json:"institution_code" bson:"institution_code"`
	InstitutionIdentifier string  `json:"institution_identifier" bson:"institution_identifier"`
	InstitutionName       string  `json:"institution_name" bson:"institution_name"`
	PaypalEmail           string  `json:"paypal_email" bson:"paypal_email"`
	PaypalOwner           string  `json:"paypal_owner" bson:"paypal_owner"`
	RoutingNumber         string  `json:"routing_number" bson:"routing_number"`
	SWIFT                 string  `json:"swift" bson:"swift"`
	Symbol                string  `json:"symbol" bson:"symbol"`
}

// Product represents a currency pair available for trading.
type Product struct {
	AuctionMode           bool          `json:"auction_mode" bson:"auction_mode"`
	BaseCurrency          string        `json:"base_currency" bson:"base_currency"`
	BaseIncrement         string        `json:"base_increment" bson:"base_increment"`
	BaseMaxSize           string        `json:"base_max_size" bson:"base_max_size"`
	BaseMinSize           string        `json:"base_min_size" bson:"base_min_size"`
	CancelOnly            bool          `json:"cancel_only" bson:"cancel_only"`
	DisplayName           string        `json:"display_name" bson:"display_name"`
	FxStablecoin          bool          `json:"fx_stablecoin" bson:"fx_stablecoin"`
	ID                    string        `json:"id" bson:"id"`
	LimitOnly             bool          `json:"limit_only" bson:"limit_only"`
	MarginEnabled         bool          `json:"margin_enabled" bson:"margin_enabled"`
	MaxMarketFunds        string        `json:"max_market_funds" bson:"max_market_funds"`
	MaxSlippagePercentage string        `json:"max_slippage_percentage" bson:"max_slippage_percentage"`
	MinMarketFunds        string        `json:"min_market_funds" bson:"min_market_funds"`
	PostOnly              bool          `json:"post_only" bson:"post_only"`
	QuoteCurrency         string        `json:"quote_currency" bson:"quote_currency"`
	QuoteIncrement        string        `json:"quote_increment" bson:"quote_increment"`
	Status                scalar.Status `json:"status" bson:"status"`
	StatusMessage         string        `json:"status_message" bson:"status_message"`
	TradingDisabled       bool          `json:"trading_disabled" bson:"trading_disabled"`
}

// ProductStats are 30day and 24hour stats for a product.
type ProductStats struct {
	High        string `json:"high" bson:"high"`
	Last        string `json:"last" bson:"last"`
	Low         string `json:"low" bson:"low"`
	Open        string `json:"open" bson:"open"`
	Volume      string `json:"volume" bson:"volume"`
	Volume30day string `json:"volume_30day" bson:"volume_30day"`
}

// ProductTicker is a snapshot information about the last trade (tick), best bid/ask and 24h volume.
type ProductTicker struct {
	Ask     string    `json:"ask" bson:"ask"`
	Bid     string    `json:"bid" bson:"bid"`
	Price   string    `json:"price" bson:"price"`
	Size    string    `json:"size" bson:"size"`
	Time    time.Time `json:"time" bson:"time"`
	TradeID int       `json:"trade_id" bson:"trade_id"`
	Volume  string    `json:"volume" bson:"volume"`
}

// Profile represents a profile to interact with the API.
type Profile struct {
	Active    bool      `json:"active" bson:"active"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	HasMargin bool      `json:"has_margin" bson:"has_margin"`
	ID        string    `json:"id" bson:"id"`
	IsDefault bool      `json:"is_default" bson:"is_default"`
	Name      string    `json:"name" bson:"name"`
	UserID    string    `json:"user_id" bson:"user_id"`
}

// RecurringOptions ??
type RecurringOptions struct {
	Label  string `json:"label" bson:"label"`
	Period string `json:"period" bson:"period"`
}

// Report represents a list of past fills/account reports.
type Report struct {
	CreatedAt time.Time         `json:"created_at" bson:"created_at"`
	ExpiresAt time.Time         `json:"expires_at" bson:"expires_at"`
	FileCount string            `json:"file_count" bson:"file_count"`
	FileURL   string            `json:"file_url" bson:"file_url"`
	ID        string            `json:"id" bson:"id"`
	Params    ReportsParams     `json:"params" bson:"params"`
	Status    scalar.Status     `json:"status" bson:"status"`
	Type      scalar.ReportType `json:"type" bson:"type"`
	UpdatedAt time.Time         `json:"updated_at" bson:"updated_at"`
	UserID    string            `json:"user_id" bson:"user_id"`
}

// TODO
type ReportsParams struct {
	AccountID    string        `json:"account_id" bson:"account_id"`
	Email        string        `json:"email" bson:"email"`
	EndDate      time.Time     `json:"end_date" bson:"end_date"`
	Format       scalar.Format `json:"format" bson:"format"`
	NewYorkState bool          `json:"new_york_state" bson:"new_york_state"`
	ProductID    string        `json:"product_id" bson:"product_id"`
	ProfileID    string        `json:"profile_id" bson:"profile_id"`
	StartDate    time.Time     `json:"start_date" bson:"start_date"`
	User         User          `json:"user" bson:"user"`
}

// TODO
type Role struct{}

// SEPADepositInformation information regarding a wallet's deposits. A SEPA credit transfer is a single transfer of
// Euros from one person or organisation to another. For example, this could be to pay the deposit for a holiday rental
// or to settle an invoice. A SEPA direct debit is a recurring payment, for example to pay monthly rent or for a service
// like a mobile phone contract.
type SEPADepositInformation struct {
	AccountAddress string      `json:"account_address" bson:"account_address"`
	AccountName    string      `json:"account_name" bson:"account_name"`
	BankAddress    string      `json:"bank_address" bson:"bank_address"`
	BankCountry    BankCountry `json:"bank_country" bson:"bank_country"`
	BankName       string      `json:"bank_name" bson:"bank_name"`
	Iban           string      `json:"iban" bson:"iban"`
	Reference      string      `json:"reference" bson:"reference"`
	SWIFT          string      `json:"swift" bson:"swift"`
}

// SWIFTDepositInformation information regarding a wallet's deposits. SWIFT stands for Society for Worldwide Interbank
// Financial Telecommunications. Basically, it's a computer network that connects over 900 banks around the world – and
// enables them to transfer money. ING is part of this network. There is no fee for accepting deposits into your account
// with ING.
type SWIFTDepositInformation struct {
	AccountAddress string      `json:"account_address" bson:"account_address"`
	AccountName    string      `json:"account_name" bson:"account_name"`
	AccountNumber  string      `json:"account_number" bson:"account_number"`
	BankAddress    string      `json:"bank_address" bson:"bank_address"`
	BankCountry    BankCountry `json:"bank_country" bson:"bank_country"`
	BankName       string      `json:"bank_name" bson:"bank_name"`
	Reference      string      `json:"reference" bson:"reference"`
}

// Ticker is real-time price updates every time a match happens. It batches updates in case of cascading matches,
// greatly reducing bandwidth requirements.
type Ticker struct {
	BestAsk   string    `json:"best_ask" bson:"best_ask"`
	BestBid   string    `json:"best_bid" bson:"best_bid"`
	LastSize  string    `json:"last_size" bson:"last_size"`
	Price     string    `json:"price" bson:"price"`
	ProductID string    `json:"product_id" bson:"product_id"`
	Sequence  int       `json:"sequence" bson:"sequence"`
	Side      string    `json:"side" bson:"side"`
	Time      time.Time `json:"time" bson:"time"`
	TradeID   int       `json:"trade_id" bson:"trade_id"`
	Type      string    `json:"type" bson:"type"`
}

// Trade is the list the latest trades for a product.
type Trade struct {
	Price   string           `json:"price" bson:"price"`
	Side    scalar.OrderSide `json:"side" bson:"side"`
	Size    string           `json:"size" bson:"size"`
	Time    time.Time        `json:"time" bson:"time"`
	TradeID int32            `json:"trade_id" bson:"trade_id"`
}

// Transfer will lists past withdrawals and deposits for an account.
type Transfer struct {
	Amount      string                 `json:"amount" bson:"amount"`
	CanceledAt  time.Time              `json:"canceled_at" bson:"canceled_at"`
	CompletedAt time.Time              `json:"completed_at" bson:"completed_at"`
	CreatedAt   time.Time              `json:"created_at" bson:"created_at"`
	Details     AccountTransferDetails `json:"details" bson:"details"`
	ID          string                 `json:"id" bson:"id"`
	ProcessedAt time.Time              `json:"processed_at" bson:"processed_at"`
	Type        string                 `json:"type" bson:"type"`
	UserNonce   string                 `json:"user_nonce" bson:"user_nonce"`
}

// TODO
type TransferLimits struct {
	ACH                   CurrencyTransferLimits `json:"ach" bson:"ach"`
	ACHNoBalance          CurrencyTransferLimits `json:"ach_no_balance" bson:"ach_no_balance"`
	Buy                   CurrencyTransferLimits `json:"buy" bson:"buy"`
	CreditDebitCard       CurrencyTransferLimits `json:"credit_debit_card" bson:"credit_debit_card"`
	ExchangeWithdraw      CurrencyTransferLimits `json:"exchange_withdraw" bson:"exchange_withdraw"`
	IdealDeposit          CurrencyTransferLimits `json:"ideal_deposit" bson:"ideal_deposit"`
	InstanceACHWithdrawal CurrencyTransferLimits `json:"instance_ach_withdrawal" bson:"instance_ach_withdrawal"`
	PaypalBuy             CurrencyTransferLimits `json:"paypal_buy" bson:"paypal_buy"`
	PaypalWithdrawal      CurrencyTransferLimits `json:"paypal_withdrawal" bson:"paypal_withdrawal"`
	Secure3dBuy           CurrencyTransferLimits `json:"secure3d_buy" bson:"secure3d_buy"`
	Sell                  CurrencyTransferLimits `json:"sell" bson:"sell"`
	SofortDeposit         CurrencyTransferLimits `json:"sofort_deposit" bson:"sofort_deposit"`
}

// UKDepositInformation information regarding a wallet's deposits.
type UKDepositInformation struct {
	AccountAddress string      `json:"account_address" bson:"account_address"`
	AccountName    string      `json:"account_name" bson:"account_name"`
	AccountNumber  string      `json:"account_number" bson:"account_number"`
	BankAddress    string      `json:"bank_address" bson:"bank_address"`
	BankCountry    BankCountry `json:"bank_country" bson:"bank_country"`
	BankName       string      `json:"bank_name" bson:"bank_name"`
	Reference      string      `json:"reference" bson:"reference"`
}

// TODO
type User struct {
	ActiveAt                  time.Time       `json:"active_at" bson:"active_at"`
	CbDataFromCache           bool            `json:"cb_data_from_cache" bson:"cb_data_from_cache"`
	CreatedAt                 time.Time       `json:"created_at" bson:"created_at"`
	Details                   UserDetails     `json:"details" bson:"details"`
	Flags                     Flags           `json:"flags" bson:"flags"`
	FulfillsNewRequirements   bool            `json:"fulfills_new_requirements" bson:"fulfills_new_requirements"`
	HasClawbackPaymentPending bool            `json:"has_clawback_payment_pending" bson:"has_clawback_payment_pending"`
	HasDefault                bool            `json:"has_default" bson:"has_default"`
	HasRestrictedAssets       bool            `json:"has_restricted_assets" bson:"has_restricted_assets"`
	ID                        string          `json:"id" bson:"id"`
	IsBanned                  bool            `json:"is_banned" bson:"is_banned"`
	LegalName                 string          `json:"legal_name" bson:"legal_name"`
	Name                      string          `json:"name" bson:"name"`
	Preferences               UserPreferences `json:"preferences" bson:"preferences"`
	Roles                     []*Role         `json:"roles" bson:"roles"`
	StateCode                 string          `json:"state_code" bson:"state_code"`
	TermsAccepted             time.Time       `json:"terms_accepted" bson:"terms_accepted"`
	TwoFactorMethod           string          `json:"two_factor_method" bson:"two_factor_method"`
	UserType                  string          `json:"user_type" bson:"user_type"`
}

// TODO
type UserDetails struct{}

// TODO
type UserPreferences struct{}

// Wallet represents a user's available Coinbase wallet (These are the wallets/accounts that are used for buying and
// selling on www.coinbase.com)
type Wallet struct {
	Active                  bool                    `json:"active" bson:"active"`
	AvailableOnConsumer     bool                    `json:"available_on_consumer" bson:"available_on_consumer"`
	Balance                 string                  `json:"balance" bson:"balance"`
	Currency                string                  `json:"currency" bson:"currency"`
	DestinationTagName      string                  `json:"destination_tag_name" bson:"destination_tag_name"`
	DestinationTagRegex     string                  `json:"destination_tag_regex" bson:"destination_tag_regex"`
	HoldBalance             string                  `json:"hold_balance" bson:"hold_balance"`
	HoldCurrency            string                  `json:"hold_currency" bson:"hold_currency"`
	ID                      string                  `json:"id" bson:"id"`
	Name                    string                  `json:"name" bson:"name"`
	Primary                 bool                    `json:"primary" bson:"primary"`
	Ready                   bool                    `json:"ready" bson:"ready"`
	SEPADepositInformation  SEPADepositInformation  `json:"sepa_deposit_information" bson:"sepa_deposit_information"`
	SWIFTDepositInformation SWIFTDepositInformation `json:"swift_deposit_information" bson:"swift_deposit_information"`
	Type                    string                  `json:"type" bson:"type"`
	UKDepositInformation    UKDepositInformation    `json:"uk_deposit_information" bson:"uk_deposit_information"`
	WireDepositInformation  WireDepositInformation  `json:"wire_deposit_information" bson:"wire_deposit_information"`
}

// WireDepositInformation information regarding a wallet's deposits
type WireDepositInformation struct {
	AccountAddress string      `json:"account_address" bson:"account_address"`
	AccountName    string      `json:"account_name" bson:"account_name"`
	AccountNumber  string      `json:"account_number" bson:"account_number"`
	BankAddress    string      `json:"bank_address" bson:"bank_address"`
	BankCountry    BankCountry `json:"bank_country" bson:"bank_country"`
	BankName       string      `json:"bank_name" bson:"bank_name"`
	Reference      string      `json:"reference" bson:"reference"`
	RoutingNumber  string      `json:"routing_number" bson:"routing_number"`
}

// Withdrawal is data concerning withdrawing funds from the specified profile_id to a www.coinbase.com wallet.
type Withdrawal struct {
	Amount   string `json:"amount" bson:"amount"`
	Currency string `json:"currency" bson:"currency"`
	Fee      string `json:"fee" bson:"fee"`
	ID       string `json:"id" bson:"id"`
	PayoutAt string `json:"payout_at" bson:"payout_at"`
	Subtotal string `json:"subtotal" bson:"subtotal"`
}

// WithdrawalFeeEstimate is a fee estimate for the crypto withdrawal to crypto address
type WithdrawalFeeEstimate struct {
	Fee float64 `json:"fee" bson:"fee"`
}

// UnmarshalJSON will deserialize bytes into a Oracle model
func (Oracle *Oracle) UnmarshalJSON(d []byte) error {
	const (
		timestampJSONTag  = "timestamp"
		messagesJSONTag   = "messages"
		signaturesJSONTag = "signatures"
		pricesJSONTag     = "prices"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	Oracle.Prices = OraclePrices{}
	if err := data.UnmarshalStruct(pricesJSONTag, &Oracle.Prices); err != nil {
		return err
	}
	data.UnmarshalStringSlice(messagesJSONTag, &Oracle.Messages)
	data.UnmarshalStringSlice(signaturesJSONTag, &Oracle.Signatures)
	err = data.UnmarshalUnixString(timestampJSONTag, &Oracle.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a Transfer model
func (Transfer *Transfer) UnmarshalJSON(d []byte) error {
	const (
		IDJSONTag          = "id"
		typeJSONTag        = "type"
		createdAtJSONTag   = "created_at"
		completedAtJSONTag = "completed_at"
		canceledAtJSONTag  = "canceled_at"
		processedAtJSONTag = "processed_at"
		amountJSONTag      = "amount"
		userNonceJSONTag   = "user_nonce"
		detailsJSONTag     = "details"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	Transfer.Details = AccountTransferDetails{}
	if err := data.UnmarshalStruct(detailsJSONTag, &Transfer.Details); err != nil {
		return err
	}
	data.UnmarshalString(IDJSONTag, &Transfer.ID)
	data.UnmarshalString(amountJSONTag, &Transfer.Amount)
	data.UnmarshalString(typeJSONTag, &Transfer.Type)
	data.UnmarshalString(userNonceJSONTag, &Transfer.UserNonce)
	err = data.UnmarshalTime(coinbaseTimeLayout1, canceledAtJSONTag, &Transfer.CanceledAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(coinbaseTimeLayout1, completedAtJSONTag, &Transfer.CompletedAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(coinbaseTimeLayout1, createdAtJSONTag, &Transfer.CreatedAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(coinbaseTimeLayout1, processedAtJSONTag, &Transfer.ProcessedAt)
	if err != nil {
		return err
	}
	return nil
}
