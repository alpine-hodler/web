package polygon

import "github.com/alpine-hodler/web/internal"

// * This is a generated file, do not edit

// Upcoming gets market holidays and their open/close times.
//
// source: https://polygon.io/docs/crypto/get_v1_marketstatus_upcoming
func (c *Client) Upcoming() (m []*Upcoming, _ error) {
	req, _ := internal.HTTPNewRequest("GET", "", nil)
	return m, internal.HTTPFetch(&m, internal.HTTPWithClient(c.Client),
		internal.HTTPWithEncoder(nil),
		internal.HTTPWithEndpoint(UpcomingPath),
		internal.HTTPWithParams(nil),
		internal.HTTPWithRatelimiter(getRateLimiter(UpcomingRatelimiter, 100)),
		internal.HTTPWithRequest(req))
}
