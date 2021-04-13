package server

import (
	"demo/model"
	"github.com/go-xorm/xorm"
)

type MemberServer interface {
	 SetMember(models model.MemberModel)(int,error)
	 GetMember()[]model.MemberModel
	 PutMember(models model.MemberModel)error
	GetPhoneMember(phone string)(bool,model.MemberModel)
}

func NewMemberXorm(e *xorm.Engine)MemberServer{
       return &memberServer{
       	 engine:e,
	   }
}

type memberServer struct {
	engine *xorm.Engine
}

//添加会员
func (m *memberServer)SetMember(models model.MemberModel)(int,error) {
      sql := "select * from member_model where phone = ?"
      mapModer, _ := m.engine.QueryString(sql,models.Phone)

	if mapModer == nil {
		_,err := m.engine.Insert(models)

		if err != nil {
			return 0,err
		}

		return 1 ,nil
	}else {
		return 2 ,nil
	}

}

//查询全部会员
func (m *memberServer)GetMember()[]model.MemberModel {
	   var models []model.MemberModel
	   err := m.engine.Find(&models)

		if err != nil {
			return nil
		}

		return models
}


//修改会员积分(增加)
func (m *memberServer)PutMember(models model.MemberModel) error {
          sql := "update member_model set integral = integral + ? where phone = ?"

          _,err := m.engine.Exec(sql,models.Integral,models.Phone)

			if err != nil {
				return err
			}

			return nil
}

//查询指定会员
func (m *memberServer)GetPhoneMember(phone string)(bool,model.MemberModel){

	    var modls model.MemberModel

	   boo,err :=  m.engine.Where("phone = ?",phone).Get(&modls)

		if err != nil {
			 return boo,modls
		}

		return boo,modls
}