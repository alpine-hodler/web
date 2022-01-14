package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/alpine-hodler/sdk/coinbase"
	"github.com/alpine-hodler/sdk/graph/generated"
	"github.com/alpine-hodler/sdk/iex"
	"github.com/alpine-hodler/sdk/kraken"
	"github.com/alpine-hodler/sdk/model"
	"github.com/alpine-hodler/sdk/opensea"
)

func (r *mutationResolver) CoinbaseAccountDeposit(ctx context.Context, opts *model.CoinbaseAccountDepositOptions) (*model.CoinbaseDeposit, error) {
	return coinbase.NewClient(coinbase.DefaultClient).CoinbaseAccountDeposit(opts)
}

func (r *mutationResolver) CoinbaseCancelAllOrders(ctx context.Context, opts *model.CoinbaseOrdersOptions) ([]*string, error) {
	return coinbase.NewClient(coinbase.DefaultClient).CancelOpenOrders(opts)
}

func (r *mutationResolver) CoinbaseCreateNewOrder(ctx context.Context, opts *model.CoinbaseNewOrderOptions) (*model.CoinbaseNewOrder, error) {
	return coinbase.NewClient(coinbase.DefaultClient).CreateOrder(opts)
}

func (r *mutationResolver) CoinbaseConvertCurrency(ctx context.Context, opts model.CoinbaseConversionsOptions) (*model.CoinbaseCurrencyConversion, error) {
	return coinbase.NewClient(coinbase.DefaultClient).Convert(opts)
}

func (r *mutationResolver) CoinbaseGenerateCryptoAddress(ctx context.Context, walletID string) (*model.CoinbaseCryptoAddress, error) {
	return coinbase.NewClient(coinbase.DefaultClient).GenerateCryptoAddress(walletID)
}

func (r *mutationResolver) CoinbasePaymentMethodDeposit(ctx context.Context, opts *model.CoinbasePaymentMethodDepositOptions) (*model.CoinbaseDeposit, error) {
	return coinbase.NewClient(coinbase.DefaultClient).PaymentMethodDeposit(opts)
}

func (r *mutationResolver) CoinbasePaymentMethodWithdrawal(ctx context.Context, opts *model.CoinbasePaymentMethodWithdrawalOptions) (*model.CoinbaseWithdrawal, error) {
	return coinbase.NewClient(coinbase.DefaultClient).PaymentMethodWithdrawal(opts)
}

func (r *mutationResolver) CoinbaseAccountWithdrawal(ctx context.Context, opts *model.CoinbaseAccountWithdrawalOptions) (*model.CoinbaseWithdrawal, error) {
	return coinbase.NewClient(coinbase.DefaultClient).AccountWithdrawal(opts)
}

func (r *mutationResolver) CoinbaseCryptoWithdrawal(ctx context.Context, opts *model.CoinbaseCryptoWithdrawalOptions) (*model.CoinbaseWithdrawal, error) {
	return coinbase.NewClient(coinbase.DefaultClient).CryptoWithdrawal(opts)
}

func (r *queryResolver) CoinbaseAccount(ctx context.Context, accountID string) (*model.CoinbaseAccount, error) {
	return coinbase.NewClient(coinbase.DefaultClient).FindAccount(accountID)
}

func (r *queryResolver) CoinbaseAccounts(ctx context.Context) ([]*model.CoinbaseAccount, error) {
	return coinbase.NewClient(coinbase.DefaultClient).Accounts()
}

func (r *queryResolver) CoinbaseAccountHolds(ctx context.Context, accountID string, opts *model.CoinbaseAccountHoldsOptions) ([]*model.CoinbaseAccountHold, error) {
	return coinbase.NewClient(coinbase.DefaultClient).AccountHolds(accountID, opts)
}

func (r *queryResolver) CoinbaseAccountLedger(ctx context.Context, accountID string, opts *model.CoinbaseAccountLedgerOptions) ([]*model.CoinbaseAccountLedger, error) {
	return coinbase.NewClient(coinbase.DefaultClient).AccountLedger(accountID, opts)
}

