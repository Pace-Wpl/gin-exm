package def

const (
	//config dir
	CONFIG_DIR = "../conf/config.yaml"
)

var (
	Conf = &Config{}
)

type Config struct {
	Httpport string
	AppName  string

	Log   LogConf
	Redis RedisConf
	Etcd  EtcdConf
	Mysql MysqlConf
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
	Addr          string
	Timeout       int
	SecKeyPrefix  string
	SecProductKey string
}

type MysqlConf struct {
	Addr     string
	User     string
	Pwd      string
	Database string
	Config   string
}
