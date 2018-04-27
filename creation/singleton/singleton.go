package singleton

// Singleton type
type singleton struct{}

var instance *singleton

// Get ...
func Get() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
