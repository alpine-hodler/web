package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/alpine-hodler/sdk/internal/graph/generated"
	"github.com/alpine-hodler/sdk/internal/protomodel"
	"github.com/alpine-hodler/sdk/pkg/coinbase"
	"github.com/alpine-hodler/sdk/pkg/iex"
	"github.com/alpine-hodler/sdk/pkg/kraken"
	"github.com/alpine-hodler/sdk/pkg/model"
	"github.com/alpine-hodler/sdk/pkg/option"
)

func (r *mutationResolver) CoinbaseAccountDeposit(ctx context.Context, opts *protomodel.CoinbaseAccountDepositOptions) (*model.CoinbaseDeposit, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).CoinbaseAccountDeposit(option.CoinbaseAccountDepositFromPrototype(opts))
}

func (r *mutationResolver) CoinbaseCancelAllOrders(ctx context.Context, opts *protomodel.CoinbaseOrdersOptions) ([]*string, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).CancelOpenOrders(option.CoinbaseOrdersFromPrototype(opts))
}

func (r *mutationResolver) CoinbaseCreateNewOrder(ctx context.Context, opts *protomodel.CoinbaseNewOrderOptions) (*model.CoinbaseNewOrder, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).CreateOrder(option.CoinbaseNewOrderFromPrototype(opts))
}

func (r *mutationResolver) CoinbaseConvertCurrency(ctx context.Context, opts protomodel.CoinbaseConversionsOptions) (*model.CoinbaseCurrencyConversion, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).Convert(option.CoinbaseConversionsFromPrototype(&opts))
}

func (r *mutationResolver) CoinbaseGenerateCryptoAddress(ctx context.Context, walletID string) (*model.CoinbaseCryptoAddress, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).GenerateCryptoAddress(walletID)
}

func (r *mutationResolver) CoinbasePaymentMethodDeposit(ctx context.Context, opts *protomodel.CoinbasePaymentMethodDepositOptions) (*model.CoinbaseDeposit, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).PaymentMethodDeposit(option.CoinbasePaymentMethodDepositFromPrototype(opts))
}

func (r *mutationResolver) CoinbasePaymentMethodWithdrawal(ctx context.Context, opts *protomodel.CoinbasePaymentMethodWithdrawalOptions) (*model.CoinbaseWithdrawal, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).PaymentMethodWithdrawal(option.CoinbasePaymentMethodWithdrawalFromPrototype(opts))
}

func (r *mutationResolver) CoinbaseAccountWithdrawal(ctx context.Context, opts *protomodel.CoinbaseAccountWithdrawalOptions) (*model.CoinbaseWithdrawal, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).AccountWithdrawal(option.CoinbaseAccountWithdrawalFromPrototype(opts))
}

func (r *mutationResolver) CoinbaseCryptoWithdrawal(ctx context.Context, opts *protomodel.CoinbaseCryptoWithdrawalOptions) (*model.CoinbaseWithdrawal, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).CryptoWithdrawal(option.CoinbaseCryptoWithdrawalFromPrototype(opts))
}

func (r *queryResolver) CoinbaseAccount(ctx context.Context, accountID string) (*model.CoinbaseAccount, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).FindAccount(accountID)
}

func (r *queryResolver) CoinbaseAccounts(ctx context.Context) ([]*model.CoinbaseAccount, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).Accounts()
}

func (r *queryResolver) CoinbaseAccountHolds(ctx context.Context, accountID string, opts *protomodel.CoinbaseAccountHoldsOptions) ([]*model.CoinbaseAccountHold, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).AccountHolds(accountID, option.CoinbaseAccountHoldsFromPrototype(opts))
}

func (r *queryResolver) CoinbaseAccountLedger(ctx context.Context, accountID string, opts *protomodel.CoinbaseAccountLedgerOptions) ([]*model.CoinbaseAccountLedger, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).AccountLedger(accountID, option.CoinbaseAccountLedgerFromPrototype(opts))
}