func (r *queryResolver) CoinbaseAccountTransfers(ctx context.Context, accountID string, opts *model.CoinbaseAccountTransferOptions) ([]*model.CoinbaseAccountTransfer, error) {
	return coinbase.NewClient(coinbase.DefaultClient).AccountTransfers(accountID, opts)
}

func (r *queryResolver) CoinbaseCurrencies(ctx context.Context) ([]*model.CoinbaseCurrency, error) {
	return coinbase.NewClient(coinbase.DefaultClient).Currencies()
}

func (r *queryResolver) CoinbaseCurrencyConversion(ctx context.Context, conversionID string, opts *model.CoinbaseConversionOptions) (*model.CoinbaseCurrencyConversion, error) {
	return coinbase.NewClient(coinbase.DefaultClient).FindConversion(conversionID, opts)
}

func (r *queryResolver) CoinbaseCurrency(ctx context.Context, currentID string) (*model.CoinbaseCurrency, error) {
	return coinbase.NewClient(coinbase.DefaultClient).FindCurrency(currentID)
}

func (r *queryResolver) CoinbaseFees(ctx context.Context) (*model.CoinbaseFees, error) {
	return coinbase.NewClient(coinbase.DefaultClient).Fees()
}

func (r *queryResolver) CoinbaseFills(ctx context.Context, opts *model.CoinbaseFillsOptions) ([]*model.CoinbaseFill, error) {
	return coinbase.NewClient(coinbase.DefaultClient).Fills(opts)
}

func (r *queryResolver) CoinbaseOrders(ctx context.Context, opts *model.CoinbaseOrdersOptions) ([]*model.CoinbaseOrder, error) {
	return coinbase.NewClient(coinbase.DefaultClient).Orders(opts)
}

func (r *queryResolver) CoinbasePaymentMethods(ctx context.Context) ([]*model.CoinbasePaymentMethod, error) {
	return coinbase.NewClient(coinbase.DefaultClient).PaymentMethods()
}

func (r *queryResolver) CoinbaseTransfers(ctx context.Context) ([]*model.CoinbaseAccountTransfer, error) {
	return coinbase.NewClient(coinbase.DefaultClient).Transfers()
}

func (r *queryResolver) CoinbaseTransfer(ctx context.Context, transferID string) (*model.CoinbaseAccountTransfer, error) {
	return coinbase.NewClient(coinbase.DefaultClient).FindTransfer(transferID)
}

func (r *queryResolver) CoinbaseWallets(ctx context.Context) ([]*model.CoinbaseWallet, error) {
	return coinbase.NewClient(coinbase.DefaultClient).Wallets()
}

func (r *queryResolver) CoinbaseWithdrawalFeeEstimate(ctx context.Context, opts *model.CoinbaseWithdrawalFeeEstimateOptions) (*model.CoinbaseWithdrawalFeeEstimate, error) {
	return coinbase.NewClient(coinbase.DefaultClient).WithdrawalFeeEstimate(opts)
}

func (r *queryResolver) IexRules(ctx context.Context, value string) ([]*model.IexRule, error) {
	return iex.NewClient(iex.DefaultClient).Rules(value)
}

func (r *queryResolver) IexRulesSchema(ctx context.Context) (*model.IexRulesSchema, error) {
	return iex.NewClient(iex.DefaultClient).RulesSchema()
}

func (r *queryResolver) KrakenServerTime(ctx context.Context) (*model.KrakenServerTime, error) {
	return kraken.NewClient(kraken.DefaultClient).ServerTime()
}

func (r *queryResolver) KrakenSystemStatus(ctx context.Context) (*model.KrakenSystemStatus, error) {
	return kraken.NewClient(kraken.DefaultClient).SystemStatus()
}

func (r *queryResolver) OpenseaAssets(ctx context.Context, opts *model.OpenseaAssetsOptions) (*model.OpenseaAssets, error) {
	return opensea.NewClient(opensea.DefaultClient).Assets(opts)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
