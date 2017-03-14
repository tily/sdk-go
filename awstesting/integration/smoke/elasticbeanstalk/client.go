// +build integration

//Package elasticbeanstalk provides gucumber integration tests support.
package elasticbeanstalk

import (
	"github.com/tily/sdk-go/awstesting/integration/smoke"
	"github.com/tily/sdk-go/service/elasticbeanstalk"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@elasticbeanstalk", func() {
		gucumber.World["client"] = elasticbeanstalk.New(smoke.Session)
	})
}
