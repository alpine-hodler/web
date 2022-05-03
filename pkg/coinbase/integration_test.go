package coinbase

import (
	"io/ioutil"
	"testing"

	"github.com/alpine-hodler/sdk/internal/env"
	"github.com/alpine-hodler/sdk/pkg/option"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestSimpleAPIIntegration(t *testing.T) {
	logrus.SetOutput(ioutil.Discard)
	env.Load(".simple-test.env")
	var (
		client        = NewClient(DefaultConnector)
		currency      = "USDC"
		paymentMethod = "Chase:****4442"
		productID     = "BTC-USD"
	)

	makeSimpleRequestAssertion(t, "Should Connect", func() error {
		return client.Connect()
	})
	makeSimpleRequestAssertion(t, "Should GET All Accounts", func() (err error) {
		findAccountID(t, client, currency)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET Account Holds", func() (err error) {
		_, err = client.AccountHolds(findAccountID(t, client, currency), nil)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET Account Ledgers", func() (err error) {
		_, err = client.AccountLedger(findAccountID(t, client, currency), nil)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET Account Transfers", func() (err error) {
		_, err = client.AccountTransfers(findAccountID(t, client, currency), nil)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET All Wallets", func() (err error) {
		findWalletID(t, client, currency)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET Curenncy by name", func() (err error) {
		require.NotNil(t, findCurrency(t, client, currency))
		return
	})
	makeSimpleRequestAssertion(t, "Should GET Fees", func() (err error) {
		_, err = client.Fees()
		return
	})
	makeSimpleRequestAssertion(t, "Should GET Fills", func() (err error) {
		_, err = client.Fills(new(option.CoinbaseFills).SetProductID(productID))
		return
	})
	makeSimpleRequestAssertion(t, "Should GET Payment Methods", func() (err error) {
		findPaymentMethod(t, client, paymentMethod)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET Specific Account", func() (err error) {
		_, err = client.FindAccount(findAccountID(t, client, currency))
		return
	})
	makeSimpleRequestAssertion(t, "Should GET Specific Currency", func() (err error) {
		_, err = client.FindCurrency(currency)
		return
	})
	makeSimpleRequestAssertion(t, "Should POST Crypto Address", func() (err error) {
		generateCryptoAddress(t, client, currency)
		return
	})
	makeSimpleRequestAssertion(t, "Should POST Account Deposit", func() (err error) {
		makeAccountDeposit(t, client, currency, 1)
		return
	})
	makeSimpleRequestAssertion(t, "Should POST Account Withdrawal", func() (err error) {
		makeAccountDeposit(t, client, currency, 1)
		makeAccountWithdrawal(t, client, currency, 1)
		return
	})
	makeSimpleRequestAssertion(t, "Should POST Crypto Withdrawal", func() (err error) {
		makeAccountDeposit(t, client, currency, 1)
		makeCryptoWithdrawal(t, client, currency, 1)
		return
	})
	makeSimpleRequestAssertion(t, "Should POST Payemnt Method Deposit", func() (err error) {
		makeAccountDeposit(t, client, currency, 1)
		makeCryptoWithdrawal(t, client, currency, 1)
		return
	})
	makeSimpleRequestAssertion(t, "Should POST USD Converstion", func() (err error) {
		fromCurrency := "USD"
		makeAccountDeposit(t, client, fromCurrency, 1)
		convertCurrency(t, client, fromCurrency, currency, 1)
		return
	})
}
