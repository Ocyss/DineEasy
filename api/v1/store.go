package v1

import (
	"DineEasy/model"
	"DineEasy/utils/errmsg"
	"DineEasy/utils/tool"
	"github.com/gin-gonic/gin"
)

// GetStores 获取门店列表
func GetStores(c *gin.Context) (int, any) {
	pageNum, pageSize := tool.PageTool(c)
	code, data, count := model.GetStores(pageNum, pageSize)
	return code, gin.H{"data": data, "count": count}
}

// GetStore 获取单个门店

// AddStore 添加分店
func AddStore(c *gin.Context) (int, any) {
	var data model.Store
	err = c.ShouldBindJSON(&data)
	if err != nil {
		code = errmsg.ERROR_BIND
		return code, err
	}
	code, storeID := model.AddStore(&data)
	return code, gin.H{"store_id": storeID}
}

// DelStore 删除门店

// UpdateStore 修改门店
