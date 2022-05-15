package coinbase_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/alpine-hodler/sdk/internal/env"
	"github.com/alpine-hodler/sdk/pkg/coinbase"
	"github.com/alpine-hodler/sdk/pkg/scalar"
	"github.com/alpine-hodler/sdk/tools"
	"github.com/stretchr/testify/require"
)

var someAccountID string
var someCurrencyID string
var somePaymentMethod string
var somePaymentMethodID string
var someProfileID string
var someReportID string
var someUSDProfileID string
var someUSDCWalletID string
var someUSDWalletID string
var someUserID string

func TestExamples(t *testing.T) {
	defer tools.Quiet()()

	env.Load(".simple-test.env")

	// ! Make sure that these tests only run on the sandbox.
	env.SetCoinbaseProURL("https://api-public.sandbox.exchange.coinbase.com")
	env.SetAlpineHodlerLogLevel("2")
	var (
		client = coinbase.NewClient(coinbase.DefaultConnector)
	)

	// helper functions
	findAccount := func(t *testing.T, client *coinbase.C, currency string) *coinbase.Account {
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

	findAccountID := func(t *testing.T, client *coinbase.C, currency string) (accountID string) {
		account := findAccount(t, client, currency)
		require.NotEmpty(t, account)
		return account.ID
	}

	findWalletID := func(t *testing.T, client *coinbase.C, currency string) (walletID string) {
		wallets, err := client.Wallets()
		require.NoError(t, err)
		require.NotEmpty(t, wallets)
		for _, wallet := range wallets {
			if wallet.Currency == currency {
				walletID = wallet.ID
				break
			}
		}
		require.NotEmpty(t, walletID)
		return
	}

	findProfileID := func(t *testing.T, client *coinbase.C, currency string) string {
		account := findAccount(t, client, currency)
		require.NotEmpty(t, account)
		return account.ProfileID
	}

	findCurrencyID := func(t *testing.T, client *coinbase.C, currency string) string {
		currencies, err := client.Currencies()
		require.NoError(t, err)
		require.NotEmpty(t, currencies)
		for _, item := range currencies {
			if item.ID == currency {
				return item.ID
			}
		}
		return ""
	}

	findProfileByName := func(t *testing.T, client *coinbase.C, name string) *coinbase.Profile {
		profiles, err := client.Profiles(nil)
		require.NoError(t, err)
		for _, profile := range profiles {
			if profile.Name == name {
				return profile
			}
		}
		return nil
	}

	findPaymentMethod := func(t *testing.T, client *coinbase.C, name string) *coinbase.PaymentMethod {
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

	randomReport := func(t *testing.T, client *coinbase.C) *coinbase.Report {
		reports, _ := client.Reports(nil)
		require.NotEmpty(t, reports)
		return reports[0]
	}

	// Set globals
	someAccountID = findAccountID(t, client, "BTC")
	someCurrencyID = findCurrencyID(t, client, "USD")
	somePaymentMethod = "Chase:****4442"
	somePaymentMethodID = findPaymentMethod(t, client, somePaymentMethod).ID
	someProfileID = findProfileID(t, client, "BTC")
	someReportID = randomReport(t, client).ID
	someUSDProfileID = findProfileID(t, client, "USD")
	someUSDCWalletID = findWalletID(t, client, "USDC")
	someUSDWalletID = findWalletID(t, client, "USD")
	someUserID = findProfileByName(t, client, "default").UserID

	// Run Examples
	t.Run("C.Account", func(t *testing.T) { ExampleC_Account() })
	t.Run("C.AccountHolds", func(t *testing.T) { ExampleC_AccountHolds() })
	t.Run("C.AccountLedger", func(t *testing.T) { ExampleC_AccountLedger() })
	t.Run("C.AccountTransfers", func(t *testing.T) { ExampleC_AccountTransfers() })
	t.Run("C.Accounts", func(t *testing.T) { ExampleC_Accounts() })
	t.Run("C.Book", func(t *testing.T) { ExampleC_Book() })
	t.Run("C.CancelOpenOrders", func(t *testing.T) { ExampleC_CancelOpenOrders() })
	t.Run("C.CancelOrder", func(t *testing.T) { ExampleC_CancelOrder() })
	t.Run("C.Candles", func(t *testing.T) { ExampleC_Candles() })
	t.Run("C.CoinbaseAccountDeposit", func(t *testing.T) { ExampleC_CoinbaseAccountDeposit() })
	t.Run("C.CoinbaseAccountWithdrawal", func(t *testing.T) { ExampleC_CoinbaseAccountWithdrawal() })
	t.Run("C.ConvertCurrency", func(t *testing.T) { ExampleC_ConvertCurrency() })
	t.Run("C.CreateOrder", func(t *testing.T) { ExampleC_CreateOrder() })
	t.Run("C.CreateProfile", func(t *testing.T) { ExampleC_CreateProfile() })
	t.Run("C.CreateProfileTransfer", func(t *testing.T) { ExampleC_CreateProfileTransfer() })
	t.Run("C.CreateReport", func(t *testing.T) { ExampleC_CreateReport() })
	t.Run("C.CryptoWithdrawal", func(t *testing.T) { ExampleC_CryptoWithdrawal() })
	t.Run("C.Currencies", func(t *testing.T) { ExampleC_Currencies() })
	t.Run("C.Currency", func(t *testing.T) { ExampleC_Currency() })
	t.Run("C.DeleteProfile", func(t *testing.T) { ExampleC_DeleteProfile() })
	t.Run("C.ExchangeLimits", func(t *testing.T) { ExampleC_ExchangeLimits() })
	t.Run("C.Fees", func(t *testing.T) { ExampleC_Fees() })
	t.Run("C.Fills", func(t *testing.T) { ExampleC_Fills() })
	t.Run("C.GenerateCryptoAddress", func(t *testing.T) { ExampleC_GenerateCryptoAddress() })
	t.Run("C.Order", func(t *testing.T) { ExampleC_Order() })
	t.Run("C.Orders", func(t *testing.T) { ExampleC_Orders() })
	t.Run("C.PaymentDepositMethod", func(t *testing.T) { ExampleC_PaymentMethodDeposit() })
	t.Run("C.PaymentWithdrawalMethod", func(t *testing.T) { ExampleC_PaymentMethodWithdrawal() })
	t.Run("C.PaymentMethods", func(t *testing.T) { ExampleC_PaymentMethods() })
	t.Run("C.Product", func(t *testing.T) { ExampleC_Product() })
	t.Run("C.ProductStats", func(t *testing.T) { ExampleC_ProductStats() })
	t.Run("C.ProductTicker", func(t *testing.T) { ExampleC_ProductTicker() })
	t.Run("C.Products", func(t *testing.T) { ExampleC_Products() })
	t.Run("C.Profile", func(t *testing.T) { ExampleC_Profile() })
	t.Run("C.Profiles", func(t *testing.T) { ExampleC_Profiles() })
	t.Run("C.RenameProfile", func(t *testing.T) { ExampleC_RenameProfile() })
	t.Run("C.Report", func(t *testing.T) { ExampleC_Report() })
	t.Run("C.Reports", func(t *testing.T) { ExampleC_Reports() })
	t.Run("C.SignedPrices", func(t *testing.T) { ExampleC_SignedPrices() })
	t.Run("C.Trades", func(t *testing.T) { ExampleC_Trades() })
	t.Run("C.Transfer", func(t *testing.T) { ExampleC_Transfer() })
	t.Run("C.Transfers", func(t *testing.T) { ExampleC_Transfers() })
	t.Run("C.Wallets", func(t *testing.T) { ExampleC_Wallets() })
	t.Run("C.WithdrawalFeeEstimate", func(t *testing.T) { ExampleC_WithdrawalFeeEstimate() })
}

func ExampleC_Account() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	account, err := client.Account(someAccountID)
	if err != nil {
		log.Fatalf("Error fetching account: %v", err)
	}
	fmt.Printf("account: %+v\n", account)
}

func ExampleC_AccountHolds() {
	client := coinbase.NewClient(coinbase.DefaultConnector)

	// Create an order that probably will never get filled.
	order, err := client.CreateOrder(new(coinbase.CreateOrderOptions).
		SetProfileID(someProfileID).
		SetType(scalar.OrderTypeLimit).
		SetSide(scalar.OrderSideSell).
		SetSTP(scalar.OrderSTP_DC).
		SetStop(scalar.OrderStopLoss).
		SetTimeInForce(scalar.TimeInForceGTC).
		SetCancelAfter(scalar.CancelAfterMin).
		SetProductID("BTC-USD").
		SetStopPrice(1.0). // Very unlikely!
		SetSize(1.0).
		SetPrice(1.0))
	if err != nil {
		log.Fatal(err)
	}

	// Since the order above will forever be in an "active" status, we should
	// have a hold value waiting for us.
	holds, err := client.AccountHolds(someAccountID, new(coinbase.AccountHoldsOptions).
		SetLimit(1).
		SetBefore("2010-01-01").
		SetAfter("2023-01-01"))
	if err != nil {
		log.Fatalf("Error fetching holds: %v", err)
	}
	fmt.Printf("holds: %+v\n", holds)

	// Since this order will never be filled, we might as well cancel it.
	if _, err := client.CancelOrder(order.ID); err != nil {
		log.Fatalf("Error canceling order: %v", err)
	}
}

func ExampleC_AccountLedger() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	ledger, err := client.AccountLedger(someAccountID, new(coinbase.AccountLedgerOptions).
		SetStartDate("2010-01-01").
		SetEndDate("2023-01-01").
		SetBefore(1652574195).
		SetAfter(1526365354).
		SetLimit(1).
		SetProfileID(someProfileID))
	if err != nil {
		log.Fatalf("Error fetching ledger: %v", err)
	}
	fmt.Printf("ledger: %+v\n", ledger)
}

