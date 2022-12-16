package model

import "DineEasy/utils/errmsg"

// Item 单品表
type Item struct {
	Model
	Name        string        `gorm:"comment:单品名;not null;" json:"name"`
	Description string        `gorm:"comment:单品描述;" json:"description"`
	Unit        string        `gorm:"comment:单位;" json:"unit"`
	MinQuantity uint          `gorm:"comment:起卖数;default:1;" json:"min_quantity"`
	Carts       []Cart        `gorm:"foreignKey:ItemsID;" json:"-"`
	ComboGroups []*ComboGroup `gorm:"many2many:combo_group_item;"`
	CategoryID  uint          `gorm:"comment:分类id;" json:"category_id"`
	Picture     string        `gorm:"comment:图片;" json:"picture"`
	Price       uint          `gorm:"comment:价格;default:0;not null;" json:"price"`
	Status      uint          `gorm:"comment:状态;default:0;not null;" json:"status"` //0未上架，1正常，2隐藏，3售空
}

// AddItem 增加单品
func AddItem(data *Item) (int, any) {
	err = Db.Create(data).Error
	if err != nil {
		return errmsg.ERROR, err
	}
	return errmsg.SUCCESS, data.ID
}

func GetItems(pageNum, pageSize int, cid int) (int, *[]Item, int64) {
	var data []Item
	var count int64
	err = Db.Where("status in (1,3) AND category_id=?", cid).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&data).Error
	if err != nil {
		return errmsg.ERROR, nil, 0
	}
	Db.Model(&Item{}).Where("status=? AND category_id=?", 1, cid).Count(&count)
	return errmsg.SUCCESS, &data, count
}
