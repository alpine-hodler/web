package coinbase

import (
	"path"
	"strings"

	"github.com/alpine-hodler/sdk/internal/client"
)

// * This is a generated file, do not edit
type endpoint uint8

const (
	_ endpoint = iota
	accountEndpoint
	accountHoldsEndpoint
	accountLedgerEndpoint
	accountTransferEndpoint
	accountTransfersEndpoint
	accountWithdrawalEndpoint
	accountsEndpoint
	bookEndpoint
	cancelOpenOrdersEndpoint
	cancelOrderEndpoint
	candlesEndpoint
	coinbaseAccountDepositEndpoint
	convertCurrencyEndpoint
	createProfileEndpoint
	createProfileTransferEndpoint
	cryptoWithdrawalEndpoint
	currenciesEndpoint
	currencyConversionEndpoint
	currencyEndpoint
	feesEndpoint
	fillsEndpoint
	generateCryptoAddressEndpoint
	newOrderEndpoint
	orderEndpoint
	ordersEndpoint
	paymentMethodDepositEndpoint
	paymentMethodWithdrawalEndpoint
	paymentMethodsEndpoint
	productEndpoint
	productStatsEndpoint
	productTickerEndpoint
	productsEndpoint
	profilesEndpoint
	signedPricesEndpoint
	tradesEndpoint
	transfersEndpoint
	walletsEndpoint
	withdrawalFeeEstimateEndpoint
)

// Get a list of trading accounts from the profile of the API key.
func accountsPath(args client.EndpointArgs) string {
	return path.Join("/accounts")
}

// Information for a single account. Use this endpoint when you know the account_id. API key must belong to the same
// profile as the account.
func accountPath(args client.EndpointArgs) string {
	return path.Join("/accounts", *args["account_id"].PathParam)
}

