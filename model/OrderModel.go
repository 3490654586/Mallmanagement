package model


type OrderModel struct {
	Id int `xorm:"pk autoincr" json:"id"  form:"id"` //商品id
	MemberPhone string `xorm:"varchar(20)" json:"MemberPhone" form:"MemberPhone"` //会员手机号
	Moyer float64  `xorm:"notnull" json:"moyer" form:"moyer"` //收款数
	SetIntegral int64 `xorm:"default 0" json:"set_integral" form:"set_integral"` //新增积分
}