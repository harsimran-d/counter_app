package main

import (
	"strconv"
	"sync"
)

type GloblaCount struct {
	mu    sync.Mutex
	value int
}

func (count *GloblaCount) byteValue() []byte {
	return []byte(strconv.Itoa(count.value))
}
