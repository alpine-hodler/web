package twitter

import (
	"time"

	"github.com/alpine-hodler/sdk/pkg/scalar"
)

// * This is a generated file, do not edit

// Annotations are details about annotations relative to the text within a Tweet.
type Annotation struct {
	// End is the end position (zero based) of the text used to annotate the Tweet. While all other end indices are
	// exclusive, this one is inclusive.
	End int `json:"end" bson:"end"`

	// NormalizedText is the text used to determine the annotation type.
	NormalizedText string `json:"normalized_text" bson:"normalized_text"`

	// Probability is the confidence score for the annotation as it correlates to the Tweet text.
	Probability float64 `json:"probability" bson:"probability"`

	// Start is the start position (zero-based) of the text used to annotate the Tweet. All start indices are inclusive.
	Start int `json:"start" bson:"start"`

	// Type is the description of the type of entity identified when the Tweet text was interpreted.
	Type string `json:"type" bson:"type"`
}

// TODO
type Attachments struct {
	// MediaKeys is a list of unique identifiers of media attached to this Tweet. These identifiers use the same media key
	// format as those returned by the Media Library. You can obtain the expanded object in includes.media by adding
	// expansions=attachments.media_keys in the request's query parameter.
	MediaKeys []string `json:"media_keys" bson:"media_keys"`

	// PollIds list of unique identifiers of polls present in the Tweets returned. These are returned as a string in order
	// to avoid complications with languages and tools that cannot handle large integers.
	PollIds []string `json:"poll_ids" bson:"poll_ids"`
}

// TODO
type Bookmark struct {
	// Attachments specifies the type of attachments (if any) present in this Tweet. To return this field, add
	// tweet.fields=attachments in the request's query parameter.
	Attachments *Attachments `json:"attachments" bson:"attachments"`

	// Authur is the inique identifier of this user. This is returned as a string in order to avoid complications with
	// languages and tools that cannot handle large integers. You can obtain the expanded object in includes.users by adding
	// expansions=author_id in the request's query parameter.
	AuthorID string `json:"author_id" bson:"author_id"`

	// Bookmarked indicates whether the user has removed the Bookmark of the specified Tweet. specified Tweet as a result of
	// this request. The returned value is false for a successful request. If the data has been created through a POST
	// method, Bookmarked indicates whether the user bookmarks the specified Tweet as a result of this request.
	Bookmarked bool `json:"bookmarked" bson:"bookmarked"`

	// ContextAnnotations are context annotations for the Tweet. To return this field, add tweet.fields=context_annotations
	// in the request's query parameter.
	ContextAnnotations []*ContextAnnotation `json:"context_annotations" bson:"context_annotations"`

	// ConversationID is the Tweet ID of the original Tweet of the conversation (which includes direct replies, replies of
	// replies). To return this field, add tweet.fields=conversation_id in the request's query parameter.
	ConversationID string `json:"conversation_id" bson:"conversation_id"`

	// CreatedAt is the creation time of the Tweet.
	CreatedAt time.Time `json:"created_at" bson:"created_at"`

	// Entities contain details about text that has a special meaning in a Tweet. To return this field, add
	// tweet.fields=entities in the request's query parameter.
	Entities *Entities `json:"entities" bson:"entities"`

	// Geo contains details about the location tagged by the user in this Tweet, if they specified one. To return this
	// field, add tweet.fields=geo in the request's query parameter.
	Geo *Geo `json:"geo" bson:"geo"`

	// ID is a unique identifier of this Tweet. This is returned as a string in order to avoid complications with languages
	// and tools that cannot handle large integers.
	ID string `json:"id" bson:"id"`

	// InReplyToUserID indicates the user ID of the parent Tweet's author. This is returned as a string in order to avoid
	// complications with languages and tools that cannot handle large integers. You can obtain the expanded object in
	// includes.users by adding expansions=in_reply_to_user_id in the request's query parameter.
	InReplyToUserID string `json:"in_reply_to_user_id" bson:"in_reply_to_user_id"`

	// NonPublicMetrics are non-public engagement metrics for the Tweet at the time of the request. This is a private
	// metric, and requires the use of OAuth 2.0 User Context authentication. To return this field, add
	// tweet.fields=non_public_metrics in the request's query parameter.
	NonPublicMetrics *NonPublicMetrics `json:"non_public_metrics" bson:"non_public_metrics"`

	// OrganicMetrics are organic engagement metrics for the Tweet at the time of the request. Requires user context
	// authentication.
	OrganicMetrics *OrganicMetrics `json:"organic_metrics" bson:"organic_metrics"`

	// PromotedMetrics are engagement metrics for the Tweet at the time of the request in a promoted context. Requires user
	// context authentication.
	PromotedMetrics *PromotedMetrics `json:"promoted_metrics" bson:"promoted_metrics"`

	// PublicMetrics are the engagement metrics for the Tweet at the time of the request. To return this field, add
	// tweet.fields=public_metrics in the request's query parameter.
	PublicMetrics *PublicMetrics `json:"public_metrics" bson:"public_metrics"`

	// ReferencedTweets is a list of Tweets this Tweet refers to. For example, if the parent Tweet is a Retweet, a Retweet
	// with comment (also known as Quoted Tweet) or a Reply, it will include the related Tweet referenced to by its parent.
	// To return this field, add tweet.fields=referenced_tweets in the request's query parameter.
	ReferencedTweets []*ReferencedTweet `json:"referenced_tweets" bson:"referenced_tweets"`

	// Text is the content of the Tweet.
	Text string `json:"text" bson:"text"`

	// Withheld contains withholding details for withheld content. To return this field, add tweet.fields=withheld in the
	// request's query parameter.
	Withheld          *Withheld              `json:"withheld" bson:"withheld"`
	Errors            map[string]interface{} `json:"errors" bson:"errors"`
	Includes          map[string]interface{} `json:"includes" bson:"includes"`
	Lang              string                 `json:"lang" bson:"lang"`
	PossiblySensitive bool                   `json:"possibly_sensitive" bson:"possibly_sensitive"`
	ReplySettings     string                 `json:"reply_settings" bson:"reply_settings"`
	Source            string                 `json:"source" bson:"source"`
}

