package coinbase

import (
	"testing"

	"github.com/alpine-hodler/sdk/pkg/model"
	"github.com/alpine-hodler/sdk/pkg/option"
	"github.com/stretchr/testify/require"
)

func convertCurrency(t *testing.T, client *C, from, to string, amount float64) {
	opts := new(option.CoinbaseConversions).
		SetProfileId(findProfileID(t, client, from)).
		SetFrom(from).
		SetTo(to).
		SetAmount(amount)
	_, err := client.Convert(opts)
	require.NoError(t, err)
}

func findAccount(t *testing.T, client *C, currency string) *model.CoinbaseAccount {
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

func findAccountID(t *testing.T, client *C, currency string) (accountID string) {
	account := findAccount(t, client, currency)
	require.NotEmpty(t, account)
	return account.Id
}

func findCurrency(t *testing.T, client *C, currency string) *model.CoinbaseCurrency {
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

func findPaymentMethod(t *testing.T, client *C, name string) *model.CoinbasePaymentMethod {
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

func findProfileID(t *testing.T, client *C, currency string) string {
	account := findAccount(t, client, currency)
	require.NotEmpty(t, account)
	return account.ProfileId
}

func findWalletID(t *testing.T, client *C, currency string) (walletID string) {
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

func generateCryptoAddress(t *testing.T, client *C, currency string) string {
	address, err := client.GenerateCryptoAddress(findWalletID(t, client, currency))
	require.NoError(t, err)
	require.NotNil(t, address)
	return address.Address
}

func makeAccountDeposit(t *testing.T, client *C, currency string, amount float64) {
	opts := new(option.CoinbaseAccountDeposit).
		SetCoinbaseAccountId(findWalletID(t, client, currency)).
		SetAmount(amount).
		SetCurrency(currency)
	_, err := client.CoinbaseAccountDeposit(opts)
	require.NoError(t, err)
}

func makeAccountWithdrawal(t *testing.T, client *C, currency string, amount float64) {
	opts := new(option.CoinbaseAccountWithdrawal).
		SetCoinbaseAccountId(findWalletID(t, client, currency)).
		SetAmount(amount).
		SetCurrency(currency)
	_, err := client.AccountWithdrawal(opts)
	require.NoError(t, err)
}

func makeCryptoWithdrawal(t *testing.T, client *C, currency string, amount float64) {
	opts := new(option.CoinbaseCryptoWithdrawal).
		SetCryptoAddress(generateCryptoAddress(t, client, currency)).
		SetAmount(amount).
		SetCurrency(currency)
	_, err := client.CryptoWithdrawal(opts)
	require.NoError(t, err)
}

func makeSimpleRequestAssertion(t *testing.T, name string, req func() error) {
	t.Run(name, func(t *testing.T) { require.NoError(t, req()) })
}
