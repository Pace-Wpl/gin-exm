package def

type ErrResp struct {
	Error     string `json:"message"`
	ErrorCode string `json:"error_code"`
}

var (
	ErrorRequestBodyPaseFailed = ErrResp{Error: "bad request!", ErrorCode: "001"}
	ErrorNotAuthUser           = ErrResp{Error: "not auth!", ErrorCode: "002"}
)
