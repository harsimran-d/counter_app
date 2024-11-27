package main

import (
	"sync"
)

type GloblaCount struct {
	mu    sync.Mutex
	value int
}
