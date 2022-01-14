# sdk

sdk is a monorepo that provides a graphql server and varios SDKs for connecting to and interacting with various APIs and websockets.

This repo is a work in progress

## Getting Started
  * Dependencies
    * [Install Bazel](https://github.com/alpine-hodler/sdk#install-bazel)
    * [Install docker](https://github.com/alpine-hodler/sdk#install-docker)
    * [Install go](https://go.dev/doc/install)
  * [Build SDK](https://github.com/alpine-hodler/sdk#build-sdk)
  * Servers (`cmd` package)
    * [Graph](https://github.com/alpine-hodler/sdk#graphql)
  * Software Development Kits
    * [Coinbase](https://github.com/alpine-hodler/sdk#coinbase)
  * Auxiliary Packages
    * [protomodel](https://github.com/alpine-hodler/sdk#protomodel)
    * [scalar](https://github.com/alpine-hodler/sdk#scalar)
  * [Resources](https://github.com/alpine-hodler/sdk#resources)

## Install Bazel

Bazel is an open-source build and test tool similar to Make, Maven, and Gradle. It uses a human-readable, high-level build language. Bazel supports projects in multiple languages and builds outputs for multiple platforms. Bazel supports large codebases across multiple repositories, and large numbers of users.

To install, following the intructions [here](https://docs.bazel.build/versions/4.2.2/bazel-overview.html#how-do-i-use-bazel)

If you're on macOS, [you can install Bazel via Homebrew](https://docs.bazel.build/versions/4.2.2/install-os-x.html#step-2-install-bazel-via-homebrew):

```sh
$ brew install bazel
```

## Build SDK

Given the installed dependencies above, just run

```
$ ./Makefile
```

## Install Docker

https://docs.docker.com/get-docker/

We will primarily use docker for generating go and graphqls files from schema/model files.

## Running the GraphQL Server Locally

change directories to the `cmd/graphql` directory and run the following:

```
go build; ./graphql start --port=8080
```

Then navigate to http://localhost:8080/ in the browser.

## Coinbase

`coinbase` is a package meant to be used as an SDK for creating a third-party connection in your code base to read and write information to coinbase pro using auth credentials.  For example:

```go
func CreateOrder() {
	requestOptions := &model.CoinbaseNewOrderOptions{
		Type:        scalar.OrderTypeLimit,
		Side:        scalar.OrderSideBuy,
		Stp:         scalar.OrderSTP_DC,
		Stop:        scalar.OrderStopLoss,
		TimeInForce: scalar.TimeInForceGTC,
		CancelAfter: scalar.CancelAfterMin,
		Price:       1.0,
		Size:        10.0,
		ProductId:   "BTC-USD",
		StopPrice:   500.0,
	}
	coinbase.NewOrders(coinbase.DefaultConnector).Create(requestOptions)
}
```

A slightly more complex example, using websockets:

```go
func Open() {
	ws := coinbase.NewProductWebsocket(coinbase.DefaultWebsocketConnector)
	ticker := ws.Ticker("ETH-USD")
	ticker.Open()
	go func() {
		for row := range ticker.Channel() {
			fmt.Println(row.ProductId, row.Time, row.Price)
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Close()
}
```

### Connecting

To use nealy any of the coinbase accessors, you first need to establish a valid connection to coinbase pro using your auth credentials.  There are curretly two ways of doing this:

#### .env File

You can create a .env file anywhere on your system and import the env variables into go to create a valid connection

```sh
CB_PRO_URL=
CB_PRO_ACCESS_PASSPHRASE=
CB_PRO_ACCESS_KEY=
CB_PRO_SECRET=
```

```go
// initialize the connection
conn, err := coinbase.EnvConnector("/path/to/.env")
if err != nil {
	panic(err)
}

client := coinbase.NewClient(conn)

// then use it for something
accounts := client.Accounts()
fmt.Printf("%+v\n", accounts)
```

### WebSocket

#### Async Ticker

The async ticker runs the coinbase product websocket asyncronously, connecting outside processes by sending the websocket message over a `AsyncTicket.channel`.

```go
	conn := coinbase.DefaultWebsocketConnector
	ws := coinbase.NewProductWebsocket(conn)

	// initialize the ticker object to channel product messages
	ticker := ws.Ticker("ETH-USD")

	// start a go routine that passes product messages concerning ETH-USD currency
	// pair to a channel on the ticker struct
	ticker.Open()
	go func() {

		// Next we range over the product message channel and print the product
		// messages
		for productMsg := range ticker.Channel() {
			fmt.Println(productMsg)
		}
	}()

	// Let the product messages print for 5 seconds
	time.Sleep(5 * time.Seconds)

	// then close the ticker channel, this will unsubscribe from the websocket
	// and close the underlying channel that the messages read to.
	ticker.Close()
```

## GraphQL

To start the graphql server go to the `cmd/graphql` directory and generate the `graphql` binary with `go build`.  Then run

```
./graphql start --port=8080
```

and visit http://localhost:8080/

Here is an example query using the coinbase SDK:

```graphql
query {
  coinbaseWallets {
    id,
    destinationTagName,
    swiftDepositInformation {
      bankName
    },
    sepaDepositInformation {
      bankName
    },
    destinationTagRegex
  }
}
```

## meta

The `meta` library holds code to generate repetative files accross the codebase:

- protomodel
- model
- model accessors
- graph/schema
- endpoint accessors

### Building and Running

Build the dockerfile:

```
docker-compose -f "meta.docker-compose.yaml" up -d --build
```

Run the generate method:

```
docker-compose -f "meta.docker-compose.yaml" run generate
```

### Testing

To test run

```
docker-compose -f "meta.docker-compose.yaml" run test_generate
```


## Protomodel

This package contains the data structures for specific API generated by the `generate.py` script in the `schema` directory. The idea is to programmaticaly add methods, accessors, etc to the data structures here, where a user cannot add code directly, and extend them elsewhere where they can. This will help us hide annoying underlying methods like `Unmarshal`, which is unique to every model.

More conceptually, if we had a schema with a model called `my_special_model` on the `hello` API, it would live in the file `protomodel/hello_my_special_model.go`, it would be encapsulated in the struct called `MySpecialModel`, it would have a test called `protomodel/hello_my_special_model_test.go`, and would be extended by the `model` package like so:

```go
package model

// This file was initialized by schema/generate.py but is open to extension

import (
	"github.com/cryptometrics/cql/protomodel"
)

type MySpecialModel struct {
	protomodel.MySpecialModel
}
```

The file above is generated only once and afterward can be extended to our hearts desire.

### Generate Schema

To re-generate a change to the schema:

```sh
$ gqlgen generate
```

## Scalar

The `scalar` package holds constant data in the style of enums, though most primarily indexed as strings rather than integers.  It inherits the name "scalar" from the [graphql resource for replacing types](https://github.com/graphql-go/graphql/blob/master/scalars.go).  This helps the client/user understand the limitation of certain requests.  For example, the coinbase http request for [creating a new order](https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_postorders) requires a "side" body param.  At the time of this README, "side" is not defined in their docs.  The scalar package removes this ambiguity completely, to create an order request using the coinbase SDK you must use the `model.CoinbaseNewOrderOptions` struct, which types the `Side` field as `scalar.OderSide` which only has two defined constants: `OrderSideBuy` and `OrderSideSell`.

_TL;DR package removes the ambiguity from questionable body/query parameters in http requests._



## Resouces

- [Coinbase Pro Asyncronous Websocket Client Documentation](https://readthedocs.org/projects/copra/downloads/pdf/latest/)
- [Coinbase Pro API Reference](https://docs.pro.coinbase.com/)
- [IEX Cloud API Reference](https://iexcloud.io/docs/api/)
- [Kraken API Reference](https://docs.kraken.com/rest/)
- [Monorepo](https://en.wikipedia.org/wiki/Monorepo)
- [WebSocket](https://en.wikipedia.org/wiki/WebSocket)
