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

	{
		{ //通用接口
			router.POST("/upload", Handler()(v1.Upload))
		}
		{ //分类接口
			router.POST("/category/add", Handler()(v1.AddCategory))
			router.GET("/categorys", Handler()(v1.GetCategorys))
		}
		{ //单品接口
			router.POST("/item/add", Handler()(v1.AddItem))
			router.GET("/items", Handler()(v1.GetItems))
		}
		{ //套餐接口
			router.POST("/combo/add", Handler()(v1.AddCombo))
			router.POST("/combo/del", Handler()(v1.DelCombo))
		}
		{ //门店配置
			router.GET("/stores", Handler()(v1.GetStores))
			router.POST("/store/add", Handler()(v1.AddStore))
		}
	}

	err := r.Run(utils.HttpPort)
	if err != nil {
		panic(fmt.Sprintf("路由启动失败，请检查配置.%v", err))
	}

}
