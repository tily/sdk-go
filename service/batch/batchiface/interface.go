// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package batchiface provides an interface to enable mocking the AWS Batch service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package batchiface

import (
	"github.com/tily/sdk-go/aws/request"
	"github.com/tily/sdk-go/service/batch"
)

// BatchAPI provides an interface to enable mocking the
// batch.Batch service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Batch.
//    func myFunc(svc batchiface.BatchAPI) bool {
//        // Make svc.CancelJob request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := batch.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockBatchClient struct {
//        batchiface.BatchAPI
//    }
//    func (m *mockBatchClient) CancelJob(input *batch.CancelJobInput) (*batch.CancelJobOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockBatchClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type BatchAPI interface {
	CancelJobRequest(*batch.CancelJobInput) (*request.Request, *batch.CancelJobOutput)

	CancelJob(*batch.CancelJobInput) (*batch.CancelJobOutput, error)

	CreateComputeEnvironmentRequest(*batch.CreateComputeEnvironmentInput) (*request.Request, *batch.CreateComputeEnvironmentOutput)

	CreateComputeEnvironment(*batch.CreateComputeEnvironmentInput) (*batch.CreateComputeEnvironmentOutput, error)

	CreateJobQueueRequest(*batch.CreateJobQueueInput) (*request.Request, *batch.CreateJobQueueOutput)

	CreateJobQueue(*batch.CreateJobQueueInput) (*batch.CreateJobQueueOutput, error)

	DeleteComputeEnvironmentRequest(*batch.DeleteComputeEnvironmentInput) (*request.Request, *batch.DeleteComputeEnvironmentOutput)

	DeleteComputeEnvironment(*batch.DeleteComputeEnvironmentInput) (*batch.DeleteComputeEnvironmentOutput, error)

	DeleteJobQueueRequest(*batch.DeleteJobQueueInput) (*request.Request, *batch.DeleteJobQueueOutput)

	DeleteJobQueue(*batch.DeleteJobQueueInput) (*batch.DeleteJobQueueOutput, error)

	DeregisterJobDefinitionRequest(*batch.DeregisterJobDefinitionInput) (*request.Request, *batch.DeregisterJobDefinitionOutput)

	DeregisterJobDefinition(*batch.DeregisterJobDefinitionInput) (*batch.DeregisterJobDefinitionOutput, error)

	DescribeComputeEnvironmentsRequest(*batch.DescribeComputeEnvironmentsInput) (*request.Request, *batch.DescribeComputeEnvironmentsOutput)

	DescribeComputeEnvironments(*batch.DescribeComputeEnvironmentsInput) (*batch.DescribeComputeEnvironmentsOutput, error)

	DescribeJobDefinitionsRequest(*batch.DescribeJobDefinitionsInput) (*request.Request, *batch.DescribeJobDefinitionsOutput)

	DescribeJobDefinitions(*batch.DescribeJobDefinitionsInput) (*batch.DescribeJobDefinitionsOutput, error)

	DescribeJobQueuesRequest(*batch.DescribeJobQueuesInput) (*request.Request, *batch.DescribeJobQueuesOutput)

	DescribeJobQueues(*batch.DescribeJobQueuesInput) (*batch.DescribeJobQueuesOutput, error)

	DescribeJobsRequest(*batch.DescribeJobsInput) (*request.Request, *batch.DescribeJobsOutput)

	DescribeJobs(*batch.DescribeJobsInput) (*batch.DescribeJobsOutput, error)

	ListJobsRequest(*batch.ListJobsInput) (*request.Request, *batch.ListJobsOutput)

	ListJobs(*batch.ListJobsInput) (*batch.ListJobsOutput, error)

	RegisterJobDefinitionRequest(*batch.RegisterJobDefinitionInput) (*request.Request, *batch.RegisterJobDefinitionOutput)

	RegisterJobDefinition(*batch.RegisterJobDefinitionInput) (*batch.RegisterJobDefinitionOutput, error)

	SubmitJobRequest(*batch.SubmitJobInput) (*request.Request, *batch.SubmitJobOutput)

	SubmitJob(*batch.SubmitJobInput) (*batch.SubmitJobOutput, error)

	TerminateJobRequest(*batch.TerminateJobInput) (*request.Request, *batch.TerminateJobOutput)

	TerminateJob(*batch.TerminateJobInput) (*batch.TerminateJobOutput, error)

	UpdateComputeEnvironmentRequest(*batch.UpdateComputeEnvironmentInput) (*request.Request, *batch.UpdateComputeEnvironmentOutput)

	UpdateComputeEnvironment(*batch.UpdateComputeEnvironmentInput) (*batch.UpdateComputeEnvironmentOutput, error)

	UpdateJobQueueRequest(*batch.UpdateJobQueueInput) (*request.Request, *batch.UpdateJobQueueOutput)

	UpdateJobQueue(*batch.UpdateJobQueueInput) (*batch.UpdateJobQueueOutput, error)
}

var _ BatchAPI = (*batch.Batch)(nil)
