package integration

import (
	"testing"

	"github.com/alpine-hodler/sdk/internal/env"
	"github.com/alpine-hodler/sdk/pkg/coinbase"
	"github.com/alpine-hodler/sdk/pkg/model"
	"github.com/stretchr/testify/require"
)

func TestCoinbaseGETIntegration(t *testing.T) {
	env.SetCoinbaseProURL("https://api-public.sandbox.exchange.coinbase.com")
	env.SetCoinbaseProAccessPassphrase("v5knpcj0bj9")
	env.SetCoinbaseProSecret("PXCkmxsYY9gaVuuSKJ3n442/+kmp5YlMY0jsHVI+otcco+s57sphcZHhQoRVeivxfhc1yYWTSOuU+nrCqsxIbw==")
	env.SetCoinbaseProAccessKey("9d5cd2a0d877787329c9446c1d64f8f7")

	client := coinbase.NewClient(coinbase.DefaultConnector)
	t.Run("connect client", func(t *testing.T) {
		require.NoError(t, client.Connect())
	})
	t.Run("validate accounts", func(t *testing.T) {
		accounts, err := client.Accounts()
		require.NoError(t, err)

		tc := newTestCase(t, "coinbase_accounts.json")

		var expected []*model.CoinbaseAccount
		require.NoError(t, tc.unmarshalModels(&expected))
		require.Equal(t, expected, accounts)
	})
	t.Run("validate account holds", func(t *testing.T) {
		accounts, err := client.Accounts()
		require.NoError(t, err)

		for _, account := range accounts {
			accountHolds, err := client.AccountHolds(account.Id, nil)
			require.NoError(t, err)
			require.Equal(t, accountHolds, []*model.CoinbaseAccountHold{})
		}
	})
	t.Run("validate account ledgers", func(t *testing.T) {
		accounts, err := client.Accounts()
		require.NoError(t, err)

		all := []*model.CoinbaseAccountLedger{}
		for _, account := range accounts {
			ledgers, err := client.AccountLedger(account.Id, nil)
			require.NoError(t, err)
			if len(ledgers) > 0 {
				all = append(all, ledgers...)
			}
		}

		tc := newTestCase(t, "coinbase_account_ledgers.json")

		var expected []*model.CoinbaseAccountLedger
		require.NoError(t, tc.unmarshalModels(&expected))
		require.Equal(t, expected, all)
	})
}
