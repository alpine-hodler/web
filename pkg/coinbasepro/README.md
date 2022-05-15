[![docs](https://img.shields.io/static/v1?label=coinbase&message=reference&color=blue)](https://pkg.go.dev/github.com/alpine-hodler/sdk@v0.1.0-alpha/pkg/coinbase)

# Coinbase Pro API Wrapper

- [Creating a Client](#creating-a-client)
  - [Default Connector](#default-connector)
  - [Environment File](#environment-file)
  - [Custom Connector](#custom-connector)
- [Development](#development)
  - [Testing](#testing)

This package wraps the references defined by the [Coinbase Cloud API](https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getaccounts), and can be installed using

```
go get github.com/alpine-hodlr/sdk/pkg/coinbase
```

## Creating a Client

There are multiple ways to setup a Coinbsae Pro environment.

### Default Connector

If your machine already has environment variables for `CB_PRO_URL`, `CB_PRO_ACCESS_PASSPHRASE`, `CB_PRO_ACCESS_KEY`, and `CB_PRO_SECRET`, then you can simply use the `DefaultConnector`.

```go
client := coinbase.NewClient(coinbasepro.DefaultConnector)
```

### Environment File

IF your machine doesn't already have environment variables for the above, you can define an environment file on your machine somewhere with the environment variables defined.

```.env
CB_PRO_URL=
CB_PRO_ACCESS_PASSPHRASE=
CB_PRO_ACCESS_KEY=
CB_PRO_SECRET=
```

Then use the `NewClient` function to create a new Coinbase Pro API client.

```go
client := coinbasepro.NewClientEnv("/path/to/.env")
```

### Custom Connector

If you don't want to pass the auth variables through the environment, you can pass them inline with a custom connector.

```go
client := coinbase.NewClient(func() *coinbase.Client {
	c := new(coinbase.Client)
	c.SetKey("key")
	c.SetPassphrase("passphrase")
	c.SetSecret("secret")
	c.SetURL("https://api-public.sandbox.exchange.coinbase.com")
	return c
})
```

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
