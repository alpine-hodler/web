package twitter

// * This is a generated file, do not edit

// TODO
//
// source: https://developer.twitter.com/en/docs/twitter-api/tweets/search/api-reference/get-tweets-search-all
func (c *C) AllTweets(opts *AllTweetsOptions) (m *Tweets, _ error) {
	return m, c.Get(AllTweetsPostAuthority).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// ComplianceJobs will return a list of recent compliance jobs.
//
// source: https://developer.twitter.com/en/docs/twitter-api/compliance/batch-compliance/introduction
func (c *C) ComplianceJobs(opts *ComplianceJobsOptions) (m []*Compliance, _ error) {
	return m, c.Get(ComplianceJobsPostAuthority).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}

// TODO
//
// source: https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-me
func (c *C) Me() (m []*User, _ error) {
	return m, c.Get(MePostAuthority).
		Fetch().Assign(&m).JoinMessages()
}

// TODO
//
// source: https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets
func (c *C) Tweets(opts *TweetsOptions) (m *Tweets, _ error) {
	return m, c.Get(TweetsPostAuthority).
		SetQueryParams(opts).
		Fetch().Assign(&m).JoinMessages()
}
