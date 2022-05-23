package twitter_test

import (
	"context"
	"log"
	"os"
	"testing"

	"fmt"

	"github.com/alpine-hodler/sdk/pkg/transport"
	"github.com/alpine-hodler/sdk/pkg/twitter"
	"github.com/alpine-hodler/sdk/tools"
	"github.com/joho/godotenv"
)

func TestExamples(t *testing.T) {
	defer tools.Quiet()()

	godotenv.Load(".simple-test.env")

	t.Run("NewClient_basic", func(t *testing.T) { ExampleNewClient_basic() })
	t.Run("NewClient_oauth1", func(t *testing.T) { ExampleNewClient_oauth1() })
	t.Run("NewClient_oauth2", func(t *testing.T) { ExampleNewClient_oauth2() })
}

func ExampleNewClient_basic() {
	// Read credentials from environment variables.
	email := os.Getenv("TWITTER_ENTERPRISE_EMAIL")
	password := os.Getenv("TWITTER_ENTERPRISE_PASSWORD")
	url := os.Getenv("TWITTER_URL")

	// Initialize an Basic client transport
	basic := transport.NewBasic().SetEmail(email).SetPassword(password).SetURL(url)

	// Initialize client using the Auth2 client transport.
	client, _ := twitter.NewClient(context.TODO(), basic)
	fmt.Printf("Twitter client: %+v\n", client)
}

func ExampleNewClient_oauth1() {
	// Read credentials from environment variables.
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")
	url := os.Getenv("TWITTER_URL")

	// Initialize an Auth1 client transport.
	oauth1 := transport.NewAuth1().
		SetAccessToken(accessToken).
		SetAccessTokenSecret(accessSecret).
		SetConsumerKey(consumerKey).
		SetConsumerSecret(consumerSecret).
		SetURL(url)

	// Initialize client using the Auth1 client transport.
	client, err := twitter.NewClient(context.TODO(), oauth1)
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

func ExampleNewClient_oauth2() {
	// Read credentials from environment variables.
	bearerToken := os.Getenv("TWITTER_BEARER_TOKEN")
	url := os.Getenv("TWITTER_URL")

	// Initialize an Auth2 client transport
	oauth2 := transport.NewAuth2().SetBearer(bearerToken).SetURL(url)

	// Initialize client using the Auth2 client transport.
	client, err := twitter.NewClient(context.TODO(), oauth2)
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
