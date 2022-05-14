package coinbase

import (
	"path"
	"strings"

	"github.com/alpine-hodler/sdk/tools"
)

// * This is a generated file, do not edit
type postAuthority uint8

const (
	_ postAuthority = iota
	accountHoldsPostAuthority
	accountLedgerPostAuthority
	accountPostAuthority
	accountTransferPostAuthority
	accountTransfersPostAuthority
	accountWithdrawalPostAuthority
	accountsPostAuthority
	bookPostAuthority
	cancelOpenOrdersPostAuthority
	cancelOrderPostAuthority
	candlesPostAuthority
	coinbaseAccountDepositPostAuthority
	convertCurrencyPostAuthority
	createProfilePostAuthority
	createProfileTransferPostAuthority
	cryptoWithdrawalPostAuthority
	currenciesPostAuthority
	currencyConversionPostAuthority
	currencyPostAuthority
	feesPostAuthority
	fillsPostAuthority
	generateCryptoAddressPostAuthority
	newOrderPostAuthority
	orderPostAuthority
	ordersPostAuthority
	paymentMethodDepositPostAuthority
	paymentMethodWithdrawalPostAuthority
	paymentMethodsPostAuthority
	productPostAuthority
	productStatsPostAuthority
	productTickerPostAuthority
	productsPostAuthority
	profilePostAuthority
	profilesPostAuthority
	renameProfilePostAuthority
	signedPricesPostAuthority
	tradesPostAuthority
	transfersPostAuthority
	walletsPostAuthority
	withdrawalFeeEstimatePostAuthority
)

// Get a list of trading accounts from the profile of the API key.
func getAccountsPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/accounts")
}

// Information for a single account. Use this endpoint when you know the account_id. API key must belong to the same
// profile as the account.
func getAccountPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/accounts", builder.Get(tools.URIBuilderComponentPath, "account_id"))
}

// AccountHolds will return the holds of an account that belong to the same profile as the API key. Holds are placed on
// an account for any active orders or pending withdraw requests. As an order is filled, the hold amount is updated. If
// an order is canceled, any remaining hold is removed. For withdrawals, the hold is removed after it is completed.
func getAccountHoldsPostAuthority(builder tools.URIBuilder) (p string) {
	p = path.Join("/accounts", builder.Get(tools.URIBuilderComponentPath, "account_id"), "holds")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(builder.QueryPath().String())
	return sb.String()
}

// Lists ledger activity for an account. This includes anything that would affect the accounts balance - transfers,
// trades, fees, etc. This endpoint requires either the "view" or "trade" permission.
func getAccountLedgerPostAuthority(builder tools.URIBuilder) (p string) {
	p = path.Join("/accounts", builder.Get(tools.URIBuilderComponentPath, "account_id"), "ledger")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(builder.QueryPath().String())
	return sb.String()
}

// AccountTransfers returns past withdrawals and deposits for an account.
func getAccountTransfersPostAuthority(builder tools.URIBuilder) (p string) {
	p = path.Join("/accounts", builder.Get(tools.URIBuilderComponentPath, "account_id"), "transfers")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(builder.QueryPath().String())
	return sb.String()
}

// Transfers is a list of in-progress and completed transfers of funds in/out of any of the user's accounts.
func getTransfersPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/transfers")
}

// AccountTransfer returns information on a single transfer.
func getAccountTransferPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/transfers", builder.Get(tools.URIBuilderComponentPath, "transfer_id"))
}

// Get a list of open orders for a product. The amount of detail shown can be customized with the level parameter.
func getBookPostAuthority(builder tools.URIBuilder) (p string) {
	p = path.Join("/products", builder.Get(tools.URIBuilderComponentPath, "product_id"), "book")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(builder.QueryPath().String())
	return sb.String()
}

// Get the historic rates for a product
func getCandlesPostAuthority(builder tools.URIBuilder) (p string) {
	p = path.Join("/products", builder.Get(tools.URIBuilderComponentPath, "product_id"), "candles")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(builder.QueryPath().String())
	return sb.String()
}

// GenerateCryptoAddress will create a one-time crypto address for depositing crypto, using a wallet account id. This
// endpoint requires the "transfer" permission. API key must belong to default profile.
func getGenerateCryptoAddressPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/coinbase-accounts", builder.Get(tools.URIBuilderComponentPath, "account_id"), "addresses")
}

// Gets a list of all known currencies. Note: Not all currencies may be currently in use for trading.
func getCurrenciesPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/currencies")
}

// Gets a single currency by id.
func getCurrencyPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/currencies", builder.Get(tools.URIBuilderComponentPath, "currency_id"))
}

// ConvertCurrency converts funds from from currency to to currency. Funds are converted on the from account in the
// profile_id profile. This endpoint requires the "trade" permission. A successful conversion will be assigned a
// conversion id. The corresponding ledger entries for a conversion will reference this conversion id
func getConvertCurrencyPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/conversions")
}