// BookmarkWrite details the results from a bookmark write operation.
type BookmarkWrite struct {
	Data Bookmarked `json:"data" bson:"data"`
	Meta Meta       `json:"meta" bson:"meta"`
}

// Bookmarked holds details about the status of a bookmark write.
type Bookmarked struct {
	Bookmarked bool `json:"Bookmarked" bson:"Bookmarked"`
}

// TODO
type Bookmarks struct {
	Data []*Bookmark `json:"data" bson:"data"`
	Meta Meta        `json:"meta" bson:"meta"`
}

// Cashtag contains details about text recognized as a Cashtag.
type Cashtag struct {
	// End is the end position (zero-based) of the recognized Cashtag within the Tweet. This end index is exclusive.
	End int `json:"end" bson:"end"`

	// Start is the start position (zero-based) of the recognized Cashtag within the Tweet. All start indices are inclusive.
	Start int `json:"start" bson:"start"`

	// Tag is the text of the Cashtag.
	Tag string `json:"tag" bson:"tag"`
}

// Compliance is some recent compliance jobs.
type Compliance struct {
	// CreatedAt is the date and time when the job was created.
	CreatedAt time.Time `json:"created_at" bson:"created_at"`

	// DownloadExpiresAt the date and time until which the download URL will be available (usually 7 days from the request
	// time).
	DownloadExpiresAt time.Time `json:"download_expires_at" bson:"download_expires_at"`

	// DownloadURL is the predefined location where to download the results from the compliance job. This URL is already
	// signed with an authentication key, so you will not need to pass any additional credential or header to authenticate
	// the request.
	DownloadURL string `json:"download_url" bson:"download_url"`

	// Error returns when jobs.status is failed. Specifies the reason why the job did not complete successfully.
	Error string `json:"error" bson:"error"`

	// ID is the unique identifier for this job.
	ID string `json:"id" bson:"id"`

	// Meta returns meta information about the request.
	Meta Meta `json:"meta" bson:"meta"`

	// Name is the user defined job name. Only returned if specified when the job was created.
	Name string `json:"name" bson:"name"`

	// Status is the status of this job.
	Status scalar.Status `json:"status" bson:"status"`

	// Type is the type of the job, whether tweets or users.
	Type scalar.ComplianceJob `json:"type" bson:"type"`

	// UploadExpiresAt represents the date and time until which the upload URL will be available (usually 15 minutes from
	// the request time).
	UploadExpiresAt time.Time `json:"upload_expires_at" bson:"upload_expires_at"`

	// UploadURL is a URL representing the location where to upload IDs consumed by your app. This URL is already signed
	// with an authentication key, so you will not need to pass any additional credentials or headers to authenticate the
	// request.
	UploadURL string `json:"upload_url" bson:"upload_url"`
}

// TODO
type ContextAnnotation struct {
	// Domain are elements which identify detailed information regarding the domain classification based on Tweet text.
	Domain Domain `json:"domain" bson:"domain"`

	// Entity are elements which identify detailed information regarding the domain classification bases on Tweet text.
	Entity Entity `json:"entity" bson:"entity"`
}

// TODO
type Coordinates struct {
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
	Type        string    `json:"type" bson:"type"`
}

