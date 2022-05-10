package coinbase

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alpine-hodler/sdk/internal/client"
	"github.com/alpine-hodler/sdk/internal/env"
	"github.com/alpine-hodler/sdk/internal/log"
	"github.com/alpine-hodler/sdk/pkg/websocket"
	"github.com/alpine-hodler/sdk/tools"
)

// C is the coinbase pro client
type C struct {
	client.Parent

	client     http.Client
	key        string
	passphrase string
	secret     string
	url        string
}

// newCoinbaseClient will populate the client auth credietions directly
func newCoinbaseClient(key, passphrase, secret, url string) *C {
	c := new(C)
	c.key = key
	c.passphrase = passphrase
	c.secret = secret
	c.url = url
	return c
}

// newCoinbaseClientEnv will populate the client auth credentials using a
// .env file
func newCoinbaseClientEnv() *C {
	c := new(C)
	c.key = env.CoinbaseProAccessKey.Get()
	c.passphrase = env.CoinbaseProAccessPassphrase.Get()
	c.secret = env.CoinbaseProSecret.Get()
	c.url = env.CoinbaseProURL.Get()
	return c
}

// DefaultConnector will pull the coinbase authentication data from the env
// variables.  See README for more information on how to set these up.
func DefaultConnector() (client.C, error) {
	c := newCoinbaseClientEnv()
	return c, nil
}

// NewAccounts will return a new accounts structure to query on trading accounts
func NewClient(conn client.Connector) *C {
	coinbaseClient := new(C)
	client.ConstructParent(&coinbaseClient.Parent, conn)
	return coinbaseClient
}

func NewClientEnv(envFilepath string) *C {
	env.Load(envFilepath)
	coinbaseClient := new(C)
	client.ConstructParent(&coinbaseClient.Parent, DefaultConnector)
	return coinbaseClient
}

// NewWebsocket will create a connection to the coinbase websocket and
// return a singleton that can be used to open channels that stream product
// data via a websocket.
func NewWebsocket(ws websocket.Creator) *ProductWebsocket {
	productWebsocket := new(ProductWebsocket)
	productWebsocket.conn, _ = ws(WebsocketURL)
	return productWebsocket
}

// generateSig generates the coinbase base64-encoded signature required to make
// requests.  In particular, the coinbaseClient-ACCESS-SIGN header is generated by creating
// a sha256 HMAC using the base64-decoded secret key on the prehash string
// timestamp + method + requestPath + body (where + represents string
// concatenation) and base64-encode the output. The timestamp value is the same
// as the coinbaseClient-ACCESS-TIMESTAMP header.
func (coinbaseClient *C) generateSig(secret, message string) (string, error) {
	key, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	signature := hmac.New(sha256.New, key)
	_, err = signature.Write([]byte(message))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature.Sum(nil)), nil
}

// generageMsg makes the message to be signed
func (coinbaseClient *C) generageMsg(creq client.Request, timestamp string) string {
	return fmt.Sprintf("%s%s%s%s", timestamp, creq.MethodStr(), creq.EndpointPath(), string(creq.GetBody().Bytes()))
}

// setHeaders sets the headers for a coinbase api request, in particular:
//
// - coinbaseClient-ACCESS-KEY The api key as a string.
// - coinbaseClient-ACCESS-SIGN The base64-encoded signature (see Signing a Message).
// - coinbaseClient-ACCESS-TIMESTAMP A timestamp for your request.
// - coinbaseClient-ACCESS-PASSPHRASE The passphrase you specified when creating the API key.
func (coinbaseClient *C) setHeaders(hreq *http.Request, creq client.Request) (e error) {
	// TODO depricate getting key/passphrase/secret with secret keeper
	var (
		timestamp = strconv.FormatInt(time.Now().Unix(), 10)
		msg       = coinbaseClient.generageMsg(creq, timestamp)
	)

	var sig string
	sig, e = coinbaseClient.generateSig(coinbaseClient.secret, msg)
	hreq.Header.Add("accept", "application/json")
	hreq.Header.Add("content-type", "application/json")
	hreq.Header.Add("cb-access-key", coinbaseClient.key)
	hreq.Header.Add("cb-access-passphrase", coinbaseClient.passphrase)
	hreq.Header.Add("cb-access-sign", sig)
	hreq.Header.Add("cb-access-timestamp", timestamp)

	// TODO wrap this in a logger
	logMsg := `{Client:{Access:{Key:%s,Passphrase:%s,Timestamp:%s,Sign:%s}}}`
	client.Logf(log.DEBUG, &creq, logMsg, coinbaseClient.key, coinbaseClient.passphrase, timestamp, sig)
	return
}

