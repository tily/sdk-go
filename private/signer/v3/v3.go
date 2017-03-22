package v3

import (
	goaws "github.com/mitchellh/goamz/aws"
	"github.com/tily/sdk-go/aws/credentials"
	"github.com/tily/sdk-go/aws/request"
)

var SignRequestHandler = request.NamedHandler{
	Name: "v3.SignRequestHandler", Fn: SignSDKRequest,
}

func SignSDKRequest(req *request.Request) {
	if req.Config.Credentials == credentials.AnonymousCredentials {
		return
	}

	credValue, err := req.Config.Credentials.Get()
	if err != nil {
		return
	}

	auth := goaws.Auth{credValue.AccessKeyID, credValue.SecretAccessKey, credValue.SessionToken}
	method := req.HTTPRequest.Method
	params := make(map[string]string)
	sign(auth, method, params)
	header := req.HTTPRequest.Header
	for k, v := range params {
		header.Set(k, v)
	}
}
