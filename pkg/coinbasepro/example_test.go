package coinbasepro_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/alpine-hodler/sdk/internal/env"
	"github.com/alpine-hodler/sdk/pkg/coinbasepro"
	"github.com/alpine-hodler/sdk/pkg/scalar"
	"github.com/alpine-hodler/sdk/pkg/websocket"
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
		client = coinbasepro.NewClient(coinbasepro.DefaultConnector)
	)

	// helper functions
	findAccount := func(t *testing.T, client *coinbasepro.C, currency string) *coinbasepro.Account {
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

	findAccountID := func(t *testing.T, client *coinbasepro.C, currency string) (accountID string) {
		account := findAccount(t, client, currency)
		require.NotEmpty(t, account)
		return account.ID
	}

	findWalletID := func(t *testing.T, client *coinbasepro.C, currency string) (walletID string) {
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

	findProfileID := func(t *testing.T, client *coinbasepro.C, currency string) string {
		account := findAccount(t, client, currency)
		require.NotEmpty(t, account)
		return account.ProfileID
	}

	findCurrencyID := func(t *testing.T, client *coinbasepro.C, currency string) string {
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

	findProfileByName := func(t *testing.T, client *coinbasepro.C, name string) *coinbasepro.Profile {
		profiles, err := client.Profiles(nil)
		require.NoError(t, err)
		for _, profile := range profiles {
			if profile.Name == name {
				return profile
			}
		}
		return nil
	}

	findPaymentMethod := func(t *testing.T, client *coinbasepro.C, name string) *coinbasepro.PaymentMethod {
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

	randomReport := func(t *testing.T, client *coinbasepro.C) *coinbasepro.Report {
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
	t.Run("ProductWebsocket.Ticker", func(t *testing.T) { ExampleProductWebsocket_Ticker() })
}

func ExampleC_Account() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	account, err := client.Account(someAccountID)
	if err != nil {
		log.Fatalf("Error fetching account: %v", err)
	}
	fmt.Printf("account: %+v\n", account)
}

func ExampleC_AccountHolds() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)

	// Create an order that probably will never get filled.
	order, err := client.CreateOrder(new(coinbasepro.CreateOrderOptions).
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
	holds, err := client.AccountHolds(someAccountID, new(coinbasepro.AccountHoldsOptions).
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
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	ledger, err := client.AccountLedger(someAccountID, new(coinbasepro.AccountLedgerOptions).
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
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)

	// Make an account deposit to look up.
	_, err := client.CoinbaseAccountDeposit(new(coinbasepro.CoinbaseAccountDepositOptions).
		SetCoinbaseAccountID(someUSDWalletID).
		SetAmount(1.0).
		SetCurrency("USD"))
	if err != nil {
		log.Fatalf("Error making deposit: %v", err)
	}

	// Get the resulting Account Transfers.
	transfers, err := client.AccountTransfers(someAccountID, new(coinbasepro.AccountTransfersOptions).
		SetBefore("2010-01-01").
		SetAfter("2023-01-01").
		SetType(scalar.TransferMethodDeposit).
		SetLimit(1))
	if err != nil {
		log.Fatalf("Error fetching account transfers: %v", err)
	}
	fmt.Printf("account transfers: %+v\n", transfers)

	// Withdraw because we don't want to get too greedy :)
	_, err = client.CoinbaseAccountWithdrawal(new(coinbasepro.CoinbaseAccountWithdrawalOptions).
		SetCoinbaseAccountID(someUSDWalletID).
		SetAmount(1.0).
		SetCurrency("USD"))
	if err != nil {
		log.Fatalf("Error making withdrawal: %v", err)
	}
}

func ExampleC_Accounts() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	accounts, err := client.Accounts()
	if err != nil {
		log.Fatalf("Error fetching accounts: %v", err)
	}
	fmt.Printf("accounts: %+v\n", accounts)
}

func ExampleC_Book() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	book, err := client.Book("BTC-USD", new(coinbasepro.BookOptions).SetLevel(1))
	if err != nil {
		log.Fatalf("Error fetching book: %v", err)
	}
	fmt.Printf("book: %+v\n", book)
}

func ExampleC_CancelOpenOrders() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)

	// Create an order that probably will never get filled.
	_, err := client.CreateOrder(new(coinbasepro.CreateOrderOptions).
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
	_, err = client.CancelOpenOrders(new(coinbasepro.CancelOpenOrdersOptions).
		SetProductID("BTC-USD").
		SetProfileID(someProfileID))
	if err != nil {
		log.Fatalf("Error canceling open orders: %v", err)
	}
}

func ExampleC_CancelOrder() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)

	// Create an order that probably will never get filled.
	order, err := client.CreateOrder(new(coinbasepro.CreateOrderOptions).
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
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)

	startTimestamp := time.Date(2018, 11, 9, 17, 11, 8, 0, time.UTC).Format(time.RFC3339)
	endTimestamp := time.Date(2018, 11, 9, 18, 11, 8, 0, time.UTC).Format(time.RFC3339)

	candles, err := client.Candles("BTC-USD", new(coinbasepro.CandlesOptions).
		SetStart(startTimestamp).
		SetEnd(endTimestamp).
		SetGranularity(scalar.Seconds60))
	if err != nil {
		log.Fatalf("Error canceling order: %v", err)
	}
	fmt.Printf("candles: %+v\n", candles)
}

