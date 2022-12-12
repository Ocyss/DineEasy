package model

// Order 订单表
type Order struct {
	Model
	MemberID uint `gorm:"comment:会员ID;not null;"  json:"member_id"`
	Member   Member
	TableID  uint    `gorm:"comment:餐桌号;not null;"  json:"table_id"`
	Amount   float64 `gorm:"comment:总金额;default:0;not null;"  json:"amount"`
	Status   uint    `gorm:"comment:订单状态;default:0;not null;"  json:"status"`
}
