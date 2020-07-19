package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"reflect"
	"time"

	"c6x.io/chassis/config"
	"c6x.io/chassis/logx"

	"c6x.io/chassis"
)

//RedisCacheStore based redis implement cache store
type RedisCacheStore struct {
	prefix      string
	valType     reflect.Type
	expiredTime time.Duration
	log         *logx.Entry
	redisCfg    *config.RedisConfig
}

//Get get value from redis cache store
func (rc *RedisCacheStore) Get(key string) (val interface{}, ok bool) {
	ctx := context.Background()
	valPtr := reflect.New(rc.valType)

	redisCfg := config.Redis()
	var data []byte
	var err error
	if redisCfg.Mode == "simple" || redisCfg.Mode == "sentinel" {
		data, err = chassis.RedisClient().Get(ctx, fmt.Sprintf("%s:%s", rc.prefix, key)).Bytes()
	}
	if redisCfg.Mode == "cluster" {
		data, err = chassis.RedisClusterClient().Get(ctx, fmt.Sprintf("%s:%s", rc.prefix, key)).Bytes()
	}

	if err != nil {
		rc.log.Errorf("get key [%s] from cache error: %s", key, err)
		return
	}

	decoder := gob.NewDecoder(bytes.NewReader(data))
	if err = decoder.Decode(valPtr.Interface()); err != nil {
		rc.log.Debugf("gob decode failed, error: %s", err)
		return
	}
	val = valPtr.Elem().Interface()
	ok = true
	return
}

//Set add or update key value
func (rc *RedisCacheStore) Set(key string, val interface{}) (ok bool) {
	t := reflect.TypeOf(val)
	if t != rc.valType {
		ok = false
		rc.log.Infof("value type wrong. should be %s actual: %+v", rc.valType, t)
		return
	}

	ctx := context.Background()
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	enc.Encode(val)
	var err error
	if rc.redisCfg.Mode == "simple" || rc.redisCfg.Mode == "sentinel" {
		err = chassis.RedisClient().Set(ctx, fmt.Sprintf("%s:%s", rc.prefix, key), buffer.Bytes(), rc.expiredTime).Err()
	} else if rc.redisCfg.Mode == "cluster" {
		err = chassis.RedisClusterClient().Set(ctx, fmt.Sprintf("%s:%s", rc.prefix, key), buffer.Bytes(), rc.expiredTime).Err()
	}
	if err != nil {
		rc.log.Errorf("set key [%s] failed, error: %s", key, err)
		return
	}
	ok = true
	return
}

//Delete delete key value in redis cache store
func (rc *RedisCacheStore) Delete(key string) {
	if rc.redisCfg.Mode == "simple" || rc.redisCfg.Mode == "sentinel" {
		chassis.RedisClient().Del(context.Background(), fmt.Sprintf("%s:%s", rc.prefix, key))
	} else if rc.redisCfg.Mode == "cluster" {
		chassis.RedisClusterClient().Del(context.Background(), fmt.Sprintf("%s:%s", rc.prefix, key))
	}
}

//Contains check key existed in redis cache store
func (rc *RedisCacheStore) Contains(key string) bool {
	var result int64
	key = fmt.Sprintf("%s:%s", rc.prefix, key)
	if rc.redisCfg.Mode == "simple" || rc.redisCfg.Mode == "sentinel" {
		result = chassis.RedisClient().Exists(context.Background(), key).Val()
	} else if rc.redisCfg.Mode == "cluster" {
		result = chassis.RedisClusterClient().Exists(context.Background(), key).Val()
	}
	return result == 1
}

//NewRedisCacheStore new redis cache store
func NewRedisCacheStore(name string, valType interface{}, expired time.Duration) (store Store, err error) {
	if name == "" {
		err = StoreNameIsEmptyErr
		return
	}
	t := reflect.TypeOf(valType)

	redisStore := &RedisCacheStore{prefix: name, valType: t, expiredTime: expired}
	redisStore.log = logx.New().Category("cache").Component(name)
	redisCfg := config.Redis()
	redisStore.redisCfg = &redisCfg
	store = redisStore

	return
}
