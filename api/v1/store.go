package v1

import (
	"DineEasy/model"
	"DineEasy/utils/errmsg"
	"github.com/gin-gonic/gin"
)

func GetStores(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    0,
		"message": "ok",
		"data":    []string{"store1", "store2"},
	})
}

// AddStore 添加分店
func AddStore(c *gin.Context) {
	var data model.Store
	err = c.ShouldBindJSON(data)
	if err != nil {
		code = errmsg.ERROR_BIND
		c.JSON(200, gin.H{
			"code":    code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	code, id := model.AddStore(&data)
	c.JSON(200, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
		"id":      id,
	})
}
