package model

// * This file was initialized by the meta-program, but is open to modification
import "github.com/alpine-hodler/sdk/internal/protomodel"

// Asset is the primary object in the OpenSea API is the, representing a unique
// digital item whose ownership is managed by the blockchain. The below
// CryptoSaga hero is an example of an asset shown on OpenSea.
type OpenseaAsset struct{ protomodel.OpenseaAsset }
