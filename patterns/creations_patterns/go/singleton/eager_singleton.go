package singleton

type EagerSingleton struct{}

var instance = &EagerSingleton{}

func GetInstance() *EagerSingleton {
	return instance
}
