package main 
  
 import ( 
 	"context" 
 	"sync" 
  
 	"github.com/ServiceWeaver/weaver" 
 ) 

type Cache interface {
	Get(context.Context, string) ([]string, error)
  Put(context.Context, string, []string) error
}

type cache struct {
		weaver.Implements[Cache]
		weaver.WithRouter[route]
    mu sync.Mutex
    emoji map[string][]string
}

type route struct {}
func (route) Get(_ context.Context, key string) string { return key }
func (route) Put(_ context.Context, key string, _ []string) string { return key }

func (c *cache) Init(context.Context) error { 
 	c.emoji = map[string][]string{} 
 	return nil 
}
func (c *cache) Get(ctx context.Context, query string) ([]string, error) {
    c.mu.Lock()
    defer c.mu.Unlock()
		c.Logger(ctx).Debug("Get", "query", query)
    return c.emoji[query], nil
}
func (c *cache) Put(ctx context.Context, query string, emoji[]string) error {
    c.mu.Lock()
    defer c.mu.Unlock()
		c.Logger(ctx).Debug("Put", "query", query)
    c.emoji[query] = emoji
    return nil
}


