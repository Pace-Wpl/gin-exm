package def

const (
	//config dir
	CONFIG_DIR = "../conf/config.yaml"
)

var (
	Conf = &Config{}
)

type Config struct {
	RequestWaitTimeOut  int
	ResponseSendTimeOut int
	ReqChannelBuffer    int
	RespChannelBuffer   int
	ReadGoroutineNum    int
	WriteGoroutineNum   int
	HandleGoroutineNum  int
	CryptoStr           string
	Log                 LogConf
	Redis               RedisConf
	Etcd                EtcdConf
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
	ControlKey string
}

type Control struct {
	BusinessSwitch int
}
