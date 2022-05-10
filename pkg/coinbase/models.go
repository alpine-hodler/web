package coinbase

import (
	"encoding/json"
	"time"

	"github.com/alpine-hodler/sdk/internal/serial"
	"github.com/alpine-hodler/sdk/pkg/scalar"
)

// * This is a generated file, do not edit

// Account holds data for trading account from the profile of the API key
type Account struct {
	Available      float64 `json:"available" bson:"available"`
	Balance        float64 `json:"balance" bson:"balance"`
	Currency       string  `json:"currency" bson:"currency"`
	Hold           float64 `json:"hold" bson:"hold"`
	Id             string  `json:"id" bson:"id"`
	ProfileId      string  `json:"profile_id" bson:"profile_id"`
	TradingEnabled bool    `json:"trading_enabled" bson:"trading_enabled"`
}

// CoinbaseHold represents the hold on an account that belong to the same profile as the API key. Holds are placed on an
// account for any active orders or pending withdraw requests. As an order is filled, the hold amount is updated. If an
// order is canceled, any remaining hold is removed. For withdrawals, the hold is removed after it is completed.
type AccountHold struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Id        string    `json:"id" bson:"id"`
	Ref       string    `json:"ref" bson:"ref"`
	Type      string    `json:"type" bson:"type"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

// AccountLedger lists ledger activity for an account. This includes anything that would affect the accounts balance -
// transfers, trades, fees, etc.
type AccountLedger struct {
	Amount    float64              `json:"amount" bson:"amount"`
	Balance   float64              `json:"balance" bson:"balance"`
	CreatedAt time.Time            `json:"created_at" bson:"created_at"`
	Details   AccountLedgerDetails `json:"details" bson:"details"`
	Id        string               `json:"id" bson:"id"`
	Type      scalar.EntryType     `json:"type" bson:"type"`
}

// AccountLedgerDetails are the details for account history.
type AccountLedgerDetails struct {
	OrderId   string `json:"order_id" bson:"order_id"`
	ProductId string `json:"product_id" bson:"product_id"`
	TradeId   string `json:"trade_id" bson:"trade_id"`
}

// AccountTransfer will lists past withdrawals and deposits for an account.
type AccountTransfer struct {
	Amount      float64                `json:"amount" bson:"amount"`
	CanceledAt  time.Time              `json:"canceled_at" bson:"canceled_at"`
	CompletedAt time.Time              `json:"completed_at" bson:"completed_at"`
	CreatedAt   time.Time              `json:"created_at" bson:"created_at"`
	Details     AccountTransferDetails `json:"details" bson:"details"`
	Id          string                 `json:"id" bson:"id"`
	ProcessedAt time.Time              `json:"processed_at" bson:"processed_at"`
	Type        string                 `json:"type" bson:"type"`
	UserNonce   string                 `json:"user_nonce" bson:"user_nonce"`
}

// AccountTransferDetails are the details for an account transfer.
type AccountTransferDetails struct {
	CoinbaseAccountId       string `json:"coinbase_account_id" bson:"coinbase_account_id"`
	CoinbasePaymentMethodId string `json:"coinbase_payment_method_id" bson:"coinbase_payment_method_id"`
	CoinbaseTransactionId   string `json:"coinbase_transaction_id" bson:"coinbase_transaction_id"`
}

// Auction is an object of data concerning a book request.
type Auction struct {
	AuctionState string    `json:"auction_state" bson:"auction_state"`
	BestAskPrice float64   `json:"best_ask_price" bson:"best_ask_price"`
	BestAskSize  float64   `json:"best_ask_size" bson:"best_ask_size"`
	BestBidPrice float64   `json:"best_bid_price" bson:"best_bid_price"`
	BestBidSize  float64   `json:"best_bid_size" bson:"best_bid_size"`
	CanOpen      string    `json:"can_open" bson:"can_open"`
	OpenPrice    float64   `json:"open_price" bson:"open_price"`
	OpenSize     float64   `json:"open_size" bson:"open_size"`
	Time         time.Time `json:"time" bson:"time"`
}

// AvailableBalance is the available balance on the coinbase account
type AvailableBalance struct {
	Amount   float64 `json:"amount" bson:"amount"`
	Currency string  `json:"currency" bson:"currency"`
	Scale    string  `json:"scale" bson:"scale"`
}

// Balance is the balance for picker data
type Balance struct {
	Amount   float64 `json:"amount" bson:"amount"`
	Currency string  `json:"currency" bson:"currency"`
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

// CryptoAccount references a crypto account that a CoinbasePaymentMethod belongs to
type CryptoAccount struct {
	Id           string `json:"id" bson:"id"`
	Resource     string `json:"resource" bson:"resource"`
	ResourcePath string `json:"resource_path" bson:"resource_path"`
}

// CryptoAddress is used for a one-time crypto address for depositing crypto.
type CryptoAddress struct {
	Address        string                  `json:"address" bson:"address"`
	AddressInfo    CryptoAddressInfo       `json:"address_info" bson:"address_info"`
	CallbackUrl    string                  `json:"callback_url" bson:"callback_url"`
	CreateAt       time.Time               `json:"create_at" bson:"create_at"`
	DepositUri     string                  `json:"deposit_uri" bson:"deposit_uri"`
	DestinationTag string                  `json:"destination_tag" bson:"destination_tag"`
	Id             string                  `json:"id" bson:"id"`
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
	ImageUrl string `json:"image_url" bson:"image_url"`
	Title    string `json:"title" bson:"title"`
}

// Currency is a currency that coinbase knows about. Not al currencies may be currently in use for trading.
type Currency struct {
	ConvertibleTo []string        `json:"convertible_to" bson:"convertible_to"`
	Details       CurrencyDetails `json:"details" bson:"details"`
	Id            string          `json:"id" bson:"id"`
	MaxPrecision  float64         `json:"max_precision" bson:"max_precision"`
	Message       string          `json:"message" bson:"message"`
	MinSize       float64         `json:"min_size" bson:"min_size"`
	Name          string          `json:"name" bson:"name"`
	Status        string          `json:"status" bson:"status"`
}

// CurrencyConversion is the response that converts funds from from currency to to currency. Funds are converted on the
// from account in the profile_id profile.
type CurrencyConversion struct {
	Amount        float64 `json:"amount" bson:"amount"`
	From          string  `json:"from" bson:"from"`
	FromAccountId string  `json:"from_account_id" bson:"from_account_id"`
	Id            string  `json:"id" bson:"id"`
	To            string  `json:"to" bson:"to"`
	ToAccountId   string  `json:"to_account_id" bson:"to_account_id"`
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

// Deposit is the response for deposited funds from a www.coinbase.com wallet to the specified profile_id.
type Deposit struct {
	Amount   float64 `json:"amount" bson:"amount"`
	Currency string  `json:"currency" bson:"currency"`
	Fee      float64 `json:"fee" bson:"fee"`
	Id       string  `json:"id" bson:"id"`
	PayoutAt string  `json:"payout_at" bson:"payout_at"`
	Subtotal float64 `json:"subtotal" bson:"subtotal"`
}

// Fees are fees rates and 30 days trailing volume.
type Fees struct {
	MakerFeeRate float64 `json:"maker_fee_rate" bson:"maker_fee_rate"`
	TakerFeeRate float64 `json:"taker_fee_rate" bson:"taker_fee_rate"`
	UsdVolume    float64 `json:"usd_volume" bson:"usd_volume"`
}

// FiatAccount references a FIAT account thata CoinbasePaymentMethod belongs to
type FiatAccount struct {
	Id           string `json:"id" bson:"id"`
	Resource     string `json:"resource" bson:"resource"`
	ResourcePath string `json:"resource_path" bson:"resource_path"`
}

// TODO: Get fill description
type Fill struct {
	Fee       float64 `json:"fee" bson:"fee"`
	Liquidity string  `json:"liquidity" bson:"liquidity"`
	OrderId   string  `json:"order_id" bson:"order_id"`
	Price     float64 `json:"price" bson:"price"`
	ProductId string  `json:"product_id" bson:"product_id"`
	ProfileId string  `json:"profile_id" bson:"profile_id"`
	Settled   bool    `json:"settled" bson:"settled"`
	Side      string  `json:"side" bson:"side"`
	Size      float64 `json:"size" bson:"size"`
	TradeId   int     `json:"trade_id" bson:"trade_id"`
	UsdVolume float64 `json:"usd_volume" bson:"usd_volume"`
	UserId    string  `json:"user_id" bson:"user_id"`
}

// Limits defines limits for a payment method
type Limits struct {
	Name string `json:"name" bson:"name"`
	Type string `json:"type" bson:"type"`
}

// NewOrder is the server's response for placing a new order.
type NewOrder struct {
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	DoneAt        time.Time          `json:"done_at" bson:"done_at"`
	DoneReason    string             `json:"done_reason" bson:"done_reason"`
	ExpireTime    time.Time          `json:"expire_time" bson:"expire_time"`
	FillFees      float64            `json:"fill_fees" bson:"fill_fees"`
	FilledSize    float64            `json:"filled_size" bson:"filled_size"`
	FundingAmount float64            `json:"funding_amount" bson:"funding_amount"`
	Funds         float64            `json:"funds" bson:"funds"`
	Id            string             `json:"id" bson:"id"`
	PostOnly      bool               `json:"post_only" bson:"post_only"`
	Price         float64            `json:"price" bson:"price"`
	ProductId     string             `json:"product_id" bson:"product_id"`
	ProfileId     string             `json:"profile_id" bson:"profile_id"`
	RejectReason  string             `json:"reject_reason" bson:"reject_reason"`
	Settled       bool               `json:"settled" bson:"settled"`
	Side          scalar.OrderSide   `json:"side" bson:"side"`
	Size          float64            `json:"size" bson:"size"`
	SpecificFunds float64            `json:"specific_funds" bson:"specific_funds"`
	Status        string             `json:"status" bson:"status"`
	Stop          scalar.OrderStop   `json:"stop" bson:"stop"`
	StopPrice     float64            `json:"stop_price" bson:"stop_price"`
	TimeInForce   scalar.TimeInForce `json:"time_in_force" bson:"time_in_force"`
	Type          scalar.OrderType   `json:"type" bson:"type"`
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
	ExecutedValue  float64            `json:"executed_value" bson:"executed_value"`
	ExpireTime     time.Time          `json:"expire_time" bson:"expire_time"`
	FillFees       float64            `json:"fill_fees" bson:"fill_fees"`
	FilledSize     float64            `json:"filled_size" bson:"filled_size"`
	FundingAmount  float64            `json:"funding_amount" bson:"funding_amount"`
	Funds          float64            `json:"funds" bson:"funds"`
	Id             string             `json:"id" bson:"id"`
	PostOnly       bool               `json:"post_only" bson:"post_only"`
	Price          float64            `json:"price" bson:"price"`
	ProductId      string             `json:"product_id" bson:"product_id"`
	RejectReason   string             `json:"reject_reason" bson:"reject_reason"`
	Settled        bool               `json:"settled" bson:"settled"`
	Side           scalar.OrderSide   `json:"side" bson:"side"`
	Size           float64            `json:"size" bson:"size"`
	SpecifiedFunds float64            `json:"specified_funds" bson:"specified_funds"`
	Status         string             `json:"status" bson:"status"`
	Stop           string             `json:"stop" bson:"stop"`
	StopPrice      float64            `json:"stop_price" bson:"stop_price"`
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
	FiatAccount        FiatAccount         `json:"fiat_account" bson:"fiat_account"`
	HoldBusinessDays   int                 `json:"hold_business_days" bson:"hold_business_days"`
	HoldDays           int                 `json:"hold_days" bson:"hold_days"`
	Id                 string              `json:"id" bson:"id"`
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
	IconUrl               string  `json:"icon_url" bson:"icon_url"`
	InstitutionCode       string  `json:"institution_code" bson:"institution_code"`
	InstitutionIdentifier string  `json:"institution_identifier" bson:"institution_identifier"`
	InstitutionName       string  `json:"institution_name" bson:"institution_name"`
	PaypalEmail           string  `json:"paypal_email" bson:"paypal_email"`
	PaypalOwner           string  `json:"paypal_owner" bson:"paypal_owner"`
	RoutingNumber         string  `json:"routing_number" bson:"routing_number"`
	Swift                 string  `json:"swift" bson:"swift"`
	Symbol                string  `json:"symbol" bson:"symbol"`
}

// Product represents a currency pair available for trading.
type Product struct {
	AuctionMode           bool          `json:"auction_mode" bson:"auction_mode"`
	BaseCurrency          string        `json:"base_currency" bson:"base_currency"`
	BaseIncrement         float64       `json:"base_increment" bson:"base_increment"`
	BaseMaxSize           float64       `json:"base_max_size" bson:"base_max_size"`
	BaseMinSize           float64       `json:"base_min_size" bson:"base_min_size"`
	CancelOnly            bool          `json:"cancel_only" bson:"cancel_only"`
	DisplayName           string        `json:"display_name" bson:"display_name"`
	FxStablecoin          bool          `json:"fx_stablecoin" bson:"fx_stablecoin"`
	Id                    string        `json:"id" bson:"id"`
	LimitOnly             bool          `json:"limit_only" bson:"limit_only"`
	MarginEnabled         bool          `json:"margin_enabled" bson:"margin_enabled"`
	MaxMarketFunds        float64       `json:"max_market_funds" bson:"max_market_funds"`
	MaxSlippagePercentage float64       `json:"max_slippage_percentage" bson:"max_slippage_percentage"`
	MinMarketFunds        float64       `json:"min_market_funds" bson:"min_market_funds"`
	PostOnly              bool          `json:"post_only" bson:"post_only"`
	QuoteCurrency         string        `json:"quote_currency" bson:"quote_currency"`
	QuoteIncrement        float64       `json:"quote_increment" bson:"quote_increment"`
	Status                scalar.Status `json:"status" bson:"status"`
	StatusMessage         string        `json:"status_message" bson:"status_message"`
	TradingDisabled       bool          `json:"trading_disabled" bson:"trading_disabled"`
}

// ProductStats are 30day and 24hour stats for a product.
type ProductStats struct {
	High        float64 `json:"high" bson:"high"`
	Last        float64 `json:"last" bson:"last"`
	Low         float64 `json:"low" bson:"low"`
	Open        float64 `json:"open" bson:"open"`
	Volume      float64 `json:"volume" bson:"volume"`
	Volume30day float64 `json:"volume_30day" bson:"volume_30day"`
}

// ProductTicker is a snapshot information about the last trade (tick), best bid/ask and 24h volume.
type ProductTicker struct {
	Ask     float64   `json:"ask" bson:"ask"`
	Bid     float64   `json:"bid" bson:"bid"`
	Price   float64   `json:"price" bson:"price"`
	Size    float64   `json:"size" bson:"size"`
	Time    time.Time `json:"time" bson:"time"`
	TradeId int       `json:"trade_id" bson:"trade_id"`
	Volume  float64   `json:"volume" bson:"volume"`
}

// Profile represents a profile to interact with the API.
type Profile struct {
	Active    bool      `json:"active" bson:"active"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	HasMargin bool      `json:"has_margin" bson:"has_margin"`
	Id        string    `json:"id" bson:"id"`
	IsDefault bool      `json:"is_default" bson:"is_default"`
	Name      string    `json:"name" bson:"name"`
	UserId    string    `json:"user_id" bson:"user_id"`
}

