package coinbase

import (
	"testing"

	"github.com/alpine-hodler/sdk/internal/env"
	"github.com/alpine-hodler/sdk/pkg/scalar"
	"github.com/stretchr/testify/require"
)

func TestSimpleIntegrations(t *testing.T) {
	env.Load(".simple-test.env")
	env.SetAlpineHodlerLogLevel("2")
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
	makeSimpleRequestAssertion(t, "Should GET client.Book", func() (err error) {
		_, err = client.Book(productID, new(BookOptions).SetLevel(1))
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.Candles", func() (err error) {
		candles, err := client.Candles(productID, new(CandlesOptions).SetGranularity(scalar.Seconds60))
		require.NotEmpty(t, candles)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.CurrencyConversion", func() (err error) {
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
	makeSimpleRequestAssertion(t, "Should GET client.Currency", func() (err error) {
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
		makeNewOrder(t, client, productID, findProfileByName(t, client, profileName))
		_, err = client.Orders(new(OrdersOptions).SetProductId(productID).SetLimit(1))
		_, cancellationError := client.CancelOpenOrders(nil)
		require.NoError(t, cancellationError)
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
	makeSimpleRequestAssertion(t, "Should GET client.ProductStats", func() (err error) {
		stats, err := client.ProductStats(productID)
		require.NotEmpty(t, stats)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.Profiles", func() (err error) {
		require.NotNil(t, findProfileByName(t, client, profileName))
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.Account", func() (err error) {
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
	makeSimpleRequestAssertion(t, "Should GET client.Profile", func() (err error) {
		_, err = client.Profile(findProfileByName(t, client, profileName).Id, nil)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.AccountTransfer", func() (err error) {
		transfer := makeAccountDeposit(t, client, "USD", 1.0)
		require.NotNil(t, transfer)

		_, err = client.AccountTransfer(transfer.Id)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.Reports", func() (err error) {
		_, err = client.Reports(nil)
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.RenameProfile", func() (err error) {
		// TODO: Fix {Message:Invalid scope Status:403 Forbidden StatusCode:Forbidden}

		// original := "Simple Test"
		// updated := "Simple Test 2: Electric Boogaloo"
		// profile := findProfileByName(t, client, original)
		// _, err = client.RenameProfile(profile.Id, new(RenameProfileOptions).SetName(updated).SetProfileId(updated))

		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.SignedPrices", func() (err error) {
		_, err = client.SignedPrices()
		return
	})
	makeSimpleRequestAssertion(t, "Should GET client.Trades", func() (err error) {
		trades, err := client.Trades(productID, nil)
		require.NotEmpty(t, trades)
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
	makeSimpleRequestAssertion(t, "Should POST client.CreateProfile", func() (err error) {
		// TODO: Fix {Message:Invalid scope Status:403 Forbidden StatusCode:Forbidden}
		// _, err = client.CreateProfile(new(CreateProfileOptions).SetName("test"))
		return
	})
	makeSimpleRequestAssertion(t, "Should POST client.CreateProfileTransfer", func() (err error) {
		// TODO: count the amount before and then after and make sure their difference is 1.
		profile1 := findProfileByName(t, client, profileName)
		profile2 := findProfileByName(t, client, "Simple Test")
		makeAccountDeposit(t, client, "USDC", 1.0)
		err = client.CreateProfileTransfer(new(CreateProfileTransferOptions).
			SetFrom(profile1.Id).
			SetTo(profile2.Id).
			SetCurrency("USD").
			SetAmount("1"))
		return
	})
	makeSimpleRequestAssertion(t, "Should POST client.CreateReport", func() (err error) {
		profile1 := findProfileByName(t, client, profileName)
		_, err = client.CreateReport(new(CreateReportOptions).
			SetType(scalar.ReportTypeFills).
			SetFormat(scalar.FormatPDF).
			SetStartDate("2020-02-25T03:11:42.164Z").
			SetProfileId(profile1.Id).
			SetProductId(productID).
			SetEndDate("2020-03-26T02:11:42.164Z"))
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

func makeNewOrder(t *testing.T, client *C, productID string, profile *Profile) *CreateOrder {
	opts := new(CreateOrderOptions).
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
	order, err := client.CreateOrder(opts)
	require.NoError(t, err)
	return order
}

func makeSimpleRequestAssertion(t *testing.T, name string, req func() error) {
	t.Run(name, func(t *testing.T) { require.NoError(t, req()) })
}
