package routes

import (
	v1 "DineEasy/api/v1"
	"DineEasy/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	//设置Gin模式
	gin.SetMode(utils.AppMode)
	//强制日志颜色
	gin.ForceConsoleColor()
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.Use(gin.Recovery())

	router := r.Group("api/v1")

	router.GET("/stores", v1.GetStores)
	router.POST("/store/add", v1.AddStore)
	err := r.Run(utils.HttpPort)
	if err != nil {
		panic(fmt.Sprintf("路由启动失败，请检查配置.%v", err))
	}

}
