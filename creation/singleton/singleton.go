package singleton

// Singleton type
type Singleton struct{}

var instance *Singleton

// GetInstance ...
func GetInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}
