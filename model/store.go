package model

import (
	"DineEasy/utils/errmsg"
)

// Store 门店表
type Store struct {
	Model
	Name      string `gorm:"comment:门店名字;not null;"  json:"name"`
	Address   string `gorm:"comment:门店地址;"  json:"address"`
	Phone     string `gorm:"comment:门店电话;"  json:"phone"`
	Picture   string `gorm:"comment:门店照片;"  json:"picture"`
	StartTime uint   `gorm:"comment:营业开始时间;" json:"start_time"`
	EndTime   uint   `gorm:"comment:营业结束时间;" json:"end_time"`
	Status    uint   `gorm:"comment:门店状态;default:0;not null;"  json:"status"`
}

func AddStore(data *Store) (int, uint) {
	err := Db.Create(data).Error
	if err != nil {
		return errmsg.ERROR, 0
	}
	return errmsg.SUCCESS, data.ID
}