func ExampleC_CoinbaseAccountDeposit() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)

	// Make the deposit.
	_, err := client.CoinbaseAccountDeposit(new(coinbasepro.CoinbaseAccountDepositOptions).
		SetProfileID(someProfileID).
		SetCoinbaseAccountID(someUSDWalletID).
		SetAmount(1).
		SetCurrency("USD"))
	if err != nil {
		log.Fatalf("Error making deposit: %v", err)
	}

	// Withdraw the deposit.
	client.CoinbaseAccountWithdrawal(new(coinbasepro.CoinbaseAccountWithdrawalOptions).
		SetProfileID(someProfileID).
		SetCoinbaseAccountID(someUSDWalletID).
		SetAmount(1).
		SetCurrency("USD"))
}

func ExampleC_CoinbaseAccountWithdrawal() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)

	// Make the deposit.
	client.CoinbaseAccountDeposit(new(coinbasepro.CoinbaseAccountDepositOptions).
		SetProfileID(someProfileID).
		SetCoinbaseAccountID(someUSDWalletID).
		SetAmount(1).
		SetCurrency("USD"))

	// Withdraw the deposit.
	_, err := client.CoinbaseAccountWithdrawal(new(coinbasepro.CoinbaseAccountWithdrawalOptions).
		SetProfileID(someProfileID).
		SetCoinbaseAccountID(someUSDWalletID).
		SetAmount(1).
		SetCurrency("USD"))
	if err != nil {
		log.Fatalf("Error making withdrawal: %v", err)
	}
}

func ExampleC_ConvertCurrency() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)

	// Make a deposit to convert into BTC.
	client.CoinbaseAccountDeposit(new(coinbasepro.CoinbaseAccountDepositOptions).
		SetProfileID(someProfileID).
		SetCoinbaseAccountID(someUSDWalletID).
		SetAmount(1).
		SetCurrency("USD"))

	// Convert the USD into USDC.
	_, err := client.ConvertCurrency(new(coinbasepro.ConvertCurrencyOptions).
		SetAmount(1).
		SetFrom("USD").
		SetTo("USDC").
		SetProfileID(someProfileID))
	if err != nil {
		log.Fatalf("Error converting currency: %v", err)
	}

	// Withdraw the USDC we just converted.
	client.CoinbaseAccountWithdrawal(new(coinbasepro.CoinbaseAccountWithdrawalOptions).
		SetProfileID(someProfileID).
		SetCoinbaseAccountID(someUSDWalletID).
		SetAmount(1).
		SetCurrency("USDC"))
}

func ExampleC_CreateOrder() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	order, err := client.CreateOrder(new(coinbasepro.CreateOrderOptions).
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

	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	_, err := client.CreateReport(new(coinbasepro.CreateReportOptions).
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
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)

	// Post one 1 USDC into the target wallet.
	client.CoinbaseAccountDeposit(new(coinbasepro.CoinbaseAccountDepositOptions).
		SetCoinbaseAccountID(someUSDCWalletID).
		SetAmount(1).
		SetCurrency("USDC"))

	// Get an address for the target wallet.
	address, _ := client.GenerateCryptoAddress(someUSDCWalletID)

	// Withdraw the funds using the generated USDC wallet address.
	_, err := client.CryptoWithdrawal(new(coinbasepro.CryptoWithdrawalOptions).
		SetCryptoAddress(address.Address).
		SetAmount(1).
		SetCurrency("USDC"))
	if err != nil {
		log.Fatalf("Error withdrawing crypto: %v", err)
	}
}

func ExampleC_Currencies() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	currencies, err := client.Currencies()
	if err != nil {
		log.Fatalf("Error listing currencies: %v", err)
	}
	fmt.Printf("currencies: %+v\n", currencies)
}

func ExampleC_Currency() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
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
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	exchangeLimits, err := client.ExchangeLimits(someUserID)
	if err != nil {
		log.Fatalf("Error fetching exchange limits: %v", err)
	}
	fmt.Printf("exchange limits: %+v\n", exchangeLimits)
}

func ExampleC_Fees() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	fees, err := client.Fees()
	if err != nil {
		log.Fatalf("Error fetching fees: %v", err)
	}
	fmt.Printf("fees: %+v\n", fees)
}

func ExampleC_Fills() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	fills, err := client.Fills(new(coinbasepro.FillsOptions).
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
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	address, err := client.GenerateCryptoAddress(someUSDCWalletID)
	if err != nil {
		log.Fatalf("Error fetching address: %v", err)
	}
	fmt.Printf("USD wallet address: %+v\n", address)
}

