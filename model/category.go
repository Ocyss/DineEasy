package model

import (
	"DineEasy/utils/errmsg"
)

// Category 分类表
type Category struct {
	Model
	Items   []Item `gorm:"foreignKey:CategoryID;" json:"items"`
	Name    string `gorm:"comment:分类名;not null;" json:"name"`
	Picture string `gorm:"comment:图片;" json:"picture"`
	Status  uint   `gorm:"comment:分类状态;default:0;not null;" json:"status"`
}

// AddCategory 增加分类
func AddCategory(data *Category) (int, any) {
	err = Db.Create(data).Error
	if err != nil {
		return errmsg.ERROR, err
	}
	return errmsg.SUCCESS, data.ID
}

func GetCategorys(load bool) (int, any) {
	var data []Category
	if load {
		err = Db.Preload("Items", "status in (1,3)").Find(&data).Error
	} else {
		err = Db.Find(&data).Error
	}
	if err != nil {
		return errmsg.ERROR, err
	}
	return errmsg.SUCCESS, data
}