func ExampleC_AccountTransfers() {
	client := coinbase.NewClient(coinbase.DefaultConnector)

	// Make an account deposit to look up.
	_, err := client.CoinbaseAccountDeposit(new(coinbase.CoinbaseAccountDepositOptions).
		SetCoinbaseAccountID(someUSDWalletID).
		SetAmount(1.0).
		SetCurrency("USD"))
	if err != nil {
		log.Fatalf("Error making deposit: %v", err)
	}

	// Get the resulting Account Transfers.
	transfers, err := client.AccountTransfers(someAccountID, new(coinbase.AccountTransfersOptions).
		SetBefore("2010-01-01").
		SetAfter("2023-01-01").
		SetType(scalar.TransferMethodDeposit).
		SetLimit(1))
	if err != nil {
		log.Fatalf("Error fetching account transfers: %v", err)
	}
	fmt.Printf("account transfers: %+v\n", transfers)

	// Withdraw because we don't want to get too greedy :)
	_, err = client.CoinbaseAccountWithdrawal(new(coinbase.CoinbaseAccountWithdrawalOptions).
		SetCoinbaseAccountID(someUSDWalletID).
		SetAmount(1.0).
		SetCurrency("USD"))
	if err != nil {
		log.Fatalf("Error making withdrawal: %v", err)
	}
}

