package chassis

import (
	"github.com/go-redis/redis/v8"
	"pgxs.io/chassis/config"
	"reflect"
	"time"
)

//RedisClient get redis client
//return for simple or sentinel(failover)mode
func RedisClient() *redis.Client {
	if config.NotNil() {
		redisCfg := config.Redis()

		switch redisCfg.Mode {
		case "simple":
			var opts redis.Options
			readRedisOptions(&opts)
			return redis.NewClient(&opts)

		case "sentinel":
			var opts redis.FailoverOptions
			readRedisOptions(&opts)
			return redis.NewFailoverClient(&opts)
		}

	}
	return nil
}

func readRedisOptions(opts interface{}) {
	if config.NotNil() {
		redisCfg := config.Redis()

		elem := reflect.ValueOf(opts).Elem()

		switch redisCfg.Mode {
		case "simple":
			elem.FieldByName("Addr").Set(reflect.ValueOf(redisCfg.Addr))
		case "sentinel":
			elem.FieldByName("SentinelUsername").Set(reflect.ValueOf(redisCfg.Sentinel.Username))
			elem.FieldByName("SentinelPassword").Set(reflect.ValueOf(redisCfg.Sentinel.Password))
			elem.FieldByName("MasterName").Set(reflect.ValueOf(redisCfg.Sentinel.Master))
			elem.FieldByName("SentinelAddrs").Set(reflect.ValueOf(redisCfg.Sentinel.Addrs))
		}

		elem.FieldByName("Username").Set(reflect.ValueOf(redisCfg.Username))
		elem.FieldByName("Password").Set(reflect.ValueOf(redisCfg.Password))
		elem.FieldByName("DB").Set(reflect.ValueOf(redisCfg.DB))
		elem.FieldByName("MaxRetries").Set(reflect.ValueOf(redisCfg.MaxRetries))

		//pool settings
		elem.FieldByName("PoolSize").Set(reflect.ValueOf(redisCfg.PoolSize))

		if redisCfg.MaxConnAge != "" {
			if t, err := time.ParseDuration(redisCfg.MaxConnAge); err == nil {
				elem.FieldByName("MaxConnAge").Set(reflect.ValueOf(t))
			}
		}

		if redisCfg.PoolTimeout != "" {
			if t, err := time.ParseDuration(redisCfg.PoolTimeout); err == nil {
				elem.FieldByName("PoolTimeout").Set(reflect.ValueOf(t))
			}
		}

		if redisCfg.IdleTimeout != "" {
			if t, err := time.ParseDuration(redisCfg.IdleTimeout); err == nil {
				elem.FieldByName("IdleTimeout").Set(reflect.ValueOf(t))
			}
		}

		if redisCfg.IdleCheckFrequency != "" {
			if t, err := time.ParseDuration(redisCfg.IdleCheckFrequency); err == nil {
				elem.FieldByName("IdleCheckFrequency").Set(reflect.ValueOf(t))
			}
		}

		//opts.Username = redisCfg.Username
		//
		//opts.PoolSize = redisCfg.PoolSize
		//opts.MinIdleConns = redisCfg.MinIdleConns

		//
		//if redisCfg.PoolTimeout != "" {
		//	if t, err := time.ParseDuration(redisCfg.PoolTimeout); err == nil {
		//		opts.PoolTimeout = t
		//	}
		//}
		//

		//
		//if redisCfg.IdleCheckFrequency != "" {
		//	if t, err := time.ParseDuration(redisCfg.IdleCheckFrequency); err == nil {
		//		opts.IdleCheckFrequency = t
		//	}
		//}

		//return opts

	}
}
