package main

import (
	"context"
	"sort"
	"strings"

	"github.com/ServiceWeaver/weaver" 
 	"golang.org/x/exp/slices"
)

type Searcher interface {
	Search(ctx context.Context, query string) ([]string, error)
}

type searcher struct {
	weaver.Implements[Searcher]
}

func (s *searcher) Search(_ context.Context, q string) ([]string, error) {
	results := []string{}
	text := strings.Fields(strings.ToLower(q))
	for emoji, labels := range emojis {
		if match(labels, text) {
			results = append(results, emoji)
		}
	}
	sort.Strings(results)
	return results, nil
}

func match(labels, text []string) bool {
	for _, t := range text {
		if slices.Contains(labels, t) {
			return true			
		} 	
	}
	return false
}
