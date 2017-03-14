// +build integration

//Package opsworks provides gucumber integration tests support.
package opsworks

import (
	"github.com/tily/sdk-go/awstesting/integration/smoke"
	"github.com/tily/sdk-go/service/opsworks"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@opsworks", func() {
		gucumber.World["client"] = opsworks.New(smoke.Session)
	})
}
