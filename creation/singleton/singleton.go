package singleton

// Singleton type
type Singleton struct{}

var instance *Singleton

// New ...
func New() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}
