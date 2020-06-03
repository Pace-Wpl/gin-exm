package def

const (
	//默认头像dir
	DEFAULT_ICON = "../icon/default.jpg"
	//session过期时间，60s * 60m * 24h
	SESSION_EXPIRED = 60 * 60 * 24
)

type RespMes struct {
	Code int    `json:"code"`
	Mes  string `json:"message"`
}

type ReqUser struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Username string `json:"name"`
	Pwd      string `json:"password"`
	Icon     string `json:"icon"`
}

type Session struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
