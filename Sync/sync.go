package main

import "sync"

// type Counter struct {
// 	value int
// }

// func (c *Counter) Inc() {
// 	c.value++
// }

type Counter struct {
	sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

