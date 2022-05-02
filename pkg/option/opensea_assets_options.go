package option

// * This is a generated file, do not edit
import "github.com/alpine-hodler/sdk/internal/protomodel"

// OpenseaAssets are options for API requests.
type OpenseaAssets struct {
	protomodel.OpenseaAssetsOptions
}

// OpenseaAssetsFromPrototype will initialize the exportable options struct from the internal prototype.
func OpenseaAssetsFromPrototype(proto *protomodel.OpenseaAssetsOptions) (opts *OpenseaAssets) {
	if proto != nil {
		opts.OpenseaAssetsOptions = *proto
	}
	return
}

// SetOwner sets the Owner field on OpenseaAssets.
func (opts *OpenseaAssets) SetOwner(owner string) *OpenseaAssets {
	opts.Owner = &owner
	return opts
}

// SetTokenIds sets the TokenIds field on OpenseaAssets.
func (opts *OpenseaAssets) SetTokenIds(tokenIds string) *OpenseaAssets {
	opts.TokenIds = &tokenIds
	return opts
}

// SetAssetContractAddress sets the AssetContractAddress field on OpenseaAssets.
func (opts *OpenseaAssets) SetAssetContractAddress(assetContractAddress string) *OpenseaAssets {
	opts.AssetContractAddress = &assetContractAddress
	return opts
}

// SetAssetContractAddresses sets the AssetContractAddresses field on OpenseaAssets.
func (opts *OpenseaAssets) SetAssetContractAddresses(assetContractAddresses []*string) *OpenseaAssets {
	opts.AssetContractAddresses = assetContractAddresses
	return opts
}

// SetOrderBy sets the OrderBy field on OpenseaAssets.
func (opts *OpenseaAssets) SetOrderBy(orderBy string) *OpenseaAssets {
	opts.OrderBy = &orderBy
	return opts
}

// SetOrderDirection sets the OrderDirection field on OpenseaAssets.
func (opts *OpenseaAssets) SetOrderDirection(orderDirection string) *OpenseaAssets {
	opts.OrderDirection = &orderDirection
	return opts
}

// SetOffset sets the Offset field on OpenseaAssets.
func (opts *OpenseaAssets) SetOffset(offset string) *OpenseaAssets {
	opts.Offset = &offset
	return opts
}

// SetLimit sets the Limit field on OpenseaAssets.
func (opts *OpenseaAssets) SetLimit(limit string) *OpenseaAssets {
	opts.Limit = &limit
	return opts
}

// SetCollection sets the Collection field on OpenseaAssets.
func (opts *OpenseaAssets) SetCollection(collection string) *OpenseaAssets {
	opts.Collection = &collection
	return opts
}
