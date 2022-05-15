package twitter

import (
	"time"

	"github.com/alpine-hodler/sdk/pkg/scalar"
)

// * This is a generated file, do not edit

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

// Meta holds metadata concerning requests
type Meta struct {
	ResultCount int `json:"result_count" bson:"result_count"`
}

// TODO
type Tweets struct {
	Data []*TweetDatum `json:"data" bson:"data"`
	Meta Meta          `json:"meta" bson:"meta"`
}

// TODO
type TweetDatum struct {
	// ID is a unique identifier of this Tweet. This is returned as a string in order to avoid complications with languages
	// and tools that cannot handle large integers.
	ID string `json:"id" bson:"id"`

	// Text is the content of the Tweet. To return this field, add tweet.fields=text in the request's query parameter.
	Text string `json:"text" bson:"text"`
}

// TODO
type User struct {
	// ID is a unique identifier of this user. This is returned as a string in order to avoid complications with languages
	// and tools that cannot handle large integers.
	ID string `json:"id" bson:"id"`
}
