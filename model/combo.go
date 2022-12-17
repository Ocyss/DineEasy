package model

import (
	"DineEasy/utils/errmsg"
	"gorm.io/gorm"
)

// Combo 套餐
type Combo struct {
	Model
	Name        string       `gorm:"comment:套餐名;not null;" json:"name"`
	Description string       `gorm:"comment:套餐描述;" json:"description"`
	MinQuantity uint         `gorm:"comment:起卖数;default:1;" json:"min_quantity"`
	ComboGroups []ComboGroup `gorm:"comment:对应套餐分组;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"combo_group"`
	Carts       []Cart       `gorm:"foreignKey:ComboID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Picture     string       `gorm:"comment:图片;" json:"picture"`
	Price       uint         `gorm:"comment:价格;default:0;not null;" json:"price"`
	Status      uint         `gorm:"comment:状态;default:0;not null;" json:"status"` //0未上架，1正常，2隐藏，3售空
}

// ComboGroup 套餐分组
type ComboGroup struct {
	Model
	Name           string  `gorm:"comment:分组名;not null" json:"name"`
	Fixed          bool    `gorm:"comment:是否固定;default:true;not null" json:"fixed"`
	Same           bool    `gorm:"comment:是否可选相同单品;default:true;not null" json:"same"`
	SelectQuantity uint    `gorm:"comment:可选数量;default:1;not null;" json:"select_quantity"`
	Items          []*Item `gorm:"many2many:combo_group_item;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items"`
	ComboID        uint    `json:"combo_id"`
}

// ComboGroupItem 自定义连接表
type ComboGroupItem struct {
	ComboGroupID uint `gorm:"primaryKey;" json:"combo_group_id"`
	ItemID       uint `gorm:"primaryKey;" json:"item_id"`
	Quantity     uint `gorm:"comment:数量;" json:"quantity"`
	Price        uint `gorm:"comment:价格;" json:"price"`
	Must         bool `gorm:"comment:必选;" json:"must"`
}

func (ComboGroupItem) BeforeCreate(db *gorm.DB) error {
	err = db.SetupJoinTable(&ComboGroup{}, "Items", &ComboGroupItem{})
	return err
}

func AddCombo(data *Combo) (int, any) {
	err = Db.Create(data).Error
	if err != nil {
		return errmsg.ERROR, err
	}
	return errmsg.SUCCESS, data.ID
}

func DelCombo(cid uint) (int, any) {
	tx := Db.Begin()
	//err = tx.Select(clause.Associations).Delete(&ComboGroup{}, "combo_id=?", cid).Error
	//if err != nil {
	//	tx.Rollback()
	//	return errmsg.ERROR, err
	//}
	err = tx.Delete(&Combo{Model: Model{ID: cid}}).Error

	if err != nil {
		tx.Rollback()
		return errmsg.ERROR, err
	}
	tx.Commit()
	return errmsg.SUCCESS, nil
}