func ExampleC_Accounts() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	accounts, err := client.Accounts()
	if err != nil {
		log.Fatalf("Error fetching accounts: %v", err)
	}
	fmt.Printf("accounts: %+v\n", accounts)
}

func ExampleC_Book() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	book, err := client.Book("BTC-USD", new(coinbase.BookOptions).SetLevel(1))
	if err != nil {
		log.Fatalf("Error fetching book: %v", err)
	}
	fmt.Printf("book: %+v\n", book)
}

func ExampleC_CancelOpenOrders() {
	client := coinbase.NewClient(coinbase.DefaultConnector)

	// Create an order that probably will never get filled.
	_, err := client.CreateOrder(new(coinbase.CreateOrderOptions).
		SetProfileID(someProfileID).
		SetType(scalar.OrderTypeLimit).
		SetSide(scalar.OrderSideSell).
		SetSTP(scalar.OrderSTP_DC).
		SetStop(scalar.OrderStopLoss).
		SetTimeInForce(scalar.TimeInForceGTC).
		SetCancelAfter(scalar.CancelAfterMin).
		SetProductID("BTC-USD").
		SetStopPrice(1.0). // Very unlikely!
		SetSize(1.0).
		SetPrice(1.0))
	if err != nil {
		log.Fatal(err)
	}

	// Cancel the order by simply cancelling all BTC-USD orders for some profile ID.
	_, err = client.CancelOpenOrders(new(coinbase.CancelOpenOrdersOptions).
		SetProductID("BTC-USD").
		SetProfileID(someProfileID))
	if err != nil {
		log.Fatalf("Error canceling open orders: %v", err)
	}
}

