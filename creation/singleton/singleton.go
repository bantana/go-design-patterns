package singleton

// Singleton type
type Singleton struct{}

var instance *Singleton

// Get ...
func Get() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}