// Domain identifies detailed information regarding the domain classification based on Tweet text.
type Domain struct {
	// Description is the Long form description of domain classification.
	Description string `json:"description" bson:"description"`

	// ID is the numeric value of the domain.
	ID string `json:"id" bson:"id"`

	// Name is the domain name based on the Tweet text.
	Name string `json:"name" bson:"name"`
}

// Entities are details about text that has a special meaning in a Tweet. To return this field, add
// tweet.fields=entities in the request's query parameter.
type Entities struct {
	// Annotations contain details about annotations relative to the text within a Tweet.
	Annotations []*Annotation `json:"annotations" bson:"annotations"`

	// Cashtags contain details about text recognized as a Cashtag.
	Cashtags []*Cashtag `json:"cashtags" bson:"cashtags"`

	// Hashtags contains details about text recognized as a Hashtag.
	Hashtags []*Hashtag `json:"hashtags" bson:"hashtags"`

	// Mentions contains details about text recognized as a user mention.
	Mentions []*Mention `json:"mentions" bson:"mentions"`

	// Urls contains details about text recognized as a URL.
	Urls []*URL `json:"urls" bson:"urls"`
}

// Entity identifies detailed information regarding the domain classification bases on Tweet text.
type Entity struct {
	// Description is additional information regarding referenced entity.
	Description string `json:"description" bson:"description"`

	// ID is a unique value which correlates to an explicitly mentioned Person, Place, Product or Organization.
	ID string `json:"id" bson:"id"`

	// Name is the name or reference of entity referenced in the Tweet.
	Name string `json:"name" bson:"name"`
}

// TODO
type Geo struct {
	Coordinates *Coordinates `json:"coordinates" bson:"coordinates"`
	PlaceID     string       `json:"place_id" bson:"place_id"`
}

// Hashtag contains details about text recognized as a Hashtag.
type Hashtag struct {
	// End is the end position (zero-based) of the recognized Hashtag within the Tweet. This end index is exclusive.
	End int `json:"end" bson:"end"`

	// Start is the start position (zero-based) of the recognized Hashtag within the Tweet. All start indices are inclusive.
	Start int `json:"start" bson:"start"`

	// Tag is the text of the Hashtag.
	Tag string `json:"tag" bson:"tag"`
}

// Mention contains details about text recognized as a user mention.
type Mention struct {
	// End is the end position (zero-based) of the recognized user mention within the Tweet. This end index is exclusive.
	End int `json:"end" bson:"end"`

	// Start is the start position (zero-based) of the recognized user mention within the Tweet. All start indices are
	// inclusive.
	Start int `json:"start" bson:"start"`

	// Username is the part of text recognized as a user mention. You can obtain the expanded object in includes.users by
	// adding expansions=entities.mentions.username in the request's query parameter.
	Username string `json:"username" bson:"username"`
}

// Meta holds metadata concerning requests
type Meta struct {
	ResultCount int `json:"result_count" bson:"result_count"`
}

// NonPublicMetrics are non-public engagement metrics for the Tweet at the time of the request. This is a private
// metric, and requires the use of OAuth 2.0 User Context authentication.
type NonPublicMetrics struct {
	// ImpressionCount are the number of times the Tweet has been viewed. This is a private metric, and requires the use of
	// OAuth 2.0 User Context authentication..
	ImpressionCount int `json:"impression_count" bson:"impression_count"`

	// URLLinkClicks are the number of times a user clicks on a URL link or URL preview card in a Tweet. This is a private
	// metric, and requires the use of OAuth 2.0 User Context authentication.
	URLLinkClicks int `json:"url_link_clicks" bson:"url_link_clicks"`

	// UserProfileClicks are the number of times a user clicks the following portions of a Tweet - display name, user name,
	// profile picture. This is a private metric, and requires the use of OAuth 2.0 User Context authentication.
	UserProfileClicks int `json:"user_profile_clicks" bson:"user_profile_clicks"`
}

// OrganicMetrics are engagement metrics for the Tweet at the time of the request. Requires user context authentication.
type OrganicMetrics struct {
	// ImpressionCount is the number of times the Tweet has been viewed organically. This is a private metric, and requires
	// the use of OAuth 2.0 User Context authentication.
	ImpressionCount int `json:"impression_count" bson:"impression_count"`

	// LikeCount is the number of likes the Tweet has received organically.
	LikeCount int `json:"like_count" bson:"like_count"`

	// ReplyCount is the number of replies the Tweet has received organically.
	ReplyCount int `json:"reply_count" bson:"reply_count"`

	// RetweetCountis the number of times the Tweet has been Retweeted organically.
	RetweetCount int `json:"retweet_count" bson:"retweet_count"`

	// URLLinkClicks is the number of times a user clicks on a URL link or URL preview card in a Tweet organically. This is
	// a private metric, and requires the use of OAuth 2.0 User Context authentication.
	URLLinkClicks int `json:"url_link_clicks" bson:"url_link_clicks"`

	// UserProfileClicks is the number of times a user clicks the following portions of a Tweet organically - display name,
	// user name, profile picture. This is a private metric, and requires the use of OAuth 2.0 User Context authentication.
	UserProfileClicks int `json:"user_profile_clicks" bson:"user_profile_clicks"`
}

