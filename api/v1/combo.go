package v1

import (
	"DineEasy/model"
	"DineEasy/utils/errmsg"
	"github.com/gin-gonic/gin"
)

func AddCombo(c *gin.Context) (int, any) {
	var data model.Combo
	err = c.ShouldBindJSON(&data)
	if err != nil {
		return errmsg.ERROR_BIND, err
	}
	code, comboID := model.AddCombo(&data)
	return code, gin.H{"combo_id": comboID}
}

func DelCombo(c *gin.Context) (int, any) {
	var data struct {
		Cid uint `json:"cid"`
	}
	err = c.ShouldBindJSON(&data)
	if err != nil {
		return errmsg.ERROR_BIND, err
	}
	return model.DelCombo(data.Cid)
}
