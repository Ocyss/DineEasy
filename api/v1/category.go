package v1

import (
	"DineEasy/model"
	"DineEasy/utils/errmsg"
	"github.com/gin-gonic/gin"
	"strconv"
)

// AddCategory 增加分类
func AddCategory(c *gin.Context) (int, any) {
	var data model.Category
	err = c.ShouldBindJSON(&data)
	if err != nil {
		return errmsg.ERROR_BIND, err
	}
	return model.AddCategory(&data)
}

func GetCategorys(c *gin.Context) (int, any) {
	load, _ := strconv.ParseBool(c.Query("load"))
	return model.GetCategorys(load)
}
