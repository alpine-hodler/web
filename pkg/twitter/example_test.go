package twitter_test

import (
	"context"
	"log"
	"testing"

	"fmt"

	"github.com/alpine-hodler/sdk/internal/env"
	"github.com/alpine-hodler/sdk/pkg/twitter"
	"github.com/alpine-hodler/sdk/tools"
)

var someAuth2BearerToken string

type testOS struct{}

func (tos testOS) Setenv(key, val string) {}

var os = testOS{}

func TestExamples(t *testing.T) {
	defer tools.Quiet()()

	env.Load(".simple-test.env")
	env.SetAlpineHodlerLogLevel("2")

	t.Run("NewClientAuth1", func(t *testing.T) { ExampleNewClientAuth1() })
	t.Run("NewClientAuth2", func(t *testing.T) { ExampleNewClientAuth2() })
}

func ExampleNewClientAuth1() {
	os.Setenv("TWITTER_ACCESS_KEY", "some-twitter-access-key")
	os.Setenv("TWITTER_SECRET", "some-twitter-secret")
	os.Setenv("TWITTER_ACCESS_TOKEN", "some-twitter-access-token")
	os.Setenv("TWITTER_ACCESS_TOKEN_SECRET", "some-twitter-token-secret")

	client, err := twitter.NewClientAuth1(context.TODO())
	if err != nil {
		log.Fatalf("Error creating new client: %v", err)
	}

	// Fetch some tweets to test the connection.
	tweet, err := client.Tweets(new(twitter.TweetsOptions).SetIds([]string{"1261326399320715264"}))
	if err != nil {
		log.Fatalf("Error fetching MongoDB tweet: %v", err)
	}

	fmt.Printf("A tweet about MongoDB: %+v\n", tweet.Data[0])
}

func ExampleNewClientAuth2() {
	os.Setenv("TWITTER_BEARER_TOKEN", "some-twitter-bearer-token")

	client, err := twitter.NewClientAuth2(context.TODO())
	if err != nil {
		log.Fatalf("Error creating new client: %v", err)
	}

	// Fetch some tweets to test the connection.
	tweet, err := client.Tweets(new(twitter.TweetsOptions).SetIds([]string{"1261326399320715264"}))
	if err != nil {
		log.Fatalf("Error fetching MongoDB tweet: %v", err)
	}

	fmt.Printf("A tweet about MongoDB: %+v\n", tweet.Data[0])
}