// CurrencyConversion returns the currency conversion by conversion id (i.e. USD -> USDC).
func getCurrencyConversionPostAuthority(builder tools.URIBuilder) (p string) {
	p = path.Join("/conversions", builder.Get(tools.URIBuilderComponentPath, "conversion_id"))
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(builder.QueryPath().String())
	return sb.String()
}

// Deposits funds from a www.coinbase.com wallet to the specified profile_id.
func getCoinbaseAccountDepositPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/deposits", "coinbase-account")
}

// Deposits funds from a linked external payment method to the specified profile_id.
func getPaymentMethodDepositPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/deposits", "payment-method")
}

// Get fees rates and 30 days trailing volume.
func getFeesPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/fees")
}

// Get a list of fills. A fill is a partial or complete match on a specific order.
func getFillsPostAuthority(builder tools.URIBuilder) (p string) {
	p = path.Join("/fills")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(builder.QueryPath().String())
	return sb.String()
}

// Create an order. You can place two types of orders: limit and market. Orders can only be placed if your account has
// sufficient funds. Once an order is placed, your account funds will be put on hold for the duration of the order. How
// much and which funds are put on hold depends on the order type and parameters specified.
func getNewOrderPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/orders")
}

// SignedPrices returns cryptographically signed prices ready to be posted on-chain using Compound's Open Oracle smart
// contract.
func getSignedPricesPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/oracle")
}

// List your current open orders. Only open or un-settled orders are returned by default. As soon as an order is no
// longer open and settled, it will no longer appear in the default request. Open orders may change state between the
// request and the response depending on market conditions.
func getOrdersPostAuthority(builder tools.URIBuilder) (p string) {
	p = path.Join("/orders")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(builder.QueryPath().String())
	return sb.String()
}

// Get a single order by id.
func getOrderPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/orders", builder.Get(tools.URIBuilderComponentPath, "order_id"))
}

// CancelOrder will cancel a single open order by order id.
func getCancelOrderPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/orders", builder.Get(tools.URIBuilderComponentPath, "order_id"))
}

// CancelOpenOrders will try with best effort to cancel all open orders. This may require you to make the request
// multiple times until all of the open orders are deleted.
func getCancelOpenOrdersPostAuthority(builder tools.URIBuilder) (p string) {
	p = path.Join("/orders")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(builder.QueryPath().String())
	return sb.String()
}

// PaymentMethods returns a list of the user's linked payment methods.
func getPaymentMethodsPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/payment-methods")
}

// Gets a list of available currency pairs for trading.
func getProductsPostAuthority(builder tools.URIBuilder) (p string) {
	p = path.Join("/products")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(builder.QueryPath().String())
	return sb.String()
}

// Get information on a single product.
func getProductPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/products", builder.Get(tools.URIBuilderComponentPath, "product_id"))
}

// Gets 30day and 24hour stats for a product.
func getProductStatsPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/products", builder.Get(tools.URIBuilderComponentPath, "product_id"), "stats")
}

// Gets snapshot information about the last trade (tick), best bid/ask and 24h volume.
func getProductTickerPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/products", builder.Get(tools.URIBuilderComponentPath, "product_id"), "ticker")
}

// Profile returns information for a single profile. Use this endpoint when you know the profile_id. This endpoint
// requires the "view" permission and is accessible by any profile's API key.
func getProfilePostAuthority(builder tools.URIBuilder) (p string) {
	p = path.Join("/profiles", builder.Get(tools.URIBuilderComponentPath, "profile_id"))
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(builder.QueryPath().String())
	return sb.String()
}

// RenameProfile will rename a profile. Names 'default' and 'margin' are reserved.
func getRenameProfilePostAuthority(builder tools.URIBuilder) (p string) {
	p = path.Join("/profiles", builder.Get(tools.URIBuilderComponentPath, "profile_id"))
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(builder.QueryPath().String())
	return sb.String()
}

// Gets a list of all of the current user's profiles.
func getProfilesPostAuthority(builder tools.URIBuilder) (p string) {
	p = path.Join("/profiles")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(builder.QueryPath().String())
	return sb.String()
}

// CreateProfile will create a new profile. Will fail if no name is provided or if user already has max number of
// profiles.
func getCreateProfilePostAuthority(builder tools.URIBuilder) string {
	return path.Join("/profiles")
}

// CreateProfileTransfer will transfer an amount of currency from one profile to another. This endpoint requires the
// "transfer" permission.
func getCreateProfileTransferPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/profiles", "transfer")
}

// Gets a list the latest trades for a product.
func getTradesPostAuthority(builder tools.URIBuilder) (p string) {
	p = path.Join("/products", builder.Get(tools.URIBuilderComponentPath, "product_id"), "trades")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(builder.QueryPath().String())
	return sb.String()
}

