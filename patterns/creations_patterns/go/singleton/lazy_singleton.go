package singleton

import "sync"

type LazySingleton struct{}

var (
	lazySingletonInstace *LazySingleton
	once                 sync.Once
)

func GetLazySingletonInstance() *LazySingleton {
	once.Do(func() {
		lazySingletonInstace = &LazySingleton{}
	})

	return lazySingletonInstace
}
