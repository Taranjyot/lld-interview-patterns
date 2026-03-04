package singleton

import "sync"

type DoubleCheckedSingleton struct{}

var (
	dcmu                   sync.Mutex
	doubleCheckedSingleton *DoubleCheckedSingleton
)

func GetDoubleCheckedSingletonInstance() *DoubleCheckedSingleton {

	if doubleCheckedSingleton == nil {
		dcmu.Lock()
		defer dcmu.Unlock()

		if doubleCheckedSingleton == nil {
			doubleCheckedSingleton = &DoubleCheckedSingleton{}
		}

	}

	return doubleCheckedSingleton
}
