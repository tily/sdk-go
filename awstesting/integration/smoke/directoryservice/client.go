// +build integration

//Package directoryservice provides gucumber integration tests support.
package directoryservice

import (
	"github.com/tily/sdk-go/awstesting/integration/smoke"
	"github.com/tily/sdk-go/service/directoryservice"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@directoryservice", func() {
		gucumber.World["client"] = directoryservice.New(smoke.Session)
	})
}
