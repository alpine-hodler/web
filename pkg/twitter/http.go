package twitter

import (
	"net/http"

	"github.com/alpine-hodler/sdk/tools"
)

// * This is a generated file, do not edit

// TODO
//
// source: https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-all
func (c *Client) AllTweets(opts *AllTweetsOptions) (m *Tweets, _ error) {
	req, _ := tools.HTTPNewRequest("GET", "", opts)
	return m, tools.HTTPFetch(http.Client(*c), req, opts, AllTweetsPath, nil, &m)
}

// Bookmarks allows you to get information about a authenticated user’s 800 most recent bookmarked Tweets. This request
// requires OAuth 2.0 Authorization Code with PKCE.
//
// source: https://developer.twitter.com/en/docs/twitter-api/tweets/bookmarks/api-reference/get-users-id-bookmarks
func (c *Client) Bookmarks(userId string, opts *BookmarksOptions) (m *Bookmarks, _ error) {
	req, _ := tools.HTTPNewRequest("GET", "", opts)
	return m, tools.HTTPFetch(http.Client(*c), req, opts, BookmarksPath, map[string]string{
		"user_id": userId,
	}, &m)
}

// ComplianceJobs will return a list of recent compliance jobs.
//
// source: https://developer.twitter.com/en/docs/twitter-api/compliance/batch-compliance/introduction
func (c *Client) ComplianceJobs(opts *ComplianceJobsOptions) (m []*Compliance, _ error) {
	req, _ := tools.HTTPNewRequest("GET", "", opts)
	return m, tools.HTTPFetch(http.Client(*c), req, opts, ComplianceJobsPath, nil, &m)
}

// CreateBookmarks causes the user ID of an authenticated user identified in the path parameter to Bookmark the target
// Tweet provided in the request body. This request requires OAuth 2.0 Authorization Code with PKCE.
//
// source: https://developer.twitter.com/en/docs/twitter-api/tweets/bookmarks/api-reference/post-users-id-bookmarks
func (c *Client) CreateBookmark(userId string, opts *CreateBookmarkOptions) (m BookmarkWrite, _ error) {
	req, _ := tools.HTTPNewRequest("POST", "", opts)
	return m, tools.HTTPFetch(http.Client(*c), req, nil, CreateBookmarkPath, map[string]string{
		"user_id": userId,
	}, &m)
}

// DeleteBookmarks are a core feature of the Twitter app that allows you to “save” Tweets and easily access them later.
// With these endpoints, you can retrieve, create, delete or build solutions to manage your Bookmarks via the API. This
// request requires OAuth 2.0 Authorization Code with PKCE
//
// source: https://developer.twitter.com/en/docs/twitter-api/tweets/bookmarks/api-reference/delete-users-id-bookmarks-tweet_id
func (c *Client) DeleteBookmark(userId string, tweetId string) (m BookmarkWrite, _ error) {
	req, _ := tools.HTTPNewRequest("DELETE", "", nil)
	return m, tools.HTTPFetch(http.Client(*c), req, nil, DeleteBookmarkPath, map[string]string{
		"user_id": userId, "tweet_id": tweetId,
	}, &m)
}

// TODO
//
// source: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-me
func (c *Client) Me() (m []*User, _ error) {
	req, _ := tools.HTTPNewRequest("GET", "", nil)
	return m, tools.HTTPFetch(http.Client(*c), req, nil, MePath, nil, &m)
}

// TODO
//
// source: https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
func (c *Client) Tweets(opts *TweetsOptions) (m *Tweets, _ error) {
	req, _ := tools.HTTPNewRequest("GET", "", opts)
	return m, tools.HTTPFetch(http.Client(*c), req, opts, TweetsPath, nil, &m)
}
