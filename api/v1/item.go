package v1

import (
	"DineEasy/model"
	"DineEasy/utils/errmsg"
	"DineEasy/utils/tool"
	"github.com/gin-gonic/gin"
	"strconv"
)

// AddItem 添加单品
func AddItem(c *gin.Context) (int, any) {
	var data model.Item
	err = c.ShouldBindJSON(&data)
	if err != nil {
		return errmsg.ERROR_BIND, err
	}
	code, itemID := model.AddItem(&data)
	return code, gin.H{"item_id": itemID}
}

// GetItems 获取单品列表
func GetItems(c *gin.Context) (int, any) {
	pageNum, pageSize := tool.PageTool(c)
	cid, _ := strconv.Atoi(c.Query("cid"))
	if cid == 0 {
		cid = 1
	}
	code, data, count := model.GetItems(pageNum, pageSize, cid)
	return code, gin.H{"data": data, "count": count}
}
