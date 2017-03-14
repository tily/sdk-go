// +build integration

//Package apigateway provides gucumber integration tests support.
package apigateway

import (
	"github.com/tily/sdk-go/awstesting/integration/smoke"
	"github.com/tily/sdk-go/service/apigateway"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@apigateway", func() {
		gucumber.World["client"] = apigateway.New(smoke.Session)
	})
}
