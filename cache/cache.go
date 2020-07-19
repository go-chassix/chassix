package cache

import "errors"

//Store cache store interface
type Store interface {
	Get(key string) (val interface{}, ok bool)
	Set(key string, val interface{}) (ok bool)
	Delete(key string)
	Contains(key string) bool
}

var (
	StoreNameIsEmptyErr = errors.New("store name cannot be empty")
)
