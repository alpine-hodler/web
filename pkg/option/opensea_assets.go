package option

// * This is a generated file, do not edit

// OpenseaAssets are options for API requests.
type OpenseaAssets struct {
	AssetContractAddress   *string   `json:"asset_contract_address" bson:"asset_contract_address"`
	AssetContractAddresses []*string `json:"asset_contract_addresses" bson:"asset_contract_addresses"`
	Collection             *string   `json:"collection" bson:"collection"`
	Limit                  *string   `json:"limit" bson:"limit"`
	Offset                 *string   `json:"offset" bson:"offset"`
	OrderBy                *string   `json:"order_by" bson:"order_by"`
	OrderDirection         *string   `json:"order_direction" bson:"order_direction"`
	Owner                  *string   `json:"owner" bson:"owner"`
	TokenIds               *string   `json:"token_ids" bson:"token_ids"`
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
