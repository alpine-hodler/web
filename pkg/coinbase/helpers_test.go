package coinbase

import (
	"testing"

	"github.com/alpine-hodler/sdk/pkg/scalar"
	"github.com/stretchr/testify/require"
)

func convertCurrency(t *testing.T, client *C, from, to string, amount float64) *CurrencyConversion {
	opts := new(ConvertCurrencyOptions).
		SetProfileId(findProfileID(t, client, from)).
		SetFrom(from).
		SetTo(to).
		SetAmount(amount)
	conversion, err := client.ConvertCurrency(opts)
	require.NoError(t, err)
	return conversion
}

func findAccount(t *testing.T, client *C, currency string) *Account {
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

func findCurrency(t *testing.T, client *C, currency string) *Currency {
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

func findPaymentMethod(t *testing.T, client *C, name string) *PaymentMethod {
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

func findProfileByName(t *testing.T, client *C, name string) *Profile {
	profiles, err := client.Profiles(nil)
	require.NoError(t, err)
	for _, profile := range profiles {
		if profile.Name == name {
			return profile
		}
	}
	return nil
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

func makeAccountDeposit(t *testing.T, client *C, currency string, amount float64) *Deposit {
	opts := new(CoinbaseAccountDepositOptions).
		SetCoinbaseAccountId(findWalletID(t, client, currency)).
		SetAmount(amount).
		SetCurrency(currency)
	deposit, err := client.CoinbaseAccountDeposit(opts)
	require.NoError(t, err)
	return deposit
}

func makeAccountWithdrawal(t *testing.T, client *C, currency string, amount float64) {
	opts := new(AccountWithdrawalOptions).
		SetCoinbaseAccountId(findWalletID(t, client, currency)).
		SetAmount(amount).
		SetCurrency(currency)
	_, err := client.AccountWithdrawal(opts)
	require.NoError(t, err)
}

func makeCryptoWithdrawal(t *testing.T, client *C, currency string, amount float64) {
	opts := new(CryptoWithdrawalOptions).
		SetCryptoAddress(generateCryptoAddress(t, client, currency)).
		SetAmount(amount).
		SetCurrency(currency)
	_, err := client.CryptoWithdrawal(opts)
	require.NoError(t, err)
}

func makeNewOrder(t *testing.T, client *C, productID string, profile *Profile) *NewOrder {
	opts := new(NewOrderOptions).
		SetProfileId(profile.Id).
		SetType(scalar.OrderTypeLimit).
		SetSide(scalar.OrderSideSell).
		SetStp(scalar.OrderSTP_DC).
		SetStop(scalar.OrderStopLoss).
		SetTimeInForce(scalar.TimeInForceGTC).
		SetCancelAfter(scalar.CancelAfterMin).
		SetProductId(productID).
		SetStopPrice(1.0).
		SetSize(1.0).
		SetPrice(1.0)
	order, err := client.NewOrder(opts)
	require.NoError(t, err)
	return order
}

func makeSimpleRequestAssertion(t *testing.T, name string, req func() error) {
	t.Run(name, func(t *testing.T) { require.NoError(t, req()) })
}
