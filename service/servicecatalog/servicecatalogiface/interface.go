// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package servicecatalogiface provides an interface to enable mocking the AWS Service Catalog service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package servicecatalogiface

import (
	"github.com/tily/sdk-go/aws/request"
	"github.com/tily/sdk-go/service/servicecatalog"
)

// ServiceCatalogAPI provides an interface to enable mocking the
// servicecatalog.ServiceCatalog service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Service Catalog.
//    func myFunc(svc servicecatalogiface.ServiceCatalogAPI) bool {
//        // Make svc.AcceptPortfolioShare request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := servicecatalog.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockServiceCatalogClient struct {
//        servicecatalogiface.ServiceCatalogAPI
//    }
//    func (m *mockServiceCatalogClient) AcceptPortfolioShare(input *servicecatalog.AcceptPortfolioShareInput) (*servicecatalog.AcceptPortfolioShareOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockServiceCatalogClient{}
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
type ServiceCatalogAPI interface {
	AcceptPortfolioShareRequest(*servicecatalog.AcceptPortfolioShareInput) (*request.Request, *servicecatalog.AcceptPortfolioShareOutput)

	AcceptPortfolioShare(*servicecatalog.AcceptPortfolioShareInput) (*servicecatalog.AcceptPortfolioShareOutput, error)

	AssociatePrincipalWithPortfolioRequest(*servicecatalog.AssociatePrincipalWithPortfolioInput) (*request.Request, *servicecatalog.AssociatePrincipalWithPortfolioOutput)

	AssociatePrincipalWithPortfolio(*servicecatalog.AssociatePrincipalWithPortfolioInput) (*servicecatalog.AssociatePrincipalWithPortfolioOutput, error)

	AssociateProductWithPortfolioRequest(*servicecatalog.AssociateProductWithPortfolioInput) (*request.Request, *servicecatalog.AssociateProductWithPortfolioOutput)

	AssociateProductWithPortfolio(*servicecatalog.AssociateProductWithPortfolioInput) (*servicecatalog.AssociateProductWithPortfolioOutput, error)

	CreateConstraintRequest(*servicecatalog.CreateConstraintInput) (*request.Request, *servicecatalog.CreateConstraintOutput)

	CreateConstraint(*servicecatalog.CreateConstraintInput) (*servicecatalog.CreateConstraintOutput, error)

	CreatePortfolioRequest(*servicecatalog.CreatePortfolioInput) (*request.Request, *servicecatalog.CreatePortfolioOutput)

	CreatePortfolio(*servicecatalog.CreatePortfolioInput) (*servicecatalog.CreatePortfolioOutput, error)

	CreatePortfolioShareRequest(*servicecatalog.CreatePortfolioShareInput) (*request.Request, *servicecatalog.CreatePortfolioShareOutput)

	CreatePortfolioShare(*servicecatalog.CreatePortfolioShareInput) (*servicecatalog.CreatePortfolioShareOutput, error)

	CreateProductRequest(*servicecatalog.CreateProductInput) (*request.Request, *servicecatalog.CreateProductOutput)

	CreateProduct(*servicecatalog.CreateProductInput) (*servicecatalog.CreateProductOutput, error)

	CreateProvisioningArtifactRequest(*servicecatalog.CreateProvisioningArtifactInput) (*request.Request, *servicecatalog.CreateProvisioningArtifactOutput)

	CreateProvisioningArtifact(*servicecatalog.CreateProvisioningArtifactInput) (*servicecatalog.CreateProvisioningArtifactOutput, error)

	DeleteConstraintRequest(*servicecatalog.DeleteConstraintInput) (*request.Request, *servicecatalog.DeleteConstraintOutput)

	DeleteConstraint(*servicecatalog.DeleteConstraintInput) (*servicecatalog.DeleteConstraintOutput, error)

	DeletePortfolioRequest(*servicecatalog.DeletePortfolioInput) (*request.Request, *servicecatalog.DeletePortfolioOutput)

	DeletePortfolio(*servicecatalog.DeletePortfolioInput) (*servicecatalog.DeletePortfolioOutput, error)

	DeletePortfolioShareRequest(*servicecatalog.DeletePortfolioShareInput) (*request.Request, *servicecatalog.DeletePortfolioShareOutput)

	DeletePortfolioShare(*servicecatalog.DeletePortfolioShareInput) (*servicecatalog.DeletePortfolioShareOutput, error)

	DeleteProductRequest(*servicecatalog.DeleteProductInput) (*request.Request, *servicecatalog.DeleteProductOutput)

	DeleteProduct(*servicecatalog.DeleteProductInput) (*servicecatalog.DeleteProductOutput, error)

	DeleteProvisioningArtifactRequest(*servicecatalog.DeleteProvisioningArtifactInput) (*request.Request, *servicecatalog.DeleteProvisioningArtifactOutput)

	DeleteProvisioningArtifact(*servicecatalog.DeleteProvisioningArtifactInput) (*servicecatalog.DeleteProvisioningArtifactOutput, error)

	DescribeConstraintRequest(*servicecatalog.DescribeConstraintInput) (*request.Request, *servicecatalog.DescribeConstraintOutput)

	DescribeConstraint(*servicecatalog.DescribeConstraintInput) (*servicecatalog.DescribeConstraintOutput, error)

	DescribePortfolioRequest(*servicecatalog.DescribePortfolioInput) (*request.Request, *servicecatalog.DescribePortfolioOutput)

	DescribePortfolio(*servicecatalog.DescribePortfolioInput) (*servicecatalog.DescribePortfolioOutput, error)

	DescribeProductRequest(*servicecatalog.DescribeProductInput) (*request.Request, *servicecatalog.DescribeProductOutput)

	DescribeProduct(*servicecatalog.DescribeProductInput) (*servicecatalog.DescribeProductOutput, error)

	DescribeProductAsAdminRequest(*servicecatalog.DescribeProductAsAdminInput) (*request.Request, *servicecatalog.DescribeProductAsAdminOutput)

	DescribeProductAsAdmin(*servicecatalog.DescribeProductAsAdminInput) (*servicecatalog.DescribeProductAsAdminOutput, error)

	DescribeProductViewRequest(*servicecatalog.DescribeProductViewInput) (*request.Request, *servicecatalog.DescribeProductViewOutput)

	DescribeProductView(*servicecatalog.DescribeProductViewInput) (*servicecatalog.DescribeProductViewOutput, error)

	DescribeProvisioningArtifactRequest(*servicecatalog.DescribeProvisioningArtifactInput) (*request.Request, *servicecatalog.DescribeProvisioningArtifactOutput)

	DescribeProvisioningArtifact(*servicecatalog.DescribeProvisioningArtifactInput) (*servicecatalog.DescribeProvisioningArtifactOutput, error)

	DescribeProvisioningParametersRequest(*servicecatalog.DescribeProvisioningParametersInput) (*request.Request, *servicecatalog.DescribeProvisioningParametersOutput)

	DescribeProvisioningParameters(*servicecatalog.DescribeProvisioningParametersInput) (*servicecatalog.DescribeProvisioningParametersOutput, error)

	DescribeRecordRequest(*servicecatalog.DescribeRecordInput) (*request.Request, *servicecatalog.DescribeRecordOutput)

	DescribeRecord(*servicecatalog.DescribeRecordInput) (*servicecatalog.DescribeRecordOutput, error)

	DisassociatePrincipalFromPortfolioRequest(*servicecatalog.DisassociatePrincipalFromPortfolioInput) (*request.Request, *servicecatalog.DisassociatePrincipalFromPortfolioOutput)

	DisassociatePrincipalFromPortfolio(*servicecatalog.DisassociatePrincipalFromPortfolioInput) (*servicecatalog.DisassociatePrincipalFromPortfolioOutput, error)

	DisassociateProductFromPortfolioRequest(*servicecatalog.DisassociateProductFromPortfolioInput) (*request.Request, *servicecatalog.DisassociateProductFromPortfolioOutput)

	DisassociateProductFromPortfolio(*servicecatalog.DisassociateProductFromPortfolioInput) (*servicecatalog.DisassociateProductFromPortfolioOutput, error)

	ListAcceptedPortfolioSharesRequest(*servicecatalog.ListAcceptedPortfolioSharesInput) (*request.Request, *servicecatalog.ListAcceptedPortfolioSharesOutput)

	ListAcceptedPortfolioShares(*servicecatalog.ListAcceptedPortfolioSharesInput) (*servicecatalog.ListAcceptedPortfolioSharesOutput, error)

	ListConstraintsForPortfolioRequest(*servicecatalog.ListConstraintsForPortfolioInput) (*request.Request, *servicecatalog.ListConstraintsForPortfolioOutput)

	ListConstraintsForPortfolio(*servicecatalog.ListConstraintsForPortfolioInput) (*servicecatalog.ListConstraintsForPortfolioOutput, error)

	ListLaunchPathsRequest(*servicecatalog.ListLaunchPathsInput) (*request.Request, *servicecatalog.ListLaunchPathsOutput)

	ListLaunchPaths(*servicecatalog.ListLaunchPathsInput) (*servicecatalog.ListLaunchPathsOutput, error)

	ListPortfolioAccessRequest(*servicecatalog.ListPortfolioAccessInput) (*request.Request, *servicecatalog.ListPortfolioAccessOutput)

	ListPortfolioAccess(*servicecatalog.ListPortfolioAccessInput) (*servicecatalog.ListPortfolioAccessOutput, error)

	ListPortfoliosRequest(*servicecatalog.ListPortfoliosInput) (*request.Request, *servicecatalog.ListPortfoliosOutput)

	ListPortfolios(*servicecatalog.ListPortfoliosInput) (*servicecatalog.ListPortfoliosOutput, error)

	ListPortfoliosForProductRequest(*servicecatalog.ListPortfoliosForProductInput) (*request.Request, *servicecatalog.ListPortfoliosForProductOutput)

	ListPortfoliosForProduct(*servicecatalog.ListPortfoliosForProductInput) (*servicecatalog.ListPortfoliosForProductOutput, error)

	ListPrincipalsForPortfolioRequest(*servicecatalog.ListPrincipalsForPortfolioInput) (*request.Request, *servicecatalog.ListPrincipalsForPortfolioOutput)

	ListPrincipalsForPortfolio(*servicecatalog.ListPrincipalsForPortfolioInput) (*servicecatalog.ListPrincipalsForPortfolioOutput, error)

	ListProvisioningArtifactsRequest(*servicecatalog.ListProvisioningArtifactsInput) (*request.Request, *servicecatalog.ListProvisioningArtifactsOutput)

	ListProvisioningArtifacts(*servicecatalog.ListProvisioningArtifactsInput) (*servicecatalog.ListProvisioningArtifactsOutput, error)

	ListRecordHistoryRequest(*servicecatalog.ListRecordHistoryInput) (*request.Request, *servicecatalog.ListRecordHistoryOutput)

	ListRecordHistory(*servicecatalog.ListRecordHistoryInput) (*servicecatalog.ListRecordHistoryOutput, error)

	ProvisionProductRequest(*servicecatalog.ProvisionProductInput) (*request.Request, *servicecatalog.ProvisionProductOutput)

	ProvisionProduct(*servicecatalog.ProvisionProductInput) (*servicecatalog.ProvisionProductOutput, error)

	RejectPortfolioShareRequest(*servicecatalog.RejectPortfolioShareInput) (*request.Request, *servicecatalog.RejectPortfolioShareOutput)

	RejectPortfolioShare(*servicecatalog.RejectPortfolioShareInput) (*servicecatalog.RejectPortfolioShareOutput, error)

	ScanProvisionedProductsRequest(*servicecatalog.ScanProvisionedProductsInput) (*request.Request, *servicecatalog.ScanProvisionedProductsOutput)

	ScanProvisionedProducts(*servicecatalog.ScanProvisionedProductsInput) (*servicecatalog.ScanProvisionedProductsOutput, error)

	SearchProductsRequest(*servicecatalog.SearchProductsInput) (*request.Request, *servicecatalog.SearchProductsOutput)

	SearchProducts(*servicecatalog.SearchProductsInput) (*servicecatalog.SearchProductsOutput, error)

	SearchProductsAsAdminRequest(*servicecatalog.SearchProductsAsAdminInput) (*request.Request, *servicecatalog.SearchProductsAsAdminOutput)

	SearchProductsAsAdmin(*servicecatalog.SearchProductsAsAdminInput) (*servicecatalog.SearchProductsAsAdminOutput, error)

	TerminateProvisionedProductRequest(*servicecatalog.TerminateProvisionedProductInput) (*request.Request, *servicecatalog.TerminateProvisionedProductOutput)

	TerminateProvisionedProduct(*servicecatalog.TerminateProvisionedProductInput) (*servicecatalog.TerminateProvisionedProductOutput, error)

	UpdateConstraintRequest(*servicecatalog.UpdateConstraintInput) (*request.Request, *servicecatalog.UpdateConstraintOutput)

	UpdateConstraint(*servicecatalog.UpdateConstraintInput) (*servicecatalog.UpdateConstraintOutput, error)

	UpdatePortfolioRequest(*servicecatalog.UpdatePortfolioInput) (*request.Request, *servicecatalog.UpdatePortfolioOutput)

	UpdatePortfolio(*servicecatalog.UpdatePortfolioInput) (*servicecatalog.UpdatePortfolioOutput, error)

	UpdateProductRequest(*servicecatalog.UpdateProductInput) (*request.Request, *servicecatalog.UpdateProductOutput)

	UpdateProduct(*servicecatalog.UpdateProductInput) (*servicecatalog.UpdateProductOutput, error)

	UpdateProvisionedProductRequest(*servicecatalog.UpdateProvisionedProductInput) (*request.Request, *servicecatalog.UpdateProvisionedProductOutput)

	UpdateProvisionedProduct(*servicecatalog.UpdateProvisionedProductInput) (*servicecatalog.UpdateProvisionedProductOutput, error)

	UpdateProvisioningArtifactRequest(*servicecatalog.UpdateProvisioningArtifactInput) (*request.Request, *servicecatalog.UpdateProvisioningArtifactOutput)

	UpdateProvisioningArtifact(*servicecatalog.UpdateProvisioningArtifactInput) (*servicecatalog.UpdateProvisioningArtifactOutput, error)
}

var _ ServiceCatalogAPI = (*servicecatalog.ServiceCatalog)(nil)
