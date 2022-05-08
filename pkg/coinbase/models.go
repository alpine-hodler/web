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
	Amount       float64              `json:"amount" bson:"amount"`
	Balance      float64              `json:"balance" bson:"balance"`
	CreatedAt    time.Time            `json:"created_at" bson:"created_at"`
	Id           string               `json:"id" bson:"id"`
	ProtoDetails AccountLedgerDetails `json:"details" bson:"details"`
	Type         scalar.EntryType     `json:"type" bson:"type"`
}

// AccountLedgerDetails are the details for account history.
type AccountLedgerDetails struct {
	OrderId   string `json:"order_id" bson:"order_id"`
	ProductId string `json:"product_id" bson:"product_id"`
	TradeId   string `json:"trade_id" bson:"trade_id"`
}

// AccountTransfer will lists past withdrawals and deposits for an account.
type AccountTransfer struct {
	Amount       float64                `json:"amount" bson:"amount"`
	CanceledAt   time.Time              `json:"canceled_at" bson:"canceled_at"`
	CompletedAt  time.Time              `json:"completed_at" bson:"completed_at"`
	CreatedAt    time.Time              `json:"created_at" bson:"created_at"`
	Id           string                 `json:"id" bson:"id"`
	ProcessedAt  time.Time              `json:"processed_at" bson:"processed_at"`
	ProtoDetails AccountTransferDetails `json:"details" bson:"details"`
	Type         string                 `json:"type" bson:"type"`
	UserNonce    string                 `json:"user_nonce" bson:"user_nonce"`
}

