// +build integration

//Package codepipeline provides gucumber integration tests support.
package codepipeline

import (
	"github.com/tily/sdk-go/awstesting/integration/smoke"
	"github.com/tily/sdk-go/service/codepipeline"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@codepipeline", func() {
		gucumber.World["client"] = codepipeline.New(smoke.Session)
	})
}