// request makes an http request to the coinbase api, given a method and an endpoint.
func (coinbaseClient *C) Do(creq client.Request) (*http.Response, error) {
	// TODO make data-compatible for non-get requests
	uri := coinbaseClient.url + creq.EndpointPath()

	client.Logf(log.DEBUG, &creq, `{Client:{URI:%s}}`, uri)

	hreq, err := http.NewRequest(creq.MethodStr(), uri, bytes.NewReader(creq.GetBody().Bytes()))
	if err != nil {
		return nil, err
	}
	if err := coinbaseClient.setHeaders(hreq, creq); err != nil {
		return nil, err
	}
	return coinbaseClient.client.Do(hreq)
}

// Connect creats a new client instance
func (coinbaseClient *C) Connect() error {
	coinbaseClient.client = http.Client{}
	return nil
}

// Identifier identifies requests
func (coinbaseClient *C) Identifier() string {
	return "Coinbase Pro"
}

// Accont returns information for a single account. Use this endpoint when you know the account_id. API key must
// belong to the same profile as the account.
func (coinbaseClient *C) Account(accountId string) (m *Account, _ error) {
	return m, coinbaseClient.Get(AccountEndpoint).PathParam("account_id", accountId).Fetch().Assign(&m).JoinMessages()
}

// Accounts lists all trading accounts from the profile of the API key.
func (coinbaseClient *C) Accounts() (m []*Account, _ error) {
	return m, coinbaseClient.Get(AccountsEndpoint).Fetch().Assign(&m).JoinMessages()
}

