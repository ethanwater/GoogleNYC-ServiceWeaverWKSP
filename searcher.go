package main

import (
	"context"
	"strings"

	"github.com/ServiceWeaver/weaver" 
	"github.com/ServiceWeaver/weaver/metrics" 
 	"golang.org/x/exp/slices"
)

var (
	search_cache_hits = metrics.NewCounter (
		"cache_hit",
		"total cache hits",
	)

	search_cache_misses = metrics.NewCounter (
		"cache_miss",
		"total cache misses",
	)

)

type Searcher interface {
	Search(context.Context, string) ([]string, error)
}

type searcher struct {
	weaver.Implements[Searcher]
	cache weaver.Ref[Cache]
}

func (s *searcher) Search(ctx context.Context, q string) ([]string, error) {
	logger := s.Logger(ctx)
	logger.Debug("Search", "query", q)

	//Get
	if emoji, err := s.cache.Get().Get(ctx, q); err != nil {
		logger.Error("cache.Get", "query", q, "err", err)
	} else if emoji != nil {
		search_cache_hits.Inc()
		return emoji, nil
	} else {
		search_cache_misses.Inc()
	}

	input := strings.Fields(strings.ToLower(q))
	results := []string{}
	for emoji, labels := range emojis {
		if match(labels, input) {
			results = append(results, emoji)
		} 	
	}

	//Put
	if err := s.cache.Get().Put(ctx, q, results); err != nil {
		logger.Error("cache.Put", "query", q, "err", err)
	} 

	return results, nil
}

func match(labels, input []string) bool {
	for _, t := range input {
		if slices.Contains(labels, t) {
			return true			
		} 	
	}
	return false
}
