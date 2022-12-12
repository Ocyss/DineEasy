package model

// Item 单品表
type Item struct {
	Model
	Name        string `gorm:"comment:单品名;not null;" json:"name"`
	Description string `gorm:"comment:单品描述;" json:"description"`
	CategoryID  uint   `gorm:"comment:分类id;" json:"category_id"`
	Category    Category
	Picture     string `gorm:"comment:图片;" json:"picture"`
	Price       uint   `gorm:"comment:价格;default:0;not null;" json:"price"`
	Status      uint   `gorm:"comment:状态;default:0;not null;" json:"status"`
}
