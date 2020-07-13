package def

const (
	//config dir
	CONFIG_DIR = "../conf/config.yaml"
)

var (
	Conf = &Config{}
)

type Config struct {
	Httpport           string
	AppName            string
	SessionExpired     int
	CookieKey1         string
	CookieKey2         string
	CookieKey3         string
	Domain             string
	DefaultIcon        string
	UserSecAccessLimit int
	IpSecAccessLimit   int
	StreamLimit        int
	ReqTimeout         int
	ReqChannelBuffer   int
	ReadGoroutineNum   int
	WriteGoroutineNum  int

	Log   LogConf
	Redis RedisConf
	Etcd  EtcdConf
	Mysql MysqlConf
}

type ProductConf struct {
	ProductID int   `json:"product_id"`
	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`
	Status    int   `json:"status"`
	Total     int   `json:"total"`
}

type LogConf struct {
	GinLogPath  string
	GinLogLevel string
	SysLogPath  string
	SysLogLevel string
}

type RedisConf struct {
	Addr         string
	MaxIdle      int
	MaxActive    int
	IdleTimeout  int
	SecReqQueue  string
	SecRespQueue string
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
