package model

// Cart 购物车表
type Cart struct {
	Model
	MemberID     uint    `gorm:"comment:会员ID;default:0;not null;"  json:"member_id"`
	Items        []Item  `gorm:"foreignKey:ID;comment:单品列表;"  json:"item_id"`
	TtemQuantity uint    `gorm:"comment:单品数量;default:0;not null;"  json:"ttem_quantity"`
	TotalPrice   float64 `gorm:"comment:总价;default:0;not null;"  json:"total_price"`
	Status       uint    `gorm:"comment:购物车状态;default:0;not null;"  json:"status"`
	Member       Member
}