// RecurringOptions ??
type RecurringOptions struct {
	Label  string `json:"label" bson:"label"`
	Period string `json:"period" bson:"period"`
}

// SepaDepositInformation information regarding a wallet's deposits. A SEPA credit transfer is a single transfer of
// Euros from one person or organisation to another. For example, this could be to pay the deposit for a holiday rental
// or to settle an invoice. A SEPA direct debit is a recurring payment, for example to pay monthly rent or for a service
// like a mobile phone contract.
type SepaDepositInformation struct {
	AccountAddress string      `json:"account_address" bson:"account_address"`
	AccountName    string      `json:"account_name" bson:"account_name"`
	BankAddress    string      `json:"bank_address" bson:"bank_address"`
	BankCountry    BankCountry `json:"bank_country" bson:"bank_country"`
	BankName       string      `json:"bank_name" bson:"bank_name"`
	Iban           string      `json:"iban" bson:"iban"`
	Reference      string      `json:"reference" bson:"reference"`
	Swift          string      `json:"swift" bson:"swift"`
}

// SwiftDepositInformation information regarding a wallet's deposits. SWIFT stands for Society for Worldwide Interbank
// Financial Telecommunications. Basically, it's a computer network that connects over 900 banks around the world – and
// enables them to transfer money. ING is part of this network. There is no fee for accepting deposits into your account
// with ING.
type SwiftDepositInformation struct {
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
	BestAsk   float64   `json:"best_ask" bson:"best_ask"`
	BestBid   float64   `json:"best_bid" bson:"best_bid"`
	LastSize  float64   `json:"last_size" bson:"last_size"`
	Price     float64   `json:"price" bson:"price"`
	ProductId string    `json:"product_id" bson:"product_id"`
	Sequence  int       `json:"sequence" bson:"sequence"`
	Side      string    `json:"side" bson:"side"`
	Time      time.Time `json:"time" bson:"time"`
	TradeId   int       `json:"trade_id" bson:"trade_id"`
	Type      string    `json:"type" bson:"type"`
}

