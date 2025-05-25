package context

type Base interface {
	Next()
}

type Abort interface {
	Abort()
	IsAborted() bool
}

type KeyValue interface {
	Get(key string) (value any, exists bool)
	Set(key string, value any)
}
