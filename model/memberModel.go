package model

import "time"

type MemberModel struct {
     Id int64 `xorm:"pk autoincr" json:"id" form:"id"` //会员id
     Phone string `xorm:"notnull" json:"phone" form:"phone"` //手机号
	 Integral int64 `xorm:"default(0)" json:"integral" form:"integral"` //会员积分
	 MemberTime time.Time `xorm:"created" json:"member_time" form:"member_time"` //会员添加时间
}
