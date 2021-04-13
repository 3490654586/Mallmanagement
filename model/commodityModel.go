package model

import "time"

/**
  商品模型
 */
type CommdityModel struct {
	Id int64 `xorm:"pk autoincr" json:"id" form:"id"`      //商品id
	Pid int64 `xorm:"notnull" json:"pid" form:"pid"`   //属于哪个分类下面
	CommdityName string `xorm:"varchar(50) notnull" json:"commdity_name" form:"commdity_name" uri:"commdity_name"`  //商品名字
	CommdityNumber int64 `xorm:"default(0)" json:"commdity_number" form:"commdity_number" uri:"commdity_number"` //商品
	CommdityMoye float64 `xorm:"notnull" json:"commdity_moye" form:"commdity_moye" uri:"commdity_moye"`//商品价格
	CommdityTime time.Time `xorm:"created" json:"commdity_time" form:"commdity_number"` //商品添加时间
}
