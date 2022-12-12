package model

// Category 分类表
type Category struct {
	Model
	Name    string `gorm:"comment:分类名;not null;" json:"name"`
	Picture string `gorm:"comment:图片;" json:"picture"`
	Status  uint   `gorm:"comment:分类状态;default:0;not null;" json:"status"`
}
