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

// ComplianceJobs will return a list of recent compliance jobs.
//
// source: https://developer.twitter.com/en/docs/twitter-api/compliance/batch-compliance/introduction
func (c *Client) ComplianceJobs(opts *ComplianceJobsOptions) (m []*Compliance, _ error) {
	req, _ := tools.HTTPNewRequest("GET", "", opts)
	return m, tools.HTTPFetch(http.Client(*c), req, opts, ComplianceJobsPath, nil, &m)
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