// AccountTransferDetails are the details for an account transfer.
type AccountTransferDetails struct {
	CoinbaseAccountId       string `json:"coinbase_account_id" bson:"coinbase_account_id"`
	CoinbasePaymentMethodId string `json:"coinbase_payment_method_id" bson:"coinbase_payment_method_id"`
	CoinbaseTransactionId   string `json:"coinbase_transaction_id" bson:"coinbase_transaction_id"`
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

// CryptoAccount references a crypto account that a CoinbasePaymentMethod belongs to
type CryptoAccount struct {
	Id           string `json:"id" bson:"id"`
	Resource     string `json:"resource" bson:"resource"`
	ResourcePath string `json:"resource_path" bson:"resource_path"`
}

// CryptoAddress is used for a one-time crypto address for depositing crypto.
type CryptoAddress struct {
	Address          string                  `json:"address" bson:"address"`
	CallbackUrl      string                  `json:"callback_url" bson:"callback_url"`
	CreateAt         time.Time               `json:"create_at" bson:"create_at"`
	DepositUri       string                  `json:"deposit_uri" bson:"deposit_uri"`
	DestinationTag   string                  `json:"destination_tag" bson:"destination_tag"`
	Id               string                  `json:"id" bson:"id"`
	LegacyAddress    string                  `json:"legacy_address" bson:"legacy_address"`
	Name             string                  `json:"name" bson:"name"`
	Network          string                  `json:"network" bson:"network"`
	ProtoAddressInfo CryptoAddressInfo       `json:"address_info" bson:"address_info"`
	ProtoWarnings    []*CryptoAddressWarning `json:"warnings" bson:"warnings"`
	Resource         string                  `json:"resource" bson:"resource"`
	ResourcePath     string                  `json:"resource_path" bson:"resource_path"`
	UpdatedAt        time.Time               `json:"updated_at" bson:"updated_at"`
	UriScheme        string                  `json:"uri_scheme" bson:"uri_scheme"`
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
	Id            string          `json:"id" bson:"id"`
	MaxPrecision  float64         `json:"max_precision" bson:"max_precision"`
	Message       string          `json:"message" bson:"message"`
	MinSize       float64         `json:"min_size" bson:"min_size"`
	Name          string          `json:"name" bson:"name"`
	ProtoDetails  CurrencyDetails `json:"details" bson:"details"`
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
	AllowBuy              bool                `json:"allow_buy" bson:"allow_buy"`
	AllowDeposit          bool                `json:"allow_deposit" bson:"allow_deposit"`
	AllowSell             bool                `json:"allow_sell" bson:"allow_sell"`
	AllowWithdraw         bool                `json:"allow_withdraw" bson:"allow_withdraw"`
	CdvStatus             string              `json:"cdv_status" bson:"cdv_status"`
	CreateAt              time.Time           `json:"create_at" bson:"create_at"`
	Currency              string              `json:"currency" bson:"currency"`
	HoldBusinessDays      int                 `json:"hold_business_days" bson:"hold_business_days"`
	HoldDays              int                 `json:"hold_days" bson:"hold_days"`
	Id                    string              `json:"id" bson:"id"`
	InstantBuy            bool                `json:"instant_buy" bson:"instant_buy"`
	InstantSale           bool                `json:"instant_sale" bson:"instant_sale"`
	Name                  string              `json:"name" bson:"name"`
	PrimaryBuy            bool                `json:"primary_buy" bson:"primary_buy"`
	PrimarySell           bool                `json:"primary_sell" bson:"primary_sell"`
	ProtoAvailableBalance AvailableBalance    `json:"available_balance" bson:"available_balance"`
	ProtoCryptoAccount    CryptoAccount       `json:"crypto_account" bson:"crypto_account"`
	ProtoFiatAccount      FiatAccount         `json:"fiat_account" bson:"fiat_account"`
	ProtoLimits           Limits              `json:"limits" bson:"limits"`
	ProtoPickerData       PickerData          `json:"picker_data" bson:"picker_data"`
	ProtoRecurringOptions []*RecurringOptions `json:"recurring_options" bson:"recurring_options"`
	Resource              string              `json:"resource" bson:"resource"`
	ResourcePath          string              `json:"resource_path" bson:"resource_path"`
	Type                  string              `json:"type" bson:"type"`
	UpdatedAt             time.Time           `json:"updated_at" bson:"updated_at"`
	VerificationMethod    string              `json:"verification_method" bson:"verification_method"`
	Verified              bool                `json:"verified" bson:"verified"`
}

// PickerData ??
type PickerData struct {
	AccountName           string  `json:"account_name" bson:"account_name"`
	AccountNumber         string  `json:"account_number" bson:"account_number"`
	AccountType           string  `json:"account_type" bson:"account_type"`
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
	ProtoBalance          Balance `json:"balance" bson:"balance"`
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
	AccountAddress   string      `json:"account_address" bson:"account_address"`
	AccountName      string      `json:"account_name" bson:"account_name"`
	BankAddress      string      `json:"bank_address" bson:"bank_address"`
	BankName         string      `json:"bank_name" bson:"bank_name"`
	Iban             string      `json:"iban" bson:"iban"`
	ProtoBankCountry BankCountry `json:"bank_country" bson:"bank_country"`
	Reference        string      `json:"reference" bson:"reference"`
	Swift            string      `json:"swift" bson:"swift"`
}

// SwiftDepositInformation information regarding a wallet's deposits. SWIFT stands for Society for Worldwide Interbank
// Financial Telecommunications. Basically, it's a computer network that connects over 900 banks around the world – and
// enables them to transfer money. ING is part of this network. There is no fee for accepting deposits into your account
// with ING.
type SwiftDepositInformation struct {
	AccountAddress   string      `json:"account_address" bson:"account_address"`
	AccountName      string      `json:"account_name" bson:"account_name"`
	AccountNumber    string      `json:"account_number" bson:"account_number"`
	BankAddress      string      `json:"bank_address" bson:"bank_address"`
	BankName         string      `json:"bank_name" bson:"bank_name"`
	ProtoBankCountry BankCountry `json:"bank_country" bson:"bank_country"`
	Reference        string      `json:"reference" bson:"reference"`
}

// UkDepositInformation information regarding a wallet's deposits.
type UkDepositInformation struct {
	AccountAddress   string      `json:"account_address" bson:"account_address"`
	AccountName      string      `json:"account_name" bson:"account_name"`
	AccountNumber    string      `json:"account_number" bson:"account_number"`
	BankAddress      string      `json:"bank_address" bson:"bank_address"`
	BankName         string      `json:"bank_name" bson:"bank_name"`
	ProtoBankCountry BankCountry `json:"bank_country" bson:"bank_country"`
	Reference        string      `json:"reference" bson:"reference"`
}

// Wallet represents a user's available Coinbase wallet (These are the wallets/accounts that are used for buying and
// selling on www.coinbase.com)
type Wallet struct {
	Active                       bool                    `json:"active" bson:"active"`
	AvailableOnConsumer          bool                    `json:"available_on_consumer" bson:"available_on_consumer"`
	Balance                      float64                 `json:"balance" bson:"balance"`
	Currency                     string                  `json:"currency" bson:"currency"`
	DestinationTagName           string                  `json:"destination_tag_name" bson:"destination_tag_name"`
	DestinationTagRegex          string                  `json:"destination_tag_regex" bson:"destination_tag_regex"`
	HoldBalance                  float64                 `json:"hold_balance" bson:"hold_balance"`
	HoldCurrency                 string                  `json:"hold_currency" bson:"hold_currency"`
	Id                           string                  `json:"id" bson:"id"`
	Name                         string                  `json:"name" bson:"name"`
	Primary                      bool                    `json:"primary" bson:"primary"`
	ProtoSepaDepositInformation  SepaDepositInformation  `json:"sepa_deposit_information" bson:"sepa_deposit_information"`
	ProtoSwiftDepositInformation SwiftDepositInformation `json:"swift_deposit_information" bson:"swift_deposit_information"`
	ProtoUkDepositInformation    UkDepositInformation    `json:"uk_deposit_information" bson:"uk_deposit_information"`
	ProtoWireDepositInformation  WireDepositInformation  `json:"wire_deposit_information" bson:"wire_deposit_information"`
	Ready                        bool                    `json:"ready" bson:"ready"`
	Type                         string                  `json:"type" bson:"type"`
}

// WebsocketTicker is real-time price updates every time a match happens. It batches updates in case of cascading
// matches, greatly reducing bandwidth requirements.
type WebsocketTicker struct {
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

// WireDepositInformation information regarding a wallet's deposits
type WireDepositInformation struct {
	AccountAddress   string      `json:"account_address" bson:"account_address"`
	AccountName      string      `json:"account_name" bson:"account_name"`
	AccountNumber    string      `json:"account_number" bson:"account_number"`
	BankAddress      string      `json:"bank_address" bson:"bank_address"`
	BankName         string      `json:"bank_name" bson:"bank_name"`
	ProtoBankCountry BankCountry `json:"bank_country" bson:"bank_country"`
	Reference        string      `json:"reference" bson:"reference"`
	RoutingNumber    string      `json:"routing_number" bson:"routing_number"`
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
	data.UnmarshalFloatFromString(availableJsonTag, &account.Available)
	data.UnmarshalFloatFromString(balanceJsonTag, &account.Balance)
	data.UnmarshalFloatFromString(holdJsonTag, &account.Hold)
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
	accountLedger.ProtoDetails = AccountLedgerDetails{}
	if err := data.UnmarshalStruct(detailsJsonTag, &accountLedger.ProtoDetails); err != nil {
		return err
	}
	data.UnmarshalEntryType(typeJsonTag, &accountLedger.Type)
	data.UnmarshalFloatFromString(amountJsonTag, &accountLedger.Amount)
	data.UnmarshalFloatFromString(balanceJsonTag, &accountLedger.Balance)
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
	accountTransfer.ProtoDetails = AccountTransferDetails{}
	if err := data.UnmarshalStruct(detailsJsonTag, &accountTransfer.ProtoDetails); err != nil {
		return err
	}
	data.UnmarshalFloatFromString(amountJsonTag, &accountTransfer.Amount)
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
	data.UnmarshalFloatFromString(amountJsonTag, &availableBalance.Amount)
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
	data.UnmarshalFloatFromString(amountJsonTag, &balance.Amount)
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
	cryptoAddress.ProtoAddressInfo = CryptoAddressInfo{}
	if err := data.UnmarshalStruct(addressInfoJsonTag, &cryptoAddress.ProtoAddressInfo); err != nil {
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
			cryptoAddress.ProtoWarnings = append(cryptoAddress.ProtoWarnings, &obj)
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
	currency.ProtoDetails = CurrencyDetails{}
	if err := data.UnmarshalStruct(detailsJsonTag, &currency.ProtoDetails); err != nil {
		return err
	}
	data.UnmarshalFloatFromString(maxPrecisionJsonTag, &currency.MaxPrecision)
	data.UnmarshalFloatFromString(minSizeJsonTag, &currency.MinSize)
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
	data.UnmarshalFloatFromString(amountJsonTag, &currencyConversion.Amount)
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
	data.UnmarshalFloatFromString(amountJsonTag, &deposit.Amount)
	data.UnmarshalFloatFromString(feeJsonTag, &deposit.Fee)
	data.UnmarshalFloatFromString(subtotalJsonTag, &deposit.Subtotal)
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
	data.UnmarshalFloatFromString(makerFeeRateJsonTag, &fees.MakerFeeRate)
	data.UnmarshalFloatFromString(takerFeeRateJsonTag, &fees.TakerFeeRate)
	data.UnmarshalFloatFromString(usdVolumeJsonTag, &fees.UsdVolume)
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
	data.UnmarshalFloatFromString(feeJsonTag, &fill.Fee)
	data.UnmarshalFloatFromString(priceJsonTag, &fill.Price)
	data.UnmarshalFloatFromString(sizeJsonTag, &fill.Size)
	data.UnmarshalFloatFromString(usdVolumeJsonTag, &fill.UsdVolume)
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
	data.UnmarshalFloatFromString(fillFeesJsonTag, &newOrder.FillFees)
	data.UnmarshalFloatFromString(filledSizeJsonTag, &newOrder.FilledSize)
	data.UnmarshalFloatFromString(fundingAmountJsonTag, &newOrder.FundingAmount)
	data.UnmarshalFloatFromString(fundsJsonTag, &newOrder.Funds)
	data.UnmarshalFloatFromString(priceJsonTag, &newOrder.Price)
	data.UnmarshalFloatFromString(sizeJsonTag, &newOrder.Size)
	data.UnmarshalFloatFromString(specificFundsJsonTag, &newOrder.SpecificFunds)
	data.UnmarshalFloatFromString(stopPriceJsonTag, &newOrder.StopPrice)
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
	data.UnmarshalFloatFromString(executedValueJsonTag, &order.ExecutedValue)
	data.UnmarshalFloatFromString(fillFeesJsonTag, &order.FillFees)
	data.UnmarshalFloatFromString(filledSizeJsonTag, &order.FilledSize)
	data.UnmarshalFloatFromString(fundingAmountJsonTag, &order.FundingAmount)
	data.UnmarshalFloatFromString(fundsJsonTag, &order.Funds)
	data.UnmarshalFloatFromString(priceJsonTag, &order.Price)
	data.UnmarshalFloatFromString(sizeJsonTag, &order.Size)
	data.UnmarshalFloatFromString(specifiedFundsJsonTag, &order.SpecifiedFunds)
	data.UnmarshalFloatFromString(stopPriceJsonTag, &order.StopPrice)
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
			paymentMethod.ProtoRecurringOptions = append(paymentMethod.ProtoRecurringOptions, &obj)
		}
	}
	paymentMethod.ProtoAvailableBalance = AvailableBalance{}
	if err := data.UnmarshalStruct(availableBalanceJsonTag, &paymentMethod.ProtoAvailableBalance); err != nil {
		return err
	}
	paymentMethod.ProtoCryptoAccount = CryptoAccount{}
	if err := data.UnmarshalStruct(cryptoAccountJsonTag, &paymentMethod.ProtoCryptoAccount); err != nil {
		return err
	}
	paymentMethod.ProtoFiatAccount = FiatAccount{}
	if err := data.UnmarshalStruct(fiatAccountJsonTag, &paymentMethod.ProtoFiatAccount); err != nil {
		return err
	}
	paymentMethod.ProtoLimits = Limits{}
	if err := data.UnmarshalStruct(limitsJsonTag, &paymentMethod.ProtoLimits); err != nil {
		return err
	}
	paymentMethod.ProtoPickerData = PickerData{}
	if err := data.UnmarshalStruct(pickerDataJsonTag, &paymentMethod.ProtoPickerData); err != nil {
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
	pickerData.ProtoBalance = Balance{}
	if err := data.UnmarshalStruct(balanceJsonTag, &pickerData.ProtoBalance); err != nil {
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
	data.UnmarshalFloatFromString(baseIncrementJsonTag, &product.BaseIncrement)
	data.UnmarshalFloatFromString(baseMaxSizeJsonTag, &product.BaseMaxSize)
	data.UnmarshalFloatFromString(baseMinSizeJsonTag, &product.BaseMinSize)
	data.UnmarshalFloatFromString(maxMarketFundsJsonTag, &product.MaxMarketFunds)
	data.UnmarshalFloatFromString(maxSlippagePercentageJsonTag, &product.MaxSlippagePercentage)
	data.UnmarshalFloatFromString(minMarketFundsJsonTag, &product.MinMarketFunds)
	data.UnmarshalFloatFromString(quoteIncrementJsonTag, &product.QuoteIncrement)
	data.UnmarshalStatus(statusJsonTag, &product.Status)
	data.UnmarshalString(baseCurrencyJsonTag, &product.BaseCurrency)
	data.UnmarshalString(displayNameJsonTag, &product.DisplayName)
	data.UnmarshalString(idJsonTag, &product.Id)
	data.UnmarshalString(quoteCurrencyJsonTag, &product.QuoteCurrency)
	data.UnmarshalString(statusMessageJsonTag, &product.StatusMessage)
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
	data.UnmarshalFloatFromString(askJsonTag, &productTicker.Ask)
	data.UnmarshalFloatFromString(bidJsonTag, &productTicker.Bid)
	data.UnmarshalFloatFromString(priceJsonTag, &productTicker.Price)
	data.UnmarshalFloatFromString(sizeJsonTag, &productTicker.Size)
	data.UnmarshalFloatFromString(volumeJsonTag, &productTicker.Volume)
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
	sepaDepositInformation.ProtoBankCountry = BankCountry{}
	if err := data.UnmarshalStruct(bankCountryJsonTag, &sepaDepositInformation.ProtoBankCountry); err != nil {
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
	swiftDepositInformation.ProtoBankCountry = BankCountry{}
	if err := data.UnmarshalStruct(bankCountryJsonTag, &swiftDepositInformation.ProtoBankCountry); err != nil {
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
	ukDepositInformation.ProtoBankCountry = BankCountry{}
	if err := data.UnmarshalStruct(bankCountryJsonTag, &ukDepositInformation.ProtoBankCountry); err != nil {
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
	data.UnmarshalFloatFromString(balanceJsonTag, &wallet.Balance)
	data.UnmarshalFloatFromString(holdBalanceJsonTag, &wallet.HoldBalance)
	data.UnmarshalString(currencyJsonTag, &wallet.Currency)
	data.UnmarshalString(destinationTagNameJsonTag, &wallet.DestinationTagName)
	data.UnmarshalString(destinationTagRegexJsonTag, &wallet.DestinationTagRegex)
	data.UnmarshalString(holdCurrencyJsonTag, &wallet.HoldCurrency)
	data.UnmarshalString(idJsonTag, &wallet.Id)
	data.UnmarshalString(nameJsonTag, &wallet.Name)
	data.UnmarshalString(typeJsonTag, &wallet.Type)
	wallet.ProtoSepaDepositInformation = SepaDepositInformation{}
	if err := data.UnmarshalStruct(sepaDepositInformationJsonTag, &wallet.ProtoSepaDepositInformation); err != nil {
		return err
	}
	wallet.ProtoSwiftDepositInformation = SwiftDepositInformation{}
	if err := data.UnmarshalStruct(swiftDepositInformationJsonTag, &wallet.ProtoSwiftDepositInformation); err != nil {
		return err
	}
	wallet.ProtoUkDepositInformation = UkDepositInformation{}
	if err := data.UnmarshalStruct(ukDepositInformationJsonTag, &wallet.ProtoUkDepositInformation); err != nil {
		return err
	}
	wallet.ProtoWireDepositInformation = WireDepositInformation{}
	if err := data.UnmarshalStruct(wireDepositInformationJsonTag, &wallet.ProtoWireDepositInformation); err != nil {
		return err
	}
	return nil
}

// UnmarshalJSON will deserialize bytes into a WebsocketTicker model
func (websocketTicker *WebsocketTicker) UnmarshalJSON(d []byte) error {
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
	data.UnmarshalFloatFromString(bestAskJsonTag, &websocketTicker.BestAsk)
	data.UnmarshalFloatFromString(bestBidJsonTag, &websocketTicker.BestBid)
	data.UnmarshalFloatFromString(lastSizeJsonTag, &websocketTicker.LastSize)
	data.UnmarshalFloatFromString(priceJsonTag, &websocketTicker.Price)
	data.UnmarshalInt(sequenceJsonTag, &websocketTicker.Sequence)
	data.UnmarshalInt(tradeIdJsonTag, &websocketTicker.TradeId)
	data.UnmarshalString(productIdJsonTag, &websocketTicker.ProductId)
	data.UnmarshalString(sideJsonTag, &websocketTicker.Side)
	data.UnmarshalString(typeJsonTag, &websocketTicker.Type)
	err = data.UnmarshalTime(time.RFC3339Nano, timeJsonTag, &websocketTicker.Time)
	if err != nil {
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
	wireDepositInformation.ProtoBankCountry = BankCountry{}
	if err := data.UnmarshalStruct(bankCountryJsonTag, &wireDepositInformation.ProtoBankCountry); err != nil {
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
	data.UnmarshalFloatFromString(amountJsonTag, &withdrawal.Amount)
	data.UnmarshalFloatFromString(feeJsonTag, &withdrawal.Fee)
	data.UnmarshalFloatFromString(subtotalJsonTag, &withdrawal.Subtotal)
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