func ExampleC_CancelOrder() {
	client := coinbase.NewClient(coinbase.DefaultConnector)

	// Create an order that probably will never get filled.
	order, err := client.CreateOrder(new(coinbase.CreateOrderOptions).
		SetProfileID(someProfileID).
		SetType(scalar.OrderTypeLimit).
		SetSide(scalar.OrderSideSell).
		SetSTP(scalar.OrderSTP_DC).
		SetStop(scalar.OrderStopLoss).
		SetTimeInForce(scalar.TimeInForceGTC).
		SetCancelAfter(scalar.CancelAfterMin).
		SetProductID("BTC-USD").
		SetStopPrice(1.0). // Very unlikely!
		SetSize(1.0).
		SetPrice(1.0))
	if err != nil {
		log.Fatal(err)
	}

	// Cancel the order by using the order's ID.
	_, err = client.CancelOrder(order.ID)
	if err != nil {
		log.Fatalf("Error canceling order: %v", err)
	}
}

func ExampleC_Candles() {
	client := coinbase.NewClient(coinbase.DefaultConnector)

	startTimestamp := time.Date(2018, 11, 9, 17, 11, 8, 0, time.UTC).Format(time.RFC3339)
	endTimestamp := time.Date(2018, 11, 9, 18, 11, 8, 0, time.UTC).Format(time.RFC3339)

	candles, err := client.Candles("BTC-USD", new(coinbase.CandlesOptions).
		SetStart(startTimestamp).
		SetEnd(endTimestamp).
		SetGranularity(scalar.Seconds60))
	if err != nil {
		log.Fatalf("Error canceling order: %v", err)
	}
	fmt.Printf("candles: %+v\n", candles)
}

func ExampleC_CoinbaseAccountDeposit() {
	client := coinbase.NewClient(coinbase.DefaultConnector)

	// Make the deposit.
	_, err := client.CoinbaseAccountDeposit(new(coinbase.CoinbaseAccountDepositOptions).
		SetProfileID(someProfileID).
		SetCoinbaseAccountID(someUSDWalletID).
		SetAmount(1).
		SetCurrency("USD"))
	if err != nil {
		log.Fatalf("Error making deposit: %v", err)
	}

	// Withdraw the deposit.
	client.CoinbaseAccountWithdrawal(new(coinbase.CoinbaseAccountWithdrawalOptions).
		SetProfileID(someProfileID).
		SetCoinbaseAccountID(someUSDWalletID).
		SetAmount(1).
		SetCurrency("USD"))
}

func ExampleC_CoinbaseAccountWithdrawal() {
	client := coinbase.NewClient(coinbase.DefaultConnector)

	// Make the deposit.
	client.CoinbaseAccountDeposit(new(coinbase.CoinbaseAccountDepositOptions).
		SetProfileID(someProfileID).
		SetCoinbaseAccountID(someUSDWalletID).
		SetAmount(1).
		SetCurrency("USD"))

	// Withdraw the deposit.
	_, err := client.CoinbaseAccountWithdrawal(new(coinbase.CoinbaseAccountWithdrawalOptions).
		SetProfileID(someProfileID).
		SetCoinbaseAccountID(someUSDWalletID).
		SetAmount(1).
		SetCurrency("USD"))
	if err != nil {
		log.Fatalf("Error making withdrawal: %v", err)
	}
}

func ExampleC_ConvertCurrency() {
	client := coinbase.NewClient(coinbase.DefaultConnector)

	// Make a deposit to convert into BTC.
	client.CoinbaseAccountDeposit(new(coinbase.CoinbaseAccountDepositOptions).
		SetProfileID(someProfileID).
		SetCoinbaseAccountID(someUSDWalletID).
		SetAmount(1).
		SetCurrency("USD"))

	// Convert the USD into USDC.
	_, err := client.ConvertCurrency(new(coinbase.ConvertCurrencyOptions).
		SetAmount(1).
		SetFrom("USD").
		SetTo("USDC").
		SetProfileID(someProfileID))
	if err != nil {
		log.Fatalf("Error converting currency: %v", err)
	}

	// Withdraw the USDC we just converted.
	client.CoinbaseAccountWithdrawal(new(coinbase.CoinbaseAccountWithdrawalOptions).
		SetProfileID(someProfileID).
		SetCoinbaseAccountID(someUSDWalletID).
		SetAmount(1).
		SetCurrency("USDC"))
}

