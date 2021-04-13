package server

import (
	"demo/model"
	"github.com/go-xorm/xorm"
)

type OrderServer interface {
	SetOrder(models model.OrderModel)error
}

func NewOrderXormEng(eng *xorm.Engine)OrderServer{
	return &orderServer{engine:eng}
}

type orderServer struct {
	engine *xorm.Engine
}

func (o *orderServer)SetOrder(models model.OrderModel)error{
	 _,err := o.engine.Insert(&models)
	if err != nil {
		return err
	}

	 if models.MemberPhone != "" {
	 	sql := "update member_model set integral = integral + ? where phone = ? "
	 	o.engine.Exec(sql,models.SetIntegral,models.MemberPhone)
	 }

	 return nil
}