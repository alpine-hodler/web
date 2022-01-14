package protomodel

import "github.com/alpine-hodler/sdk/internal/serial"

// * This is a generated file, do not edit

// Asset contracts contain data about the contract itself, such as the
// CryptoKitties contract or the CoolCats contract.
type OpenseaAssetContract struct {
	Address                     string  `json:"address"`
	AssetContractType           string  `json:"asset_contract_type"`
	BuyerFeeBasisPoints         int     `json:"buyer_fee_basis_points"`
	CreatedDate                 string  `json:"created_date"`
	DefaultToFiat               bool    `json:"default_to_fiat"`
	Description                 string  `json:"description"`
	DevBuyFeeBasisPoints        int     `json:"dev_buy_fee_basis_points"`
	DevSellerFeeBasisPoints     int     `json:"dev_seller_fee_basis_points"`
	ExternalLink                string  `json:"external_link"`
	ImageUrl                    string  `json:"image_url"`
	Name                        string  `json:"name"`
	NftVersion                  string  `json:"nft_version"`
	OnlyProxiedTransfers        bool    `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int     `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int     `json:"opensea_seller_fee_basis_points"`
	OpenseaVersion              string  `json:"opensea_version"`
	Owner                       int     `json:"owner"`
	PayoutAddress               string  `json:"payout_address"`
	SchemaName                  string  `json:"schema_name"`
	SellerFeeBasisPoints        int     `json:"seller_Fee_basis_points"`
	Symbol                      string  `json:"symbol"`
	TotalSupply                 float64 `json:"total_supply"`
}

// UnmarshalJSON will deserialize bytes into a OpenseaAssetContract model
func (openseaAssetContract *OpenseaAssetContract) UnmarshalJSON(d []byte) error {
	const (
		addressJsonTag                     = "address"
		assetContractTypeJsonTag           = "asset_contract_type"
		createdDateJsonTag                 = "created_date"
		nameJsonTag                        = "name"
		nftVersionJsonTag                  = "nft_version"
		openseaVersionJsonTag              = "opensea_version"
		ownerJsonTag                       = "owner"
		schemaNameJsonTag                  = "schema_name"
		symbolJsonTag                      = "symbol"
		totalSupplyJsonTag                 = "total_supply"
		descriptionJsonTag                 = "description"
		externalLinkJsonTag                = "external_link"
		imageUrlJsonTag                    = "image_url"
		defaultToFiatJsonTag               = "default_to_fiat"
		devBuyFeeBasisPointsJsonTag        = "dev_buy_fee_basis_points"
		devSellerFeeBasisPointsJsonTag     = "dev_seller_fee_basis_points"
		onlyProxiedTransfersJsonTag        = "only_proxied_transfers"
		openseaBuyerFeeBasisPointsJsonTag  = "opensea_buyer_fee_basis_points"
		openseaSellerFeeBasisPointsJsonTag = "opensea_seller_fee_basis_points"
		buyerFeeBasisPointsJsonTag         = "buyer_fee_basis_points"
		sellerFeeBasisPointsJsonTag        = "seller_Fee_basis_points"
		payoutAddressJsonTag               = "payout_address"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(defaultToFiatJsonTag, &openseaAssetContract.DefaultToFiat)
	data.UnmarshalBool(onlyProxiedTransfersJsonTag, &openseaAssetContract.OnlyProxiedTransfers)
	data.UnmarshalFloatFromString(totalSupplyJsonTag, &openseaAssetContract.TotalSupply)
	data.UnmarshalInt(buyerFeeBasisPointsJsonTag, &openseaAssetContract.BuyerFeeBasisPoints)
	data.UnmarshalInt(devBuyFeeBasisPointsJsonTag, &openseaAssetContract.DevBuyFeeBasisPoints)
	data.UnmarshalInt(devSellerFeeBasisPointsJsonTag, &openseaAssetContract.DevSellerFeeBasisPoints)
	data.UnmarshalInt(openseaBuyerFeeBasisPointsJsonTag, &openseaAssetContract.OpenseaBuyerFeeBasisPoints)
	data.UnmarshalInt(openseaSellerFeeBasisPointsJsonTag, &openseaAssetContract.OpenseaSellerFeeBasisPoints)
	data.UnmarshalInt(ownerJsonTag, &openseaAssetContract.Owner)
	data.UnmarshalInt(sellerFeeBasisPointsJsonTag, &openseaAssetContract.SellerFeeBasisPoints)
	data.UnmarshalString(addressJsonTag, &openseaAssetContract.Address)
	data.UnmarshalString(assetContractTypeJsonTag, &openseaAssetContract.AssetContractType)
	data.UnmarshalString(createdDateJsonTag, &openseaAssetContract.CreatedDate)
	data.UnmarshalString(descriptionJsonTag, &openseaAssetContract.Description)
	data.UnmarshalString(externalLinkJsonTag, &openseaAssetContract.ExternalLink)
	data.UnmarshalString(imageUrlJsonTag, &openseaAssetContract.ImageUrl)
	data.UnmarshalString(nameJsonTag, &openseaAssetContract.Name)
	data.UnmarshalString(nftVersionJsonTag, &openseaAssetContract.NftVersion)
	data.UnmarshalString(openseaVersionJsonTag, &openseaAssetContract.OpenseaVersion)
	data.UnmarshalString(payoutAddressJsonTag, &openseaAssetContract.PayoutAddress)
	data.UnmarshalString(schemaNameJsonTag, &openseaAssetContract.SchemaName)
	data.UnmarshalString(symbolJsonTag, &openseaAssetContract.Symbol)
	return nil
}