func ExampleC_CreateOrder() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	order, err := client.CreateOrder(new(coinbase.CreateOrderOptions).
		SetProfileID(someProfileID).
		SetType(scalar.OrderTypeLimit).
		SetSide(scalar.OrderSideSell).
		SetSTP(scalar.OrderSTP_DC).
		SetStop(scalar.OrderStopLoss).
		SetTimeInForce(scalar.TimeInForceGTC).
		SetCancelAfter(scalar.CancelAfterMin).
		SetProductID("BTC-USD").
		SetStopPrice(1.0).
		SetSize(1.0).
		SetPrice(1.0))
	if err != nil {
		log.Fatal(err)
	}

	// Cancel the order since it will almost definitely never get filled.
	client.CancelOrder(order.ID)
}

func ExampleC_CreateProfile() {
	// TODO: Figure out why we get a 403 HTTP Status
}

func ExampleC_CreateProfileTransfer() {
	// TODO: Figure out why we get a 403 HTTP Status
}

func ExampleC_CreateReport() {
	startTimestamp := time.Date(2018, 11, 9, 17, 11, 8, 0, time.UTC).Format(time.RFC3339)
	endTimestamp := time.Date(2018, 11, 9, 18, 11, 8, 0, time.UTC).Format(time.RFC3339)

	client := coinbase.NewClient(coinbase.DefaultConnector)
	_, err := client.CreateReport(new(coinbase.CreateReportOptions).
		SetType(scalar.ReportTypeFills).
		SetFormat(scalar.FormatPDF).
		SetProfileID(someProfileID).
		SetProductID("BTC-USD").
		SetStartDate(startTimestamp).
		SetEndDate(endTimestamp))
	if err != nil {
		log.Fatalf("Error creating report: %v", err)
	}
}

func ExampleC_CryptoWithdrawal() {
	client := coinbase.NewClient(coinbase.DefaultConnector)

	// Post one 1 USDC into the target wallet.
	client.CoinbaseAccountDeposit(new(coinbase.CoinbaseAccountDepositOptions).
		SetCoinbaseAccountID(someUSDCWalletID).
		SetAmount(1).
		SetCurrency("USDC"))

	// Get an address for the target wallet.
	address, _ := client.GenerateCryptoAddress(someUSDCWalletID)

	// Withdraw the funds using the generated USDC wallet address.
	_, err := client.CryptoWithdrawal(new(coinbase.CryptoWithdrawalOptions).
		SetCryptoAddress(address.Address).
		SetAmount(1).
		SetCurrency("USDC"))
	if err != nil {
		log.Fatalf("Error withdrawing crypto: %v", err)
	}
}

func ExampleC_Currencies() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	currencies, err := client.Currencies()
	if err != nil {
		log.Fatalf("Error listing currencies: %v", err)
	}
	fmt.Printf("currencies: %+v\n", currencies)
}

func ExampleC_Currency() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	currency, err := client.Currency(someCurrencyID)
	if err != nil {
		log.Fatalf("Error fetching currency by ID: %v", err)
	}
	fmt.Printf("currency: %+v\n", currency)
}

func ExampleC_DeleteProfile() {
	// TODO: Figure out why we get a 403 HTTP Status
}

func ExampleC_ExchangeLimits() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	exchangeLimits, err := client.ExchangeLimits(someUserID)
	if err != nil {
		log.Fatalf("Error fetching exchange limits: %v", err)
	}
	fmt.Printf("exchange limits: %+v\n", exchangeLimits)
}

func ExampleC_Fees() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	fees, err := client.Fees()
	if err != nil {
		log.Fatalf("Error fetching fees: %v", err)
	}
	fmt.Printf("fees: %+v\n", fees)
}

func ExampleC_Fills() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	fills, err := client.Fills(new(coinbase.FillsOptions).
		SetAfter(1526365354).
		SetBefore(1652574195).
		SetLimit(1).
		SetProductID("BTC-USD").
		SetProfileID(someProfileID))
	if err != nil {
		log.Fatalf("Error fetching fills: %v", err)
	}
	fmt.Printf("fills: %+v\n", fills)
}

