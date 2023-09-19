package main

import (
	"context"
	"sort"
	"strings"

	"github.com/ServiceWeaver/weaver" 
 	"golang.org/x/exp/slices"
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
		return emoji, nil
	}

	text := strings.Fields(strings.ToLower(q))
	results := []string{}
	for emoji, labels := range emojis {
		if match(labels, text) {
			results = append(results, emoji)
		} 	
	}
	sort.Strings(results)


	//Put
	if err := s.cache.Get().Put(ctx, q, results); err != nil {
		logger.Error("cache.Put", "query", q, "err", err)
	} 

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
