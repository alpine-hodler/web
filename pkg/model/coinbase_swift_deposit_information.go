package model

// * This file was initialized by the meta-program, but is open to modification
import "github.com/alpine-hodler/sdk/internal/protomodel"

// CoinbaseSwiftDepositInformation information regarding a wallet's deposits.
// SWIFT stands for Society for Worldwide Interbank Financial
// Telecommunications. Basically, it's a computer network that connects over 900
// banks around the world – and enables them to transfer money. ING is part of
// this network. There is no fee for accepting deposits into your account with
// ING.
type CoinbaseSwiftDepositInformation struct {
	protomodel.CoinbaseSwiftDepositInformation
}
