package main

import (
		"context"
		"fmt"
		_ "embed"
		"net/http"
		"encoding/json"

		"github.com/ServiceWeaver/weaver"
)

func main(){
	if err := weaver.Run(context.Background(), run); err != nil {
		panic(err)
	}
}

type app struct {
	weaver.Implements[weaver.Main]
	searcher weaver.Ref[Searcher]
	listener weaver.Listener `weaver:"emoji"`
}

func run(ctx context.Context, a *app) error {
	a.Logger(ctx).Info("listener active", "addr", a.listener)


	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		emojis, err := a.searcher.Get().Search(ctx, query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
		}
		bytes, err := json.Marshal(emojis) 
 		if err != nil { 
 			http.Error(w, err.Error(), http.StatusInternalServerError) 
 			return 
 		} 
 		fmt.Fprintln(w, string(bytes))
	})

	return http.Serve(a.listener, nil) 
}
