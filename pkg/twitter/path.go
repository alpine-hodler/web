package twitter

import "path"

// * This is a generated file, do not edit
type rawPath uint8

const (
	_ rawPath = iota
	AllTweetsPath
	ComplianceJobsPath
	MePath
	TweetsPath
)

// TODO
func getAllTweetsPath(params map[string]string) string {
	return path.Join("/2", "tweets", "search", "all")
}

// ComplianceJobs will return a list of recent compliance jobs.
func getComplianceJobsPath(params map[string]string) string {
	return path.Join("/2", "compliance", "jobs")
}

// TODO
func getMePath(params map[string]string) string {
	return path.Join("/2", "me")
}

// TODO
func getTweetsPath(params map[string]string) string {
	return path.Join("/2", "tweets")
}

// Get takes an rawPath const and rawPath arguments to parse the URL rawPath path.
func (p rawPath) Path(params map[string]string) string {
	return map[rawPath]func(map[string]string) string{
		AllTweetsPath:      getAllTweetsPath,
		ComplianceJobsPath: getComplianceJobsPath,
		MePath:             getMePath,
		TweetsPath:         getTweetsPath,
	}[p](params)
}
