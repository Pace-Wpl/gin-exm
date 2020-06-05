package def

import "sync"

const (
	//config dir
	CONFIG_DIR = "../conf/config.yaml"
)

var (
	ProductConfig sync.Map
	Conf          = &Config{}
)

type Config struct {
	Httpport       string
	AppName        string
	SessionExpired int
	CookieKey      string
	Domain         string

	Log   LogConf
	Redis RedisConf
	Etcd  EtcdConf
	Mysql MysqlConf
}

type ProductConf struct {
	ProductID int
	StartTime int64
	EndTime   int64
	Status    int
	Total     int
}

type LogConf struct {
	GinLogPath  string
	GinLogLevel string
	SysLogPath  string
	SysLogLevel string
}

type RedisConf struct {
	Addr        string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

type EtcdConf struct {
	Addr       string
	Timeout    int
	PrefixKey  string
	ProductKey string
}

type MysqlConf struct {
	Addr     string
	User     string
	Pwd      string
	Database string
	Config   string
}
