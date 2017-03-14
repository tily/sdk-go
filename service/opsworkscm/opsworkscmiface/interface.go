// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package opsworkscmiface provides an interface to enable mocking the AWS OpsWorks for Chef Automate service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package opsworkscmiface

import (
	"github.com/tily/sdk-go/aws/request"
	"github.com/tily/sdk-go/service/opsworkscm"
)

// OpsWorksCMAPI provides an interface to enable mocking the
// opsworkscm.OpsWorksCM service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS OpsWorks for Chef Automate.
//    func myFunc(svc opsworkscmiface.OpsWorksCMAPI) bool {
//        // Make svc.AssociateNode request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := opsworkscm.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockOpsWorksCMClient struct {
//        opsworkscmiface.OpsWorksCMAPI
//    }
//    func (m *mockOpsWorksCMClient) AssociateNode(input *opsworkscm.AssociateNodeInput) (*opsworkscm.AssociateNodeOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockOpsWorksCMClient{}
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
type OpsWorksCMAPI interface {
	AssociateNodeRequest(*opsworkscm.AssociateNodeInput) (*request.Request, *opsworkscm.AssociateNodeOutput)

	AssociateNode(*opsworkscm.AssociateNodeInput) (*opsworkscm.AssociateNodeOutput, error)

	CreateBackupRequest(*opsworkscm.CreateBackupInput) (*request.Request, *opsworkscm.CreateBackupOutput)

	CreateBackup(*opsworkscm.CreateBackupInput) (*opsworkscm.CreateBackupOutput, error)

	CreateServerRequest(*opsworkscm.CreateServerInput) (*request.Request, *opsworkscm.CreateServerOutput)

	CreateServer(*opsworkscm.CreateServerInput) (*opsworkscm.CreateServerOutput, error)

	DeleteBackupRequest(*opsworkscm.DeleteBackupInput) (*request.Request, *opsworkscm.DeleteBackupOutput)

	DeleteBackup(*opsworkscm.DeleteBackupInput) (*opsworkscm.DeleteBackupOutput, error)

	DeleteServerRequest(*opsworkscm.DeleteServerInput) (*request.Request, *opsworkscm.DeleteServerOutput)

	DeleteServer(*opsworkscm.DeleteServerInput) (*opsworkscm.DeleteServerOutput, error)

	DescribeAccountAttributesRequest(*opsworkscm.DescribeAccountAttributesInput) (*request.Request, *opsworkscm.DescribeAccountAttributesOutput)

	DescribeAccountAttributes(*opsworkscm.DescribeAccountAttributesInput) (*opsworkscm.DescribeAccountAttributesOutput, error)

	DescribeBackupsRequest(*opsworkscm.DescribeBackupsInput) (*request.Request, *opsworkscm.DescribeBackupsOutput)

	DescribeBackups(*opsworkscm.DescribeBackupsInput) (*opsworkscm.DescribeBackupsOutput, error)

	DescribeEventsRequest(*opsworkscm.DescribeEventsInput) (*request.Request, *opsworkscm.DescribeEventsOutput)

	DescribeEvents(*opsworkscm.DescribeEventsInput) (*opsworkscm.DescribeEventsOutput, error)

	DescribeNodeAssociationStatusRequest(*opsworkscm.DescribeNodeAssociationStatusInput) (*request.Request, *opsworkscm.DescribeNodeAssociationStatusOutput)

	DescribeNodeAssociationStatus(*opsworkscm.DescribeNodeAssociationStatusInput) (*opsworkscm.DescribeNodeAssociationStatusOutput, error)

	DescribeServersRequest(*opsworkscm.DescribeServersInput) (*request.Request, *opsworkscm.DescribeServersOutput)

	DescribeServers(*opsworkscm.DescribeServersInput) (*opsworkscm.DescribeServersOutput, error)

	DisassociateNodeRequest(*opsworkscm.DisassociateNodeInput) (*request.Request, *opsworkscm.DisassociateNodeOutput)

	DisassociateNode(*opsworkscm.DisassociateNodeInput) (*opsworkscm.DisassociateNodeOutput, error)

	RestoreServerRequest(*opsworkscm.RestoreServerInput) (*request.Request, *opsworkscm.RestoreServerOutput)

	RestoreServer(*opsworkscm.RestoreServerInput) (*opsworkscm.RestoreServerOutput, error)

	StartMaintenanceRequest(*opsworkscm.StartMaintenanceInput) (*request.Request, *opsworkscm.StartMaintenanceOutput)

	StartMaintenance(*opsworkscm.StartMaintenanceInput) (*opsworkscm.StartMaintenanceOutput, error)

	UpdateServerRequest(*opsworkscm.UpdateServerInput) (*request.Request, *opsworkscm.UpdateServerOutput)

	UpdateServer(*opsworkscm.UpdateServerInput) (*opsworkscm.UpdateServerOutput, error)

	UpdateServerEngineAttributesRequest(*opsworkscm.UpdateServerEngineAttributesInput) (*request.Request, *opsworkscm.UpdateServerEngineAttributesOutput)

	UpdateServerEngineAttributes(*opsworkscm.UpdateServerEngineAttributesInput) (*opsworkscm.UpdateServerEngineAttributesOutput, error)
}

var _ OpsWorksCMAPI = (*opsworkscm.OpsWorksCM)(nil)
