package model

// Member 会员表
type Member struct {
	Model
	Name     string  `gorm:"comment:会员姓名;not null;" json:"name"`
	Phone    string  `gorm:"comment:会员手机号;" json:"phone"`
	Password string  `gorm:"comment:支付密码;" json:"password"`
	Points   uint    `gorm:"comment:会员积分;default:0;not null;" json:"points"`
	Balance  float64 `gorm:"comment:会员余额;default:0;not null;" json:"balance"`
	Level    uint    `gorm:"comment:会员等级;default:0;not null;" json:"level"`
	Status   uint    `gorm:"comment:会员状态;default:0;not null;" json:"status"`
}
