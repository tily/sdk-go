package s3

import (
	"github.com/tily/sdk-go/aws/credentials"
	"github.com/tily/sdk-go/aws/request"
	"time"
)

var (
	debug = false
)

var SignRequestHandler = request.NamedHandler{
	Name: "s3.SignRequestHandler", Fn: SignSDKRequest,
}

type Auth struct {
	AccessKey string
	SecretKey string
	Token     string
}

func SignSDKRequest(req *request.Request) {
	if req.Config.Credentials == credentials.AnonymousCredentials {
		return
	}

	credValue, err := req.Config.Credentials.Get()
	if err != nil {
		return
	}

	auth := Auth{credValue.AccessKeyID, credValue.SecretAccessKey, credValue.SessionToken}
	method := req.HTTPRequest.Method
	canonicalPath := req.HTTPRequest.URL.Path
	params := req.HTTPRequest.URL.Query()
	header := req.HTTPRequest.Header
	header["Date"] = []string{time.Now().String()}
	sign(auth, method, canonicalPath, params, header)
}
