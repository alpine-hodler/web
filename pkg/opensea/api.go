package opensea

import (
	"strings"

	"github.com/alpine-hodler/sdk/pkg/model"
)

// Assets retrieves assets from the NFT API, call the /assets endpoint with the
// desired filter parameters.
//
// * source: https://docs.opensea.io/reference/getting-assets
func (client *C) Assets(opts *model.OpenseaAssetsOptions) (m *model.OpenseaAssets, err error) {
	return m, client.Get(AssetsEndpoint).
		QueryParam("owner", func() (i *string) {
			if opts != nil {
				i = opts.Owner
			}
			return
		}()).
		QueryParam("token_ids", func() (i *string) {
			if opts != nil {
				i = opts.TokenIds
			}
			return
		}()).
		QueryParam("asset_contract_address", func() (i *string) {
			if opts != nil {
				i = opts.AssetContractAddress
			}
			return
		}()).
		QueryParam("asset_contract_addresses", func() (i *string) {
			slice := []string{}
			if opts != nil {
				for _, v := range opts.AssetContractAddresses {
					slice = append(slice, *v)
				}
				tmp := strings.Join(slice, ",")
				i = &tmp
			}
			return
		}()).
		QueryParam("order_by", func() (i *string) {
			if opts != nil {
				i = opts.OrderBy
			}
			return
		}()).
		QueryParam("order_direction", func() (i *string) {
			if opts != nil {
				i = opts.OrderDirection
			}
			return
		}()).
		QueryParam("offset", func() (i *string) {
			if opts != nil {
				i = opts.Offset
			}
			return
		}()).
		QueryParam("limit", func() (i *string) {
			if opts != nil {
				i = opts.Limit
			}
			return
		}()).
		QueryParam("collection", func() (i *string) {
			if opts != nil {
				i = opts.Collection
			}
			return
		}()).
		Fetch().Assign(&m).JoinMessages()
}
