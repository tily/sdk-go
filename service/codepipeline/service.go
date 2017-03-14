// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package codepipeline

import (
	"github.com/tily/sdk-go/aws"
	"github.com/tily/sdk-go/aws/client"
	"github.com/tily/sdk-go/aws/client/metadata"
	"github.com/tily/sdk-go/aws/request"
	"github.com/tily/sdk-go/aws/signer/v4"
	"github.com/tily/sdk-go/private/protocol/jsonrpc"
)

// Overview
//
// This is the AWS CodePipeline API Reference. This guide provides descriptions
// of the actions and data types for AWS CodePipeline. Some functionality for
// your pipeline is only configurable through the API. For additional information,
// see the AWS CodePipeline User Guide (http://docs.aws.amazon.com/codepipeline/latest/userguide/welcome.html).
//
// You can use the AWS CodePipeline API to work with pipelines, stages, actions,
// gates, and transitions, as described below.
//
// Pipelines are models of automated release processes. Each pipeline is uniquely
// named, and consists of actions, gates, and stages.
//
// You can work with pipelines by calling:
//
//    * CreatePipeline, which creates a uniquely-named pipeline.
//
//    * DeletePipeline, which deletes the specified pipeline.
//
//    * GetPipeline, which returns information about a pipeline structure.
//
//    * GetPipelineExecution, which returns information about a specific execution
//    of a pipeline.
//
//    * GetPipelineState, which returns information about the current state
//    of the stages and actions of a pipeline.
//
//    * ListPipelines, which gets a summary of all of the pipelines associated
//    with your account.
//
//    * StartPipelineExecution, which runs the the most recent revision of an
//    artifact through the pipeline.
//
//    * UpdatePipeline, which updates a pipeline with edits or changes to the
//    structure of the pipeline.
//
// Pipelines include stages, which are logical groupings of gates and actions.
// Each stage contains one or more actions that must complete before the next
// stage begins. A stage will result in success or failure. If a stage fails,
// then the pipeline stops at that stage and will remain stopped until either
// a new version of an artifact appears in the source location, or a user takes
// action to re-run the most recent artifact through the pipeline. You can call
// GetPipelineState, which displays the status of a pipeline, including the
// status of stages in the pipeline, or GetPipeline, which returns the entire
// structure of the pipeline, including the stages of that pipeline. For more
// information about the structure of stages and actions, also refer to the
// AWS CodePipeline Pipeline Structure Reference (http://docs.aws.amazon.com/codepipeline/latest/userguide/pipeline-structure.html).
//
// Pipeline stages include actions, which are categorized into categories such
// as source or build actions performed within a stage of a pipeline. For example,
// you can use a source action to import artifacts into a pipeline from a source
// such as Amazon S3. Like stages, you do not work with actions directly in
// most cases, but you do define and interact with actions when working with
// pipeline operations such as CreatePipeline and GetPipelineState.
//
// Pipelines also include transitions, which allow the transition of artifacts
// from one stage to the next in a pipeline after the actions in one stage complete.
//
// You can work with transitions by calling:
//
//    * DisableStageTransition, which prevents artifacts from transitioning
//    to the next stage in a pipeline.
//
//    * EnableStageTransition, which enables transition of artifacts between
//    stages in a pipeline.
//
// Using the API to integrate with AWS CodePipeline
//
// For third-party integrators or developers who want to create their own integrations
// with AWS CodePipeline, the expected sequence varies from the standard API
// user. In order to integrate with AWS CodePipeline, developers will need to
// work with the following items:
//
// Jobs, which are instances of an action. For example, a job for a source action
// might import a revision of an artifact from a source.
//
// You can work with jobs by calling:
//
//    * AcknowledgeJob, which confirms whether a job worker has received the
//    specified job,
//
//    * GetJobDetails, which returns the details of a job,
//
//    * PollForJobs, which determines whether there are any jobs to act upon,
//
//
//    * PutJobFailureResult, which provides details of a job failure, and
//
//    * PutJobSuccessResult, which provides details of a job success.
//
// Third party jobs, which are instances of an action created by a partner action
// and integrated into AWS CodePipeline. Partner actions are created by members
// of the AWS Partner Network.
//
// You can work with third party jobs by calling:
//
//    * AcknowledgeThirdPartyJob, which confirms whether a job worker has received
//    the specified job,
//
//    * GetThirdPartyJobDetails, which requests the details of a job for a partner
//    action,
//
//    * PollForThirdPartyJobs, which determines whether there are any jobs to
//    act upon,
//
//    * PutThirdPartyJobFailureResult, which provides details of a job failure,
//    and
//
//    * PutThirdPartyJobSuccessResult, which provides details of a job success.
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09
type CodePipeline struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "codepipeline" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName    // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the CodePipeline client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a CodePipeline client from just a session.
//     svc := codepipeline.New(mySession)
//
//     // Create a CodePipeline client with additional configuration
//     svc := codepipeline.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *CodePipeline {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *CodePipeline {
	svc := &CodePipeline{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-07-09",
				JSONVersion:   "1.1",
				TargetPrefix:  "CodePipeline_20150709",
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

// newRequest creates a new request for a CodePipeline operation and runs any
// custom request initialization.
func (c *CodePipeline) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
