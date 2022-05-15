[![docs](https://img.shields.io/static/v1?label=coinbase&message=reference&color=blue)](https://pkg.go.dev/github.com/alpine-hodler/sdk@v0.1.0-alpha/pkg/twitter)

# Twitter API Wrapper

- [Creating a Client](#creating-a-client)
  - [Environment File](#environment-file)
  - [Custom Connector](#custom-connector)
- [Development](#development)
  - [Testing](#testing)

This package wraps the references defined by the [Twitter API](https://developer.twitter.com/en/docs/api-reference-index), and can be installed using

```
go get github.com/alpine-hodlr/sdk/pkg/twitter
```

## Creating a Client

There are multiple ways to setup a Twitter environment.  See the go documentation examples for more concrete usages.

### Environment File

IF your machine doesn't already have environment variables for the above, you can define an environment file on your machine somewhere with the environment variables defined.

```.env
# .my-env.env
TWITTER_URL=https://api.twitter.com
TWITTER_ACCESS_KEY=
TWITTER_SECRET=
TWITTER_AUTH2_BEARER_TOKEN=
```

Then use the appropriate `New*Client` function to create the client.

```go
env.Load(".my-env.env")
auth2Client, _ := coinbase.NewAuth2Client()
```

### Custom Connector

If you don't want to pass the auth variables through the environment, you can pass them inline with a custom connector.

```go
auth2Client, _ := coinbase.NewClientConnector(func() (*twitter.Client, error) {
		c := new(twitter.C).
			SetAuth2BearerToken(someAuth2BearerToken).
			SetAuthenticationMethod(twitter.Auth2BearerToken).
			SetURL("https://api.twitter.com")
	return c, nil
})
```

## Development

Notes on developing in this package.

### Testing

You will need to [Sign up for access to the Twitter API](https://developer.twitter.com/en/docs/api-reference-index) and generate the APP keys.  Then populate the following data in `pkg/twitter/.simple-test.env`:
```.env
TWITTER_URL=https://api.twitter.com
TWITTER_ACCESS_KEY=
TWITTER_SECRET=
TWITTER_BEARER_TOKEN=

```

Note that `pkg/twitter/.simple-test.env` is an ignored file and should not be commitable to the repository.
