package twitter_test

import (
	"context"
	"log"
	"testing"

	"fmt"

	"github.com/alpine-hodler/sdk/internal/env"
	"github.com/alpine-hodler/sdk/pkg/twitter"
)

var someAuth2BearerToken string

func TestExamples(t *testing.T) {
	// defer tools.Quiet()()

	env.Load(".simple-test.env")

	someAuth2BearerToken = env.TwitterAuth2BearerToken.Get()

	// ! Make sure that these tests only run on the sandbox.
	// env.SetCoinbaseProURL("https://ads-api-sandbox.twitter.com")
	env.SetAlpineHodlerLogLevel("2")

	// Mock the env lambdas to avoid actually setting this data in test.
	env.SetTwitterAuth2BearerToken = func(_ string) error { return nil }
	env.SetTwitterURL = func(_ string) error { return nil }

	t.Run("NewAuth1aUserContextClient", func(t *testing.T) { ExampleNewAuth1aUserContextClient() })
	// t.Run("NewAuth2Client", func(t *testing.T) { ExampleNewAuth2Client() })
	// t.Run("NewClientConnector", func(t *testing.T) { ExampleNewClientConnector() })
}

func ExampleNewAuth1aUserContextClient() {
	client, err := twitter.NewClientAuth1(context.Background())
	if err != nil {
		log.Fatalf("Error creating new client: %v", err)
	}

	// Fetch the compliance jobs to test the connection.
	tweet, err := client.Tweets(new(twitter.TweetsOptions).SetIds([]string{"1261326399320715264"}))
	if err != nil {
		log.Fatalf("Error fetching tweets: %v", err)
	}

	fmt.Printf("A tweet: %+v\n", tweet.Data[0])
}

// func ExampleNewAuth2Client() {
// 	env.SetTwitterAuth2BearerToken(someAuth2BearerToken)
// 	env.SetTwitterURL("https://api.twitter.com")

// 	client, err := twitter.NewAuth2Client(context.Background())
// 	if err != nil {
// 		log.Fatalf("Error creating new client: %v", err)
// 	}

// 	// Fetch the compliance jobs to test the connection.
// 	tweet, err := client.Tweets(nil)
// 	if err != nil {
// 		log.Fatalf("Error fetching compliance jobs: %v", err)
// 	}

// 	fmt.Printf("A tweet: %+v\n", tweet.Data[0])
// }

// func ExampleNewClientConnector() {
// 	client, err := twitter.NewClientConnector(context.Background(), func() (client.C, error) {
// 		c := new(twitter.C).
// 			SetAuth2BearerToken(someAuth2BearerToken).
// 			SetAuthenticationMethod(twitter.Auth2BearerToken).
// 			SetURL("https://api.twitter.com")
// 		return c, nil
// 	})
// 	if err != nil {
// 		log.Fatalf("Error creating new client: %v", err)
// 	}

// 	// Fetch the compliance jobs to test the connection.
// 	tweet, err := client.Tweets(new(twitter.TweetsOptions).SetIds([]string{"1261326399320715264"}))
// 	if err != nil {
// 		log.Fatalf("Error fetching compliance jobs: %v", err)
// 	}

// 	fmt.Printf("A tweet: %+v\n", tweet.Data[0].Text)
// }
