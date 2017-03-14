// +build integration

//Package swf provides gucumber integration tests support.
package swf

import (
	"github.com/tily/sdk-go/awstesting/integration/smoke"
	"github.com/tily/sdk-go/service/swf"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@swf", func() {
		gucumber.World["client"] = swf.New(smoke.Session)
	})
}
