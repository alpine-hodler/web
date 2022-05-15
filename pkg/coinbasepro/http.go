package coinbasepro

import "github.com/alpine-hodler/sdk/internal/client"

// * This is a generated file, do not edit

// Account will return data for a single account. Use this endpoint when you know the account_id. API key must belong to
// the same profile as the account.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getaccount
func (c *C) Account(accountId string) (m *Account, _ error) {
	return m, c.Get(AccountPostAuthority).
		SetPathParam("account_id", accountId).
		Fetch().Assign(&m).JoinMessages()
}

// AccountHolds will return the holds of an account that belong to the same profile as the API key. Holds are placed on
// an account for any active orders or pending withdraw requests. As an order is filled, the hold amount is updated. If
// an order is canceled, any remaining hold is removed. For withdrawals, the hold is removed after it is completed.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getaccountholds
func (c *C) AccountHolds(accountId string, opts *AccountHoldsOptions) (m []*AccountHold, _ error) {
	return m, c.Get(AccountHoldsPostAuthority).
		SetPathParam("account_id", accountId).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// AccountLedger returns ledger activity for an account. This includes anything that would affect the accounts balance -
// transfers, trades, fees, etc. This endpoint requires either the "view" or "trade" permission.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getaccountledger
func (c *C) AccountLedger(accountId string, opts *AccountLedgerOptions) (m []*AccountLedger, _ error) {
	return m, c.Get(AccountLedgerPostAuthority).
		SetPathParam("account_id", accountId).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// AccountTransfers returns past withdrawals and deposits for an account.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getaccounttransfers
func (c *C) AccountTransfers(accountId string, opts *AccountTransfersOptions) (m []*Transfer, _ error) {
	return m, c.Get(AccountTransfersPostAuthority).
		SetPathParam("account_id", accountId).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// Accounts will return a list of trading accounts from the profile of the API key.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getaccounts
func (c *C) Accounts() (m []*Account, _ error) {
	return m, c.Get(AccountsPostAuthority).
		Fetch().Assign(&m).JoinMessages()
}

// Book will return a list of open orders for a product. The amount of detail shown can be customized with the level
// parameter.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getproductbook
func (c *C) Book(productId string, opts *BookOptions) (m *Book, _ error) {
	return m, c.Get(BookPostAuthority).
		SetPathParam("product_id", productId).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// CancelOpenOrders will try with best effort to cancel all open orders. This may require you to make the request
// multiple times until all of the open orders are deleted.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_deleteorders
func (c *C) CancelOpenOrders(opts *CancelOpenOrdersOptions) (m []*string, _ error) {
	return m, c.Delete(CancelOpenOrdersPostAuthority).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// CancelOrder will cancel a single open order by order id.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_deleteorder
func (c *C) CancelOrder(orderId string) (m string, _ error) {
	return m, c.Delete(CancelOrderPostAuthority).
		SetPathParam("order_id", orderId).
		Fetch().Assign(&m).JoinMessages()
}

// Candles will return historic rates for a product.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getproductcandles
func (c *C) Candles(productId string, opts *CandlesOptions) (m *Candles, _ error) {
	return m, c.Get(CandlesPostAuthority).
		SetPathParam("product_id", productId).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// CoinbaseAccountDeposit funds from a www.coinbase.com wallet to the specified profile_id.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postdepositcoinbaseaccount
func (c *C) CoinbaseAccountDeposit(opts *CoinbaseAccountDepositOptions) (m *Deposit, _ error) {
	return m, c.Post(CoinbaseAccountDepositPostAuthority).
		SetBody(client.BodyTypeJSON, opts).
		Fetch().Assign(&m).JoinMessages()
}

// AccountWithdraws funds from the specified profile_id to a www.coinbase.com wallet. Withdraw funds to a coinbase
// account. You can move funds between your Coinbase accounts and your Coinbase Exchange trading accounts within your
// daily limits. Moving funds between Coinbase and Coinbase Exchange is instant and free. See the Coinbase Accounts
// section for retrieving your Coinbase accounts. This endpoint requires the "transfer" permission.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postwithdrawcoinbaseaccount
func (c *C) CoinbaseAccountWithdrawal(opts *CoinbaseAccountWithdrawalOptions) (m *Withdrawal, _ error) {
	return m, c.Post(CoinbaseAccountWithdrawalPostAuthority).
		SetBody(client.BodyTypeJSON, opts).
		Fetch().Assign(&m).JoinMessages()
}

// ConvertCurrency converts funds from from currency to to currency. Funds are converted on the from account in the
// profile_id profile. This endpoint requires the "trade" permission. A successful conversion will be assigned a
// conversion id. The corresponding ledger entries for a conversion will reference this conversion id
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postconversion
func (c *C) ConvertCurrency(opts *ConvertCurrencyOptions) (m *CurrencyConversion, _ error) {
	return m, c.Post(ConvertCurrencyPostAuthority).
		SetBody(client.BodyTypeJSON, opts).
		Fetch().Assign(&m).JoinMessages()
}

// CreateOrder will create a new an order. You can place two types of orders: limit and market. Orders can only be
// placed if your account has sufficient funds. Once an order is placed, your account funds will be put on hold for the
// duration of the order. How much and which funds are put on hold depends on the order type and parameters specified.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postorders
func (c *C) CreateOrder(opts *CreateOrderOptions) (m *CreateOrder, _ error) {
	return m, c.Post(CreateOrderPostAuthority).
		SetBody(client.BodyTypeJSON, opts).
		Fetch().Assign(&m).JoinMessages()
}

// CreateProfile will create a new profile. Will fail if no name is provided or if user already has max number of
// profiles.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postprofile
func (c *C) CreateProfile(opts *CreateProfileOptions) (m *Profile, _ error) {
	return m, c.Post(CreateProfilePostAuthority).
		SetBody(client.BodyTypeJSON, opts).
		Fetch().Assign(&m).JoinMessages()
}

// CreateProfileTransfer will transfer an amount of currency from one profile to another. This endpoint requires the
// "transfer" permission.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postprofiletransfer
func (c *C) CreateProfileTransfer(opts *CreateProfileTransferOptions) error {
	return c.Post(CreateProfileTransferPostAuthority).
		SetBody(client.BodyTypeJSON, opts).
		Fetch().NoAssignment().JoinMessages()
}

// CreateReport generates a report. Reports are either for past account history or past fills on either all accounts or
// one product's account.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postreports
func (c *C) CreateReport(opts *CreateReportOptions) (m *CreateReport, _ error) {
	return m, c.Post(CreateReportPostAuthority).
		SetBody(client.BodyTypeJSON, opts).
		Fetch().Assign(&m).JoinMessages()
}

// CryptoWithdrawal funds from the specified profile_id to an external crypto address. This endpoint requires the
// "transfer" permission. API key must belong to default profile.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postwithdrawcrypto
func (c *C) CryptoWithdrawal(opts *CryptoWithdrawalOptions) (m *Withdrawal, _ error) {
	return m, c.Post(CryptoWithdrawalPostAuthority).
		SetBody(client.BodyTypeJSON, opts).
		Fetch().Assign(&m).JoinMessages()
}

// Currencies returns a list of all known currencies. Note: Not all currencies may be currently in use for trading.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getcurrencies
func (c *C) Currencies() (m []*Currency, _ error) {
	return m, c.Get(CurrenciesPostAuthority).
		Fetch().Assign(&m).JoinMessages()
}

// Currency returns a single currency by id.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getcurrency
func (c *C) Currency(currencyId string) (m *Currency, _ error) {
	return m, c.Get(CurrencyPostAuthority).
		SetPathParam("currency_id", currencyId).
		Fetch().Assign(&m).JoinMessages()
}

// CurrencyConversion returns the currency conversion by conversion id (i.e. USD -> USDC).
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getconversion
func (c *C) CurrencyConversion(conversionId string, opts *CurrencyConversionOptions) (m *CurrencyConversion, _ error) {
	return m, c.Get(CurrencyConversionPostAuthority).
		SetPathParam("conversion_id", conversionId).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// DeleteProfile deletes the profile specified by profile_id and transfers all funds to the profile specified by to.
// Fails if there are any open orders on the profile to be deleted.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_putprofiledeactivate
func (c *C) DeleteProfile(profileId string, opts *DeleteProfileOptions) error {
	return c.Put(DeleteProfilePostAuthority).
		SetPathParam("profile_id", profileId).
		SetQueryParams(opts).
		Fetch().NoAssignment().JoinMessages()
}

// ExchangeLimits returns exchange limits information for a single user.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getuserexchangelimits
func (c *C) ExchangeLimits(userId string) (m *ExchangeLimits, _ error) {
	return m, c.Get(ExchangeLimitsPostAuthority).
		SetPathParam("user_id", userId).
		Fetch().Assign(&m).JoinMessages()
}

// Fees returns fees rates and 30 days trailing volume.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getfees
func (c *C) Fees() (m *Fees, _ error) {
	return m, c.Get(FeesPostAuthority).
		Fetch().Assign(&m).JoinMessages()
}

// Fills returns a list of fills. A fill is a partial or complete match on a specific order.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getfills
func (c *C) Fills(opts *FillsOptions) (m []*Fill, _ error) {
	return m, c.Get(FillsPostAuthority).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// GenerateCryptoAddress will create a one-time crypto address for depositing crypto, using a wallet account id. This
// endpoint requires the "transfer" permission. API key must belong to default profile.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postcoinbaseaccountaddresses
func (c *C) GenerateCryptoAddress(accountId string) (m *CryptoAddress, _ error) {
	return m, c.Post(GenerateCryptoAddressPostAuthority).
		SetPathParam("account_id", accountId).
		Fetch().Assign(&m).JoinMessages()
}

// Order returns a single order by id.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getorder
func (c *C) Order(orderId string) (m *Order, _ error) {
	return m, c.Get(OrderPostAuthority).
		SetPathParam("order_id", orderId).
		Fetch().Assign(&m).JoinMessages()
}

// Orders will return your current open orders. Only open or un-settled orders are returned by default. As soon as an
// order is no longer open and settled, it will no longer appear in the default request. Open orders may change state
// between the request and the response depending on market conditions.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getorders
func (c *C) Orders(opts *OrdersOptions) (m []*Order, _ error) {
	return m, c.Get(OrdersPostAuthority).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// PaymentMethodDeposit will fund from a linked external payment method to the specified profile_id.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postdepositpaymentmethod
func (c *C) PaymentMethodDeposit(opts *PaymentMethodDepositOptions) (m *Deposit, _ error) {
	return m, c.Post(PaymentMethodDepositPostAuthority).
		SetBody(client.BodyTypeJSON, opts).
		Fetch().Assign(&m).JoinMessages()
}

// PaymentMethodWithdrawal will fund from the specified profile_id to a linked external payment method. This endpoint
// requires the "transfer" permission. API key is restricted to the default profile.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postwithdrawpaymentmethod
func (c *C) PaymentMethodWithdrawal(opts *PaymentMethodWithdrawalOptions) (m *Withdrawal, _ error) {
	return m, c.Post(PaymentMethodWithdrawalPostAuthority).
		SetBody(client.BodyTypeJSON, opts).
		Fetch().Assign(&m).JoinMessages()
}

// PaymentMethods returns a list of the user's linked payment methods.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getpaymentmethods
func (c *C) PaymentMethods() (m []*PaymentMethod, _ error) {
	return m, c.Get(PaymentMethodsPostAuthority).
		Fetch().Assign(&m).JoinMessages()
}

// Product will return information on a single product.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getproduct
func (c *C) Product(productId string) (m *Product, _ error) {
	return m, c.Get(ProductPostAuthority).
		SetPathParam("product_id", productId).
		Fetch().Assign(&m).JoinMessages()
}

// ProductStats will return 30day and 24hour stats for a product.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getproductstats
func (c *C) ProductStats(productId string) (m *ProductStats, _ error) {
	return m, c.Get(ProductStatsPostAuthority).
		SetPathParam("product_id", productId).
		Fetch().Assign(&m).JoinMessages()
}

// ProductTicker will return snapshot information about the last trade (tick), best bid/ask and 24h volume.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getproductticker
func (c *C) ProductTicker(productId string) (m *ProductTicker, _ error) {
	return m, c.Get(ProductTickerPostAuthority).
		SetPathParam("product_id", productId).
		Fetch().Assign(&m).JoinMessages()
}

// Products will return a list of available currency pairs for trading.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getproducts
func (c *C) Products(opts *ProductsOptions) (m []*Product, _ error) {
	return m, c.Get(ProductsPostAuthority).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// Profile returns information for a single profile. Use this endpoint when you know the profile_id. This endpoint
// requires the "view" permission and is accessible by any profile's API key.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getprofile
func (c *C) Profile(profileId string, opts *ProfileOptions) (m *Profile, _ error) {
	return m, c.Get(ProfilePostAuthority).
		SetPathParam("profile_id", profileId).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// Profiles returns a list of all of the current user's profiles.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getprofiles
func (c *C) Profiles(opts *ProfilesOptions) (m []*Profile, _ error) {
	return m, c.Get(ProfilesPostAuthority).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// RenameProfile will rename a profile. Names 'default' and 'margin' are reserved.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_putprofile
func (c *C) RenameProfile(profileId string, opts *RenameProfileOptions) (m *Profile, _ error) {
	return m, c.Put(RenameProfilePostAuthority).
		SetPathParam("profile_id", profileId).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// Report will return a specific report by report_id.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getreport
func (c *C) Report(reportId string) (m *Report, _ error) {
	return m, c.Get(ReportPostAuthority).
		SetPathParam("report_id", reportId).
		Fetch().Assign(&m).JoinMessages()
}

// Reports returns a list of past fills/account reports.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getreports
func (c *C) Reports(opts *ReportsOptions) (m []*Report, _ error) {
	return m, c.Get(ReportsPostAuthority).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// SignedPrices returns cryptographically signed prices ready to be posted on-chain using Compound's Open Oracle smart
// contract.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getcoinbasepriceoracle
func (c *C) SignedPrices() (m *Oracle, _ error) {
	return m, c.Get(SignedPricesPostAuthority).
		Fetch().Assign(&m).JoinMessages()
}

// Trades retruns a list the latest trades for a product.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getproducttrades
func (c *C) Trades(productId string, opts *TradesOptions) (m []*Trade, _ error) {
	return m, c.Get(TradesPostAuthority).
		SetPathParam("product_id", productId).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// AccountTransfer returns information on a single transfer.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_gettransfer
func (c *C) Transfer(transferId string) (m *Transfer, _ error) {
	return m, c.Get(TransferPostAuthority).
		SetPathParam("transfer_id", transferId).
		Fetch().Assign(&m).JoinMessages()
}

// Transfers is a list of in-progress and completed transfers of funds in/out of any of the user's accounts.
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_gettransfers
func (c *C) Transfers() (m []*Transfer, _ error) {
	return m, c.Get(TransfersPostAuthority).
		Fetch().Assign(&m).JoinMessages()
}

// Wallets will return all the user's available Coinbase wallets (These are the wallets/accounts that are used for
// buying and selling on www.coinbase.com)
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getcoinbaseaccounts
func (c *C) Wallets() (m []*Wallet, _ error) {
	return m, c.Get(WalletsPostAuthority).
		Fetch().Assign(&m).JoinMessages()
}

// WithdrawalFeeEstimate will return the fee estimate for the crypto withdrawal to crypto address
//
// source: https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getwithdrawfeeestimate
func (c *C) WithdrawalFeeEstimate(opts *WithdrawalFeeEstimateOptions) (m *WithdrawalFeeEstimate, _ error) {
	return m, c.Get(WithdrawalFeeEstimatePostAuthority).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}
