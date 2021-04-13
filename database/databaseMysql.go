package database

import (
	"demo/model"
	"demo/utlls"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var DataEngine  *xorm.Engine

func NewEngine (){
	//_conf := utlls.OpenConf("../config/config.json")
		_conf := utlls.GetConf()
	fmt.Println(1)
	fmt.Println("_conf=",_conf.Mysqls.Databases)
	engine,err := xorm.NewEngine(_conf.Mysqls.Databases,fmt.Sprint(_conf.Mysqls.User+":"+_conf.Mysqls.Passwad+"@tcp("+_conf.Mysqls.Host+":"+_conf.Mysqls.Port+")/"+_conf.Mysqls.Db_Name))
	if err != nil{
		fmt.Println("err=",err)
		return
	}
	err = engine.Sync2(new(model.UserModer),
		new(model.MenuModel),
		new(model.CommdityModel),
		new(model.MemberModel),
		new(model.OrderModel),
		)
	if err!=nil{
		fmt.Println("err=",err)
	}

	DataEngine = engine
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	engine.SetMaxOpenConns(10)
}
