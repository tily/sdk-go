// +build integration

//Package cloudformation provides gucumber integration tests support.
package cloudformation

import (
	"github.com/tily/sdk-go/awstesting/integration/smoke"
	"github.com/tily/sdk-go/service/cloudformation"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@cloudformation", func() {
		gucumber.World["client"] = cloudformation.New(smoke.Session)
	})
}
