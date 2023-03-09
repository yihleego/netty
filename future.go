package netty

type Future[T any] interface {
	// Sync waits for this listener until it is done
	Sync() T
	// Async nonblock waits for this listener
	Async(func(T))
	Cancel() error
}
