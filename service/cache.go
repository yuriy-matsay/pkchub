package service

type CacheInteface interface {
	Set(key string, value interface{}) (err error)
	Get(key string) (value string, err error)
	DeleteCache()
}
