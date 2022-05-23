[![docs](https://img.shields.io/static/v1?label=coinbase&message=reference&color=blue)](https://pkg.go.dev/github.com/alpine-hodler/sdk@v0.1.0-alpha/pkg/twitter)

# Twitter API Wrapper

- [Creating a Client](#creating-a-client)
- [Development](#development)
  - [Testing](#testing)

This package wraps the references defined by the [Twitter API](https://developer.twitter.com/en/docs/api-reference-index), and can be installed using

```
go get github.com/alpine-hodlr/sdk/pkg/twitter
```

## Creating a Client

The `twitter.Client` type is a wrapper for the Go [net/http](https://pkg.go.dev/net/http) standard library package.  An [`http.RoundTripper`](https://pkg.go.dev/net/http#RoundTripper) is required to authenticate for Twitter requests.  All Twitter [authentication methods](https://developer.twitter.com/en/docs/authentication/overview) are supported.  See the documentation for examples.

## Development

Notes on developing in this package.

### Testing

You will need to [Sign up for access to the Twitter API](https://developer.twitter.com/en/docs/api-reference-index) and generate the APP keys.  Then populate the following data in `pkg/twitter/.simple-test.env`:
```.env
TWITTER_URL=https://api.twitter.com

# OAuth2
TWITTER_BEARER_TOKEN=

# OAuth 1
TWITTER_ACCESS_TOKEN=
TWITTER_ACCESS_SECRET=
TWITTER_CONSUMER_KEY=
TWITTER_CONSUMER_SECRET=

```

Note that `pkg/twitter/.simple-test.env` is an ignored file and should not be commitable to the repository.
