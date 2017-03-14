// +build integration

//Package codedeploy provides gucumber integration tests support.
package codedeploy

import (
	"github.com/tily/sdk-go/awstesting/integration/smoke"
	"github.com/tily/sdk-go/service/codedeploy"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@codedeploy", func() {
		gucumber.World["client"] = codedeploy.New(smoke.Session)
	})
}
