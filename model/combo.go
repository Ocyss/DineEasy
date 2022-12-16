package model

import (
	"time"
)

// Combo 套餐
type Combo struct {
	Model
	Name        string       `gorm:"comment:套餐名;not null;" json:"name"`
	Description string       `gorm:"comment:套餐描述;" json:"description"`
	MinQuantity uint         `gorm:"comment:起卖数;default:1;" json:"min_quantity"`
	ComboGroups []ComboGroup `gorm:"comment:对应套餐分组;" json:"combo_group"`
	Carts       []Cart       `gorm:"foreignKey:ComboID;" json:"-"`
	Picture     string       `gorm:"comment:图片;" json:"picture"`
	Price       uint         `gorm:"comment:价格;default:0;not null;" json:"price"`
	Status      uint         `gorm:"comment:状态;default:0;not null;" json:"status"` //0未上架，1正常，2隐藏，3售空
}

// ComboGroup 套餐分组
type ComboGroup struct {
	Model
	Name           string  `gorm:"comment:分组名;not null"`
	Fixed          bool    `gorm:"comment:是否固定;default:true;not null"`
	Same           bool    `gorm:"comment:是否可选相同单品;default:true;not null"`
	SelectQuantity uint    `gorm:"comment:可选数量;default:1;not null;"`
	Items          []*Item `gorm:"many2many:combo_group_item;"`
	ComboID        uint
}

// ComboGroupItem 自定义连接表
type ComboGroupItem struct {
	ComboGroupID uint `gorm:"primaryKey"`
	ItemID       uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	Price        uint `gorm:"comment:价格;"`
	Must         bool `gorm:"comment:必选;"`
}