// AccountHolds will return the holds of an account that belong to the same profile as the API key. Holds are placed on
// an account for any active orders or pending withdraw requests. As an order is filled, the hold amount is updated. If
// an order is canceled, any remaining hold is removed. For withdrawals, the hold is removed after it is completed.
func accountHoldsPath(args client.EndpointArgs) (p string) {
	p = path.Join("/accounts", *args["account_id"].PathParam, "holds")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// Lists ledger activity for an account. This includes anything that would affect the accounts balance - transfers,
// trades, fees, etc. This endpoint requires either the "view" or "trade" permission.
func accountLedgerPath(args client.EndpointArgs) (p string) {
	p = path.Join("/accounts", *args["account_id"].PathParam, "ledger")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// AccountTransfers returns past withdrawals and deposits for an account.
func accountTransfersPath(args client.EndpointArgs) (p string) {
	p = path.Join("/accounts", *args["account_id"].PathParam, "transfers")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// Transfers is a list of in-progress and completed transfers of funds in/out of any of the user's accounts.
func transfersPath(args client.EndpointArgs) string {
	return path.Join("/transfers")
}

// AccountTransfer returns information on a single transfer.
func accountTransferPath(args client.EndpointArgs) string {
	return path.Join("/transfers", *args["transfer_id"].PathParam)
}

// Get a list of open orders for a product. The amount of detail shown can be customized with the level parameter.
func bookPath(args client.EndpointArgs) (p string) {
	p = path.Join("/products", *args["product_id"].PathParam, "book")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// Get the historic rates for a product
func candlesPath(args client.EndpointArgs) (p string) {
	p = path.Join("/products", *args["product_id"].PathParam, "candles")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// GenerateCryptoAddress will create a one-time crypto address for depositing crypto, using a wallet account id. This
// endpoint requires the "transfer" permission. API key must belong to default profile.
func generateCryptoAddressPath(args client.EndpointArgs) string {
	return path.Join("/coinbase-accounts", *args["account_id"].PathParam, "addresses")
}

// Gets a list of all known currencies. Note: Not all currencies may be currently in use for trading.
func currenciesPath(args client.EndpointArgs) string {
	return path.Join("/currencies")
}

// Gets a single currency by id.
func currencyPath(args client.EndpointArgs) string {
	return path.Join("/currencies", *args["currency_id"].PathParam)
}

// ConvertCurrency converts funds from from currency to to currency. Funds are converted on the from account in the
// profile_id profile. This endpoint requires the "trade" permission. A successful conversion will be assigned a
// conversion id. The corresponding ledger entries for a conversion will reference this conversion id
func convertCurrencyPath(args client.EndpointArgs) string {
	return path.Join("/conversions")
}

// CurrencyConversion returns the currency conversion by conversion id (i.e. USD -> USDC).
func currencyConversionPath(args client.EndpointArgs) (p string) {
	p = path.Join("/conversions", *args["conversion_id"].PathParam)
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// Deposits funds from a www.coinbase.com wallet to the specified profile_id.
func coinbaseAccountDepositPath(args client.EndpointArgs) string {
	return path.Join("/deposits", "coinbase-account")
}

// Deposits funds from a linked external payment method to the specified profile_id.
func paymentMethodDepositPath(args client.EndpointArgs) string {
	return path.Join("/deposits", "payment-method")
}

// Get fees rates and 30 days trailing volume.
func feesPath(args client.EndpointArgs) string {
	return path.Join("/fees")
}

// Get a list of fills. A fill is a partial or complete match on a specific order.
func fillsPath(args client.EndpointArgs) (p string) {
	p = path.Join("/fills")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// Create an order. You can place two types of orders: limit and market. Orders can only be placed if your account has
// sufficient funds. Once an order is placed, your account funds will be put on hold for the duration of the order. How
// much and which funds are put on hold depends on the order type and parameters specified.
func newOrderPath(args client.EndpointArgs) string {
	return path.Join("/orders")
}

// SignedPrices returns cryptographically signed prices ready to be posted on-chain using Compound's Open Oracle smart
// contract.
func signedPricesPath(args client.EndpointArgs) string {
	return path.Join("/oracle")
}

// List your current open orders. Only open or un-settled orders are returned by default. As soon as an order is no
// longer open and settled, it will no longer appear in the default request. Open orders may change state between the
// request and the response depending on market conditions.
func ordersPath(args client.EndpointArgs) (p string) {
	p = path.Join("/orders")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// Get a single order by id.
func orderPath(args client.EndpointArgs) string {
	return path.Join("/orders", *args["order_id"].PathParam)
}

// CancelOrder will cancel a single open order by order id.
func cancelOrderPath(args client.EndpointArgs) string {
	return path.Join("/orders", *args["order_id"].PathParam)
}

// CancelOpenOrders will try with best effort to cancel all open orders. This may require you to make the request
// multiple times until all of the open orders are deleted.
func cancelOpenOrdersPath(args client.EndpointArgs) (p string) {
	p = path.Join("/orders")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// PaymentMethods returns a list of the user's linked payment methods.
func paymentMethodsPath(args client.EndpointArgs) string {
	return path.Join("/payment-methods")
}

// Gets a list of available currency pairs for trading.
func productsPath(args client.EndpointArgs) (p string) {
	p = path.Join("/products")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// Get information on a single product.
func productPath(args client.EndpointArgs) string {
	return path.Join("/products", *args["product_id"].PathParam)
}

// Gets 30day and 24hour stats for a product.
func productStatsPath(args client.EndpointArgs) string {
	return path.Join("/products", *args["product_id"].PathParam, "stats")
}

// Gets snapshot information about the last trade (tick), best bid/ask and 24h volume.
func productTickerPath(args client.EndpointArgs) string {
	return path.Join("/products", *args["product_id"].PathParam, "ticker")
}

// Gets a list of all of the current user's profiles.
func profilesPath(args client.EndpointArgs) (p string) {
	p = path.Join("/profiles")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// CreateProfile will create a new profile. Will fail if no name is provided or if user already has max number of
// profiles.
func createProfilePath(args client.EndpointArgs) string {
	return path.Join("/profiles")
}

// CreateProfileTransfer will transfer an amount of currency from one profile to another. This endpoint requires the
// "transfer" permission.
func createProfileTransferPath(args client.EndpointArgs) string {
	return path.Join("/profiles", "transfer")
}

// Gets a list the latest trades for a product.
func tradesPath(args client.EndpointArgs) (p string) {
	p = path.Join("/products", *args["product_id"].PathParam, "trades")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// Gets all the user's available Coinbase wallets (These are the wallets/accounts that are used for buying and selling
// on www.coinbase.com)
func walletsPath(args client.EndpointArgs) string {
	return path.Join("/coinbase-accounts")
}

// Withdraws funds from the specified profile_id to a www.coinbase.com wallet. Withdraw funds to a coinbase account. You
// can move funds between your Coinbase accounts and your Coinbase Exchange trading accounts within your daily limits.
// Moving funds between Coinbase and Coinbase Exchange is instant and free. See the Coinbase Accounts section for
// retrieving your Coinbase accounts. This endpoint requires the "transfer" permission.
func accountWithdrawalPath(args client.EndpointArgs) string {
	return path.Join("/withdrawals", "coinbase-account")
}

// Withdraws funds from the specified profile_id to an external crypto address. This endpoint requires the "transfer"
// permission. API key must belong to default profile.
func cryptoWithdrawalPath(args client.EndpointArgs) string {
	return path.Join("/withdrawals", "crypto")
}

// Withdraws funds from the specified profile_id to a linked external payment method. This endpoint requires the
// "transfer" permission. API key is restricted to the default profile.
func paymentMethodWithdrawalPath(args client.EndpointArgs) string {
	return path.Join("/withdrawals", "payment-method")
}

// Gets the fee estimate for the crypto withdrawal to crypto address
func withdrawalFeeEstimatePath(args client.EndpointArgs) (p string) {
	p = path.Join("/withdrawals", "fee-estimate")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(args.QueryPath().String())
	return sb.String()
}

// Get takes an endpoint const and endpoint arguments to parse the URL endpoint path.
func (ep endpoint) Path(args client.EndpointArgs) string {
	return map[endpoint]func(args client.EndpointArgs) string{
		accountsEndpoint:                accountsPath,
		accountEndpoint:                 accountPath,
		accountHoldsEndpoint:            accountHoldsPath,
		accountLedgerEndpoint:           accountLedgerPath,
		accountTransfersEndpoint:        accountTransfersPath,
		transfersEndpoint:               transfersPath,
		accountTransferEndpoint:         accountTransferPath,
		bookEndpoint:                    bookPath,
		candlesEndpoint:                 candlesPath,
		generateCryptoAddressEndpoint:   generateCryptoAddressPath,
		currenciesEndpoint:              currenciesPath,
		currencyEndpoint:                currencyPath,
		convertCurrencyEndpoint:         convertCurrencyPath,
		currencyConversionEndpoint:      currencyConversionPath,
		coinbaseAccountDepositEndpoint:  coinbaseAccountDepositPath,
		paymentMethodDepositEndpoint:    paymentMethodDepositPath,
		feesEndpoint:                    feesPath,
		fillsEndpoint:                   fillsPath,
		newOrderEndpoint:                newOrderPath,
		signedPricesEndpoint:            signedPricesPath,
		ordersEndpoint:                  ordersPath,
		orderEndpoint:                   orderPath,
		cancelOrderEndpoint:             cancelOrderPath,
		cancelOpenOrdersEndpoint:        cancelOpenOrdersPath,
		paymentMethodsEndpoint:          paymentMethodsPath,
		productsEndpoint:                productsPath,
		productEndpoint:                 productPath,
		productStatsEndpoint:            productStatsPath,
		productTickerEndpoint:           productTickerPath,
		profilesEndpoint:                profilesPath,
		createProfileEndpoint:           createProfilePath,
		createProfileTransferEndpoint:   createProfileTransferPath,
		tradesEndpoint:                  tradesPath,
		walletsEndpoint:                 walletsPath,
		accountWithdrawalEndpoint:       accountWithdrawalPath,
		cryptoWithdrawalEndpoint:        cryptoWithdrawalPath,
		paymentMethodWithdrawalEndpoint: paymentMethodWithdrawalPath,
		withdrawalFeeEstimateEndpoint:   withdrawalFeeEstimatePath,
	}[ep](args)
}
