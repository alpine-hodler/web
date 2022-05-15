[![docs](https://img.shields.io/static/v1?label=coinbase&message=reference&color=blue)](https://pkg.go.dev/github.com/alpine-hodler/sdk@v0.1.0-alpha/pkg/coinbase)

# Coinbase API Wrapper
This package wraps the references defined by the [Coinbase Cloud API](https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getaccounts). This package can be installed using

```
go get github.com/alpine-hodlr/sdk/pkg/coinbase
```

## Creating a Client

There are multiple ways to setup a coinbase environment.

### Default Connector

If your machine already has environment variables for `CB_PRO_URL`, `CB_PRO_ACCESS_PASSPHRASE`, `CB_PRO_ACCESS_KEY`, and `CB_PRO_SECRET`, then you can simply use the `DefaultConnector`

```go
client := coinbase.NewClient(coinbase.DefaultConnector)
```

### Environment File

Define an environment file on your machine somewhere with the following environment variables defined:

```.env
CB_PRO_URL=
CB_PRO_ACCESS_PASSPHRASE=
CB_PRO_ACCESS_KEY=
CB_PRO_SECRET=
```

Then use the `NewClient` function to create a new Coinbase API client.

```go
client := coinbase.NewClientEnv("/path/to/.env")
```

### Custom Connector

If you want don't want to add the auth variables through the environment, you can add them with a custom connector.

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
