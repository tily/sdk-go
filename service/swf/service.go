// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package swf

import (
	"github.com/tily/sdk-go/aws"
	"github.com/tily/sdk-go/aws/client"
	"github.com/tily/sdk-go/aws/client/metadata"
	"github.com/tily/sdk-go/aws/request"
	"github.com/tily/sdk-go/aws/signer/v4"
	"github.com/tily/sdk-go/private/protocol/jsonrpc"
)

// The Amazon Simple Workflow Service (Amazon SWF) makes it easy to build applications
// that use Amazon's cloud to coordinate work across distributed components.
// In Amazon SWF, a task represents a logical unit of work that is performed
// by a component of your workflow. Coordinating tasks in a workflow involves
// managing intertask dependencies, scheduling, and concurrency in accordance
// with the logical flow of the application.
//
// Amazon SWF gives you full control over implementing tasks and coordinating
// them without worrying about underlying complexities such as tracking their
// progress and maintaining their state.
//
// This documentation serves as reference only. For a broader overview of the
// Amazon SWF programming model, see the Amazon SWF Developer Guide (http://docs.aws.amazon.com/amazonswf/latest/developerguide/).
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type SWF struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "swf"       // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the SWF client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a SWF client from just a session.
//     svc := swf.New(mySession)
//
//     // Create a SWF client with additional configuration
//     svc := swf.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *SWF {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *SWF {
	svc := &SWF{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2012-01-25",
				JSONVersion:   "1.0",
				TargetPrefix:  "SimpleWorkflowService",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a SWF operation and runs any
// custom request initialization.
func (c *SWF) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
