package service

type CacheInteface interface {
	Set(key string, value interface{})
	Get(key string) (value string)
}
