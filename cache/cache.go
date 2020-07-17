package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"reflect"
	"time"

	"pgxs.io/chassis/config"
	"pgxs.io/chassis/logx"

	"pgxs.io/chassis"
)

//Store cache store interface
type Store interface {
	Get(key string) (val interface{}, ok bool)
	Set(key string, val interface{}) (ok bool)
	Delete(key string)
	Contains(key string) bool
}

type RedisCacheStore struct {
	prefix      string
	valType     reflect.Type
	expiredTime time.Duration
	log         *logx.Entry
	redisCfg    *config.RedisConfig
}

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
func (rc *RedisCacheStore) Set(key string, val interface{}) (ok bool) {
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
func (rc *RedisCacheStore) Delete(key string) {
	if rc.redisCfg.Mode == "simple" || rc.redisCfg.Mode == "sentinel" {
		chassis.RedisClient().Del(context.Background(), fmt.Sprintf("%s:%s", rc.prefix, key))
	} else if rc.redisCfg.Mode == "cluster" {
		chassis.RedisClusterClient().Del(context.Background(), fmt.Sprintf("%s:%s", rc.prefix, key))
	}
}

func (rc *RedisCacheStore) Contains(key string) bool {
	var result int64
	if rc.redisCfg.Mode == "simple" || rc.redisCfg.Mode == "sentinel" {
		result = chassis.RedisClient().Exists(context.Background(), fmt.Sprintf("%s:%s", rc.prefix, key)).Val()
	} else if rc.redisCfg.Mode == "cluster" {
		result = chassis.RedisClusterClient().Exists(context.Background(), fmt.Sprintf("%s:%s", rc.prefix, key)).Val()
	}
	return result == 1
}
func NewRedisCache(name string, valType interface{}, expired time.Duration) Store {
	gob.Register(valType)
	t := reflect.TypeOf(valType)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	//todo check if key existed return error
	store := &RedisCacheStore{prefix: name, valType: t, expiredTime: expired}
	store.log = logx.New().Category("cache").Component(name)
	cfg := config.Redis()
	store.redisCfg = &cfg
	return store
}
