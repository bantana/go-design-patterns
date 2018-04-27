package singleton

import (
	"sync"
)

// Singleton type
type Singleton struct {
	m     sync.Mutex
	count int
}

var (
	singleton *Singleton
	once      sync.Once
)

// New provide Export access
func New() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}

// AddOne provide Singleton modify number
func (i *Singleton) AddOne() int {
	i.m.Lock()
	defer i.m.Unlock()
	i.count++
	return i.count
}

// GetCount provide Get Singleton properties.
func (i *Singleton) GetCount() int {
	return i.count
}
