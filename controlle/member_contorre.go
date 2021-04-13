package controlle

import (
	"demo/model"
	"demo/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MemberContolle interface {
	Router()
	Setmember(c *gin.Context)
	Getmember(c *gin.Context)
	Putmember(c *gin.Context)
	GetPhone(c *gin.Context)
}

func NewMemberEng(e *gin.Engine,s server.MemberServer)MemberContolle{
	return &memberContolle{
		engine:e,
		servers:s,
	}
}

type memberContolle struct {
	engine *gin.Engine
	servers server.MemberServer
}

func (m *memberContolle)Router(){
	member := m.engine.Group("/member")
	{
		member.POST("/setmember",m.Setmember)
		member.GET("/getmember",m.Getmember)
		member.POST("/putmember",m.Putmember)
		member.GET("/getphonemember/:phone",m.GetPhone)
	}
}

/*
 *添加会员
 */
func (m *memberContolle)Setmember(c *gin.Context) {
           var models model.MemberModel
          fmt.Println("添加会员=",models)
           c.ShouldBind(&models)

        int, err :=  m.servers.SetMember(models)
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"msg":"添加会员失败",
				"start":201,
			})
			return
		}

		if int == 2 {
			c.JSON(http.StatusOK,gin.H{
				"msg":"会员已经存在",
				"start":202,
			})
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"msg":"添加会员成功",
			"start":200,
		})
}

/*
 *查询全部会员
 */
func (m *memberContolle)Getmember(c *gin.Context){
	models := m.servers.GetMember()

	if models == nil {
		c.JSON(http.StatusOK,gin.H{
			"msg":"获取会员列表失败",
			"strat":201,
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"data":models,
		"msg":"获取会员列表成功",
		"strat":200,
	})
}

/*
 *修改会员积分
 */
func (m *memberContolle)Putmember(c *gin.Context){
	var models model.MemberModel

	c.ShouldBind(&models)
	fmt.Println("会员=",models.Integral)
   err :=  m.servers.PutMember(models)

	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"msg":"修改积分失败",
			"start":201,
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"msg":"修改积分成功",
		"start":200,
	})
}

/**
 查询指定会员
 */
func (m *memberContolle)GetPhone(c *gin.Context){
	phone := c.Param("phone")


          boo,models :=  m.servers.GetPhoneMember(phone)

          if boo != true {
          	c.JSON(http.StatusOK,gin.H{
          		"msg":"没有该会员",
          		"start":201,
			})
			  return
		  }
		c.JSON(http.StatusOK,gin.H{
			"data":models,
			"msg":"没有该会员",
			"start":200,
		})
}