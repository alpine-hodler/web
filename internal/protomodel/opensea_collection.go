package protomodel

import (
	"time"

	"github.com/alpine-hodler/sdk/internal/serial"
)

// * This is a generated file, do not edit

// Asset contracts contain data about the contract itself, such as the CryptoKitties contract or the CoolCats contract.
type OpenseaCollection struct {
	BannerImageUrl              string             `json:"banner_image_url" bson:"banner_image_url"`
	ChatUrl                     string             `json:"chat_url" bson:"chat_url"`
	CreatedAt                   time.Time          `json:"created_at" bson:"created_at"`
	DefaultToFiat               bool               `json:"default_to_fiat" bson:"default_to_fiat"`
	Description                 string             `json:"description" bson:"description"`
	DevBuyerFeeBasisPoints      string             `json:"dev_buyer_fee_basis_points" bson:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     string             `json:"dev_seller_fee_basis_points" bson:"dev_seller_fee_basis_points"`
	DiscordUrl                  string             `json:"discord_url" bson:"discord_url"`
	ExternalUrl                 string             `json:"external_url" bson:"external_url"`
	Featured                    bool               `json:"featured" bson:"featured"`
	FeaturedImageUrl            string             `json:"featured_image_url" bson:"featured_image_url"`
	Hidden                      bool               `json:"hidden" bson:"hidden"`
	ImageUrl                    string             `json:"image_url" bson:"image_url"`
	InstagramUsername           string             `json:"instagram_username" bson:"instagram_username"`
	IsSubjectToWhitelist        bool               `json:"is_subject_to_whitelist" bson:"is_subject_to_whitelist"`
	LargeImageUrl               string             `json:"large_image_url" bson:"large_image_url"`
	MediumUsernam               string             `json:"medium_usernam" bson:"medium_usernam"`
	Name                        string             `json:"name" bson:"name"`
	OnlyProxiedTransfers        bool               `json:"only_proxied_transfers" bson:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string             `json:"opensea_buyer_fee_basis_points" bson:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string             `json:"opensea_seller_fee_basis_points" bson:"opensea_seller_fee_basis_points"`
	PayoutAddress               string             `json:"payout_address" bson:"payout_address"`
	ProtoDisplayData            OpenseaDisplayData `json:"display_data" bson:"display_data"`
	RequireEmail                bool               `json:"require_email" bson:"require_email"`
	SafelistRequestStatus       string             `json:"safelist_request_status" bson:"safelist_request_status"`
	ShortDescription            string             `json:"short_description" bson:"short_description"`
	Slub                        string             `json:"slub" bson:"slub"`
	TelegramUrl                 string             `json:"telegram_url" bson:"telegram_url"`
	TwitterUsername             string             `json:"twitter_username" bson:"twitter_username"`
	WikiUrl                     string             `json:"wiki_url" bson:"wiki_url"`
}

// UnmarshalJSON will deserialize bytes into a OpenseaCollection model
func (openseaCollection *OpenseaCollection) UnmarshalJSON(d []byte) error {
	const (
		bannerImageUrlJsonTag              = "banner_image_url"
		chatUrlJsonTag                     = "chat_url"
		createdAtJsonTag                   = "created_at"
		defaultToFiatJsonTag               = "default_to_fiat"
		descriptionJsonTag                 = "description"
		devBuyerFeeBasisPointsJsonTag      = "dev_buyer_fee_basis_points"
		devSellerFeeBasisPointsJsonTag     = "dev_seller_fee_basis_points"
		discordUrlJsonTag                  = "discord_url"
		displayDataJsonTag                 = "display_data"
		externalUrlJsonTag                 = "external_url"
		featuredJsonTag                    = "featured"
		featuredImageUrlJsonTag            = "featured_image_url"
		hiddenJsonTag                      = "hidden"
		safelistRequestStatusJsonTag       = "safelist_request_status"
		imageUrlJsonTag                    = "image_url"
		isSubjectToWhitelistJsonTag        = "is_subject_to_whitelist"
		largeImageUrlJsonTag               = "large_image_url"
		mediumUsernamJsonTag               = "medium_usernam"
		nameJsonTag                        = "name"
		onlyProxiedTransfersJsonTag        = "only_proxied_transfers"
		openseaBuyerFeeBasisPointsJsonTag  = "opensea_buyer_fee_basis_points"
		openseaSellerFeeBasisPointsJsonTag = "opensea_seller_fee_basis_points"
		payoutAddressJsonTag               = "payout_address"
		requireEmailJsonTag                = "require_email"
		shortDescriptionJsonTag            = "short_description"
		slubJsonTag                        = "slub"
		telegramUrlJsonTag                 = "telegram_url"
		twitterUsernameJsonTag             = "twitter_username"
		instagramUsernameJsonTag           = "instagram_username"
		wikiUrlJsonTag                     = "wiki_url"
	)
	data, err := serial.NewJSONTransform(d)
	if err != nil {
		return err
	}
	data.UnmarshalBool(defaultToFiatJsonTag, &openseaCollection.DefaultToFiat)
	data.UnmarshalBool(featuredJsonTag, &openseaCollection.Featured)
	data.UnmarshalBool(hiddenJsonTag, &openseaCollection.Hidden)
	data.UnmarshalBool(isSubjectToWhitelistJsonTag, &openseaCollection.IsSubjectToWhitelist)
	data.UnmarshalBool(onlyProxiedTransfersJsonTag, &openseaCollection.OnlyProxiedTransfers)
	data.UnmarshalBool(requireEmailJsonTag, &openseaCollection.RequireEmail)
	data.UnmarshalString(bannerImageUrlJsonTag, &openseaCollection.BannerImageUrl)
	data.UnmarshalString(chatUrlJsonTag, &openseaCollection.ChatUrl)
	data.UnmarshalString(descriptionJsonTag, &openseaCollection.Description)
	data.UnmarshalString(devBuyerFeeBasisPointsJsonTag, &openseaCollection.DevBuyerFeeBasisPoints)
	data.UnmarshalString(devSellerFeeBasisPointsJsonTag, &openseaCollection.DevSellerFeeBasisPoints)
	data.UnmarshalString(discordUrlJsonTag, &openseaCollection.DiscordUrl)
	data.UnmarshalString(externalUrlJsonTag, &openseaCollection.ExternalUrl)
	data.UnmarshalString(featuredImageUrlJsonTag, &openseaCollection.FeaturedImageUrl)
	data.UnmarshalString(imageUrlJsonTag, &openseaCollection.ImageUrl)
	data.UnmarshalString(instagramUsernameJsonTag, &openseaCollection.InstagramUsername)
	data.UnmarshalString(largeImageUrlJsonTag, &openseaCollection.LargeImageUrl)
	data.UnmarshalString(mediumUsernamJsonTag, &openseaCollection.MediumUsernam)
	data.UnmarshalString(nameJsonTag, &openseaCollection.Name)
	data.UnmarshalString(openseaBuyerFeeBasisPointsJsonTag, &openseaCollection.OpenseaBuyerFeeBasisPoints)
	data.UnmarshalString(openseaSellerFeeBasisPointsJsonTag, &openseaCollection.OpenseaSellerFeeBasisPoints)
	data.UnmarshalString(payoutAddressJsonTag, &openseaCollection.PayoutAddress)
	data.UnmarshalString(safelistRequestStatusJsonTag, &openseaCollection.SafelistRequestStatus)
	data.UnmarshalString(shortDescriptionJsonTag, &openseaCollection.ShortDescription)
	data.UnmarshalString(slubJsonTag, &openseaCollection.Slub)
	data.UnmarshalString(telegramUrlJsonTag, &openseaCollection.TelegramUrl)
	data.UnmarshalString(twitterUsernameJsonTag, &openseaCollection.TwitterUsername)
	data.UnmarshalString(wikiUrlJsonTag, &openseaCollection.WikiUrl)
	err = data.UnmarshalTime(time.RFC3339Nano, createdAtJsonTag, &openseaCollection.CreatedAt)
	if err != nil {
		return err
	}
	openseaCollection.ProtoDisplayData = OpenseaDisplayData{}
	if err := data.UnmarshalStruct(displayDataJsonTag, &openseaCollection.ProtoDisplayData); err != nil {
		return err
	}
	return nil
}
