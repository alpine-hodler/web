package client

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func Logf(req *Request, msg string, args ...interface{}) string {
	return fmt.Sprintf(`%s : %v : %s`, req.client, req.slug, fmt.Sprintf(msg, args...))
}

func LogHTTPRequest(req *Request) {
	logrus.Info(Logf(req, "/%s %v", req.MethodStr(), req.EndpointPath()))
}

func LogResponseStatus(req *Request, res *http.Response) {
	logrus.Debug(Logf(req, `{Response:{StatusCode:%v,Status:%s}}`, res.StatusCode, res.Status))
}
