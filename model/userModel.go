package model


type UserModer struct {
	Id int64 `xorm:"pk autoincr" json:"id" form:"id"'`
	Username string `xorm:"varchar(20) notnull" json:"username" form:"username"`
	Password string `xorm:"varchar(20) notnull" json:"password" form:"Password"`
}