// Gets all the user's available Coinbase wallets (These are the wallets/accounts that are used for buying and selling
// on www.coinbase.com)
func getWalletsPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/coinbase-accounts")
}

// Withdraws funds from the specified profile_id to a www.coinbase.com wallet. Withdraw funds to a coinbase account. You
// can move funds between your Coinbase accounts and your Coinbase Exchange trading accounts within your daily limits.
// Moving funds between Coinbase and Coinbase Exchange is instant and free. See the Coinbase Accounts section for
// retrieving your Coinbase accounts. This endpoint requires the "transfer" permission.
func getAccountWithdrawalPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/withdrawals", "coinbase-account")
}

// Withdraws funds from the specified profile_id to an external crypto address. This endpoint requires the "transfer"
// permission. API key must belong to default profile.
func getCryptoWithdrawalPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/withdrawals", "crypto")
}

// Withdraws funds from the specified profile_id to a linked external payment method. This endpoint requires the
// "transfer" permission. API key is restricted to the default profile.
func getPaymentMethodWithdrawalPostAuthority(builder tools.URIBuilder) string {
	return path.Join("/withdrawals", "payment-method")
}

// Gets the fee estimate for the crypto withdrawal to crypto address
func getWithdrawalFeeEstimatePostAuthority(builder tools.URIBuilder) (p string) {
	p = path.Join("/withdrawals", "fee-estimate")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(builder.QueryPath().String())
	return sb.String()
}

// Get takes an postAuthority const and postAuthority arguments to parse the URL postAuthority path.
func (pa postAuthority) PostAuthority(builder tools.URIBuilder) string {
	return map[postAuthority]func(builder tools.URIBuilder) string{
		accountsPostAuthority:                getAccountsPostAuthority,
		accountPostAuthority:                 getAccountPostAuthority,
		accountHoldsPostAuthority:            getAccountHoldsPostAuthority,
		accountLedgerPostAuthority:           getAccountLedgerPostAuthority,
		accountTransfersPostAuthority:        getAccountTransfersPostAuthority,
		transfersPostAuthority:               getTransfersPostAuthority,
		accountTransferPostAuthority:         getAccountTransferPostAuthority,
		bookPostAuthority:                    getBookPostAuthority,
		candlesPostAuthority:                 getCandlesPostAuthority,
		generateCryptoAddressPostAuthority:   getGenerateCryptoAddressPostAuthority,
		currenciesPostAuthority:              getCurrenciesPostAuthority,
		currencyPostAuthority:                getCurrencyPostAuthority,
		convertCurrencyPostAuthority:         getConvertCurrencyPostAuthority,
		currencyConversionPostAuthority:      getCurrencyConversionPostAuthority,
		coinbaseAccountDepositPostAuthority:  getCoinbaseAccountDepositPostAuthority,
		paymentMethodDepositPostAuthority:    getPaymentMethodDepositPostAuthority,
		feesPostAuthority:                    getFeesPostAuthority,
		fillsPostAuthority:                   getFillsPostAuthority,
		newOrderPostAuthority:                getNewOrderPostAuthority,
		signedPricesPostAuthority:            getSignedPricesPostAuthority,
		ordersPostAuthority:                  getOrdersPostAuthority,
		orderPostAuthority:                   getOrderPostAuthority,
		cancelOrderPostAuthority:             getCancelOrderPostAuthority,
		cancelOpenOrdersPostAuthority:        getCancelOpenOrdersPostAuthority,
		paymentMethodsPostAuthority:          getPaymentMethodsPostAuthority,
		productsPostAuthority:                getProductsPostAuthority,
		productPostAuthority:                 getProductPostAuthority,
		productStatsPostAuthority:            getProductStatsPostAuthority,
		productTickerPostAuthority:           getProductTickerPostAuthority,
		profilePostAuthority:                 getProfilePostAuthority,
		renameProfilePostAuthority:           getRenameProfilePostAuthority,
		profilesPostAuthority:                getProfilesPostAuthority,
		createProfilePostAuthority:           getCreateProfilePostAuthority,
		createProfileTransferPostAuthority:   getCreateProfileTransferPostAuthority,
		tradesPostAuthority:                  getTradesPostAuthority,
		walletsPostAuthority:                 getWalletsPostAuthority,
		accountWithdrawalPostAuthority:       getAccountWithdrawalPostAuthority,
		cryptoWithdrawalPostAuthority:        getCryptoWithdrawalPostAuthority,
		paymentMethodWithdrawalPostAuthority: getPaymentMethodWithdrawalPostAuthority,
		withdrawalFeeEstimatePostAuthority:   getWithdrawalFeeEstimatePostAuthority,
	}[pa](builder)
}
