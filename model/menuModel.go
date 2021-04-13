package model



type MenuModel struct {
     Id int64 `xorm:"pk autoincr" json:"id" form:"id"`   //一级菜单id
     Pid int64 `xorm:"default(0)" json:"pid" form:"pid"` //是否有二级菜单 0代表没有
	 Label string `xorm:"varchar(50) notnull" json:"label" form:"label"`// menu名称
	 Path string `xorm:"default('null')" json:"path" form:"path"`
}
