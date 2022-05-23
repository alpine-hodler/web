[![docs](https://img.shields.io/static/v1?label=coinbase&message=reference&color=blue)](https://pkg.go.dev/github.com/alpine-hodler/sdk@v0.1.0-alpha/pkg/coinbase)

# Coinbase Pro API Wrapper

- [Creating a Client](#creating-a-client)
- [Development](#development)
  - [Testing](#testing)

This package wraps the references defined by the [Coinbase Cloud API](https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getaccounts), and can be installed using

```
go get github.com/alpine-hodlr/sdk/pkg/coinbase
```

## Creating a Client

The `coinbasepro.Client` type is a wrapper for the Go [net/http](https://pkg.go.dev/net/http) standard library package.  An [`http.RoundTripper`](https://pkg.go.dev/net/http#RoundTripper) is required to authenticate for Coinbase Pro requests.  Currently the only method supported for Coinbase Pro authentication is [API key authentication](https://docs.cloud.coinbase.com/sign-in-with-coinbase/docs/api-key-authentication).  See the examples for examples.

## Development

Notes on developing in this package.

### Testing

You will need to create an account for the [Coinbase Pro Sandbox]("https://api-public.sandbox.exchange.coinbase.com") and [create a new API key](https://docs.cloud.coinbase.com/exchange/docs/sandbox#creating-api-keys) for the `Default Portfolio` with `View/Trade/Transfer` permissions.  Then populate the following data in `pkg/coinbase/.simple-test.env`:
```.env
CB_PRO_ACCESS_PASSPHRASE=
CB_PRO_ACCESS_KEY=
CB_PRO_SECRET=
```

Note that `pkg/coinbase/.simple-test.env` is an ignored file and should not be commitable to the repository.  The Coinbase Pro Sandbox can be accessed [here](https://public.sandbox.pro.coinbase.com).
