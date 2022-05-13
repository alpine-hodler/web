package coinbase

import "github.com/alpine-hodler/sdk/internal/client"

// * This is a generated file, do not edit

// Information for a single account. Use this endpoint when you know the account_id. API key must belong to the same
// profile as the account.
func (c *C) Account(accountId string) (m *Account, _ error) {
	return m, c.Get(accountEndpoint).
		PathParam("account_id", accountId).
		Fetch().Assign(&m).JoinMessages()
}

// AccountHolds will return the holds of an account that belong to the same profile as the API key. Holds are placed on
// an account for any active orders or pending withdraw requests. As an order is filled, the hold amount is updated. If
// an order is canceled, any remaining hold is removed. For withdrawals, the hold is removed after it is completed.
func (c *C) AccountHolds(accountId string, opts *AccountHoldsOptions) (m []*AccountHold, _ error) {
	return m, c.Get(accountHoldsEndpoint).
		PathParam("account_id", accountId).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// Lists ledger activity for an account. This includes anything that would affect the accounts balance - transfers,
// trades, fees, etc. This endpoint requires either the "view" or "trade" permission.
func (c *C) AccountLedger(accountId string, opts *AccountLedgerOptions) (m []*AccountLedger, _ error) {
	return m, c.Get(accountLedgerEndpoint).
		PathParam("account_id", accountId).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// AccountTransfer returns information on a single transfer.
func (c *C) AccountTransfer(transferId string) (m *AccountTransfer, _ error) {
	return m, c.Get(accountTransferEndpoint).
		PathParam("transfer_id", transferId).
		Fetch().Assign(&m).JoinMessages()
}

// AccountTransfers returns past withdrawals and deposits for an account.
func (c *C) AccountTransfers(accountId string, opts *AccountTransfersOptions) (m []*AccountTransfer, _ error) {
	return m, c.Get(accountTransfersEndpoint).
		PathParam("account_id", accountId).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// Withdraws funds from the specified profile_id to a www.coinbase.com wallet. Withdraw funds to a coinbase account. You
// can move funds between your Coinbase accounts and your Coinbase Exchange trading accounts within your daily limits.
// Moving funds between Coinbase and Coinbase Exchange is instant and free. See the Coinbase Accounts section for
// retrieving your Coinbase accounts. This endpoint requires the "transfer" permission.
func (c *C) AccountWithdrawal(opts *AccountWithdrawalOptions) (m *Withdrawal, _ error) {
	return m, c.Post(accountWithdrawalEndpoint).
		SetBody(client.BodyTypeJSON, opts).
		Fetch().Assign(&m).JoinMessages()
}

// Get a list of trading accounts from the profile of the API key.
func (c *C) Accounts() (m []*Account, _ error) {
	return m, c.Get(accountsEndpoint).
		Fetch().Assign(&m).JoinMessages()
}

// Get a list of open orders for a product. The amount of detail shown can be customized with the level parameter.
func (c *C) Book(productId string, opts *BookOptions) (m *Book, _ error) {
	return m, c.Get(bookEndpoint).
		PathParam("product_id", productId).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// CancelOpenOrders will try with best effort to cancel all open orders. This may require you to make the request
// multiple times until all of the open orders are deleted.
func (c *C) CancelOpenOrders(opts *CancelOpenOrdersOptions) (m []*string, _ error) {
	return m, c.Delete(cancelOpenOrdersEndpoint).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// CancelOrder will cancel a single open order by order id.
func (c *C) CancelOrder(orderId string) (m string, _ error) {
	return m, c.Delete(cancelOrderEndpoint).
		PathParam("order_id", orderId).
		Fetch().Assign(&m).JoinMessages()
}

// Get the historic rates for a product
func (c *C) Candles(productId string, opts *CandlesOptions) (m *Candles, _ error) {
	return m, c.Get(candlesEndpoint).
		PathParam("product_id", productId).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// Deposits funds from a www.coinbase.com wallet to the specified profile_id.
func (c *C) CoinbaseAccountDeposit(opts *CoinbaseAccountDepositOptions) (m *Deposit, _ error) {
	return m, c.Post(coinbaseAccountDepositEndpoint).
		SetBody(client.BodyTypeJSON, opts).
		Fetch().Assign(&m).JoinMessages()
}

// ConvertCurrency converts funds from from currency to to currency. Funds are converted on the from account in the
// profile_id profile. This endpoint requires the "trade" permission. A successful conversion will be assigned a
// conversion id. The corresponding ledger entries for a conversion will reference this conversion id
func (c *C) ConvertCurrency(opts *ConvertCurrencyOptions) (m *CurrencyConversion, _ error) {
	return m, c.Post(convertCurrencyEndpoint).
		SetBody(client.BodyTypeJSON, opts).
		Fetch().Assign(&m).JoinMessages()
}

// Withdraws funds from the specified profile_id to an external crypto address. This endpoint requires the "transfer"
// permission. API key must belong to default profile.
func (c *C) CryptoWithdrawal(opts *CryptoWithdrawalOptions) (m *Withdrawal, _ error) {
	return m, c.Post(cryptoWithdrawalEndpoint).
		SetBody(client.BodyTypeJSON, opts).
		Fetch().Assign(&m).JoinMessages()
}

// Gets a list of all known currencies. Note: Not all currencies may be currently in use for trading.
func (c *C) Currencies() (m []*Currency, _ error) {
	return m, c.Get(currenciesEndpoint).
		Fetch().Assign(&m).JoinMessages()
}

// Gets a single currency by id.
func (c *C) Currency(currencyId string) (m *Currency, _ error) {
	return m, c.Get(currencyEndpoint).
		PathParam("currency_id", currencyId).
		Fetch().Assign(&m).JoinMessages()
}

// CurrencyConversion returns the currency conversion by conversion id (i.e. USD -> USDC).
func (c *C) CurrencyConversion(conversionId string, opts *CurrencyConversionOptions) (m *CurrencyConversion, _ error) {
	return m, c.Get(currencyConversionEndpoint).
		PathParam("conversion_id", conversionId).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// Get fees rates and 30 days trailing volume.
func (c *C) Fees() (m *Fees, _ error) {
	return m, c.Get(feesEndpoint).
		Fetch().Assign(&m).JoinMessages()
}

// Get a list of fills. A fill is a partial or complete match on a specific order.
func (c *C) Fills(opts *FillsOptions) (m []*Fill, _ error) {
	return m, c.Get(fillsEndpoint).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// GenerateCryptoAddress will create a one-time crypto address for depositing crypto, using a wallet account id. This
// endpoint requires the "transfer" permission. API key must belong to default profile.
func (c *C) GenerateCryptoAddress(accountId string) (m *CryptoAddress, _ error) {
	return m, c.Post(generateCryptoAddressEndpoint).
		PathParam("account_id", accountId).
		Fetch().Assign(&m).JoinMessages()
}

// Create an order. You can place two types of orders: limit and market. Orders can only be placed if your account has
// sufficient funds. Once an order is placed, your account funds will be put on hold for the duration of the order. How
// much and which funds are put on hold depends on the order type and parameters specified.
func (c *C) NewOrder(opts *NewOrderOptions) (m *NewOrder, _ error) {
	return m, c.Post(newOrderEndpoint).
		SetBody(client.BodyTypeJSON, opts).
		Fetch().Assign(&m).JoinMessages()
}

// Get a single order by id.
func (c *C) Order(orderId string) (m *Order, _ error) {
	return m, c.Get(orderEndpoint).
		PathParam("order_id", orderId).
		Fetch().Assign(&m).JoinMessages()
}

// List your current open orders. Only open or un-settled orders are returned by default. As soon as an order is no
// longer open and settled, it will no longer appear in the default request. Open orders may change state between the
// request and the response depending on market conditions.
func (c *C) Orders(opts *OrdersOptions) (m []*Order, _ error) {
	return m, c.Get(ordersEndpoint).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// Deposits funds from a linked external payment method to the specified profile_id.
func (c *C) PaymentMethodDeposit(opts *PaymentMethodDepositOptions) (m *Deposit, _ error) {
	return m, c.Post(paymentMethodDepositEndpoint).
		SetBody(client.BodyTypeJSON, opts).
		Fetch().Assign(&m).JoinMessages()
}

// Withdraws funds from the specified profile_id to a linked external payment method. This endpoint requires the
// "transfer" permission. API key is restricted to the default profile.
func (c *C) PaymentMethodWithdrawal(opts *PaymentMethodWithdrawalOptions) (m *Withdrawal, _ error) {
	return m, c.Post(paymentMethodWithdrawalEndpoint).
		SetBody(client.BodyTypeJSON, opts).
		Fetch().Assign(&m).JoinMessages()
}

// PaymentMethods returns a list of the user's linked payment methods.
func (c *C) PaymentMethods() (m []*PaymentMethod, _ error) {
	return m, c.Get(paymentMethodsEndpoint).
		Fetch().Assign(&m).JoinMessages()
}

// Get information on a single product.
func (c *C) Product(productId string) (m *Product, _ error) {
	return m, c.Get(productEndpoint).
		PathParam("product_id", productId).
		Fetch().Assign(&m).JoinMessages()
}

// Gets 30day and 24hour stats for a product.
func (c *C) ProductStats(productId string) (m *ProductStats, _ error) {
	return m, c.Get(productStatsEndpoint).
		PathParam("product_id", productId).
		Fetch().Assign(&m).JoinMessages()
}

// Gets snapshot information about the last trade (tick), best bid/ask and 24h volume.
func (c *C) ProductTicker(productId string) (m *ProductTicker, _ error) {
	return m, c.Get(productTickerEndpoint).
		PathParam("product_id", productId).
		Fetch().Assign(&m).JoinMessages()
}

// Gets a list of available currency pairs for trading.
func (c *C) Products(opts *ProductsOptions) (m []*Product, _ error) {
	return m, c.Get(productsEndpoint).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// Gets a list of all of the current user's profiles.
func (c *C) Profiles(opts *ProfilesOptions) (m []*Profile, _ error) {
	return m, c.Get(profilesEndpoint).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// SignedPrices returns cryptographically signed prices ready to be posted on-chain using Compound's Open Oracle smart
// contract.
func (c *C) SignedPrices() (m *Oracle, _ error) {
	return m, c.Get(signedPricesEndpoint).
		Fetch().Assign(&m).JoinMessages()
}

// Gets a list the latest trades for a product.
func (c *C) Trades(productId string, opts *TradesOptions) (m []*Trade, _ error) {
	return m, c.Get(tradesEndpoint).
		PathParam("product_id", productId).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// Transfers is a list of in-progress and completed transfers of funds in/out of any of the user's accounts.
func (c *C) Transfers() (m []*AccountTransfer, _ error) {
	return m, c.Get(transfersEndpoint).
		Fetch().Assign(&m).JoinMessages()
}

// Gets all the user's available Coinbase wallets (These are the wallets/accounts that are used for buying and selling
// on www.coinbase.com)
func (c *C) Wallets() (m []*Wallet, _ error) {
	return m, c.Get(walletsEndpoint).
		Fetch().Assign(&m).JoinMessages()
}

// Gets the fee estimate for the crypto withdrawal to crypto address
func (c *C) WithdrawalFeeEstimate(opts *WithdrawalFeeEstimateOptions) (m *WithdrawalFeeEstimate, _ error) {
	return m, c.Get(withdrawalFeeEstimateEndpoint).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}
