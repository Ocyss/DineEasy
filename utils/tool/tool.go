package tool

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// PageTool 分页通用获取
func PageTool(c *gin.Context) (int, int) {
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))   //分页偏移量
	pageSize, _ := strconv.Atoi(c.Query("pagesize")) //分页最大数
	if pageNum <= 0 {
		pageNum = -1
	}
	if pageSize <= 0 || pageSize > 20 {
		pageSize = 10
	}
	return pageNum, pageSize
}
