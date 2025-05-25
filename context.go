package context

import (
	"math"
	"sync"
)

const abortIndex int8 = math.MaxInt8 >> 1

type (
	HandlerFunc   func(*Context)
	HandlersChain []HandlerFunc
)

type Context struct {
	index int8
	mu    sync.RWMutex

	handlers HandlersChain

	Keys map[string]any
}

func (c *Context) Abort() {
	c.index = abortIndex
}

func (c *Context) IsAborted() bool {
	return c.index >= abortIndex
}

func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.Keys == nil {
		c.Keys = make(map[string]any)
	}

	c.Keys[key] = value
}

// Get returns the value for the given key, ie: (value, true).
// If the value does not exist it returns (nil, false)
func (c *Context) Get(key string) (value any, exists bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, exists = c.Keys[key]
	return
}
