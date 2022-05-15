package twitter

import (
	"path"
	"strings"

	"github.com/alpine-hodler/sdk/tools"
)

// * This is a generated file, do not edit
type postAuthority uint8

const (
	_ postAuthority = iota
	AllTweetsPostAuthority
	ComplianceJobsPostAuthority
	MePostAuthority
	TweetsPostAuthority
)

// TODO
func getAllTweetsPostAuthority(builder tools.URIBuilder) (p string) {
	p = path.Join("/2", "tweets", "search", "all")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(builder.QueryPath().String())
	return sb.String()
}

// ComplianceJobs will return a list of recent compliance jobs.
func getComplianceJobsPostAuthority(builder tools.URIBuilder) (p string) {
	p = path.Join("/2", "compliance", "jobs")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(builder.QueryPath().String())
	return sb.String()
}

// TODO
func getMePostAuthority(builder tools.URIBuilder) string {
	return path.Join("/2", "me")
}

// TODO
func getTweetsPostAuthority(builder tools.URIBuilder) (p string) {
	p = path.Join("/2", "tweets")
	var sb strings.Builder
	sb.WriteString(p)
	sb.WriteString(builder.QueryPath().String())
	return sb.String()
}

// Get takes an postAuthority const and postAuthority arguments to parse the URL postAuthority path.
func (pa postAuthority) PostAuthority(builder tools.URIBuilder) string {
	return map[postAuthority]func(builder tools.URIBuilder) string{
		AllTweetsPostAuthority:      getAllTweetsPostAuthority,
		ComplianceJobsPostAuthority: getComplianceJobsPostAuthority,
		MePostAuthority:             getMePostAuthority,
		TweetsPostAuthority:         getTweetsPostAuthority,
	}[pa](builder)
}
