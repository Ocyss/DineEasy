package errmsg

const (
	SUCCESS    = 200
	ERROR      = 500
	ERROR_BIND = 501

	ERROR_FILE_WRONG  = 601
	ERROR_CLASS_WRONG = 602
)

var codemsg = map[int]string{
	SUCCESS:           "OK",
	ERROR:             "ERROR",
	ERROR_BIND:        "参数错误",
	ERROR_FILE_WRONG:  "文件错误!",
	ERROR_CLASS_WRONG: "文件类别出错",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
