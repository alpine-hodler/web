package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/alpine-hodler/sdk/internal/graph/generated"
	"github.com/alpine-hodler/sdk/pkg/model"
)

func (r *mutationResolver) CoinbaseAccountDeposit(ctx context.Context, opts *model.CoinbaseAccountDepositOptions) (*model.CoinbaseDeposit, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CoinbaseCancelAllOrders(ctx context.Context, opts *model.CoinbaseOrdersOptions) ([]*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CoinbaseCreateNewOrder(ctx context.Context, opts *model.CoinbaseNewOrderOptions) (*model.CoinbaseNewOrder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CoinbaseConvertCurrency(ctx context.Context, opts model.CoinbaseConversionsOptions) (*model.CoinbaseCurrencyConversion, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CoinbaseGenerateCryptoAddress(ctx context.Context, walletID string) (*model.CoinbaseCryptoAddress, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CoinbasePaymentMethodDeposit(ctx context.Context, opts *model.CoinbasePaymentMethodDepositOptions) (*model.CoinbaseDeposit, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CoinbasePaymentMethodWithdrawal(ctx context.Context, opts *model.CoinbasePaymentMethodWithdrawalOptions) (*model.CoinbaseWithdrawal, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CoinbaseAccountWithdrawal(ctx context.Context, opts *model.CoinbaseAccountWithdrawalOptions) (*model.CoinbaseWithdrawal, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CoinbaseCryptoWithdrawal(ctx context.Context, opts *model.CoinbaseCryptoWithdrawalOptions) (*model.CoinbaseWithdrawal, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseAccount(ctx context.Context, accountID string) (*model.CoinbaseAccount, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseAccounts(ctx context.Context) ([]*model.CoinbaseAccount, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseAccountHolds(ctx context.Context, accountID string, opts *model.CoinbaseAccountHoldsOptions) ([]*model.CoinbaseAccountHold, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseAccountLedger(ctx context.Context, accountID string, opts *model.CoinbaseAccountLedgerOptions) ([]*model.CoinbaseAccountLedger, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseAccountTransfers(ctx context.Context, accountID string, opts *model.CoinbaseAccountTransferOptions) ([]*model.CoinbaseAccountTransfer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseCurrencies(ctx context.Context) ([]*model.CoinbaseCurrency, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseCurrencyConversion(ctx context.Context, conversionID string, opts *model.CoinbaseConversionOptions) (*model.CoinbaseCurrencyConversion, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseCurrency(ctx context.Context, currentID string) (*model.CoinbaseCurrency, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseFees(ctx context.Context) (*model.CoinbaseFees, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseFills(ctx context.Context, opts *model.CoinbaseFillsOptions) ([]*model.CoinbaseFill, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseOrders(ctx context.Context, opts *model.CoinbaseOrdersOptions) ([]*model.CoinbaseOrder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbasePaymentMethods(ctx context.Context) ([]*model.CoinbasePaymentMethod, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseTransfers(ctx context.Context) ([]*model.CoinbaseAccountTransfer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseTransfer(ctx context.Context, transferID string) (*model.CoinbaseAccountTransfer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseWallets(ctx context.Context) ([]*model.CoinbaseWallet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CoinbaseWithdrawalFeeEstimate(ctx context.Context, opts *model.CoinbaseWithdrawalFeeEstimateOptions) (*model.CoinbaseWithdrawalFeeEstimate, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) IexRules(ctx context.Context, value string) ([]*model.IexRule, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) IexRulesSchema(ctx context.Context) (*model.IexRulesSchema, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) KrakenServerTime(ctx context.Context) (*model.KrakenServerTime, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) KrakenSystemStatus(ctx context.Context) (*model.KrakenSystemStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) OpenseaAssets(ctx context.Context, opts *model.OpenseaAssetsOptions) (*model.OpenseaAssets, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
