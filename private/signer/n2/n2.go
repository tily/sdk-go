package n2

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/tily/sdk-go/aws"
	"github.com/tily/sdk-go/aws/credentials"
	"github.com/tily/sdk-go/aws/request"
)

var (
	errInvalidMethod = errors.New("n2 signer only handles HTTP POST")
)

const (
	signatureVersion = "2"
	signatureMethod  = "HmacSHA256"
	timeFormat       = "2006-01-02T15:04:05Z"
)

type signer struct {
	// Values that must be populated from the request
	Request     *http.Request
	Time        time.Time
	Credentials *credentials.Credentials
	Debug       aws.LogLevelType
	Logger      aws.Logger

	Query        url.Values
	stringToSign string
	signature    string
}

// SignRequestHandler is a named request handler the SDK will use to sign
// service client request with using the V4 signature.
var SignRequestHandler = request.NamedHandler{
	Name: "n2.SignRequestHandler", Fn: SignSDKRequest,
}

// SignSDKRequest requests with signature version 2.
//
// Will sign the requests with the service config's Credentials object
// Signing is skipped if the credentials is the credentials.AnonymousCredentials
// object.
func SignSDKRequest(req *request.Request) {
	// If the request does not need to be signed ignore the signing of the
	// request if the AnonymousCredentials object is used.
	if req.Config.Credentials == credentials.AnonymousCredentials {
		return
	}

	if req.HTTPRequest.Method != "POST" && req.HTTPRequest.Method != "GET" {
		// The V2 signer only supports GET and POST
		req.Error = errInvalidMethod
		return
	}

	n2 := signer{
		Request:     req.HTTPRequest,
		Time:        req.Time,
		Credentials: req.Config.Credentials,
		Debug:       req.Config.LogLevel.Value(),
		Logger:      req.Config.Logger,
	}

	req.Error = n2.Sign()

	if req.Error != nil {
		return
	}

	if req.HTTPRequest.Method == "POST" {
		// Set the body of the request based on the modified query parameters
		req.SetStringBody(n2.Query.Encode())

		// Now that the body has changed, remove any Content-Length header,
		// because it will be incorrect
		req.HTTPRequest.ContentLength = 0
		req.HTTPRequest.Header.Del("Content-Length")
	} else {
		req.HTTPRequest.URL.RawQuery = n2.Query.Encode()
	}
}

func (n2 *signer) Sign() error {
	credValue, err := n2.Credentials.Get()
	if err != nil {
		return err
	}

	if n2.Request.Method == "POST" {
		// Parse the HTTP request to obtain the query parameters that will
		// be used to build the string to sign. Note that because the HTTP
		// request will need to be modified, the PostForm and Form properties
		// are reset to nil after parsing.
		n2.Request.ParseForm()
		n2.Query = n2.Request.PostForm
		n2.Request.PostForm = nil
		n2.Request.Form = nil
	} else {
		n2.Query = n2.Request.URL.Query()
	}

	// Set new query parameters
	n2.Query.Set("AccessKeyId", credValue.AccessKeyID)
	n2.Query.Set("SignatureVersion", signatureVersion)
	n2.Query.Set("SignatureMethod", signatureMethod)
	n2.Query.Set("Timestamp", n2.Time.UTC().Format(timeFormat))
	if credValue.SessionToken != "" {
		n2.Query.Set("SecurityToken", credValue.SessionToken)
	}

	// in case this is a retry, ensure no signature present
	n2.Query.Del("Signature")

	method := n2.Request.Method
	host := n2.Request.URL.Host
	path := n2.Request.URL.Path
	if path == "" {
		path = "/"
	}

	// obtain all of the query keys and sort them
	queryKeys := make([]string, 0, len(n2.Query))
	for key := range n2.Query {
		queryKeys = append(queryKeys, key)
	}
	sort.Strings(queryKeys)

	// build URL-encoded query keys and values
	queryKeysAndValues := make([]string, len(queryKeys))
	for i, key := range queryKeys {
		k := strings.Replace(url.QueryEscape(key), "+", "%20", -1)
		v := strings.Replace(url.QueryEscape(n2.Query.Get(key)), "+", "%20", -1)
		queryKeysAndValues[i] = k + "=" + v
	}

	// join into one query string
	query := strings.Join(queryKeysAndValues, "&")

	// build the canonical string for the V2 signature
	n2.stringToSign = strings.Join([]string{
		method,
		host,
		path,
		query,
	}, "\n")

	hash := hmac.New(sha256.New, []byte(credValue.SecretAccessKey))
	hash.Write([]byte(n2.stringToSign))
	n2.signature = base64.StdEncoding.EncodeToString(hash.Sum(nil))
	n2.Query.Set("Signature", n2.signature)

	if n2.Debug.Matches(aws.LogDebugWithSigning) {
		n2.logSigningInfo()
	}

	return nil
}

const logSignInfoMsg = `DEBUG: Request Signature:
---[ STRING TO SIGN ]--------------------------------
%s
---[ SIGNATURE ]-------------------------------------
%s
-----------------------------------------------------`

func (n2 *signer) logSigningInfo() {
	msg := fmt.Sprintf(logSignInfoMsg, n2.stringToSign, n2.Query.Get("Signature"))
	n2.Logger.Log(msg)
}
