package polygon

import "path"

// * This is a generated file, do not edit

type rawPath uint8

const (
	_ rawPath = iota
	UpcomingPath
)

// Upcoming gets market holidays and their open/close times.
func getUpcomingPath(params map[string]string) string {
	return path.Join("/v1", "marketstatus", "upcoming")
}

// Get takes an rawPath const and rawPath arguments to parse the URL rawPath path.
func (p rawPath) Path(params map[string]string) string {
	return map[rawPath]func(map[string]string) string{
		UpcomingPath: getUpcomingPath,
	}[p](params)
}

func (p rawPath) Scope() string {
	return map[rawPath]string{}[p]
}
