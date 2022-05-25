package twitter

import (
	"time"

	"github.com/alpine-hodler/sdk/pkg/scalar"
)

// * This is a generated file, do not edit

// TODO
type Attachment struct {
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
	Attachments []Attachment `json:"attachments" bson:"attachments"`

	// Authur is the inique identifier of this user. This is returned as a string in order to avoid complications with
	// languages and tools that cannot handle large integers. You can obtain the expanded object in includes.users by adding
	// expansions=author_id in the request's query parameter.
	AuthorID string `json:"author_id" bson:"author_id"`

	// Bookmarked indicates whether the user has removed the Bookmark of the specified Tweet. specified Tweet as a result of
	// this request. The returned value is false for a successful request. If the data has been created through a POST
	// method, Bookmarked indicates whether the user bookmarks the specified Tweet as a result of this request.
	Bookmarked bool `json:"bookmarked" bson:"bookmarked"`

	// ConversationID is the Tweet ID of the original Tweet of the conversation (which includes direct replies, replies of
	// replies). To return this field, add tweet.fields=conversation_id in the request's query parameter.
	ConversationID string `json:"conversation_id" bson:"conversation_id"`

	// CreatedAt is the creation time of the Tweet.
	CreatedAt time.Time `json:"created_at" bson:"created_at"`

	// Geo contains details about the location tagged by the user in this Tweet, if they specified one. To return this
	// field, add tweet.fields=geo in the request's query parameter.
	Geo Geo `json:"geo" bson:"geo"`

	// ID is a unique identifier of this Tweet. This is returned as a string in order to avoid complications with languages
	// and tools that cannot handle large integers.
	ID string `json:"id" bson:"id"`

	// InReplyToUserID indicates the user ID of the parent Tweet's author. This is returned as a string in order to avoid
	// complications with languages and tools that cannot handle large integers. You can obtain the expanded object in
	// includes.users by adding expansions=in_reply_to_user_id in the request's query parameter.
	InReplyToUserID string `json:"in_reply_to_user_id" bson:"in_reply_to_user_id"`

	// ReferencedTweets is a list of Tweets this Tweet refers to. For example, if the parent Tweet is a Retweet, a Retweet
	// with comment (also known as Quoted Tweet) or a Reply, it will include the related Tweet referenced to by its parent.
	// To return this field, add tweet.fields=referenced_tweets in the request's query parameter.
	ReferencedTweets []*ReferencedTweet `json:"referenced_tweets" bson:"referenced_tweets"`

	// Text is the content of the Tweet.
	Text string `json:"text" bson:"text"`
}

// TODO
type Bookmarks struct {
	Data []*Bookmark `json:"data" bson:"data"`
	Meta Meta        `json:"meta" bson:"meta"`
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
type Coordinates struct {
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
	Type        string    `json:"type" bson:"type"`
}

// TODO
type Geo struct {
	Coordinates *Coordinates `json:"coordinates" bson:"coordinates"`
	PlaceID     string       `json:"place_id" bson:"place_id"`
}

// Meta holds metadata concerning requests
type Meta struct {
	ResultCount int `json:"result_count" bson:"result_count"`
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

// TODO
type User struct {
	// ID is a unique identifier of this user. This is returned as a string in order to avoid complications with languages
	// and tools that cannot handle large integers.
	ID string `json:"id" bson:"id"`
}
