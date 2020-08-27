package cache

import (
	"errors"
	"fmt"
	"reflect"

	lru "github.com/hashicorp/golang-lru"

	"c6x.io/chassix.v2/logx"
)

//MemoryCache implements cache store based or lru cache
type MemoryCacheStore struct {
	cache     *lru.Cache
	log       *logx.Entry
	valType   reflect.Type
	keyPrefix string
}

//NewMemoryCacheStore new memory cache store stored spec size
func NewMemoryCacheStore(name string, valType interface{}, size int) (cache *MemoryCacheStore, err error) {

	if name == "" {
		err = StoreNameIsEmptyErr
		return
	}
	t := reflect.TypeOf(valType)
	log := logx.New().Category("mem-cache").Component(name)

	mCache, lErr := lru.New(size)
	if lErr != nil {
		err = errors.New("new lru cache failed")
		log.Errorf("%s,error:%s", err, lErr)
		return
	}
	cache = &MemoryCacheStore{log: log, cache: mCache, valType: t}
	return
}

//Set set a key/value to memory cache store
func (mcs *MemoryCacheStore) Set(key string, val interface{}) (ok bool) {
	t := reflect.TypeOf(val)
	if t != mcs.valType {
		ok = false
		mcs.log.Infof("value type wrong. should be %s actual: %+v", mcs.valType, t)
		return
	}
	mcs.cache.Add(mcs.realKey(key), val)
	return true
}

//Get get a key/value from memory cache store
func (mcs *MemoryCacheStore) Get(key string) (val interface{}, ok bool) {
	return mcs.cache.Get(mcs.realKey(key))
}

//Contains check existed in a memory cache store
func (mcs *MemoryCacheStore) Contains(key string) bool {
	return mcs.cache.Contains(mcs.realKey(key))
}

//Delete delete key in a memory cache store
func (mcs *MemoryCacheStore) Delete(key string) {
	mcs.cache.Remove(mcs.realKey(key))
}

func (mcs *MemoryCacheStore) realKey(key string) string {
	return fmt.Sprintf("%s:%s", mcs.keyPrefix, key)
}
