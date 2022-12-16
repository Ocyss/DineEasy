package model

// Cart 购物车表
type Cart struct {
	Model
	MemberID   uint    `gorm:"comment:会员ID;not null;"  json:"member"`
	ItemsID    uint    `gorm:"comment:对应单品;" json:"items"`
	ComboID    uint    `gorm:"comment:对应套餐;" json:"combos"`
	Quantity   uint    `gorm:"comment:数量;default:0;not null;"  json:"item_quantity"`
	TotalPrice float64 `gorm:"comment:总价;default:0;not null;"  json:"total_price"`
	Status     uint    `gorm:"comment:购物车状态;default:0;not null;"  json:"status"`
	Member     Member
}
