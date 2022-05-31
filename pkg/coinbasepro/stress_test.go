package coinbasepro

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/alpine-hodler/web/pkg/transport"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestStress(t *testing.T) {
	godotenv.Load(".simple-test.env")
	os.Setenv("CB_PRO_URL", "https://api-public.sandbox.exchange.coinbase.com") // safety check

	url := os.Getenv("CB_PRO_URL")
	passphrase := os.Getenv("CB_PRO_ACCESS_PASSPHRASE")
	key := os.Getenv("CB_PRO_ACCESS_KEY")
	secret := os.Getenv("CB_PRO_SECRET")

	client, err := NewClient(context.TODO(), transport.NewAPIKey().
		SetKey(key).
		SetPassphrase(passphrase).
		SetSecret(secret).
		SetURL(url))
	require.NoError(t, err)

	t.Run("Client.Candles intensive looping 21600", func(t *testing.T) {
		t0 := time.Date(2019, 01, 01, 0, 0, 0, 0, time.UTC)
		for t0.Before(time.Now()) {
			next := t0.AddDate(0, 0, 75)
			opts := new(CandlesOptions).
				SetGranularity(Granularity21600).
				SetStart(t0.Format(time.RFC3339)).
				SetEnd(next.Format(time.RFC3339))

			t0 = next
			_, err := client.Candles("BTC-USD", opts)
			require.NoError(t, err)
		}
	})
	// t.Run("Client.Candles intensive looping 60", func(t *testing.T) {
	// 	t0 := time.Date(2022, 05, 10, 0, 0, 0, 0, time.UTC)
	// 	for t0.Before(time.Now()) {
	// 		next := t0.Add(time.Second * 18000)
	// 		opts := new(CandlesOptions).
	// 			SetGranularity(Granularity60).
	// 			SetStart(t0.Format(time.RFC3339)).
	// 			SetEnd(next.Format(time.RFC3339))

	// 		t0 = next
	// 		_, err := client.Candles("BTC-USD", opts)
	// 		require.NoError(t, err)
	// 	}
	// })
}
