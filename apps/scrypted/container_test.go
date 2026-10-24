package main

import (
	"context"
	"testing"

	"github.com/gabrielcosi/containers/testhelpers"
)

func Test(t *testing.T) {
	ctx := context.Background()
	image := testhelpers.GetTestImage("ghcr.io/gabrielcosi/containers/scrypted:rolling")

	t.Run("HTTP endpoint test", func(t *testing.T) {
		testhelpers.TestHTTPEndpoint(t, ctx, image, testhelpers.HTTPTestConfig{
			Port: "11080",
		}, nil)
	})
}
