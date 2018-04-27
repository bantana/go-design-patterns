package singleton

import (
	"sync"
)

// Singleton type
type singleton struct {
	m     sync.Mutex
	count int64
}

var (
	instance *singleton
	once     sync.Once
)

// Get provide Export access
func Get() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}

// AddOne provide Singleton modify number
func (i *singleton) Add() int64 {
	i.m.Lock()
	defer i.m.Unlock()
	i.count++
	return i.count
}

// GetCount provide Get Singleton properties.
func (i *singleton) Count() int64 {
	return i.count
}
