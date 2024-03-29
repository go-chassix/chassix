package config

//RedisConfig redis配置
type RedisConfig struct {
	RedisSimpleConfig `yaml:",inline"`
	RedisCommonConfig `yaml:",inline"`
	Mode              string              `yaml:"mode"` // 3种模式 1 simple (单机/主从) 2 sentinel 哨兵模式 3 cluster 集群模式
	Sentinel          RedisSentinelConfig `yaml:"sentinel"`
	Cluster           RedisClusterConfig  `yaml:"cluster"`
}

//RedisSentinelConfig redis sentinel server config
type RedisSentinelConfig struct {
	Master   string   `yaml:"master"`     // The master name.
	Addrs    []string `yaml:"addrs,flow"` // A seed list of host:port addresses of sentinel nodes.
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
}

//RedisSimpleConfig redis simple config for stand-alone or master/slave mode
type RedisSimpleConfig struct {
	// host:port address.
	Addr string
}

//RedisCommonConfig redis common config
//for simple,sentinel,cluster
type RedisCommonConfig struct {
	// Use the specified Username to authenticate the current connection
	// with one of the connections defined in the ACL list when connecting
	// to a Redis 6.0 instance, or greater, that is using the Redis ACL system.
	Username string `yaml:"username"`
	// Optional password. Must match the password specified in the
	// requirepass server configuration option (if connecting to a Redis 5.0 instance, or lower),
	// or the User Password when connecting to a Redis 6.0 instance, or greater,
	// that is using the Redis ACL system.
	Password string `yaml:"password"`

	// Database to be selected after connecting to the server.
	DB int `yaml:"db"`

	// Maximum number of retries before giving up.
	// Default is to not retry failed commands.
	MaxRetries int `yaml:"max-retries"`
	// Minimum backoff between each retry.
	// Default is 8 milliseconds; -1 disables backoff.
	MinRetryBackoff string `yaml:"min-retry-backoff"`
	// Maximum backoff between each retry.
	// Default is 512 milliseconds; -1 disables backoff.
	MaxRetryBackoff string `yaml:"max-retry-backoff"`

	// Dial timeout for establishing new connections.
	// Default is 5 seconds.
	DialTimeout string `yaml:"dial-timeout"`
	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Use value -1 for no timeout and 0 for default.
	// Default is 3 seconds.
	ReadTimeout string `yaml:"read-timeout"`
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.
	// Default is ReadTimeout.
	WriteTimeout string `yaml:"write-timeout"`

	// PoolSize applies per cluster node and not for the whole cluster.

	// Maximum number of socket connections.
	// Default is 10 connections per every CPU as reported by runtime.NumCPU.
	PoolSize int `yaml:"pool-size"`
	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	MinIdleConns int `yaml:"min-idle-conns"`
	// Connection age at which client retires (closes) the connection.
	// Default is to not close aged connections.
	MaxConnAge string `yaml:"max-conn-age"`
	// Amount of time client waits for connection if all connections
	// are busy before returning an error.
	// Default is ReadTimeout + 1 second.
	PoolTimeout string `yaml:"pool-timeout"`
	// Amount of time after which client closes idle connections.
	// Should be less than server's timeout.
	// Default is 5 minutes. -1 disables idle timeout check.
	IdleTimeout string `yaml:"idle-timeout"`
	// Frequency of idle checks made by idle connections reaper.
	// Default is 1 minute. -1 disables idle connections reaper,
	// but idle connections are still discarded by the client
	// if IdleTimeout is set.
	IdleCheckFrequency string `yaml:"idle-check-frequency"`
}

//RedisClusterConfig 集群模式配置
type RedisClusterConfig struct {
	Addrs []string

	// The maximum number of retries before giving up. Command is retried
	// on network errors and MOVED/ASK redirects.
	// Default is 8 retries.
	MaxRedirects int

	// Enables read-only commands on slave nodes.
	ReadOnly bool
	// Allows routing read-only commands to the closest master or slave node.
	// It automatically enables ReadOnly.
	RouteByLatency bool
	// Allows routing read-only commands to the random master or slave node.
	// It automatically enables ReadOnly.
	RouteRandomly bool
}