func (r *queryResolver) CoinbaseAccountTransfers(ctx context.Context, accountID string, opts *protomodel.CoinbaseAccountTransfersOptions) ([]*model.CoinbaseAccountTransfer, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).AccountTransfers(accountID, option.CoinbaseAccountTransfersFromPrototype(opts))
}

func (r *queryResolver) CoinbseCancelOrder(ctx context.Context, orderID string) (string, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).CancelOrder(orderID)
}

func (r *queryResolver) CoinbaseCurrencies(ctx context.Context) ([]*model.CoinbaseCurrency, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).Currencies()
}

func (r *queryResolver) CoinbaseCurrencyConversion(ctx context.Context, conversionID string, opts *protomodel.CoinbaseConversionOptions) (*model.CoinbaseCurrencyConversion, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).FindConversion(conversionID, option.CoinbaseConversionFromPrototype(opts))
}

func (r *queryResolver) CoinbaseCurrency(ctx context.Context, currentID string) (*model.CoinbaseCurrency, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).FindCurrency(currentID)
}

func (r *queryResolver) CoinbaseFees(ctx context.Context) (*model.CoinbaseFees, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).Fees()
}

func (r *queryResolver) CoinbaseFills(ctx context.Context, opts *protomodel.CoinbaseFillsOptions) ([]*model.CoinbaseFill, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).Fills(option.CoinbaseFillsFromPrototype(opts))
}

func (r *queryResolver) CoinbaseOrders(ctx context.Context, opts *protomodel.CoinbaseOrdersOptions) ([]*model.CoinbaseOrder, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).Orders(option.CoinbaseOrdersFromPrototype(opts))
}

func (r *queryResolver) CoinbaseOrder(ctx context.Context, orderID string) (*model.CoinbaseOrder, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).Order(orderID)
}

func (r *queryResolver) CoinbasePaymentMethods(ctx context.Context) ([]*model.CoinbasePaymentMethod, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).PaymentMethods()
}

func (r *queryResolver) CoinbaseProductTicker(ctx context.Context, productID string) (*model.CoinbaseProductTicker, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).ProductTicker(productID)
}

func (r *queryResolver) CoinbaseTransfers(ctx context.Context) ([]*model.CoinbaseAccountTransfer, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).Transfers()
}

func (r *queryResolver) CoinbaseTransfer(ctx context.Context, transferID string) (*model.CoinbaseAccountTransfer, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).FindTransfer(transferID)
}

func (r *queryResolver) CoinbaseWallets(ctx context.Context) ([]*model.CoinbaseWallet, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).Wallets()
}

func (r *queryResolver) CoinbaseWithdrawalFeeEstimate(ctx context.Context, opts *protomodel.CoinbaseWithdrawalFeeEstimateOptions) (*model.CoinbaseWithdrawalFeeEstimate, error) {
	return coinbase.NewClient(coinbase.DefaultConnector).WithdrawalFeeEstimate(option.CoinbaseWithdrawalFeeEstimateFromPrototype(opts))
}

func (r *queryResolver) IexRules(ctx context.Context, value string) ([]*model.IexRule, error) {
	return iex.NewClient(iex.DefaultConnector).Rules(value)
}

func (r *queryResolver) IexRulesSchema(ctx context.Context) (*model.IexRulesSchema, error) {
	return iex.NewClient(iex.DefaultConnector).RulesSchema()
}

func (r *queryResolver) KrakenServerTime(ctx context.Context) (*model.KrakenServerTime, error) {
	return kraken.NewClient(kraken.DefaultConnector).ServerTime()
}

func (r *queryResolver) KrakenSystemStatus(ctx context.Context) (*model.KrakenSystemStatus, error) {
	return kraken.NewClient(kraken.DefaultConnector).SystemStatus()
}

func (r *queryResolver) OpenseaAssets(ctx context.Context, opts *protomodel.OpenseaAssetsOptions) (*model.OpenseaAssets, error) {
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
