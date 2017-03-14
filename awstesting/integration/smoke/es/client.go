// +build integration

//Package es provides gucumber integration tests support.
package es

import (
	"github.com/tily/sdk-go/awstesting/integration/smoke"
	"github.com/tily/sdk-go/service/elasticsearchservice"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@es", func() {
		gucumber.World["client"] = elasticsearchservice.New(smoke.Session)
	})
}
