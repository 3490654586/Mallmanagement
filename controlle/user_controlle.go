package controlle

import (
	"demo/model"
	"demo/server"
	"demo/utlls"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserContro interface {
	Router()
	LoginUser(c *gin.Context)
	Gettoke(c *gin.Context)
}

func NewUsereng(e *gin.Engine,x server.UserServer)UserContro  {
	return &userContro{
		gineng:  e,
		xormeng: x,
	}
}

type userContro struct {
	gineng *gin.Engine
	xormeng server.UserServer
}

func (u *userContro)Router() {
	user := u.gineng.Group("user")

	{
		//用户
		user.POST("/LoginUser",u.LoginUser)
		user.Use(utlls.JwtToken())
		user.GET("/token",u.Gettoke)
		//侧边栏
		user.GET("/menu",)
	}
}

//用户登录
func (u *userContro)LoginUser(c *gin.Context) {
	    var user model.UserModer
        c.ShouldBind(&user)

	    res :=  u.xormeng.LoginUser(user.Username,user.Password)

	    if res == "账号不存在" {
	    	utlls.ResponseJson(c,"没有该用户")
		} else if res == "登录成功"{
              token :=  utlls.SetToken(user.Username)
              c.JSON(http.StatusOK,gin.H{
              	 "data":"登录成功",
              	 "token":token,
			  })
		}else {
			utlls.ResponseJson(c,"账号密码错误")
		}
}

//
func (u *userContro)Gettoke(c *gin.Context) {
	utlls.ResponseJson(c,"41")
}
