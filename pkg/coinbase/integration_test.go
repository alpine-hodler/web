package coinbase

import (
	"testing"

	"github.com/alpine-hodler/sdk/internal/env"
	"github.com/stretchr/testify/require"
)

func TestSimpleAPIIntegration(t *testing.T) {
	env.Load(".simple-test.env")
	var (
		client        = NewClient(DefaultConnector)
		currency      = "USDC"
		paymentMethod = "Chase:****4442"
		productID     = "BTC-USD"
		profileName   = "default"
	)

	makeSimpleRequestAssertion(t, "Should Connect", func() error {
		return client.Connect()
	})
	makeSimpleRequestAssertion(t, "Should DELETE client.CancelOpenOrders", func() (err error) {
		profile := findProfileByName(t, client, profileName)
		makeNewOrder(t, client, "BTC-USD", profile)
		_, err = client.CancelOpenOrders(nil)
		return
	})
	makeSimpleRequestAssertion(t, "Should DELETE client.CancelOrder", func() (err error) {
		profile := findProfileByName(t, client, profileName)
		order := makeNewOrder(t, client, "BTC-USD", profile)
		_, err = client.CancelOrder(order.Id)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.Accounts", func() (err error) {
		findAccountID(t, client, currency)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.AccountHolds", func() (err error) {
		_, err = client.AccountHolds(findAccountID(t, client, currency), nil)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.AccountLedger", func() (err error) {
		_, err = client.AccountLedger(findAccountID(t, client, currency), nil)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.AccountTransfers", func() (err error) {
		_, err = client.AccountTransfers(findAccountID(t, client, currency), nil)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.FindConversion", func() (err error) {
		fromCurrency := "USD"
		makeAccountDeposit(t, client, fromCurrency, 1)
		conversion := convertCurrency(t, client, fromCurrency, currency, 1)
		_, err = client.CurrencyConversion(conversion.Id, nil)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.Currencies", func() (err error) {
		_, err = client.Currencies()
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.FindCurrency", func() (err error) {
		require.NotNil(t, findCurrency(t, client, currency))
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.Fees", func() (err error) {
		_, err = client.Fees()
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.Fills", func() (err error) {
		_, err = client.Fills(new(FillsOptions).SetProductId(productID))
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.Order", func() (err error) {
		order := makeNewOrder(t, client, "BTC-USD", findProfileByName(t, client, profileName))
		_, err = client.Order(order.Id)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.Orders", func() (err error) {
		_, err = client.Orders(new(OrdersOptions).SetProductId("BTC-USD").SetLimit(1))
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.PaymentMethods", func() (err error) {
		findPaymentMethod(t, client, paymentMethod)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.Products", func() (err error) {
		products, err := client.Products(nil)
		require.NotEmpty(t, products)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.Profiles", func() (err error) {
		require.NotNil(t, findProfileByName(t, client, profileName))
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.FindAccount", func() (err error) {
		_, err = client.Account(findAccountID(t, client, currency))
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.Product", func() (err error) {
		products, err := client.Products(nil)
		require.NotEmpty(t, products)

		product, err := client.Product(products[0].Id)
		require.NotNil(t, product)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.FindCurrency", func() (err error) {
		_, err = client.Currency(currency)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.FindTransfer", func() (err error) {
		transfer := makeAccountDeposit(t, client, "USD", 1.0)
		require.NotNil(t, transfer)

		_, err = client.AccountTransfer(transfer.Id)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.Transfers", func() (err error) {
		makeAccountDeposit(t, client, "USD", 1.0)
		transfers, err := client.Transfers()
		require.NotEmpty(t, transfers)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET clientWallets", func() (err error) {
		findWalletID(t, client, currency)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.WithdrawalFeeEstimate", func() (err error) {
		_, err = client.WithdrawalFeeEstimate(new(WithdrawalFeeEstimateOptions).
			SetCryptoAddress(generateCryptoAddress(t, client, "USDC")).
			SetCurrency("USDC"))
		return
	})
	makeSimpleRequestAssertion(t, "Should POST client.GenerateCryptoAddress", func() (err error) {
		generateCryptoAddress(t, client, currency)
		return
	})
	makeSimpleRequestAssertion(t, "Should POST client.CoinbaseAccountDeposit", func() (err error) {
		makeAccountDeposit(t, client, currency, 1)
		return
	})
	makeSimpleRequestAssertion(t, "Should POST client.AccountWithdrawal", func() (err error) {
		makeAccountDeposit(t, client, currency, 1)
		makeAccountWithdrawal(t, client, currency, 1)
		return
	})
	makeSimpleRequestAssertion(t, "Should POST client.CryptoWithdrawal", func() (err error) {
		makeAccountDeposit(t, client, currency, 1)
		makeCryptoWithdrawal(t, client, currency, 1)
		return
	})
	makeSimpleRequestAssertion(t, "Should POST client.CreateOrder", func() (err error) {
		profile := findProfileByName(t, client, profileName)
		makeNewOrder(t, client, "BTC-USD", profile)
		return
	})
	makeSimpleRequestAssertion(t, "Should POST client.PaymentMethodDeposit", func() (err error) {
		makeAccountDeposit(t, client, currency, 1)
		makeCryptoWithdrawal(t, client, currency, 1)
		return
	})
	makeSimpleRequestAssertion(t, "Should POST client.PaymentMethodWithdrawal", func() (err error) {
		// TODO: Fix {Message:Invalid scope Status:403 Forbidden StatusCode:Forbidden}
		// makeAccountDeposit(t, client, "USDC", 1)
		// _, err = client.PaymentMethodWithdrawal(new(option.CoinbasePaymentMethodWithdrawal).
		// 	SetAmount(1).
		// 	SetPaymentMethodId(findPaymentMethod(t, client, paymentMethod).Id).
		// 	SetCurrency("USDC").
		// 	SetProfileId(findProfileByName(t, client, profileName).Id))
		return
	})
	makeSimpleRequestAssertion(t, "Should POST client.Convert", func() (err error) {
		fromCurrency := "USD"
		makeAccountDeposit(t, client, fromCurrency, 1)
		convertCurrency(t, client, fromCurrency, currency, 1)
		return
	})
}
