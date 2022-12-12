package errmsg

const (
	SUCCESS    = 200
	ERROR      = 500
	ERROR_BIND = 501
)

var codemsg = map[int]string{
	SUCCESS:    "OK",
	ERROR:      "ERROR",
	ERROR_BIND: "参数错误",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
