package singleton

import "sync"

type ThreadSafeSingleton struct{}

var (
	threadSafeSingleton *ThreadSafeSingleton
	mu                  sync.Mutex
)

func GetThreadSafeInstance() *ThreadSafeSingleton {
	mu.Lock()
	defer mu.Unlock()

	if threadSafeSingleton == nil {
		return &ThreadSafeSingleton{}
	}

	return threadSafeSingleton
}
