package storage

type Engine interface {
	Get(key string) string
	Set(key, value string)
	Del(key string)
}
