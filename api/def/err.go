package def

type ErrResp struct {
	Error     string `json:"message"`
	ErrorCode string `json:"error_code"`
}

var (
	ErrorRequestBodyPaseFailed = ErrResp{Error: "bad request!", ErrorCode: "001"}
	ErrorNotAuthUser           = ErrResp{Error: "not auth!", ErrorCode: "002"}
	ErrorInternalError         = ErrResp{Error: "internal error!", ErrorCode: "003"}
	ErrorServerBusy            = ErrResp{Error: "server busy!", ErrorCode: "004"}
	ErrorUserNotLogin          = ErrResp{Error: "user not login!", ErrorCode: "005"}
	ErrorInvalidReq            = ErrResp{Error: "invalid request!", ErrorCode: "006"}
	ErrorRequestTimeOut        = ErrResp{Error: "request time out!", ErrorCode: "007"}
)
