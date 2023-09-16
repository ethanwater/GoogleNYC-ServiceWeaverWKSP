package main

import (
	"context"
	"testing"
	"fmt"

	"github.com/ServiceWeaver/weaver/weavertest"
)

func SearchTest(t *testing.T) {
	runner := weavertest.Local
	runner.Test(t, func(t *testing.T, searcher Searcher) {
		ctx := context.Background()
		got, err := searcher.Search(ctx, "black cat")
		if err != nil {
			t.Fatal(err)
		}
	})
}
