package client

import (
	"fmt"
	"net/http"

	"github.com/alpine-hodler/sdk/internal/log"
)

func Logf(level log.Level, req *Request, msg string, args ...interface{}) {
	log.Logf(level, fmt.Sprintf(`(%s,%v) %s`, req.client, req.slug, fmt.Sprintf(msg, args...)))
}

func LogHTTPRequest(req *Request) {
	Logf(log.DEBUG, req, "/%s %v", req.MethodStr(), req.URIPostAuthority())
}

func LogResponseStatus(req *Request, res *http.Response) {
	Logf(log.INFO, req, `{Response:{StatusCode:%v,Status:%s}}`, res.StatusCode, res.Status)
}
