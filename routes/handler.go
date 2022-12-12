package routes

import (
	"DineEasy/utils"
	"DineEasy/utils/errmsg"
	"github.com/gin-gonic/gin"
)

type MyHandler func(c *gin.Context) (int, any)

// Handler 统一Gin返回格式
func Handler() func(h MyHandler) gin.HandlerFunc {
	return func(h MyHandler) gin.HandlerFunc {
		return func(c *gin.Context) {
			code, data := h(c)
			req := gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			}
			if code == 200 {
				//判断数据类型
				if val, ok := data.(gin.H); ok {
					for k, v := range val {
						req[k] = v
					}
				} else {
					req["data"] = data
				}
				c.JSON(200, req)
			} else {
				//判断是否debug模式，是的话返回错误信息
				if utils.AppMode == "debug" {
					req["errmsg"] = data
				}
				c.JSON(400, req)
			}

		}
	}
}
