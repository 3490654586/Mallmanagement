package server

import (
	"demo/model"
	"github.com/go-xorm/xorm"
)

type MenuServer interface {
	GetLabse()[]TreeList
}

func NewMenXormeng(eng *xorm.Engine)MenuServer  {
	return &menuServer{
		 engine:eng,
	}
}

type menuServer struct {
	engine *xorm.Engine
}


type TreeList struct {
	Id int64   `json:"id"`
	Pid int64  `json:"pid"`
	Label string `json:"label"`
	Path string   `json:"path"`
	Children []*TreeList
}
//查询节点树
func (m *menuServer) GetLabse()[]TreeList{
      //sql := "select a.id,a.pid,a.label,a.path,b.id,b.pid,b.label,b.path from menu_model AS a menu_model AS b where a.id = b.pid "
	//sql1 := "select * from menu_model AS a, menu_model AS b where a.id = b.pid "

    //查询父节点
	var menu []model.MenuModel
	m.engine.Where("pid = ?", 0).Find(&menu)

    var Node []TreeList
	for _,v:= range menu{
		//父节点赋值给Tree结构体
		node := TreeList{
			Id:       v.Id,
			Pid:      v.Pid,
			Label:    v.Label,
			Path:     v.Path,
			Children: []*TreeList{},
		}

		//查询子节点
		var chenr  []model.MenuModel
		m.engine.Where("pid = ?", v.Id).Find(&chenr)

		for _,c := range chenr {
			 //赋值给Tree
			 chid := TreeList{c.Id,c.Pid,c.Label,c.Path,[]*TreeList{}}
			 //添加到子节点
			node.Children=append(node.Children,&chid)
		}
		Node = append(Node,node)
	}
	return Node
}