func ExampleC_GenerateCryptoAddress() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	address, err := client.GenerateCryptoAddress(someUSDCWalletID)
	if err != nil {
		log.Fatalf("Error fetching address: %v", err)
	}
	fmt.Printf("USD wallet address: %+v\n", address)
}

func ExampleC_Order() {
	client := coinbase.NewClient(coinbase.DefaultConnector)

	// Create an order that probably will never get filled.
	unfillableOrder, _ := client.CreateOrder(new(coinbase.CreateOrderOptions).
		SetProfileID(someProfileID).
		SetType(scalar.OrderTypeLimit).
		SetSide(scalar.OrderSideSell).
		SetSTP(scalar.OrderSTP_DC).
		SetStop(scalar.OrderStopLoss).
		SetTimeInForce(scalar.TimeInForceGTC).
		SetCancelAfter(scalar.CancelAfterMin).
		SetProductID("BTC-USD").
		SetStopPrice(1.0).
		SetSize(1.0).
		SetPrice(1.0))

	// Lookup the order.
	order, err := client.Order(unfillableOrder.ID)
	if err != nil {
		log.Fatalf("Error fetching order: %v", order)
	}
	fmt.Printf("order: %+v\n", order)

	// Cancel the order since it will almost definitely never get filled.
	client.CancelOrder(order.ID)
}

func ExampleC_Orders() {
	client := coinbase.NewClient(coinbase.DefaultConnector)

	// Create an order that probably will never get filled.
	unfillableOrder, _ := client.CreateOrder(new(coinbase.CreateOrderOptions).
		SetProfileID(someProfileID).
		SetType(scalar.OrderTypeLimit).
		SetSide(scalar.OrderSideSell).
		SetSTP(scalar.OrderSTP_DC).
		SetStop(scalar.OrderStopLoss).
		SetTimeInForce(scalar.TimeInForceGTC).
		SetCancelAfter(scalar.CancelAfterMin).
		SetProductID("BTC-USD").
		SetStopPrice(1.0).
		SetSize(1.0).
		SetPrice(1.0))

	startTimestamp := time.Date(2018, 11, 9, 17, 11, 8, 0, time.UTC).Format(time.RFC3339)
	endTimestamp := time.Date(2025, 11, 9, 18, 11, 8, 0, time.UTC).Format(time.RFC3339)

	// Lookup the order.
	order, err := client.Orders(new(coinbase.OrdersOptions).
		SetAfter("2023-01-01").
		SetBefore("2010-01-01").
		SetStartDate(startTimestamp).
		SetEndDate(endTimestamp).
		SetLimit(1).
		SetProductID("BTC-USD").
		SetProfileID(someProfileID))
	if err != nil {
		log.Fatalf("Error fetching order: %v", err)
	}
	fmt.Printf("order: %+v\n", order)

	// Cancel the order since it will almost definitely never get filled.
	client.CancelOrder(unfillableOrder.ID)
}

func ExampleC_PaymentMethodDeposit() {
	// TODO: Figure out why we get a 403 HTTP Status

	// client := coinbase.NewClient(coinbase.DefaultConnector)
	// _, err := client.PaymentMethodDeposit(new(coinbase.PaymentMethodDepositOptions).
	// 	SetAmount(100).
	// 	SetCurrency("USD").
	// 	SetPaymentMethodID(somePaymentMethodID).
	// 	SetProfileID(someUSDProfileID))
	// if err != nil {
	// 	log.Fatalf("Error depositing from payment method: %v", err)
	// }
}

func ExampleC_PaymentMethodWithdrawal() {
	// TODO: Figure out why we get a 403 HTTP Status
}

func ExampleC_PaymentMethods() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	paymentMethods, err := client.PaymentMethods()
	if err != nil {
		log.Fatalf("Error fetching payment methods: %v", err)
	}
	fmt.Printf("payment methods: %+v\n", paymentMethods)
}

func ExampleC_Product() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	product, err := client.Product("BTC-USD")
	if err != nil {
		log.Fatalf("Error fetching product: %v", err)
	}
	fmt.Printf("product: %+v\n", product)
}

func ExampleC_ProductStats() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	stats, err := client.ProductStats("BTC-USD")
	if err != nil {
		log.Fatalf("Error fetching stats: %v", err)
	}
	fmt.Printf("stats: %+v\n", stats)
}

