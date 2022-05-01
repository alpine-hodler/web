package coinbase

import (
	"io/ioutil"
	"testing"

	"github.com/alpine-hodler/sdk/internal/env"
	"github.com/alpine-hodler/sdk/pkg/model"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func makeSimpleRequestAssertion(t *testing.T, name string, req func() error) {
	t.Run(name, func(t *testing.T) { require.NoError(t, req()) })
}

func testGetAccount(t *testing.T, client *C, currency string) *model.CoinbaseAccount {
	accounts, err := client.Accounts()
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	for _, account := range accounts {
		if account.Currency == currency {
			return account
		}
	}
	return nil
}

func testGetAccountID(t *testing.T, client *C, currency string) (accountID string) {
	account := testGetAccount(t, client, currency)
	require.NotEmpty(t, account)
	return account.Id
}

func testGetCryptoAddress(t *testing.T, client *C, currency string) string {
	address, err := client.GenerateCryptoAddress(testGetWalletID(t, client, currency))
	require.NoError(t, err)
	require.NotNil(t, address)
	return address.Address
}

func testGetCurrency(t *testing.T, client *C, currency string) *model.CoinbaseCurrency {
	currencies, err := client.Currencies()
	require.NoError(t, err)
	require.NotEmpty(t, currencies)
	for _, item := range currencies {
		if item.Id == currency {
			return item
		}
	}
	return nil
}

func testGetPaymentMethod(t *testing.T, client *C, name string) *model.CoinbasePaymentMethod {
	methods, err := client.PaymentMethods()
	require.NoError(t, err)
	require.NotEmpty(t, methods)
	for _, method := range methods {
		if method.Name == name {
			return method
		}
	}
	return nil
}

func testGetProfileID(t *testing.T, client *C, currency string) string {
	account := testGetAccount(t, client, currency)
	require.NotEmpty(t, account)
	return account.ProfileId
}

func testGetWalletID(t *testing.T, client *C, currency string) (walletID string) {
	wallets, err := client.Wallets()
	require.NoError(t, err)
	require.NotEmpty(t, wallets)
	for _, wallet := range wallets {
		if wallet.Currency == currency {
			walletID = wallet.Id
			break
		}
	}
	require.NotEmpty(t, walletID)
	return
}

func testPostAccountDeposit(t *testing.T, client *C, currency string, amount float64) {
	_, err := client.CoinbaseAccountDeposit(&model.CoinbaseAccountDepositOptions{
		CoinbaseAccountID: testGetWalletID(t, client, currency),
		Amount:            amount,
		Currency:          currency,
	})
	require.NoError(t, err)
}

func testPostAccountWithdrawal(t *testing.T, client *C, currency string, amount float64) {
	_, err := client.AccountWithdrawal(&model.CoinbaseAccountWithdrawalOptions{
		CoinbaseAccountID: testGetWalletID(t, client, currency),
		Amount:            amount,
		Currency:          currency,
	})
	require.NoError(t, err)
}

func testPostConvert(t *testing.T, client *C, from, to string, amount float64) {
	profileID := testGetProfileID(t, client, from)
	_, err := client.Convert(model.CoinbaseConversionsOptions{
		ProfileID: &profileID,
		From:      from,
		To:        to,
		Amount:    amount,
	})
	require.NoError(t, err)
}

func testPostCryptoWithdrawal(t *testing.T, client *C, currency string, amount float64) {
	_, err := client.CryptoWithdrawal(&model.CoinbaseCryptoWithdrawalOptions{
		CryptoAddress: testGetCryptoAddress(t, client, currency),
		Amount:        amount,
		Currency:      currency,
	})
	require.NoError(t, err)
}

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
		testGetAccountID(t, client, currency)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET Account Holds", func() (err error) {
		_, err = client.AccountHolds(testGetAccountID(t, client, currency), nil)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET Account Ledgers", func() (err error) {
		_, err = client.AccountLedger(testGetAccountID(t, client, currency), nil)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET Account Transfers", func() (err error) {
		_, err = client.AccountTransfers(testGetAccountID(t, client, currency), nil)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET All Wallets", func() (err error) {
		testGetWalletID(t, client, currency)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET Curenncy by name", func() (err error) {
		require.NotNil(t, testGetCurrency(t, client, currency))
		return
	})
	makeSimpleRequestAssertion(t, "Should GET Fees", func() (err error) {
		_, err = client.Fees()
		return
	})
	makeSimpleRequestAssertion(t, "Should GET Fills", func() (err error) {
		_, err = client.Fills(&model.CoinbaseFillsOptions{ProductID: &productID})
		return
	})
	makeSimpleRequestAssertion(t, "Should GET Payment Methods", func() (err error) {
		testGetPaymentMethod(t, client, paymentMethod)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET Specific Account", func() (err error) {
		_, err = client.FindAccount(testGetAccountID(t, client, currency))
		return
	})
	makeSimpleRequestAssertion(t, "Should GET Specific Currency", func() (err error) {
		_, err = client.FindCurrency(currency)
		return
	})
	makeSimpleRequestAssertion(t, "Should POST Crypto Address", func() (err error) {
		testGetCryptoAddress(t, client, currency)
		return
	})
	makeSimpleRequestAssertion(t, "Should POST Account Deposit", func() (err error) {
		testPostAccountDeposit(t, client, currency, 1)
		return
	})
	makeSimpleRequestAssertion(t, "Should POST Account Withdrawal", func() (err error) {
		testPostAccountDeposit(t, client, currency, 1)
		testPostAccountWithdrawal(t, client, currency, 1)
		return
	})
	makeSimpleRequestAssertion(t, "Should POST Crypto Withdrawal", func() (err error) {
		testPostAccountDeposit(t, client, currency, 1)
		testPostCryptoWithdrawal(t, client, currency, 1)
		return
	})
	makeSimpleRequestAssertion(t, "Should POST Payemnt Method Deposit", func() (err error) {
		testPostAccountDeposit(t, client, currency, 1)
		testPostCryptoWithdrawal(t, client, currency, 1)
		return
	})
	makeSimpleRequestAssertion(t, "Should POST USD Converstion", func() (err error) {
		fromCurrency := "USD"
		testPostAccountDeposit(t, client, fromCurrency, 1)
		testPostConvert(t, client, fromCurrency, currency, 1)
		return
	})
}