// UkDepositInformation information regarding a wallet's deposits.
type UkDepositInformation struct {
	AccountAddress string      `json:"account_address" bson:"account_address"`
	AccountName    string      `json:"account_name" bson:"account_name"`
	AccountNumber  string      `json:"account_number" bson:"account_number"`
	BankAddress    string      `json:"bank_address" bson:"bank_address"`
	BankCountry    BankCountry `json:"bank_country" bson:"bank_country"`
	BankName       string      `json:"bank_name" bson:"bank_name"`
	Reference      string      `json:"reference" bson:"reference"`
}

// Wallet represents a user's available Coinbase wallet (These are the wallets/accounts that are used for buying and
// selling on www.coinbase.com)
type Wallet struct {
	Active                  bool                    `json:"active" bson:"active"`
	AvailableOnConsumer     bool                    `json:"available_on_consumer" bson:"available_on_consumer"`
	Balance                 float64                 `json:"balance" bson:"balance"`
	Currency                string                  `json:"currency" bson:"currency"`
	DestinationTagName      string                  `json:"destination_tag_name" bson:"destination_tag_name"`
	DestinationTagRegex     string                  `json:"destination_tag_regex" bson:"destination_tag_regex"`
	HoldBalance             float64                 `json:"hold_balance" bson:"hold_balance"`
	HoldCurrency            string                  `json:"hold_currency" bson:"hold_currency"`
	Id                      string                  `json:"id" bson:"id"`
	Name                    string                  `json:"name" bson:"name"`
	Primary                 bool                    `json:"primary" bson:"primary"`
	Ready                   bool                    `json:"ready" bson:"ready"`
	SepaDepositInformation  SepaDepositInformation  `json:"sepa_deposit_information" bson:"sepa_deposit_information"`
	SwiftDepositInformation SwiftDepositInformation `json:"swift_deposit_information" bson:"swift_deposit_information"`
	Type                    string                  `json:"type" bson:"type"`
	UkDepositInformation    UkDepositInformation    `json:"uk_deposit_information" bson:"uk_deposit_information"`
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
	Amount   float64 `json:"amount" bson:"amount"`
	Currency string  `json:"currency" bson:"currency"`
	Fee      float64 `json:"fee" bson:"fee"`
	Id       string  `json:"id" bson:"id"`
	PayoutAt string  `json:"payout_at" bson:"payout_at"`
	Subtotal float64 `json:"subtotal" bson:"subtotal"`
}

// WithdrawalFeeEstimate is a fee estimate for the crypto withdrawal to crypto address
type WithdrawalFeeEstimate struct {
	Fee float64 `json:"fee" bson:"fee"`
}

// UnmarshalJSON will deserialize bytes into a Account model
func (account *Account) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag             = "id"
		currencyJsonTag       = "currency"
		balanceJsonTag        = "balance"
		availableJsonTag      = "available"
		holdJsonTag           = "hold"
		profileIdJsonTag      = "profile_id"
		tradingEnabledJsonTag = "trading_enabled"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(tradingEnabledJsonTag, &account.TradingEnabled)
	data.UnmarshalFloatString(availableJsonTag, &account.Available)
	data.UnmarshalFloatString(balanceJsonTag, &account.Balance)
	data.UnmarshalFloatString(holdJsonTag, &account.Hold)
	data.UnmarshalString(currencyJsonTag, &account.Currency)
	data.UnmarshalString(idJsonTag, &account.Id)
	data.UnmarshalString(profileIdJsonTag, &account.ProfileId)
	return nil
}