func ExampleC_ProductTicker() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	ticker, err := client.ProductTicker("BTC-USD")
	if err != nil {
		log.Fatalf("Error fetching ticker: %v", err)
	}
	fmt.Printf("ticker: %+v\n", ticker)
}

func ExampleC_Products() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	products, err := client.Products(new(coinbase.ProductsOptions).SetType("USD-BTC"))
	if err != nil {
		log.Fatalf("Error fetching products: %v", err)
	}
	fmt.Printf("products: %+v\n", products)
}

func ExampleC_Profile() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	profile, err := client.Profile(someProfileID, new(coinbase.ProfileOptions).SetActive(true))
	if err != nil {
		log.Fatalf("Error fetching profile: %v", err)
	}
	fmt.Printf("profile: %+v\n", profile)
}

func ExampleC_Profiles() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	profiles, err := client.Profiles(new(coinbase.ProfilesOptions).SetActive(true))
	if err != nil {
		log.Fatalf("Error fetching profiles: %v", err)
	}
	fmt.Printf("profiles: %+v\n", profiles)
}

func ExampleC_RenameProfile() {
	// TODO: Figure out why we get a 403 HTTP Status
}

func ExampleC_Report() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	report, err := client.Report(someReportID)
	if err != nil {
		log.Fatalf("Error fetching report: %v", err)
	}
	fmt.Printf("report: %+v\n", report)
}

func ExampleC_Reports() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	afterTimestamp := time.Date(2018, 11, 9, 17, 11, 8, 0, time.UTC).Format(time.RFC3339)

	reports, err := client.Reports(new(coinbase.ReportsOptions).
		SetAfter(afterTimestamp).
		SetIgnoredExpired(true).
		SetLimit(1).
		SetPortfolioID(someProfileID).
		SetType(scalar.ReportTypeAcccount))
	if err != nil {
		log.Fatalf("Error fetching reports: %v", err)
	}
	fmt.Printf("reports: %+v\n", reports)
}

func ExampleC_SignedPrices() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	prices, err := client.SignedPrices()
	if err != nil {
		log.Fatalf("Error fetching prices: %v", err)
	}
	fmt.Printf("prices: %+v\n", prices)
}

func ExampleC_Trades() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	trades, err := client.Trades("BTC-USD", new(coinbase.TradesOptions).
		SetAfter(1526365354).
		SetBefore(1652574165).
		SetLimit(1))
	if err != nil {
		log.Fatalf("Error fetching prices: %v", err)
	}
	fmt.Printf("prices: %+v\n", trades)
}

func ExampleC_Transfer() {
	client := coinbase.NewClient(coinbase.DefaultConnector)

	// Make a deposit.
	deposit, err := client.CoinbaseAccountDeposit(new(coinbase.CoinbaseAccountDepositOptions).
		SetCoinbaseAccountID(someUSDWalletID).
		SetAmount(1).
		SetCurrency("USD"))
	if err != nil {
		log.Fatalf("Error making deposit: %v", err)
	}

	// Lookup the deposit.
	transfer, err := client.Transfer(deposit.ID)
	if err != nil {
		log.Fatalf("Error fetching transfer: %v", err)
	}
	fmt.Printf("transfer: %+v\n", transfer)
}

func ExampleC_Transfers() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	transfers, err := client.Transfers()
	if err != nil {
		log.Fatalf("Error fetching transfers: %v", err)
	}
	fmt.Printf("transfers: %+v\n", transfers)
}

func ExampleC_Wallets() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	wallets, err := client.Wallets()
	if err != nil {
		log.Fatalf("Error fetching wallets: %v", err)
	}
	fmt.Printf("wallets: %+v\n", wallets)
}

func ExampleC_WithdrawalFeeEstimate() {
	client := coinbase.NewClient(coinbase.DefaultConnector)
	address, _ := client.GenerateCryptoAddress(someUSDCWalletID)

	estimates, err := client.WithdrawalFeeEstimate(new(coinbase.WithdrawalFeeEstimateOptions).
		SetCryptoAddress(address.Address).
		SetCurrency("USDC"))
	if err != nil {
		log.Fatalf("Error fetching estimates: %v", err)
	}
	fmt.Printf("estimates: %+v\n", estimates)
}
