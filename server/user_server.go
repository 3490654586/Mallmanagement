package server

import (
	"fmt"
	"github.com/go-xorm/xorm"
)

type UserServer interface {
	//QuerServer() //查询
	LoginUser(username string, password string)string //登录
}

func NewUserXormEngine(eng *xorm.Engine) UserServer {
	return &userServer{engine: eng}
}

type userServer struct {
	engine *xorm.Engine
}

//登录
func (u *userServer)LoginUser(username string, password string)string{
	   sql := "select * from user_moder where username = ?"
       res ,_ :=  u.engine.QueryString(sql,username)
        fmt.Println("res",res)
       if res == nil {
       	   return "账号不存在"
	   }

	   if res[0]["username"] == username && res[0]["password"] == password {
	   	        return "登录成功"
	   }else {
	   	       return "账号密码不对"
	   }
}


