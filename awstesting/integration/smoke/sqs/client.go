// +build integration

//Package sqs provides gucumber integration tests support.
package sqs

import (
	"github.com/tily/sdk-go/awstesting/integration/smoke"
	"github.com/tily/sdk-go/service/sqs"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@sqs", func() {
		gucumber.World["client"] = sqs.New(smoke.Session)
	})
}
