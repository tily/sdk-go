// +build integration

//Package ec2 provides gucumber integration tests support.
package ec2

import (
	"github.com/tily/sdk-go/awstesting/integration/smoke"
	"github.com/tily/sdk-go/service/ec2"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@ec2", func() {
		gucumber.World["client"] = ec2.New(smoke.Session)
	})
}
