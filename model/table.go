package model

// Table 餐桌表
type Table struct {
	Model
	Number          uint `gorm:"comment:桌号;not null;"  json:"number"`
	StoreID         uint `gorm:"comment:门店id;not null;"  json:"Store_id"`
	Store           Store
	Capacity        uint `gorm:"comment:最大容量;default:1;not null;"  json:"capacity"`
	CurrentCapacity uint `gorm:"comment:当前人数;default:0;not null;"  json:"current_capacity"`
	Status          uint `gorm:"comment:餐桌状态;default:0;not null;"  json:"status"`
}
