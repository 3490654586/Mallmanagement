package controlle

import (
	"demo/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MenuContro interface {
	Router()
	GetMenu(c *gin.Context)
}

func NewMengeng(g *gin.Engine,e server.MenuServer)MenuContro  {
	return &menuContro{
		gineng:g,
		xormeng: e,
	}
}

type menuContro struct {
	gineng *gin.Engine
	xormeng server.MenuServer
}

func (m *menuContro)Router()  {
	menu := m.gineng.Group("/menu")
	{
		menu.GET("/terr",m.GetMenu)
	}
}

//查询子节点
func (m *menuContro)GetMenu(c *gin.Context) {
	  treeMenu := m.xormeng.GetLabse()
	  c.JSON(http.StatusOK,gin.H{
	  	 "date":treeMenu,
	  	  "ment":"获取菜单栏成功",
	  })
}
