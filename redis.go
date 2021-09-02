package chassix

import (
	"reflect"
	"sync"
	"time"

	"github.com/go-chassix/chassix/v2/config"
	"github.com/go-redis/redis/v8"
)

var (
	redisClient         *redis.Client
	redisClientInitOnce sync.Once

	redisClusterClient         *redis.ClusterClient
	redisClusterClientInitOnce sync.Once
)

//RedisClient get redis client
//return for simple or sentinel(failover)mode
func RedisClient() *redis.Client {
	redisClientInitOnce.Do(func() {
		if config.NotNil() {
			redisCfg := config.Redis()

			switch redisCfg.Mode {
			case "simple":
				var opts redis.Options
				readRedisOptions(&opts)
				redisClient = redis.NewClient(&opts)

			case "sentinel":
				var opts redis.FailoverOptions
				readRedisOptions(&opts)
				redisClient = redis.NewFailoverClient(&opts)
			}
		}
	})

	return redisClient
}

//RedisClusterClient redis cluster client
func RedisClusterClient() *redis.ClusterClient {
	redisClusterClientInitOnce.Do(func() {
		var opts redis.ClusterOptions
		readRedisOptions(&opts)
		redisClusterClient = redis.NewClusterClient(&opts)
	})
	return redisClusterClient
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
		case "cluster":
			elem.FieldByName("Addrs").Set(reflect.ValueOf(redisCfg.Cluster.Addrs))
			elem.FieldByName("MaxRedirects").Set(reflect.ValueOf(redisCfg.Cluster.MaxRedirects))
			elem.FieldByName("ReadOnly").Set(reflect.ValueOf(redisCfg.Cluster.ReadOnly))
			elem.FieldByName("RouteByLatency").Set(reflect.ValueOf(redisCfg.Cluster.RouteByLatency))
			elem.FieldByName("RouteRandomly").Set(reflect.ValueOf(redisCfg.Cluster.RouteRandomly))
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

	}
}
