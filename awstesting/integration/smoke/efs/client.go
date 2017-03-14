// +build integration

//Package efs provides gucumber integration tests support.
package efs

import (
	"github.com/tily/sdk-go/aws"
	"github.com/tily/sdk-go/awstesting/integration/smoke"
	"github.com/tily/sdk-go/service/efs"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@efs", func() {
		// FIXME remove custom region
		gucumber.World["client"] = efs.New(smoke.Session,
			aws.NewConfig().WithRegion("us-west-2"))
	})
}
