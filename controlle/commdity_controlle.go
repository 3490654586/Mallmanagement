package controlle

import (
	"demo/model"
	"demo/server"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommdityeContro interface {
	Router()
	GetCommdity(c *gin.Context)
	SerCommdity(c *gin.Context)
	Getid(c *gin.Context)
	PutCommdity(c *gin.Context)
	DelectCommdity(c *gin.Context)
}

func NewCommeng(g *gin.Engine, s server.CommdityServer)CommdityeContro  {
return &commdityeContro{
	gineg:   g,
	xormeng: s ,
}
}

type commdityeContro struct {
	gineg *gin.Engine
	xormeng server.CommdityServer
}

func (c *commdityeContro) Router() {
	comm := c.gineg.Group("/comm")
	{
		comm.GET("/commdity",c.GetCommdity)
		comm.POST("/setcommdity",c.SerCommdity)
		comm.GET("/getidcommdity",c.Getid)
        comm.PUT("/putcommdity/:id",c.PutCommdity)
		comm.DELETE("/delcommdity/:id",c.DelectCommdity)
	}
}

/*
 *查询商品
 */
func (cc *commdityeContro)GetCommdity(c *gin.Context) {
	c.JSON(200,gin.H{})
	start :=  c.Query("start")
	pageSize := c.Query("pageSize")
     name := c.Query("query")
	   commModel := cc.xormeng.GetCommdity(name,start,pageSize)
		if commModel == nil {
			c.JSON(http.StatusOK,gin.H{
				"msg":"获取商品失败",
				 "stats":300,
			})
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"data":commModel,
			"msg":"获取商品成功",
			"stats":200,
		})
}

/*
 *添加酒水
 */
func (cc *commdityeContro)SerCommdity(c *gin.Context)  {
          var models model.CommdityModel
          c.ShouldBind(&models)

		nubber :=  c.PostForm("commdity_number")
		newNumber,_ :=strconv.ParseInt(nubber,10,10)
		models.CommdityNumber = newNumber

		moye :=c.PostForm("commdity_moye")
		newMoye,_ := strconv.ParseFloat(moye,10)
		models.CommdityMoye = newMoye

         code := cc.xormeng.SetCommdity(models)
		if code == 0 {
			c.JSON(http.StatusOK,gin.H{
				"statr":201,
				"msg":"酒水已经存在",
			})
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"statr":200,
			"msg":"酒水添加成功",
		})
}

/*
 *根据id查询商品
 */

func (cc *commdityeContro) Getid(c *gin.Context) {
	    id := c.Query("id")
	    newid,_ := strconv.Atoi(id)

	  data,err :=  cc.xormeng.Getid(newid)
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"mes":"获取商品失败",
				"stres":201,
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"data":data,
			"mes":"获取商品成功",
			"stres":200,
		})
}

/*
 *修改商品
 */
func (cc *commdityeContro)PutCommdity(c *gin.Context)  {
	      var model model.CommdityModel


           id := c.Param("id")
	       newid,_ := strconv.Atoi(id)
	       model.Id = int64(newid)

         buf := make([]byte,1024)
	     n, _ := c.Request.Body.Read(buf)
	    err:= json.Unmarshal(buf[0:n],&model)

	    err = cc.xormeng.PutCommdity(model)
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"msg":"修改商品失败",
				"stare":201,
			})
			return
		}

	c.JSON(http.StatusOK,gin.H{
		"msg":"修改商品成功",
		"stare":200,
	})
}

/*
 *删除商品
 */
func (cc *commdityeContro)DelectCommdity(c *gin.Context){
	   id := c.Param("id")

	  idInt,_ := strconv.Atoi(id)
	  err := cc.xormeng.DelectCommdity(idInt)

	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"msg":"删除商品失败",
			"stars":201,
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"msg":"删除商品成功",
		"stars":200,
	})
}