// UnmarshalJSON will deserialize bytes into a AccountHold model
func (accountHold *AccountHold) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag        = "id"
		createdAtJsonTag = "created_at"
		updatedAtJsonTag = "updated_at"
		typeJsonTag      = "type"
		refJsonTag       = "ref"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(idJsonTag, &accountHold.Id)
	data.UnmarshalString(refJsonTag, &accountHold.Ref)
	data.UnmarshalString(typeJsonTag, &accountHold.Type)
	err = data.UnmarshalTime(time.RFC3339Nano, createdAtJsonTag, &accountHold.CreatedAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, updatedAtJsonTag, &accountHold.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a AccountLedger model
func (accountLedger *AccountLedger) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag        = "id"
		amountJsonTag    = "amount"
		createdAtJsonTag = "created_at"
		balanceJsonTag   = "balance"
		typeJsonTag      = "type"
		detailsJsonTag   = "details"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	accountLedger.Details = AccountLedgerDetails{}
	if err := data.UnmarshalStruct(detailsJsonTag, &accountLedger.Details); err != nil {
		return err
	}
	data.UnmarshalEntryType(typeJsonTag, &accountLedger.Type)
	data.UnmarshalFloatString(amountJsonTag, &accountLedger.Amount)
	data.UnmarshalFloatString(balanceJsonTag, &accountLedger.Balance)
	data.UnmarshalString(idJsonTag, &accountLedger.Id)
	err = data.UnmarshalTime(time.RFC3339Nano, createdAtJsonTag, &accountLedger.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a AccountLedgerDetails model
func (accountLedgerDetails *AccountLedgerDetails) UnmarshalJSON(d []byte) error {
	const (
		orderIdJsonTag   = "order_id"
		tradeIdJsonTag   = "trade_id"
		productIdJsonTag = "product_id"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(orderIdJsonTag, &accountLedgerDetails.OrderId)
	data.UnmarshalString(productIdJsonTag, &accountLedgerDetails.ProductId)
	data.UnmarshalString(tradeIdJsonTag, &accountLedgerDetails.TradeId)
	return nil
}

// UnmarshalJSON will deserialize bytes into a AccountTransfer model
func (accountTransfer *AccountTransfer) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag          = "id"
		typeJsonTag        = "type"
		createdAtJsonTag   = "created_at"
		completedAtJsonTag = "completed_at"
		canceledAtJsonTag  = "canceled_at"
		processedAtJsonTag = "processed_at"
		amountJsonTag      = "amount"
		userNonceJsonTag   = "user_nonce"
		detailsJsonTag     = "details"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	accountTransfer.Details = AccountTransferDetails{}
	if err := data.UnmarshalStruct(detailsJsonTag, &accountTransfer.Details); err != nil {
		return err
	}
	data.UnmarshalFloatString(amountJsonTag, &accountTransfer.Amount)
	data.UnmarshalString(idJsonTag, &accountTransfer.Id)
	data.UnmarshalString(typeJsonTag, &accountTransfer.Type)
	data.UnmarshalString(userNonceJsonTag, &accountTransfer.UserNonce)
	err = data.UnmarshalTime(CoinbaseTimeLayout1, canceledAtJsonTag, &accountTransfer.CanceledAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(CoinbaseTimeLayout1, completedAtJsonTag, &accountTransfer.CompletedAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(CoinbaseTimeLayout1, createdAtJsonTag, &accountTransfer.CreatedAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(CoinbaseTimeLayout1, processedAtJsonTag, &accountTransfer.ProcessedAt)
	if err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a AccountTransferDetails model
func (accountTransferDetails *AccountTransferDetails) UnmarshalJSON(d []byte) error {
	const (
		coinbaseAccountIdJsonTag       = "coinbase_account_id"
		coinbaseTransactionIdJsonTag   = "coinbase_transaction_id"
		coinbasePaymentMethodIdJsonTag = "coinbase_payment_method_id"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(coinbaseAccountIdJsonTag, &accountTransferDetails.CoinbaseAccountId)
	data.UnmarshalString(coinbasePaymentMethodIdJsonTag, &accountTransferDetails.CoinbasePaymentMethodId)
	data.UnmarshalString(coinbaseTransactionIdJsonTag, &accountTransferDetails.CoinbaseTransactionId)
	return nil
}

// UnmarshalJSON will deserialize bytes into a Auction model
func (auction *Auction) UnmarshalJSON(d []byte) error {
	const (
		openPriceJsonTag    = "open_price"
		openSizeJsonTag     = "open_size"
		bestBidPriceJsonTag = "best_bid_price"
		bestBidSizeJsonTag  = "best_bid_size"
		bestAskPriceJsonTag = "best_ask_price"
		bestAskSizeJsonTag  = "best_ask_size"
		auctionStateJsonTag = "auction_state"
		canOpenJsonTag      = "can_open"
		timeJsonTag         = "time"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloatString(bestAskPriceJsonTag, &auction.BestAskPrice)
	data.UnmarshalFloatString(bestAskSizeJsonTag, &auction.BestAskSize)
	data.UnmarshalFloatString(bestBidPriceJsonTag, &auction.BestBidPrice)
	data.UnmarshalFloatString(bestBidSizeJsonTag, &auction.BestBidSize)
	data.UnmarshalFloatString(openPriceJsonTag, &auction.OpenPrice)
	data.UnmarshalFloatString(openSizeJsonTag, &auction.OpenSize)
	data.UnmarshalString(auctionStateJsonTag, &auction.AuctionState)
	data.UnmarshalString(canOpenJsonTag, &auction.CanOpen)
	err = data.UnmarshalTime(time.RFC3339Nano, timeJsonTag, &auction.Time)
	if err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a AvailableBalance model
func (availableBalance *AvailableBalance) UnmarshalJSON(d []byte) error {
	const (
		amountJsonTag   = "amount"
		currencyJsonTag = "currency"
		scaleJsonTag    = "scale"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloatString(amountJsonTag, &availableBalance.Amount)
	data.UnmarshalString(currencyJsonTag, &availableBalance.Currency)
	data.UnmarshalString(scaleJsonTag, &availableBalance.Scale)
	return nil
}

// UnmarshalJSON will deserialize bytes into a Balance model
func (balance *Balance) UnmarshalJSON(d []byte) error {
	const (
		amountJsonTag   = "amount"
		currencyJsonTag = "currency"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloatString(amountJsonTag, &balance.Amount)
	data.UnmarshalString(currencyJsonTag, &balance.Currency)
	return nil
}

// UnmarshalJSON will deserialize bytes into a BankCountry model
func (bankCountry *BankCountry) UnmarshalJSON(d []byte) error {
	const (
		nameJsonTag = "name"
		codeJsonTag = "code"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(codeJsonTag, &bankCountry.Code)
	data.UnmarshalString(nameJsonTag, &bankCountry.Name)
	return nil
}

// UnmarshalJSON will deserialize bytes into a Book model
func (book *Book) UnmarshalJSON(d []byte) error {
	const (
		bidsJsonTag        = "bids"
		asksJsonTag        = "asks"
		sequenceJsonTag    = "sequence"
		auctionModeJsonTag = "auction_mode"
		auctionJsonTag     = "auction"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	book.Auction = Auction{}
	if err := data.UnmarshalStruct(auctionJsonTag, &book.Auction); err != nil {
		return err
	}
	data.UnmarshalBidAsk(asksJsonTag, &book.Asks)
	data.UnmarshalBidAsk(bidsJsonTag, &book.Bids)
	data.UnmarshalBool(auctionModeJsonTag, &book.AuctionMode)
	data.UnmarshalFloat(sequenceJsonTag, &book.Sequence)
	return nil
}

// UnmarshalJSON will deserialize bytes into a CryptoAccount model
func (cryptoAccount *CryptoAccount) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag           = "id"
		resourceJsonTag     = "resource"
		resourcePathJsonTag = "resource_path"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(idJsonTag, &cryptoAccount.Id)
	data.UnmarshalString(resourceJsonTag, &cryptoAccount.Resource)
	data.UnmarshalString(resourcePathJsonTag, &cryptoAccount.ResourcePath)
	return nil
}

// UnmarshalJSON will deserialize bytes into a CryptoAddress model
func (cryptoAddress *CryptoAddress) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag             = "id"
		addressJsonTag        = "address"
		addressInfoJsonTag    = "address_info"
		nameJsonTag           = "name"
		createAtJsonTag       = "create_at"
		updatedAtJsonTag      = "updated_at"
		networkJsonTag        = "network"
		uriSchemeJsonTag      = "uri_scheme"
		resourceJsonTag       = "resource"
		resourcePathJsonTag   = "resource_path"
		warningsJsonTag       = "warnings"
		legacyAddressJsonTag  = "legacy_address"
		destinationTagJsonTag = "destination_tag"
		depositUriJsonTag     = "deposit_uri"
		callbackUrlJsonTag    = "callback_url"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	cryptoAddress.AddressInfo = CryptoAddressInfo{}
	if err := data.UnmarshalStruct(addressInfoJsonTag, &cryptoAddress.AddressInfo); err != nil {
		return err
	}
	data.UnmarshalString(addressJsonTag, &cryptoAddress.Address)
	data.UnmarshalString(callbackUrlJsonTag, &cryptoAddress.CallbackUrl)
	data.UnmarshalString(depositUriJsonTag, &cryptoAddress.DepositUri)
	data.UnmarshalString(destinationTagJsonTag, &cryptoAddress.DestinationTag)
	data.UnmarshalString(idJsonTag, &cryptoAddress.Id)
	data.UnmarshalString(legacyAddressJsonTag, &cryptoAddress.LegacyAddress)
	data.UnmarshalString(nameJsonTag, &cryptoAddress.Name)
	data.UnmarshalString(networkJsonTag, &cryptoAddress.Network)
	data.UnmarshalString(resourceJsonTag, &cryptoAddress.Resource)
	data.UnmarshalString(resourcePathJsonTag, &cryptoAddress.ResourcePath)
	data.UnmarshalString(uriSchemeJsonTag, &cryptoAddress.UriScheme)
	err = data.UnmarshalTime(time.RFC3339Nano, createAtJsonTag, &cryptoAddress.CreateAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, updatedAtJsonTag, &cryptoAddress.UpdatedAt)
	if err != nil {
		return err
	}
	if v := data.Value(warningsJsonTag); v != nil {
		for _, item := range data.Value(warningsJsonTag).([]interface{}) {
			bytes, _ := json.Marshal(item)
			obj := CryptoAddressWarning{}
			if err := json.Unmarshal(bytes, &obj); err != nil {
				return err
			}
			cryptoAddress.Warnings = append(cryptoAddress.Warnings, &obj)
		}
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a CryptoAddressInfo model
func (cryptoAddressInfo *CryptoAddressInfo) UnmarshalJSON(d []byte) error {
	const (
		addressJsonTag        = "address"
		destinationTagJsonTag = "destination_tag"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(addressJsonTag, &cryptoAddressInfo.Address)
	data.UnmarshalString(destinationTagJsonTag, &cryptoAddressInfo.DestinationTag)
	return nil
}

// UnmarshalJSON will deserialize bytes into a CryptoAddressWarning model
func (cryptoAddressWarning *CryptoAddressWarning) UnmarshalJSON(d []byte) error {
	const (
		titleJsonTag    = "title"
		detailsJsonTag  = "details"
		imageUrlJsonTag = "image_url"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(detailsJsonTag, &cryptoAddressWarning.Details)
	data.UnmarshalString(imageUrlJsonTag, &cryptoAddressWarning.ImageUrl)
	data.UnmarshalString(titleJsonTag, &cryptoAddressWarning.Title)
	return nil
}

// UnmarshalJSON will deserialize bytes into a Currency model
func (currency *Currency) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag            = "id"
		nameJsonTag          = "name"
		minSizeJsonTag       = "min_size"
		statusJsonTag        = "status"
		messageJsonTag       = "message"
		maxPrecisionJsonTag  = "max_precision"
		convertibleToJsonTag = "convertible_to"
		detailsJsonTag       = "details"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	currency.Details = CurrencyDetails{}
	if err := data.UnmarshalStruct(detailsJsonTag, &currency.Details); err != nil {
		return err
	}
	data.UnmarshalFloatString(maxPrecisionJsonTag, &currency.MaxPrecision)
	data.UnmarshalFloatString(minSizeJsonTag, &currency.MinSize)
	data.UnmarshalString(idJsonTag, &currency.Id)
	data.UnmarshalString(messageJsonTag, &currency.Message)
	data.UnmarshalString(nameJsonTag, &currency.Name)
	data.UnmarshalString(statusJsonTag, &currency.Status)
	data.UnmarshalStringSlice(convertibleToJsonTag, &currency.ConvertibleTo)
	return nil
}

// UnmarshalJSON will deserialize bytes into a CurrencyConversion model
func (currencyConversion *CurrencyConversion) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag            = "id"
		amountJsonTag        = "amount"
		fromAccountIdJsonTag = "from_account_id"
		toAccountIdJsonTag   = "to_account_id"
		fromJsonTag          = "from"
		toJsonTag            = "to"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloatString(amountJsonTag, &currencyConversion.Amount)
	data.UnmarshalString(fromAccountIdJsonTag, &currencyConversion.FromAccountId)
	data.UnmarshalString(fromJsonTag, &currencyConversion.From)
	data.UnmarshalString(idJsonTag, &currencyConversion.Id)
	data.UnmarshalString(toAccountIdJsonTag, &currencyConversion.ToAccountId)
	data.UnmarshalString(toJsonTag, &currencyConversion.To)
	return nil
}

// UnmarshalJSON will deserialize bytes into a CurrencyDetails model
func (currencyDetails *CurrencyDetails) UnmarshalJSON(d []byte) error {
	const (
		typeJsonTag                  = "type"
		symbolJsonTag                = "symbol"
		networkConfirmationsJsonTag  = "network_confirmations"
		sortOrderJsonTag             = "sort_order"
		cryptoAddressLinkJsonTag     = "crypto_address_link"
		cryptoTransactionLinkJsonTag = "crypto_transaction_link"
		pushPaymentMethodsJsonTag    = "push_payment_methods"
		groupTypesJsonTag            = "group_types"
		displayNameJsonTag           = "display_name"
		processingTimeSecondsJsonTag = "processing_time_seconds"
		minWithdrawalAmountJsonTag   = "min_withdrawal_amount"
		maxWithdrawalAmountJsonTag   = "max_withdrawal_amount"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloat(maxWithdrawalAmountJsonTag, &currencyDetails.MaxWithdrawalAmount)
	data.UnmarshalFloat(minWithdrawalAmountJsonTag, &currencyDetails.MinWithdrawalAmount)
	data.UnmarshalFloat(processingTimeSecondsJsonTag, &currencyDetails.ProcessingTimeSeconds)
	data.UnmarshalInt(networkConfirmationsJsonTag, &currencyDetails.NetworkConfirmations)
	data.UnmarshalInt(sortOrderJsonTag, &currencyDetails.SortOrder)
	data.UnmarshalString(cryptoAddressLinkJsonTag, &currencyDetails.CryptoAddressLink)
	data.UnmarshalString(cryptoTransactionLinkJsonTag, &currencyDetails.CryptoTransactionLink)
	data.UnmarshalString(displayNameJsonTag, &currencyDetails.DisplayName)
	data.UnmarshalString(symbolJsonTag, &currencyDetails.Symbol)
	data.UnmarshalString(typeJsonTag, &currencyDetails.Type)
	data.UnmarshalStringSlice(groupTypesJsonTag, &currencyDetails.GroupTypes)
	data.UnmarshalStringSlice(pushPaymentMethodsJsonTag, &currencyDetails.PushPaymentMethods)
	return nil
}

// UnmarshalJSON will deserialize bytes into a Deposit model
func (deposit *Deposit) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag       = "id"
		amountJsonTag   = "amount"
		currencyJsonTag = "currency"
		payoutAtJsonTag = "payout_at"
		feeJsonTag      = "fee"
		subtotalJsonTag = "subtotal"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloatString(amountJsonTag, &deposit.Amount)
	data.UnmarshalFloatString(feeJsonTag, &deposit.Fee)
	data.UnmarshalFloatString(subtotalJsonTag, &deposit.Subtotal)
	data.UnmarshalString(currencyJsonTag, &deposit.Currency)
	data.UnmarshalString(idJsonTag, &deposit.Id)
	data.UnmarshalString(payoutAtJsonTag, &deposit.PayoutAt)
	return nil
}

// UnmarshalJSON will deserialize bytes into a Fees model
func (fees *Fees) UnmarshalJSON(d []byte) error {
	const (
		takerFeeRateJsonTag = "taker_fee_rate"
		makerFeeRateJsonTag = "maker_fee_rate"
		usdVolumeJsonTag    = "usd_volume"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloatString(makerFeeRateJsonTag, &fees.MakerFeeRate)
	data.UnmarshalFloatString(takerFeeRateJsonTag, &fees.TakerFeeRate)
	data.UnmarshalFloatString(usdVolumeJsonTag, &fees.UsdVolume)
	return nil
}

// UnmarshalJSON will deserialize bytes into a FiatAccount model
func (fiatAccount *FiatAccount) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag           = "id"
		resourceJsonTag     = "resource"
		resourcePathJsonTag = "resource_path"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(idJsonTag, &fiatAccount.Id)
	data.UnmarshalString(resourceJsonTag, &fiatAccount.Resource)
	data.UnmarshalString(resourcePathJsonTag, &fiatAccount.ResourcePath)
	return nil
}

// UnmarshalJSON will deserialize bytes into a Fill model
func (fill *Fill) UnmarshalJSON(d []byte) error {
	const (
		tradeIdJsonTag   = "trade_id"
		productIdJsonTag = "product_id"
		orderIdJsonTag   = "order_id"
		userIdJsonTag    = "user_id"
		profileIdJsonTag = "profile_id"
		liquidityJsonTag = "liquidity"
		priceJsonTag     = "price"
		sizeJsonTag      = "size"
		feeJsonTag       = "fee"
		sideJsonTag      = "side"
		settledJsonTag   = "settled"
		usdVolumeJsonTag = "usd_volume"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(settledJsonTag, &fill.Settled)
	data.UnmarshalFloatString(feeJsonTag, &fill.Fee)
	data.UnmarshalFloatString(priceJsonTag, &fill.Price)
	data.UnmarshalFloatString(sizeJsonTag, &fill.Size)
	data.UnmarshalFloatString(usdVolumeJsonTag, &fill.UsdVolume)
	data.UnmarshalInt(tradeIdJsonTag, &fill.TradeId)
	data.UnmarshalString(liquidityJsonTag, &fill.Liquidity)
	data.UnmarshalString(orderIdJsonTag, &fill.OrderId)
	data.UnmarshalString(productIdJsonTag, &fill.ProductId)
	data.UnmarshalString(profileIdJsonTag, &fill.ProfileId)
	data.UnmarshalString(sideJsonTag, &fill.Side)
	data.UnmarshalString(userIdJsonTag, &fill.UserId)
	return nil
}

// UnmarshalJSON will deserialize bytes into a Limits model
func (limits *Limits) UnmarshalJSON(d []byte) error {
	const (
		typeJsonTag = "type"
		nameJsonTag = "name"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(nameJsonTag, &limits.Name)
	data.UnmarshalString(typeJsonTag, &limits.Type)
	return nil
}

// UnmarshalJSON will deserialize bytes into a NewOrder model
func (newOrder *NewOrder) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag            = "id"
		priceJsonTag         = "price"
		sizeJsonTag          = "size"
		productIdJsonTag     = "product_id"
		profileIdJsonTag     = "profile_id"
		sideJsonTag          = "side"
		fundsJsonTag         = "funds"
		specificFundsJsonTag = "specific_funds"
		typeJsonTag          = "type"
		timeInForceJsonTag   = "time_in_force"
		expireTimeJsonTag    = "expire_time"
		postOnlyJsonTag      = "post_only"
		createdAtJsonTag     = "created_at"
		doneAtJsonTag        = "done_at"
		doneReasonJsonTag    = "done_reason"
		rejectReasonJsonTag  = "reject_reason"
		fillFeesJsonTag      = "fill_fees"
		filledSizeJsonTag    = "filled_size"
		statusJsonTag        = "status"
		settledJsonTag       = "settled"
		stopJsonTag          = "stop"
		stopPriceJsonTag     = "stop_price"
		fundingAmountJsonTag = "funding_amount"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(postOnlyJsonTag, &newOrder.PostOnly)
	data.UnmarshalBool(settledJsonTag, &newOrder.Settled)
	data.UnmarshalFloatString(fillFeesJsonTag, &newOrder.FillFees)
	data.UnmarshalFloatString(filledSizeJsonTag, &newOrder.FilledSize)
	data.UnmarshalFloatString(fundingAmountJsonTag, &newOrder.FundingAmount)
	data.UnmarshalFloatString(fundsJsonTag, &newOrder.Funds)
	data.UnmarshalFloatString(priceJsonTag, &newOrder.Price)
	data.UnmarshalFloatString(sizeJsonTag, &newOrder.Size)
	data.UnmarshalFloatString(specificFundsJsonTag, &newOrder.SpecificFunds)
	data.UnmarshalFloatString(stopPriceJsonTag, &newOrder.StopPrice)
	data.UnmarshalOrderSide(sideJsonTag, &newOrder.Side)
	data.UnmarshalOrderStop(stopJsonTag, &newOrder.Stop)
	data.UnmarshalOrderType(typeJsonTag, &newOrder.Type)
	data.UnmarshalString(doneReasonJsonTag, &newOrder.DoneReason)
	data.UnmarshalString(idJsonTag, &newOrder.Id)
	data.UnmarshalString(productIdJsonTag, &newOrder.ProductId)
	data.UnmarshalString(profileIdJsonTag, &newOrder.ProfileId)
	data.UnmarshalString(rejectReasonJsonTag, &newOrder.RejectReason)
	data.UnmarshalString(statusJsonTag, &newOrder.Status)
	data.UnmarshalTimeInForce(timeInForceJsonTag, &newOrder.TimeInForce)
	err = data.UnmarshalTime(time.RFC3339Nano, createdAtJsonTag, &newOrder.CreatedAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, doneAtJsonTag, &newOrder.DoneAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, expireTimeJsonTag, &newOrder.ExpireTime)
	if err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a Oracle model
func (oracle *Oracle) UnmarshalJSON(d []byte) error {
	const (
		timestampJsonTag  = "timestamp"
		messagesJsonTag   = "messages"
		signaturesJsonTag = "signatures"
		pricesJsonTag     = "prices"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalStringSlice(messagesJsonTag, &oracle.Messages)
	data.UnmarshalStringSlice(signaturesJsonTag, &oracle.Signatures)
	oracle.Prices = OraclePrices{}
	if err := data.UnmarshalStruct(pricesJsonTag, &oracle.Prices); err != nil {
		return err
	}
	err = data.UnmarshalUnixString(timestampJsonTag, &oracle.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a OraclePrices model
func (oraclePrices *OraclePrices) UnmarshalJSON(d []byte) error {
	const (
		additionalPropJsonTag = "additionalProp"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(additionalPropJsonTag, &oraclePrices.AdditionalProp)
	return nil
}

// UnmarshalJSON will deserialize bytes into a Order model
func (order *Order) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag             = "id"
		priceJsonTag          = "price"
		sizeJsonTag           = "size"
		productIdJsonTag      = "product_id"
		sideJsonTag           = "side"
		fundsJsonTag          = "funds"
		specifiedFundsJsonTag = "specified_funds"
		typeJsonTag           = "type"
		timeInForceJsonTag    = "time_in_force"
		expireTimeJsonTag     = "expire_time"
		postOnlyJsonTag       = "post_only"
		createdAtJsonTag      = "created_at"
		doneAtJsonTag         = "done_at"
		doneReasonJsonTag     = "done_reason"
		rejectReasonJsonTag   = "reject_reason"
		fillFeesJsonTag       = "fill_fees"
		filledSizeJsonTag     = "filled_size"
		executedValueJsonTag  = "executed_value"
		statusJsonTag         = "status"
		settledJsonTag        = "settled"
		stopJsonTag           = "stop"
		stopPriceJsonTag      = "stop_price"
		fundingAmountJsonTag  = "funding_amount"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(postOnlyJsonTag, &order.PostOnly)
	data.UnmarshalBool(settledJsonTag, &order.Settled)
	data.UnmarshalFloatString(executedValueJsonTag, &order.ExecutedValue)
	data.UnmarshalFloatString(fillFeesJsonTag, &order.FillFees)
	data.UnmarshalFloatString(filledSizeJsonTag, &order.FilledSize)
	data.UnmarshalFloatString(fundingAmountJsonTag, &order.FundingAmount)
	data.UnmarshalFloatString(fundsJsonTag, &order.Funds)
	data.UnmarshalFloatString(priceJsonTag, &order.Price)
	data.UnmarshalFloatString(sizeJsonTag, &order.Size)
	data.UnmarshalFloatString(specifiedFundsJsonTag, &order.SpecifiedFunds)
	data.UnmarshalFloatString(stopPriceJsonTag, &order.StopPrice)
	data.UnmarshalOrderSide(sideJsonTag, &order.Side)
	data.UnmarshalOrderType(typeJsonTag, &order.Type)
	data.UnmarshalString(doneReasonJsonTag, &order.DoneReason)
	data.UnmarshalString(idJsonTag, &order.Id)
	data.UnmarshalString(productIdJsonTag, &order.ProductId)
	data.UnmarshalString(rejectReasonJsonTag, &order.RejectReason)
	data.UnmarshalString(statusJsonTag, &order.Status)
	data.UnmarshalString(stopJsonTag, &order.Stop)
	data.UnmarshalTimeInForce(timeInForceJsonTag, &order.TimeInForce)
	err = data.UnmarshalTime(time.RFC3339Nano, createdAtJsonTag, &order.CreatedAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, doneAtJsonTag, &order.DoneAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, expireTimeJsonTag, &order.ExpireTime)
	if err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a PaymentMethod model
func (paymentMethod *PaymentMethod) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag                 = "id"
		typeJsonTag               = "type"
		nameJsonTag               = "name"
		currencyJsonTag           = "currency"
		primaryBuyJsonTag         = "primary_buy"
		primarySellJsonTag        = "primary_sell"
		instantBuyJsonTag         = "instant_buy"
		instantSaleJsonTag        = "instant_sale"
		createAtJsonTag           = "create_at"
		updatedAtJsonTag          = "updated_at"
		resourceJsonTag           = "resource"
		resourcePathJsonTag       = "resource_path"
		verifiedJsonTag           = "verified"
		allowBuyJsonTag           = "allow_buy"
		allowSellJsonTag          = "allow_sell"
		allowDepositJsonTag       = "allow_deposit"
		allowWithdrawJsonTag      = "allow_withdraw"
		holdBusinessDaysJsonTag   = "hold_business_days"
		holdDaysJsonTag           = "hold_days"
		verificationMethodJsonTag = "verification_method"
		cdvStatusJsonTag          = "cdv_status"
		limitsJsonTag             = "limits"
		fiatAccountJsonTag        = "fiat_account"
		cryptoAccountJsonTag      = "crypto_account"
		recurringOptionsJsonTag   = "recurring_options"
		availableBalanceJsonTag   = "available_balance"
		pickerDataJsonTag         = "picker_data"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(allowBuyJsonTag, &paymentMethod.AllowBuy)
	data.UnmarshalBool(allowDepositJsonTag, &paymentMethod.AllowDeposit)
	data.UnmarshalBool(allowSellJsonTag, &paymentMethod.AllowSell)
	data.UnmarshalBool(allowWithdrawJsonTag, &paymentMethod.AllowWithdraw)
	data.UnmarshalBool(instantBuyJsonTag, &paymentMethod.InstantBuy)
	data.UnmarshalBool(instantSaleJsonTag, &paymentMethod.InstantSale)
	data.UnmarshalBool(primaryBuyJsonTag, &paymentMethod.PrimaryBuy)
	data.UnmarshalBool(primarySellJsonTag, &paymentMethod.PrimarySell)
	data.UnmarshalBool(verifiedJsonTag, &paymentMethod.Verified)
	data.UnmarshalInt(holdBusinessDaysJsonTag, &paymentMethod.HoldBusinessDays)
	data.UnmarshalInt(holdDaysJsonTag, &paymentMethod.HoldDays)
	data.UnmarshalString(cdvStatusJsonTag, &paymentMethod.CdvStatus)
	data.UnmarshalString(currencyJsonTag, &paymentMethod.Currency)
	data.UnmarshalString(idJsonTag, &paymentMethod.Id)
	data.UnmarshalString(nameJsonTag, &paymentMethod.Name)
	data.UnmarshalString(resourceJsonTag, &paymentMethod.Resource)
	data.UnmarshalString(resourcePathJsonTag, &paymentMethod.ResourcePath)
	data.UnmarshalString(typeJsonTag, &paymentMethod.Type)
	data.UnmarshalString(verificationMethodJsonTag, &paymentMethod.VerificationMethod)
	err = data.UnmarshalTime(time.RFC3339Nano, createAtJsonTag, &paymentMethod.CreateAt)
	if err != nil {
		return err
	}
	err = data.UnmarshalTime(time.RFC3339Nano, updatedAtJsonTag, &paymentMethod.UpdatedAt)
	if err != nil {
		return err
	}
	if v := data.Value(recurringOptionsJsonTag); v != nil {
		for _, item := range data.Value(recurringOptionsJsonTag).([]interface{}) {
			bytes, _ := json.Marshal(item)
			obj := RecurringOptions{}
			if err := json.Unmarshal(bytes, &obj); err != nil {
				return err
			}
			paymentMethod.RecurringOptions = append(paymentMethod.RecurringOptions, &obj)
		}
	}
	paymentMethod.AvailableBalance = AvailableBalance{}
	if err := data.UnmarshalStruct(availableBalanceJsonTag, &paymentMethod.AvailableBalance); err != nil {
		return err
	}
	paymentMethod.CryptoAccount = CryptoAccount{}
	if err := data.UnmarshalStruct(cryptoAccountJsonTag, &paymentMethod.CryptoAccount); err != nil {
		return err
	}
	paymentMethod.FiatAccount = FiatAccount{}
	if err := data.UnmarshalStruct(fiatAccountJsonTag, &paymentMethod.FiatAccount); err != nil {
		return err
	}
	paymentMethod.Limits = Limits{}
	if err := data.UnmarshalStruct(limitsJsonTag, &paymentMethod.Limits); err != nil {
		return err
	}
	paymentMethod.PickerData = PickerData{}
	if err := data.UnmarshalStruct(pickerDataJsonTag, &paymentMethod.PickerData); err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a PickerData model
func (pickerData *PickerData) UnmarshalJSON(d []byte) error {
	const (
		symbolJsonTag                = "symbol"
		customerNameJsonTag          = "customer_name"
		accountNameJsonTag           = "account_name"
		accountNumberJsonTag         = "account_number"
		accountTypeJsonTag           = "account_type"
		institutionCodeJsonTag       = "institution_code"
		institutionNameJsonTag       = "institution_name"
		ibanJsonTag                  = "iban"
		swiftJsonTag                 = "swift"
		paypalEmailJsonTag           = "paypal_email"
		paypalOwnerJsonTag           = "paypal_owner"
		routingNumberJsonTag         = "routing_number"
		institutionIdentifierJsonTag = "institution_identifier"
		bankNameJsonTag              = "bank_name"
		branchNameJsonTag            = "branch_name"
		iconUrlJsonTag               = "icon_url"
		balanceJsonTag               = "balance"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(accountNameJsonTag, &pickerData.AccountName)
	data.UnmarshalString(accountNumberJsonTag, &pickerData.AccountNumber)
	data.UnmarshalString(accountTypeJsonTag, &pickerData.AccountType)
	data.UnmarshalString(bankNameJsonTag, &pickerData.BankName)
	data.UnmarshalString(branchNameJsonTag, &pickerData.BranchName)
	data.UnmarshalString(customerNameJsonTag, &pickerData.CustomerName)
	data.UnmarshalString(ibanJsonTag, &pickerData.Iban)
	data.UnmarshalString(iconUrlJsonTag, &pickerData.IconUrl)
	data.UnmarshalString(institutionCodeJsonTag, &pickerData.InstitutionCode)
	data.UnmarshalString(institutionIdentifierJsonTag, &pickerData.InstitutionIdentifier)
	data.UnmarshalString(institutionNameJsonTag, &pickerData.InstitutionName)
	data.UnmarshalString(paypalEmailJsonTag, &pickerData.PaypalEmail)
	data.UnmarshalString(paypalOwnerJsonTag, &pickerData.PaypalOwner)
	data.UnmarshalString(routingNumberJsonTag, &pickerData.RoutingNumber)
	data.UnmarshalString(swiftJsonTag, &pickerData.Swift)
	data.UnmarshalString(symbolJsonTag, &pickerData.Symbol)
	pickerData.Balance = Balance{}
	if err := data.UnmarshalStruct(balanceJsonTag, &pickerData.Balance); err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a Product model
func (product *Product) UnmarshalJSON(d []byte) error {
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
	data.UnmarshalBool(auctionModeJsonTag, &product.AuctionMode)
	data.UnmarshalBool(cancelOnlyJsonTag, &product.CancelOnly)
	data.UnmarshalBool(fxStablecoinJsonTag, &product.FxStablecoin)
	data.UnmarshalBool(limitOnlyJsonTag, &product.LimitOnly)
	data.UnmarshalBool(marginEnabledJsonTag, &product.MarginEnabled)
	data.UnmarshalBool(postOnlyJsonTag, &product.PostOnly)
	data.UnmarshalBool(tradingDisabledJsonTag, &product.TradingDisabled)
	data.UnmarshalFloatString(baseIncrementJsonTag, &product.BaseIncrement)
	data.UnmarshalFloatString(baseMaxSizeJsonTag, &product.BaseMaxSize)
	data.UnmarshalFloatString(baseMinSizeJsonTag, &product.BaseMinSize)
	data.UnmarshalFloatString(maxMarketFundsJsonTag, &product.MaxMarketFunds)
	data.UnmarshalFloatString(maxSlippagePercentageJsonTag, &product.MaxSlippagePercentage)
	data.UnmarshalFloatString(minMarketFundsJsonTag, &product.MinMarketFunds)
	data.UnmarshalFloatString(quoteIncrementJsonTag, &product.QuoteIncrement)
	data.UnmarshalStatus(statusJsonTag, &product.Status)
	data.UnmarshalString(baseCurrencyJsonTag, &product.BaseCurrency)
	data.UnmarshalString(displayNameJsonTag, &product.DisplayName)
	data.UnmarshalString(idJsonTag, &product.Id)
	data.UnmarshalString(quoteCurrencyJsonTag, &product.QuoteCurrency)
	data.UnmarshalString(statusMessageJsonTag, &product.StatusMessage)
	return nil
}

// UnmarshalJSON will deserialize bytes into a ProductStats model
func (productStats *ProductStats) UnmarshalJSON(d []byte) error {
	const (
		openJsonTag        = "open"
		highJsonTag        = "high"
		lowJsonTag         = "low"
		lastJsonTag        = "last"
		volumeJsonTag      = "volume"
		volume30dayJsonTag = "volume_30day"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloatString(highJsonTag, &productStats.High)
	data.UnmarshalFloatString(lastJsonTag, &productStats.Last)
	data.UnmarshalFloatString(lowJsonTag, &productStats.Low)
	data.UnmarshalFloatString(openJsonTag, &productStats.Open)
	data.UnmarshalFloatString(volume30dayJsonTag, &productStats.Volume30day)
	data.UnmarshalFloatString(volumeJsonTag, &productStats.Volume)
	return nil
}

// UnmarshalJSON will deserialize bytes into a ProductTicker model
func (productTicker *ProductTicker) UnmarshalJSON(d []byte) error {
	const (
		askJsonTag     = "ask"
		bidJsonTag     = "bid"
		volumeJsonTag  = "volume"
		tradeIdJsonTag = "trade_id"
		priceJsonTag   = "price"
		sizeJsonTag    = "size"
		timeJsonTag    = "time"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloatString(askJsonTag, &productTicker.Ask)
	data.UnmarshalFloatString(bidJsonTag, &productTicker.Bid)
	data.UnmarshalFloatString(priceJsonTag, &productTicker.Price)
	data.UnmarshalFloatString(sizeJsonTag, &productTicker.Size)
	data.UnmarshalFloatString(volumeJsonTag, &productTicker.Volume)
	data.UnmarshalInt(tradeIdJsonTag, &productTicker.TradeId)
	err = data.UnmarshalTime(time.RFC3339Nano, timeJsonTag, &productTicker.Time)
	if err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a Profile model
func (profile *Profile) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag        = "id"
		userIdJsonTag    = "user_id"
		nameJsonTag      = "name"
		activeJsonTag    = "active"
		isDefaultJsonTag = "is_default"
		hasMarginJsonTag = "has_margin"
		createdAtJsonTag = "created_at"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(activeJsonTag, &profile.Active)
	data.UnmarshalBool(hasMarginJsonTag, &profile.HasMargin)
	data.UnmarshalBool(isDefaultJsonTag, &profile.IsDefault)
	data.UnmarshalString(idJsonTag, &profile.Id)
	data.UnmarshalString(nameJsonTag, &profile.Name)
	data.UnmarshalString(userIdJsonTag, &profile.UserId)
	err = data.UnmarshalTime(time.RFC3339Nano, createdAtJsonTag, &profile.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a RecurringOptions model
func (recurringOptions *RecurringOptions) UnmarshalJSON(d []byte) error {
	const (
		periodJsonTag = "period"
		labelJsonTag  = "label"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(labelJsonTag, &recurringOptions.Label)
	data.UnmarshalString(periodJsonTag, &recurringOptions.Period)
	return nil
}

// UnmarshalJSON will deserialize bytes into a SepaDepositInformation model
func (sepaDepositInformation *SepaDepositInformation) UnmarshalJSON(d []byte) error {
	const (
		ibanJsonTag           = "iban"
		swiftJsonTag          = "swift"
		bankNameJsonTag       = "bank_name"
		bankAddressJsonTag    = "bank_address"
		bankCountryJsonTag    = "bank_country"
		accountNameJsonTag    = "account_name"
		accountAddressJsonTag = "account_address"
		referenceJsonTag      = "reference"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(accountAddressJsonTag, &sepaDepositInformation.AccountAddress)
	data.UnmarshalString(accountNameJsonTag, &sepaDepositInformation.AccountName)
	data.UnmarshalString(bankAddressJsonTag, &sepaDepositInformation.BankAddress)
	data.UnmarshalString(bankNameJsonTag, &sepaDepositInformation.BankName)
	data.UnmarshalString(ibanJsonTag, &sepaDepositInformation.Iban)
	data.UnmarshalString(referenceJsonTag, &sepaDepositInformation.Reference)
	data.UnmarshalString(swiftJsonTag, &sepaDepositInformation.Swift)
	sepaDepositInformation.BankCountry = BankCountry{}
	if err := data.UnmarshalStruct(bankCountryJsonTag, &sepaDepositInformation.BankCountry); err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a SwiftDepositInformation model
func (swiftDepositInformation *SwiftDepositInformation) UnmarshalJSON(d []byte) error {
	const (
		accountNumberJsonTag  = "account_number"
		bankNameJsonTag       = "bank_name"
		bankAddressJsonTag    = "bank_address"
		bankCountryJsonTag    = "bank_country"
		accountNameJsonTag    = "account_name"
		accountAddressJsonTag = "account_address"
		referenceJsonTag      = "reference"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(accountAddressJsonTag, &swiftDepositInformation.AccountAddress)
	data.UnmarshalString(accountNameJsonTag, &swiftDepositInformation.AccountName)
	data.UnmarshalString(accountNumberJsonTag, &swiftDepositInformation.AccountNumber)
	data.UnmarshalString(bankAddressJsonTag, &swiftDepositInformation.BankAddress)
	data.UnmarshalString(bankNameJsonTag, &swiftDepositInformation.BankName)
	data.UnmarshalString(referenceJsonTag, &swiftDepositInformation.Reference)
	swiftDepositInformation.BankCountry = BankCountry{}
	if err := data.UnmarshalStruct(bankCountryJsonTag, &swiftDepositInformation.BankCountry); err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a Ticker model
func (ticker *Ticker) UnmarshalJSON(d []byte) error {
	const (
		typeJsonTag      = "type"
		productIdJsonTag = "product_id"
		tradeIdJsonTag   = "trade_id"
		sequenceJsonTag  = "sequence"
		timeJsonTag      = "time"
		sideJsonTag      = "side"
		priceJsonTag     = "price"
		lastSizeJsonTag  = "last_size"
		bestBidJsonTag   = "best_bid"
		bestAskJsonTag   = "best_ask"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloatString(bestAskJsonTag, &ticker.BestAsk)
	data.UnmarshalFloatString(bestBidJsonTag, &ticker.BestBid)
	data.UnmarshalFloatString(lastSizeJsonTag, &ticker.LastSize)
	data.UnmarshalFloatString(priceJsonTag, &ticker.Price)
	data.UnmarshalInt(sequenceJsonTag, &ticker.Sequence)
	data.UnmarshalInt(tradeIdJsonTag, &ticker.TradeId)
	data.UnmarshalString(productIdJsonTag, &ticker.ProductId)
	data.UnmarshalString(sideJsonTag, &ticker.Side)
	data.UnmarshalString(typeJsonTag, &ticker.Type)
	err = data.UnmarshalTime(time.RFC3339Nano, timeJsonTag, &ticker.Time)
	if err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a UkDepositInformation model
func (ukDepositInformation *UkDepositInformation) UnmarshalJSON(d []byte) error {
	const (
		accountNumberJsonTag  = "account_number"
		bankNameJsonTag       = "bank_name"
		bankAddressJsonTag    = "bank_address"
		bankCountryJsonTag    = "bank_country"
		accountNameJsonTag    = "account_name"
		accountAddressJsonTag = "account_address"
		referenceJsonTag      = "reference"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(accountAddressJsonTag, &ukDepositInformation.AccountAddress)
	data.UnmarshalString(accountNameJsonTag, &ukDepositInformation.AccountName)
	data.UnmarshalString(accountNumberJsonTag, &ukDepositInformation.AccountNumber)
	data.UnmarshalString(bankAddressJsonTag, &ukDepositInformation.BankAddress)
	data.UnmarshalString(bankNameJsonTag, &ukDepositInformation.BankName)
	data.UnmarshalString(referenceJsonTag, &ukDepositInformation.Reference)
	ukDepositInformation.BankCountry = BankCountry{}
	if err := data.UnmarshalStruct(bankCountryJsonTag, &ukDepositInformation.BankCountry); err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a Wallet model
func (wallet *Wallet) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag                      = "id"
		nameJsonTag                    = "name"
		balanceJsonTag                 = "balance"
		currencyJsonTag                = "currency"
		typeJsonTag                    = "type"
		primaryJsonTag                 = "primary"
		activeJsonTag                  = "active"
		availableOnConsumerJsonTag     = "available_on_consumer"
		readyJsonTag                   = "ready"
		wireDepositInformationJsonTag  = "wire_deposit_information"
		swiftDepositInformationJsonTag = "swift_deposit_information"
		sepaDepositInformationJsonTag  = "sepa_deposit_information"
		ukDepositInformationJsonTag    = "uk_deposit_information"
		destinationTagNameJsonTag      = "destination_tag_name"
		destinationTagRegexJsonTag     = "destination_tag_regex"
		holdBalanceJsonTag             = "hold_balance"
		holdCurrencyJsonTag            = "hold_currency"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(activeJsonTag, &wallet.Active)
	data.UnmarshalBool(availableOnConsumerJsonTag, &wallet.AvailableOnConsumer)
	data.UnmarshalBool(primaryJsonTag, &wallet.Primary)
	data.UnmarshalBool(readyJsonTag, &wallet.Ready)
	data.UnmarshalFloatString(balanceJsonTag, &wallet.Balance)
	data.UnmarshalFloatString(holdBalanceJsonTag, &wallet.HoldBalance)
	data.UnmarshalString(currencyJsonTag, &wallet.Currency)
	data.UnmarshalString(destinationTagNameJsonTag, &wallet.DestinationTagName)
	data.UnmarshalString(destinationTagRegexJsonTag, &wallet.DestinationTagRegex)
	data.UnmarshalString(holdCurrencyJsonTag, &wallet.HoldCurrency)
	data.UnmarshalString(idJsonTag, &wallet.Id)
	data.UnmarshalString(nameJsonTag, &wallet.Name)
	data.UnmarshalString(typeJsonTag, &wallet.Type)
	wallet.SepaDepositInformation = SepaDepositInformation{}
	if err := data.UnmarshalStruct(sepaDepositInformationJsonTag, &wallet.SepaDepositInformation); err != nil {
		return err
	}
	wallet.SwiftDepositInformation = SwiftDepositInformation{}
	if err := data.UnmarshalStruct(swiftDepositInformationJsonTag, &wallet.SwiftDepositInformation); err != nil {
		return err
	}
	wallet.UkDepositInformation = UkDepositInformation{}
	if err := data.UnmarshalStruct(ukDepositInformationJsonTag, &wallet.UkDepositInformation); err != nil {
		return err
	}
	wallet.WireDepositInformation = WireDepositInformation{}
	if err := data.UnmarshalStruct(wireDepositInformationJsonTag, &wallet.WireDepositInformation); err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a WireDepositInformation model
func (wireDepositInformation *WireDepositInformation) UnmarshalJSON(d []byte) error {
	const (
		accountNumberJsonTag  = "account_number"
		routingNumberJsonTag  = "routing_number"
		bankNameJsonTag       = "bank_name"
		bankAddressJsonTag    = "bank_address"
		bankCountryJsonTag    = "bank_country"
		accountNameJsonTag    = "account_name"
		accountAddressJsonTag = "account_address"
		referenceJsonTag      = "reference"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalString(accountAddressJsonTag, &wireDepositInformation.AccountAddress)
	data.UnmarshalString(accountNameJsonTag, &wireDepositInformation.AccountName)
	data.UnmarshalString(accountNumberJsonTag, &wireDepositInformation.AccountNumber)
	data.UnmarshalString(bankAddressJsonTag, &wireDepositInformation.BankAddress)
	data.UnmarshalString(bankNameJsonTag, &wireDepositInformation.BankName)
	data.UnmarshalString(referenceJsonTag, &wireDepositInformation.Reference)
	data.UnmarshalString(routingNumberJsonTag, &wireDepositInformation.RoutingNumber)
	wireDepositInformation.BankCountry = BankCountry{}
	if err := data.UnmarshalStruct(bankCountryJsonTag, &wireDepositInformation.BankCountry); err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a Withdrawal model
func (withdrawal *Withdrawal) UnmarshalJSON(d []byte) error {
	const (
		idJsonTag       = "id"
		amountJsonTag   = "amount"
		currencyJsonTag = "currency"
		payoutAtJsonTag = "payout_at"
		feeJsonTag      = "fee"
		subtotalJsonTag = "subtotal"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloatString(amountJsonTag, &withdrawal.Amount)
	data.UnmarshalFloatString(feeJsonTag, &withdrawal.Fee)
	data.UnmarshalFloatString(subtotalJsonTag, &withdrawal.Subtotal)
	data.UnmarshalString(currencyJsonTag, &withdrawal.Currency)
	data.UnmarshalString(idJsonTag, &withdrawal.Id)
	data.UnmarshalString(payoutAtJsonTag, &withdrawal.PayoutAt)
	return nil
}

// UnmarshalJSON will deserialize bytes into a WithdrawalFeeEstimate model
func (withdrawalFeeEstimate *WithdrawalFeeEstimate) UnmarshalJSON(d []byte) error {
	const (
		feeJsonTag = "fee"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalFloat(feeJsonTag, &withdrawalFeeEstimate.Fee)
	return nil
}
