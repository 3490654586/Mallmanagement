package controlle

import (
	"demo/model"
	"demo/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)



type OrderContro interface {
	Router()
	SetOder(c *gin.Context)
}

func NewOrderen(eng *gin.Engine, s server.OrderServer)OrderContro{
     return &orderContro{
     	Engine:eng,
     	servers:s,
	 }
}

type orderContro struct {
	Engine *gin.Engine
	servers server.OrderServer
}

func (o *orderContro)Router() {
	order := o.Engine.Group("/order")
	{
		order.POST("/setorder",o.SetOder)
	}
}

//添加订单
func (o *orderContro)SetOder(c *gin.Context)  {
          var modes model.OrderModel
          c.ShouldBind(&modes)
          fmt.Println("model=",modes)

         err := o.servers.SetOrder(modes)
			if err != nil {
				c.JSON(http.StatusOK,gin.H{
					"msg":"付款失败",
					 "start":201,
				})
				return
			}

		c.JSON(http.StatusOK,gin.H{
			"msg":"付款成功",
			"start":200,
		})
}