// PromotedMetrics are engagement metrics for the Tweet at the time of the request in a promoted context. Requires user
// context authentication.
type PromotedMetrics struct {
	ImpressionCount   int `json:"impression_count" bson:"impression_count"`
	LikeCount         int `json:"like_count" bson:"like_count"`
	ReplyCount        int `json:"reply_count" bson:"reply_count"`
	RetweetCount      int `json:"retweet_count" bson:"retweet_count"`
	URLLinkClicks     int `json:"url_link_clicks" bson:"url_link_clicks"`
	UserProfileClicks int `json:"user_profile_clicks" bson:"user_profile_clicks"`
}

// PublicMetrics are the engagement metrics for the Tweet at the time of the request.
type PublicMetrics struct {
	// LikeCount is the number of Likes of this Tweet.
	LikeCount int `json:"like_count" bson:"like_count"`

	// QuoteCount is the number of times this Tweet has been Retweeted with a comment (also known as Quote).
	QuoteCount int `json:"quote_count" bson:"quote_count"`

	// ReplyCount is the number of Replies of this Tweet.
	ReplyCount int `json:"reply_count" bson:"reply_count"`

	// RetweetCount is the number of times this Tweet has been Retweeted.
	RetweetCount int `json:"retweet_count" bson:"retweet_count"`
}

// ReferencedTweet is a tweet referenced by another tweet.
type ReferencedTweet struct {
	// ID is the unique identifier of the referenced Tweet. You can obtain the expanded object in includes.tweets by adding
	// expansions=referenced_tweets.id in the request's query parameter.
	ID string `json:"id" bson:"id"`

	// Type indicates the type of relationship between this Tweet and the Tweet returned in the response: retweeted (this
	// Tweet is a Retweet), quoted (a Retweet with comment, also known as Quoted Tweet), or replied_to (this Tweet is a
	// reply).
	Type string `json:"type" bson:"type"`
}

// TODO
type Tweet struct {
	// Creation time of the Tweet. To return this field, add tweet.fields=created_at in the request's query parameter.
	CreatedAt time.Time `json:"created_at" bson:"created_at"`

	// ID is a unique identifier of this Tweet. This is returned as a string in order to avoid complications with languages
	// and tools that cannot handle large integers.
	ID string `json:"id" bson:"id"`

	// Text is the content of the Tweet. To return this field, add tweet.fields=text in the request's query parameter.
	Text string `json:"text" bson:"text"`
}

// TODO
type Tweets struct {
	Data []*Tweet `json:"data" bson:"data"`
	Meta Meta     `json:"meta" bson:"meta"`
}

// Entity identifies detailed information regarding the domain classification bases on Tweet text.
type URL struct {
	// DisplayUrl is the URL as displayed in the Twitter client.
	DisplayURL string `json:"display_url" bson:"display_url"`

	// End is the end position (zero-based) of the recognized URL within the Tweet. This end index is exclusive.
	End int `json:"end" bson:"end"`

	// ExpandedUrl is the The fully resolved URL.
	ExpandedURL string `json:"expanded_url" bson:"expanded_url"`

	// Start is the start position (zero-based) of the recognized URL within the Tweet. All start indices are inclusive.
	Start int `json:"start" bson:"start"`

	// UnwoundUrl is the full destination URL.
	UnwoundURL string `json:"unwound_url" bson:"unwound_url"`

	// Url is the URL in the format tweeted by the user.
	URL string `json:"url" bson:"url"`
}

// TODO
type User struct {
	// ID is a unique identifier of this user. This is returned as a string in order to avoid complications with languages
	// and tools that cannot handle large integers.
	ID string `json:"id" bson:"id"`
}

// Withheld contains withholding details for withheld content.
type Withheld struct {
	// Copyright indicates if the content is being withheld for on the basis of copyright infringement.
	Copyright bool `json:"copyright" bson:"copyright"`

	// CountryCodes is a list of countries where this content is not available.
	CountryCodes []string `json:"country_codes" bson:"country_codes"`

	// Scope indicates whether the content being withheld is a Tweet or a user.
	Scope WithholdingScope `json:"scope" bson:"scope"`
}
