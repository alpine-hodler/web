package coinbase

import (
	"strconv"
	"strings"

	"github.com/alpine-hodler/sdk/internal/client"
	"github.com/alpine-hodler/sdk/pkg/model"
	"github.com/alpine-hodler/sdk/tools"
)

// Accounts lists all trading accounts from the profile of the API key.
func (coinbaseClient *C) Accounts() (m []*model.CoinbaseAccount, err error) {
	req := coinbaseClient.Get(AccountsEndpoint)
	return m, req.Fetch().Assign(&m).JoinMessages()
}

// AccountHolds returns a list of holds of an account that belong to the same
// profile as the API key. Holds are placed for any active orders or pending
// withdraw requests. As an order is filled, the hold amount is updated.
func (coinbaseClient *C) AccountHolds(accountId string, opts *model.CoinbaseAccountHoldsOptions) (m []*model.CoinbaseAccountHold, err error) {
	return m, coinbaseClient.Get(AccountHoldsEndpoint).
		PathParam("account_id", accountId).
		QueryParam("before", func() (i *string) {
			if opts != nil {
				i = opts.Before
			}
			return
		}()).
		QueryParam("after", func() (i *string) {
			if opts != nil {
				i = opts.After
			}
			return
		}()).
		QueryParam("limit", func() (i *string) {
			if opts != nil {
				i = tools.IntPtrStringPtr(opts.Limit)
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// AccountLedger lists ledger activity for an account. This includes anything
// that would affect the accounts balance - transfers, trades, fees, etc
func (coinbaseClient *C) AccountLedger(accountId string, opts *model.CoinbaseAccountLedgerOptions) (m []*model.CoinbaseAccountLedger, err error) {
	return m, coinbaseClient.Get(AccountLedgerEndpoint).
		PathParam("account_id", accountId).
		QueryParam("before", func() (i *string) {
			if opts != nil {
				i = opts.Before
			}
			return
		}()).
		QueryParam("after", func() (i *string) {
			if opts != nil {
				i = opts.After
			}
			return
		}()).
		QueryParam("start_date", func() (i *string) {
			if opts != nil {
				i = opts.StartDate
			}
			return
		}()).
		QueryParam("end_date", func() (i *string) {
			if opts != nil {
				i = opts.EndDate
			}
			return
		}()).
		QueryParam("limit", func() (i *string) {
			if opts != nil {
				i = tools.IntPtrStringPtr(opts.Limit)
			}
			return
		}()).
		QueryParam("profile_id", func() (i *string) {
			if opts != nil {
				i = opts.ProfileID
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// AccountTransfers lists past withdrawals and deposits for an account.
func (coinbaseClient *C) AccountTransfers(accountId string, opts *model.CoinbaseAccountTransferOptions) (m []*model.CoinbaseAccountTransfer, err error) {
	return m, coinbaseClient.Get(AccountTransfersEndpoint).
		PathParam("account_id", accountId).
		QueryParam("before", func() (i *string) {
			if opts != nil {
				i = opts.Before
			}
			return
		}()).
		QueryParam("after", func() (i *string) {
			if opts != nil {
				i = opts.After
			}
			return
		}()).
		QueryParam("limit", func() (i *string) {
			if opts != nil {
				i = tools.IntPtrStringPtr(opts.Limit)
			}
			return
		}()).
		QueryParam("limit", func() (i *string) {
			if opts != nil {
				i = tools.IntPtrStringPtr(opts.Limit)
			}
			return
		}()).
		QueryParam("type", func() (i *string) {
			if opts != nil {
				i = opts.After
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// AccountWithdrawal Withdraws funds from the specified profile_id to a
// www.coinbase.com wallet.
func (coinbaseClient *C) AccountWithdrawal(opts *model.CoinbaseAccountWithdrawalOptions) (m *model.CoinbaseWithdrawal, err error) {
	return m, coinbaseClient.Post(AccountWithdrawalEndpoint).
		Body(client.NewBody(client.BodyTypeJSON).
			SetString("profile_id", opts.ProfileID).
			SetFloat("amount", &opts.Amount).
			SetString("coinbase_account_id", &opts.CoinbaseAccountID).
			SetString("currency", &opts.Currency)).
		Fetch().Assign(&m).JoinMessages()
}

// CryptoWithdrawal withdraws funds from the specified profile_id to an external
// crypto address
func (coinbaseClient *C) CryptoWithdrawal(opts *model.CoinbaseCryptoWithdrawalOptions) (m *model.CoinbaseWithdrawal, err error) {
	return m, coinbaseClient.Post(CryptoWithdrawalEndpoint).
		Body(client.NewBody(client.BodyTypeJSON).
			SetString("profile_id", opts.ProfileID).
			SetFloat("amount", &opts.Amount).
			SetString("crypto_address", &opts.CryptoAddress).
			SetString("currency", &opts.Currency).
			SetString("destination_tag", opts.DestinationTag).
			SetBool("no_destination_tag", opts.NoDestinationTag).
			SetString("two_factor_code", opts.TwoFactorCode).
			SetInt("nonce", opts.Nonce).
			SetFloat("fee", opts.Fee)).
		Fetch().Assign(&m).JoinMessages()
}

// CoinbaseAccountDeposit will deposit funds from a www.coinbase.com wallet to
// the specified profile_id.
func (coinbaseClient *C) CoinbaseAccountDeposit(opts *model.CoinbaseAccountDepositOptions) (m *model.CoinbaseDeposit, err error) {
	return m, coinbaseClient.Post(AccountDepositEndpoint).
		Body(client.NewBody(client.BodyTypeJSON).
			SetString("profile_id", opts.ProfileID).
			SetString("coinbase_account_id", &opts.CoinbaseAccountID).
			SetString("currency", &opts.Currency).
			SetFloat("amount", &opts.Amount)).
		Fetch().Assign(&m).JoinMessages()
}

// CanceCancelOpenOrderslAll will with best effort, cancel all open orders. This
// may require you to make the request multiple times until all of the open
// orders are deleted.
func (coinbaseClient *C) CancelOpenOrders(opts *model.CoinbaseOrdersOptions) (m []*string, err error) {
	return m, coinbaseClient.Delete(OrdersEndpoint).
		QueryParam("profile_id", func() (i *string) {
			if opts != nil {
				i = opts.ProfileID
			}
			return
		}()).
		QueryParam("product_id", func() (i *string) {
			if opts != nil && opts.ProductID != nil {
				i = opts.ProductID
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// CancelOrder will cancel a single open order by {id}.
func (coinbaseClient *C) CancelOrder(orderID string) (str string, err error) {
	return str, coinbaseClient.Delete(OrderEndpoint).Fetch().Assign(&str).JoinMessages()
}

// CreateOrder will create an order. You can place two types of orders: limit
func (coinbaseClient *C) CreateOrder(opts *model.CoinbaseNewOrderOptions) (m *model.CoinbaseNewOrder, err error) {
	return m, coinbaseClient.Post(NewOrderEndpoint).
		Body(client.NewBody(client.BodyTypeJSON).
			SetString("profile_id", opts.ProfileID).
			SetString("product_id", &opts.ProductID).
			SetFloat("stop_price", opts.StopPrice).
			SetFloat("size", opts.Size).
			SetFloat("price", opts.Price).
			SetFloat("funds", opts.Funds).
			SetBool("post_only", opts.PostOnly).
			SetString("client_oid", opts.ClientOid).
			SetString("type", func() *string { str := opts.Type.String(); return &str }()).
			SetString("side", func() *string { str := string(opts.Side); return &str }()).
			SetString("stp", func() *string { str := opts.Stp.String(); return &str }()).
			SetString("stop", func() *string { str := opts.Stop.String(); return &str }()).
			SetString("time_in_force", func() *string { str := opts.TimeInForce.String(); return &str }()).
			SetString("cancel_after", func() *string { str := opts.CancelAfter.String(); return &str }())).
		Fetch().Assign(&m).JoinMessages()
}

// Convert converts funds from from currency to to currency. Funds are
// converted on the from account in the profile_id profile.
func (coinbaseClient *C) Convert(opts model.CoinbaseConversionsOptions) (m *model.CoinbaseCurrencyConversion, err error) {
	return m, coinbaseClient.Post(ConversionsEndpoint).
		Body(client.NewBody(client.BodyTypeJSON).
			SetString("from", &opts.From).
			SetString("to", &opts.To).
			SetFloat("amount", &opts.Amount).
			SetString("profile_id", func() (s *string) {
				if opts.ProfileID != nil {
					s = opts.ProfileID
				}
				return
			}()).
			SetString("nonce", func() (s *string) {
				if opts.Nonce != nil {
					s = opts.Nonce
				}
				return
			}())).
		Fetch().Assign(&m).JoinMessages()
}

// Currencies gets a list of all known currencies
func (coinbaseClient *C) Currencies() (m []*model.CoinbaseCurrency, err error) {
	return m, coinbaseClient.Get(CurrenciesEndpoint).Fetch().Assign(&m).JoinMessages()
}

// FindAccount returns information for a single account. Use this endpoint when you
// know the account_id. API key must belong to the same profile as the account.
func (coinbaseClient *C) FindAccount(accountId string) (m *model.CoinbaseAccount, err error) {
	req := coinbaseClient.Get(AccountEndpoint)
	return m, req.PathParam("account_id", accountId).Fetch().Assign(&m).JoinMessages()
}

// GenerateCryptoAddress will generates a one-time crypto address for depositing
// crypto.
func (coinbaseClient *C) GenerateCryptoAddress(walletId string) (m *model.CoinbaseCryptoAddress, err error) {
	req := coinbaseClient.Post(AddressesEndpoint)
	return m, req.PathParam("account_id", walletId).Fetch().Assign(&m).JoinMessages()
}

// Fees will get fees rates and 30 days trailing volume.
func (coinbaseClient *C) Fees() (m *model.CoinbaseFees, err error) {
	req := coinbaseClient.Get(FeesEndpoint)
	return m, req.Fetch().Assign(&m).JoinMessages()
}

// ## API Key Permissions
// This endpoint requires either the "view" or "trade" permission.
//
// ## Settlement and Fees
// Fees are recorded in two stages. Immediately after the matching engine
// completes a match, the fill is inserted into our datastore. Once the fill is
// recorded, a settlement process will settle the fill and credit both trading
// counterparties.
//
// The fee field indicates the fees charged for this individual fill.
//
// ### Liquidity
// The liquidity field indicates if the fill was the result of a liquidity
// provider or liquidity taker. M indicates Maker and T indicates Taker.
//
// ### Pagination
// Fills are returned sorted by descending trade_id from the largest trade_id to
// the smallest trade_id. The CB-BEFORE header will have this first trade id so
// that future requests using the cb-before parameter will fetch fills with a
// greater trade id (newer fills).
func (coinbaseClient *C) Fills(opts *model.CoinbaseFillsOptions) (m []*model.CoinbaseFill, err error) {
	return m, coinbaseClient.Get(FillsEndpoint).
		QueryParam("order_id", func() (i *string) {
			if opts != nil {
				i = opts.OrderID
			}
			return
		}()).
		QueryParam("product_id", func() (i *string) {
			if opts != nil {
				i = opts.ProductID
			}
			return
		}()).
		QueryParam("profile_id", func() (i *string) {
			if opts != nil {
				i = opts.ProfileID
			}
			return
		}()).
		QueryParam("limit", func() (i *string) {
			if opts != nil && opts.Limit != nil {
				tmp := strconv.Itoa(*opts.Limit)
				i = &tmp
			}
			return
		}()).
		QueryParam("before", func() (i *string) {
			if opts != nil && opts.Before != nil {
				tmp := strconv.Itoa(*opts.Before)
				i = &tmp
			}
			return
		}()).
		QueryParam("after", func() (i *string) {
			if opts != nil && opts.After != nil {
				tmp := strconv.Itoa(*opts.After)
				i = &tmp
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// FindConversion gets currency conversion by id (i.e. USD -> USDC).
func (coinbaseClient *C) FindConversion(conversionId string, opts *model.CoinbaseConversionOptions) (m *model.CoinbaseCurrencyConversion, err error) {
	return m, coinbaseClient.Get(ConversionEndpoint).
		PathParam("conversion_id", conversionId).
		QueryParam("profile_id", func() (i *string) {
			if opts != nil {
				i = opts.ProfileID
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// FindCurrency gets a single currency by id.
func (coinbaseClient *C) FindCurrency(currencyId string) (m *model.CoinbaseCurrency, err error) {
	req := coinbaseClient.Get(CurrencyEndpoint)
	return m, req.PathParam("currency_id", currencyId).Fetch().Assign(&m).JoinMessages()
}

// FindTransfer get information on a single coinbaseClient.
func (coinbaseClient *C) FindTransfer(id string) (m *model.CoinbaseAccountTransfer, err error) {
	req := coinbaseClient.Get(TransferEndpoint)
	return m, req.PathParam("transfer_id", id).Fetch().Assign(&m).JoinMessages()
}

// MakePaymentMethodDeposit will deposits funds from a linked external payment
// method to the specified profile_id.
func (coinbaseClient *C) PaymentMethodDeposit(opts *model.CoinbasePaymentMethodDepositOptions) (m *model.CoinbaseDeposit, err error) {
	return m, coinbaseClient.Post(PaymentMethodDepositEndpoint).
		Body(client.NewBody(client.BodyTypeJSON).
			SetString("profile_id", opts.ProfileID).
			SetString("payment_method_id", &opts.PaymentMethodID).
			SetString("currency", &opts.Currency).
			SetFloat("amount", &opts.Amount)).
		Fetch().Assign(&m).JoinMessages()
}

// PaymentMethods gets a list of the user's linked payment methods
func (coinbaseClient *C) PaymentMethods() (m []*model.CoinbasePaymentMethod, err error) {
	req := coinbaseClient.Get(PaymentMethodEndpoint)
	return m, req.Fetch().Assign(&m).JoinMessages()
}

// PaymentMethodWithdrawal ithdraws funds from the specified profile_id to a
// linked external payment method
func (coinbaseClient *C) PaymentMethodWithdrawal(opts *model.CoinbasePaymentMethodWithdrawalOptions) (m *model.CoinbaseWithdrawal, err error) {
	return m, coinbaseClient.Post(PaymentMethodWithdrawalEndpoint).
		Body(client.NewBody(client.BodyTypeJSON).
			SetString("profile_id", opts.ProfileID).
			SetFloat("amount", &opts.Amount).
			SetString("payment_method_id", &opts.PaymentMethodID).
			SetString("currency", &opts.Currency)).
		Fetch().Assign(&m).JoinMessages()
}

// Orders will list your current open orders. Only open or un-settled orders are
// returned by default. As soon as an order is no longer open and settled, it
// will no longer appear in the default request. Open orders may change state
// between the request and the response depending on market conditions.
func (coinbaseClient *C) Orders(opts *model.CoinbaseOrdersOptions) (m []*model.CoinbaseOrder, err error) {
	return m, coinbaseClient.Get(FillsEndpoint).
		QueryParam("profile_id", func() (i *string) {
			if opts != nil {
				i = opts.ProfileID
			}
			return
		}()).
		QueryParam("product_id", func() (i *string) {
			if opts != nil && opts.ProductID != nil {
				i = opts.ProductID
			}
			return
		}()).
		QueryParam("sortedBy", func() (i *string) {
			if opts != nil && opts.SortedBy != nil {
				i = opts.SortedBy
			}
			return
		}()).
		QueryParam("sorting", func() (i *string) {
			if opts != nil && opts.Sorting != nil {
				i = opts.Sorting
			}
			return
		}()).
		QueryParam("start_date", func() (i *string) {
			if opts != nil && opts.StartDate != nil {
				tmp := opts.StartDate.String()
				i = &tmp
			}
			return
		}()).
		QueryParam("end_date", func() (i *string) {
			if opts != nil && opts.EndDate != nil {
				tmp := opts.EndDate.String()
				i = &tmp
			}
			return
		}()).
		QueryParam("before", func() (i *string) {
			if opts != nil && opts.Before != nil {
				i = opts.Before
			}
			return
		}()).
		QueryParam("after", func() (i *string) {
			if opts != nil && opts.After != nil {
				i = opts.After
			}
			return
		}()).
		QueryParam("limit", func() (i *string) {
			if opts != nil {
				tmp := strconv.Itoa(opts.Limit)
				i = &tmp
			}
			return
		}()).
		QueryParam("status", func() (i *string) {
			if opts != nil && opts.Status != nil {
				slice := []string{}
				for _, v := range opts.Status {
					slice = append(slice, *v)
				}
				tmp := strings.Join(slice, ", ")
				i = &tmp
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// Order will get a single order by order id.
func (coinbaseClient *C) Order(orderID string) (m *model.CoinbaseOrder, err error) {
	req := coinbaseClient.Get(OrderEndpoint)
	return m, req.PathParam("order_id", orderID).Fetch().Assign(&m).JoinMessages()
}

// Accounts lists all trading accounts from the profile of the API key.
func (coinbaseClient *C) ProductTicker(productID string) (m *model.CoinbaseProductTicker, err error) {
	req := coinbaseClient.Get(ProductTickerEndpoint)
	return m, req.PathParam("product_id", productID).Fetch().Assign(&m).JoinMessages()
}

// Transfers gets a list of in-progress and completed transfers of funds in/out of any
// of the user's accounts.
func (coinbaseClient *C) Transfers() (m []*model.CoinbaseAccountTransfer, err error) {
	return m, coinbaseClient.Get(TransfersEndpoint).Fetch().Assign(&m).JoinMessages()
}

// Wallets lists all the user's available Coinbase wallets (These are the
// wallets/accounts that are used for buying and selling on www.coinbase.com)
func (coinbaseClient *C) Wallets() (m []*model.CoinbaseWallet, err error) {
	req := coinbaseClient.Get(WalletsEndpoint)
	return m, req.Fetch().Assign(&m).JoinMessages()
}

// WithdrawalFeeEstimate gets the fee estimate for the crypto withdrawal to
// crypto address
func (coinbaseClient *C) WithdrawalFeeEstimate(opts *model.CoinbaseWithdrawalFeeEstimateOptions) (m *model.CoinbaseWithdrawalFeeEstimate, err error) {
	return m, coinbaseClient.Get(WithdrawalFeeEstimateEndpoint).
		QueryParam("currency", func() (i *string) {
			if opts != nil {
				i = opts.Currency
			}
			return
		}()).
		QueryParam("crypto_address", func() (i *string) {
			if opts != nil {
				i = opts.CryptoAddress
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}
