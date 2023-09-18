package main

import (
		"context"
		"fmt"

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
	lis weaver.Listener `weaver:"emoji"`
}

func run(ctx context.Context, a *app) error {
	a.Logger(ctx).Info("listener active", "addr", a.lis)
	emojis, err := a.searcher.Get().Search(ctx, "city")
	if err != nil {
		return err
	}

	fmt.Println(emojis)
	return nil
}
