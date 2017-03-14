// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudsearch

import (
	"github.com/tily/sdk-go/aws"
	"github.com/tily/sdk-go/aws/client"
	"github.com/tily/sdk-go/aws/client/metadata"
	"github.com/tily/sdk-go/aws/request"
	"github.com/tily/sdk-go/aws/signer/v4"
	"github.com/tily/sdk-go/private/protocol/query"
)

// You use the Amazon CloudSearch configuration service to create, configure,
// and manage search domains. Configuration service requests are submitted using
// the AWS Query protocol. AWS Query requests are HTTP or HTTPS requests submitted
// via HTTP GET or POST with a query parameter named Action.
//
// The endpoint for configuration service requests is region-specific: cloudsearch.region.amazonaws.com.
// For example, cloudsearch.us-east-1.amazonaws.com. For a current list of supported
// regions and endpoints, see Regions and Endpoints (http://docs.aws.amazon.com/general/latest/gr/rande.html#cloudsearch_region).
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type CloudSearch struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "cloudsearch" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName   // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the CloudSearch client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a CloudSearch client from just a session.
//     svc := cloudsearch.New(mySession)
//
//     // Create a CloudSearch client with additional configuration
//     svc := cloudsearch.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *CloudSearch {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *CloudSearch {
	svc := &CloudSearch{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2013-01-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a CloudSearch operation and runs any
// custom request initialization.
func (c *CloudSearch) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