func ExampleC_Order() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)

	// Create an order that probably will never get filled.
	unfillableOrder, _ := client.CreateOrder(new(coinbasepro.CreateOrderOptions).
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
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)

	// Create an order that probably will never get filled.
	unfillableOrder, _ := client.CreateOrder(new(coinbasepro.CreateOrderOptions).
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
	order, err := client.Orders(new(coinbasepro.OrdersOptions).
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

	// client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	// _, err := client.PaymentMethodDeposit(new(coinbasepro.PaymentMethodDepositOptions).
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
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	paymentMethods, err := client.PaymentMethods()
	if err != nil {
		log.Fatalf("Error fetching payment methods: %v", err)
	}
	fmt.Printf("payment methods: %+v\n", paymentMethods)
}

func ExampleC_Product() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	product, err := client.Product("BTC-USD")
	if err != nil {
		log.Fatalf("Error fetching product: %v", err)
	}
	fmt.Printf("product: %+v\n", product)
}

func ExampleC_ProductStats() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	stats, err := client.ProductStats("BTC-USD")
	if err != nil {
		log.Fatalf("Error fetching stats: %v", err)
	}
	fmt.Printf("stats: %+v\n", stats)
}

func ExampleC_ProductTicker() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	ticker, err := client.ProductTicker("BTC-USD")
	if err != nil {
		log.Fatalf("Error fetching ticker: %v", err)
	}
	fmt.Printf("ticker: %+v\n", ticker)
}

func ExampleC_Products() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	products, err := client.Products(new(coinbasepro.ProductsOptions).SetType("USD-BTC"))
	if err != nil {
		log.Fatalf("Error fetching products: %v", err)
	}
	fmt.Printf("products: %+v\n", products)
}

func ExampleC_Profile() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	profile, err := client.Profile(someProfileID, new(coinbasepro.ProfileOptions).SetActive(true))
	if err != nil {
		log.Fatalf("Error fetching profile: %v", err)
	}
	fmt.Printf("profile: %+v\n", profile)
}

func ExampleC_Profiles() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	profiles, err := client.Profiles(new(coinbasepro.ProfilesOptions).SetActive(true))
	if err != nil {
		log.Fatalf("Error fetching profiles: %v", err)
	}
	fmt.Printf("profiles: %+v\n", profiles)
}

func ExampleC_RenameProfile() {
	// TODO: Figure out why we get a 403 HTTP Status
}

func ExampleC_Report() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	report, err := client.Report(someReportID)
	if err != nil {
		log.Fatalf("Error fetching report: %v", err)
	}
	fmt.Printf("report: %+v\n", report)
}

func ExampleC_Reports() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	afterTimestamp := time.Date(2018, 11, 9, 17, 11, 8, 0, time.UTC).Format(time.RFC3339)

	reports, err := client.Reports(new(coinbasepro.ReportsOptions).
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
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	prices, err := client.SignedPrices()
	if err != nil {
		log.Fatalf("Error fetching prices: %v", err)
	}
	fmt.Printf("prices: %+v\n", prices)
}

func ExampleC_Trades() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	trades, err := client.Trades("BTC-USD", new(coinbasepro.TradesOptions).
		SetAfter(1526365354).
		SetBefore(1652574165).
		SetLimit(1))
	if err != nil {
		log.Fatalf("Error fetching prices: %v", err)
	}
	fmt.Printf("prices: %+v\n", trades)
}

func ExampleC_Transfer() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)

	// Make a deposit.
	deposit, err := client.CoinbaseAccountDeposit(new(coinbasepro.CoinbaseAccountDepositOptions).
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
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	transfers, err := client.Transfers()
	if err != nil {
		log.Fatalf("Error fetching transfers: %v", err)
	}
	fmt.Printf("transfers: %+v\n", transfers)
}

func ExampleC_Wallets() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	wallets, err := client.Wallets()
	if err != nil {
		log.Fatalf("Error fetching wallets: %v", err)
	}
	fmt.Printf("wallets: %+v\n", wallets)
}

func ExampleC_WithdrawalFeeEstimate() {
	client := coinbasepro.NewClient(coinbasepro.DefaultConnector)
	address, _ := client.GenerateCryptoAddress(someUSDCWalletID)

	estimates, err := client.WithdrawalFeeEstimate(new(coinbasepro.WithdrawalFeeEstimateOptions).
		SetCryptoAddress(address.Address).
		SetCurrency("USDC"))
	if err != nil {
		log.Fatalf("Error fetching estimates: %v", err)
	}
	fmt.Printf("estimates: %+v\n", estimates)
}

func ExampleProductWebsocket_Ticker() {
	ws := coinbasepro.NewWebsocket(websocket.DefaultConnector)

	// initialize the ticker object to channel product messages
	ticker := ws.Ticker("ETH-USD")

	// start a go routine that passes product messages concerning ETH-USD currency pair to a channel on the ticker struct.
	ticker.Open()
	go func() {
		// Next we range over the product message channel and print the product messages.
		for productMsg := range ticker.Channel() {
			fmt.Printf("ETH-USD Price @ %v: %v\n", productMsg.Time, productMsg.Price)
		}
	}()

	// Let the product messages print for 5 seconds.
	time.Sleep(5 * time.Second)

	// Then close the ticker channel, this will unsubscribe from the websocket and close the underlying channel that the
	// messages read to.
	ticker.Close()
}