// AccountHolds returns a list of holds of an account that belong to the same profile as the API key. Holds are placed
// for any active orders or pending withdraw requests. As an order is filled, the hold amount is updated.
func (coinbaseClient *C) AccountHolds(accountId string, opts *AccountHoldsOptions) (m []*AccountHold, _ error) {
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

// AccountLedger lists ledger activity for an account. This includes anything that would affect the accounts balance -
// transfers, trades, fees, etc
func (coinbaseClient *C) AccountLedger(accountId string, opts *AccountLedgerOptions) (m []*AccountLedger, _ error) {
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
				i = opts.ProfileId
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// AccountTransfer get information on a single coinbaseClient.
func (coinbaseClient *C) AccountTransfer(id string) (m *AccountTransfer, _ error) {
	return m, coinbaseClient.Get(TransferEndpoint).PathParam("transfer_id", id).Fetch().Assign(&m).JoinMessages()
}

// AccountTransfers lists past withdrawals and deposits for an account.
func (coinbaseClient *C) AccountTransfers(accountId string,
	opts *AccountTransfersOptions) (m []*AccountTransfer, _ error) {
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

// AccountWithdrawal Withdraws funds from the specified profile_id to a www.coinbase.com wallet.
func (coinbaseClient *C) AccountWithdrawal(opts *AccountWithdrawalOptions) (m *Withdrawal, _ error) {
	return m, coinbaseClient.Post(AccountWithdrawalEndpoint).
		Body(client.NewBody(client.BodyTypeJSON).
			SetString("profile_id", opts.ProfileId).
			SetFloat("amount", &opts.Amount).
			SetString("coinbase_account_id", &opts.CoinbaseAccountId).
			SetString("currency", &opts.Currency)).
		Fetch().Assign(&m).JoinMessages()
}

// Accont returns information for a single account. Use this endpoint when you know the account_id. API key must
// belong to the same profile as the account.
func (coinbaseClient *C) Book(productID string, opts *BookOptions) (m *Book, _ error) {
	return m, coinbaseClient.Get(BookEndpoint).PathParam("product_id", productID).
		QueryParam("level", func() (i *string) {
			if opts != nil {
				i = tools.Int32PtrStringPtr(opts.Level)
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// Candles gets the historic rates for a product. Rates are returned in grouped buckets.
func (coinbaseClient *C) Candles(productID string, opts *CandlesOptions) (m *Candles, _ error) {
	return m, coinbaseClient.Get(CandlesEndpoint).
		PathParam("product_id", productID).
		QueryParam("granularity", func() (i *string) {
			if opts != nil {
				granularity := opts.Granularity.Int()
				i = tools.IntPtrStringPtr(&granularity)
			}
			return
		}()).
		QueryParam("start", func() (i *string) {
			if opts != nil && opts.Start != nil {
				start := opts.Start.Format(CoinbaseTimeLayout1)
				i = &start
			}
			return
		}()).
		QueryParam("end", func() (i *string) {
			if opts != nil && opts.End != nil {
				end := opts.End.Format(CoinbaseTimeLayout1)
				i = &end
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// CryptoWithdrawal withdraws funds from the specified profile_id to an external crypto address
func (coinbaseClient *C) CryptoWithdrawal(opts *CryptoWithdrawalOptions) (m *Withdrawal, _ error) {
	return m, coinbaseClient.Post(CryptoWithdrawalEndpoint).
		Body(client.NewBody(client.BodyTypeJSON).
			SetString("profile_id", opts.ProfileId).
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

// CoinbaseAccountDeposit will deposit funds from a www.coinbase.com wallet to the specified profile_id.
func (coinbaseClient *C) CoinbaseAccountDeposit(opts *AccountDepositOptions) (m *Deposit, _ error) {
	return m, coinbaseClient.Post(AccountDepositEndpoint).
		Body(client.NewBody(client.BodyTypeJSON).
			SetString("profile_id", opts.ProfileId).
			SetString("coinbase_account_id", &opts.CoinbaseAccountId).
			SetString("currency", &opts.Currency).
			SetFloat("amount", &opts.Amount)).
		Fetch().Assign(&m).JoinMessages()
}

// CanceCancelOpenOrderslAll will with best effort, cancel all open orders. This may require you to make the request
// multiple times until all of the open orders are deleted.
func (coinbaseClient *C) CancelOpenOrders(opts *OrdersOptions) (m []*string, _ error) {
	return m, coinbaseClient.Delete(OrdersEndpoint).
		QueryParam("profile_id", func() (i *string) {
			if opts != nil {
				i = opts.ProfileId
			}
			return
		}()).
		QueryParam("product_id", func() (i *string) {
			if opts != nil && opts.ProductId != nil {
				i = opts.ProductId
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// CancelOrder will cancel a single open order by {id}.
func (coinbaseClient *C) CancelOrder(orderID string) (str string, _ error) {
	return str, coinbaseClient.Delete(OrderEndpoint).PathParam("order_id", orderID).Fetch().Assign(&str).JoinMessages()
}

// CreateOrder will create an order. You can place two types of orders: limit
func (coinbaseClient *C) CreateOrder(opts *NewOrderOptions) (m *NewOrder, _ error) {
	return m, coinbaseClient.Post(NewOrderEndpoint).
		Body(client.NewBody(client.BodyTypeJSON).
			SetString("profile_id", opts.ProfileId).
			SetString("product_id", &opts.ProductId).
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

// Convert converts funds from from currency to to currency. Funds are converted on the from account  in the profile_id
// profile.
func (coinbaseClient *C) Convert(opts *ConversionsOptions) (m *CurrencyConversion, _ error) {
	return m, coinbaseClient.Post(ConversionsEndpoint).
		Body(client.NewBody(client.BodyTypeJSON).
			SetString("from", &opts.From).
			SetString("to", &opts.To).
			SetFloat("amount", &opts.Amount).
			SetString("profile_id", func() (s *string) {
				if opts.ProfileId != nil {
					s = opts.ProfileId
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
func (coinbaseClient *C) Currencies() (m []*Currency, _ error) {
	return m, coinbaseClient.Get(CurrenciesEndpoint).Fetch().Assign(&m).JoinMessages()
}

// Currency gets a single currency by id.
func (coinbaseClient *C) Currency(currencyId string) (m *Currency, _ error) {
	return m, coinbaseClient.Get(CurrencyEndpoint).PathParam("currency_id", currencyId).Fetch().Assign(&m).JoinMessages()
}

// CurrencyConversion gets currency conversion by id (i.e. USD -> USDC).
func (coinbaseClient *C) CurrencyConversion(conversionId string,
	opts *ConversionOptions) (m *CurrencyConversion, _ error) {
	return m, coinbaseClient.Get(ConversionEndpoint).
		PathParam("conversion_id", conversionId).
		QueryParam("profile_id", func() (i *string) {
			if opts != nil {
				i = opts.ProfileId
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// GenerateCryptoAddress will generates a one-time crypto address for depositing crypto.
func (coinbaseClient *C) GenerateCryptoAddress(walletId string) (m *CryptoAddress, _ error) {
	return m, coinbaseClient.Post(AddressesEndpoint).PathParam("account_id", walletId).Fetch().Assign(&m).JoinMessages()
}

// Fees will get fees rates and 30 days trailing volume.
func (coinbaseClient *C) Fees() (m *Fees, _ error) {
	return m, coinbaseClient.Get(FeesEndpoint).Fetch().Assign(&m).JoinMessages()
}

// ## API Key Permissions
// This endpoint requires either the "view" or "trade" permission.
//
// ## Settlement and Fees
// Fees are recorded in two stages. Immediately after the matching engine completes a match, the fill is inserted into
// our datastore. Once the fill is recorded, a settlement process will settle the fill and credit both trading
// counterparties.
//
// The fee field indicates the fees charged for this individual fill.
//
// ### Liquidity
// The liquidity field indicates if the fill was the result of a liquidity provider or liquidity taker. M indicates
// Maker and T indicates Taker.
//
// ### Pagination
// Fills are returned sorted by descending trade_id from the largest trade_id to the smallest trade_id. The CB-BEFORE
// header will have this first trade id so that future requests using the cb-before parameter will fetch fills with a
// greater trade id (newer fills).
func (coinbaseClient *C) Fills(opts *FillsOptions) (m []*Fill, _ error) {
	return m, coinbaseClient.Get(FillsEndpoint).
		QueryParam("order_id", func() (i *string) {
			if opts != nil {
				i = opts.OrderId
			}
			return
		}()).
		QueryParam("product_id", func() (i *string) {
			if opts != nil {
				i = opts.ProductId
			}
			return
		}()).
		QueryParam("profile_id", func() (i *string) {
			if opts != nil {
				i = opts.ProfileId
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

// MakePaymentMethodDeposit will deposits funds from a linked external payment method to the
// specified profile_id.
func (coinbaseClient *C) PaymentMethodDeposit(opts *PaymentMethodDepositOptions) (m *Deposit, _ error) {
	return m, coinbaseClient.Post(PaymentMethodDepositEndpoint).
		Body(client.NewBody(client.BodyTypeJSON).
			SetString("profile_id", opts.ProfileId).
			SetString("payment_method_id", &opts.PaymentMethodId).
			SetString("currency", &opts.Currency).
			SetFloat("amount", &opts.Amount)).
		Fetch().Assign(&m).JoinMessages()
}

// PaymentMethods gets a list of the user's linked payment methods
func (coinbaseClient *C) PaymentMethods() (m []*PaymentMethod, _ error) {
	return m, coinbaseClient.Get(PaymentMethodEndpoint).Fetch().Assign(&m).JoinMessages()
}

// PaymentMethodWithdrawal ithdraws funds from the specified profile_id to a linked external payment method
func (coinbaseClient *C) PaymentMethodWithdrawal(opts *PaymentMethodWithdrawalOptions) (m *Withdrawal, _ error) {
	return m, coinbaseClient.Post(PaymentMethodWithdrawalEndpoint).
		Body(client.NewBody(client.BodyTypeJSON).
			SetString("profile_id", opts.ProfileId).
			SetFloat("amount", &opts.Amount).
			SetString("payment_method_id", &opts.PaymentMethodId).
			SetString("currency", &opts.Currency)).
		Fetch().Assign(&m).JoinMessages()
}

// Orders will list your current open orders. Only open or un-settled orders are returned by default. As soon as an
// order is no longer open and settled, it will no longer appear in the default request. Open orders may change state
// between the request and the response depending on market conditions.
func (coinbaseClient *C) Orders(opts *OrdersOptions) (m []*Order, _ error) {
	return m, coinbaseClient.Get(OrdersEndpoint).
		QueryParam("profile_id", func() (i *string) {
			if opts != nil {
				i = opts.ProfileId
			}
			return
		}()).
		QueryParam("product_id", func() (i *string) {
			if opts != nil && opts.ProductId != nil {
				i = opts.ProductId
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
					slice = append(slice, v)
				}
				tmp := strings.Join(slice, ", ")
				i = &tmp
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// Order will get a single order by order id.
func (coinbaseClient *C) Order(orderID string) (m *Order, _ error) {
	return m, coinbaseClient.Get(OrderEndpoint).PathParam("order_id", orderID).Fetch().Assign(&m).JoinMessages()
}

// Products will return a slce of all available currency pairs for trading.
func (coinbaseClient *C) Products(opts *ProductsOptions) (m []*Product, _ error) {
	return m, coinbaseClient.Get(ProductsEndpoint).
		QueryParam("type", func() (i *string) {
			if opts != nil {
				i = opts.Type
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// Product will get information on a single product using a productID.
func (coinbaseClient *C) Product(productID string) (m *Product, _ error) {
	return m, coinbaseClient.Get(ProductEndpoint).PathParam("product_id", productID).Fetch().Assign(&m).JoinMessages()
}

// Accounts lists all trading accounts from the profile of the API key.
func (coinbaseClient *C) ProductTicker(
	productID string) (m *ProductTicker, _ error) {
	req := coinbaseClient.Get(ProductTickerEndpoint)
	return m, req.PathParam("product_id", productID).Fetch().Assign(&m).JoinMessages()
}

// Profiles will return a slice of all of the current user's profiles.
func (coinbaseClient *C) Profiles(opts *ProfilesOptions) (m []*Profile, _ error) {
	return m, coinbaseClient.Get(ProfilesEndpoint).
		QueryParam("active", func() (i *string) {
			if opts != nil && opts.Active != nil {
				boolStr := strconv.FormatBool(*opts.Active)
				i = &boolStr
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}

// SignedPrices returns cryptographically signed prices ready to be posted on-chain using Compound's Open Oracle smart
// contract.
func (coinbaseClient *C) SignedPrices() (m *Oracle, _ error) {
	return m, coinbaseClient.Get(OracleEndpoint).Fetch().Assign(&m).JoinMessages()
}

// Transfers gets a list of in-progress and completed transfers of funds in/out of any of the user's accounts.
func (coinbaseClient *C) Transfers() (m []*AccountTransfer, _ error) {
	return m, coinbaseClient.Get(TransfersEndpoint).Fetch().Assign(&m).JoinMessages()
}

// Wallets lists all the user's available Coinbase wallets (These are the wallets/accounts that are  used for buying
// and selling on www.coinbase.com)
func (coinbaseClient *C) Wallets() (m []*Wallet, _ error) {
	return m, coinbaseClient.Get(WalletsEndpoint).Fetch().Assign(&m).JoinMessages()
}

// WithdrawalFeeEstimate gets the fee estimate for the crypto withdrawal to crypto address
func (coinbaseClient *C) WithdrawalFeeEstimate(
	opts *WithdrawalFeeEstimateOptions) (m *WithdrawalFeeEstimate, _ error) {
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
