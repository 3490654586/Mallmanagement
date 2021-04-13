package server

import (
	"demo/model"
	"fmt"
	"github.com/go-xorm/xorm"
	"strconv"
	"time"
)

type CommdityServer interface {
	 SetCommdity(modelcomm model.CommdityModel)int
	 GetCommdity(name string,start string,pageSize string)[]model.CommdityModel
	 Getid(id int)(model.CommdityModel,error)
	 PutCommdity(modelcomm model.CommdityModel)error
	 DelectCommdity(id int)error
}

func NewCommXormeng(e *xorm.Engine)CommdityServer  {
	return &commdityServer{engine:e}
}

type commdityServer struct {
	engine *xorm.Engine
}

/**
 *插入商品
 */
func (c *commdityServer) SetCommdity(modelcomm model.CommdityModel)int{
	sql := "select * From commdity_model where commdity_name = ?"
	res ,_ := c.engine.QueryString(sql,modelcomm.CommdityName)
	if res != nil{
		return 0
	}
	_,err := c.engine.Insert(&modelcomm)
	if err != nil {
		fmt.Println(err)
	}
 return 1
}

/**
 *查询商品
 */
func (c *commdityServer)GetCommdity(name string,start string,pageSize string) []model.CommdityModel {

	var modes []model.CommdityModel
	   if pageSize == ""{
		   err := c.engine.Find(&modes)
		   if err != nil {
			   return nil
		   }
	   }else {

		   if  name != ""{
			   c.engine.Where("commdity_name=?",name).Find(&modes)
			   fmt.Println(modes)
			   return modes
		   }

	   	   sql := "SELECT * FROM commdity_model  LIMIT ?,? "
	       res,_:= c.engine.QueryString(sql,start,pageSize)
		   var modess model.CommdityModel

		   for _,v := range res{
			   id,_  := strconv.ParseInt(v["id"],10,10)
			   pid,_ := strconv.ParseInt(v["pid"],10,10)
			   number,_ := strconv.ParseInt(v["commdity_number"],10,10)
			   moye,_ := strconv.ParseFloat(v["commdity_moye"],10)
			   times,_:=   time.Parse(v["commdity_time"],"2006-01-02 15:04:05")
			      modess.Id =  id
                  modess.Pid = pid
                  modess.CommdityName = v["commdity_name"]
                  modess.CommdityNumber = number
                  modess.CommdityMoye = moye
                  modess.CommdityTime = times
			   modes =append(modes,modess)
		   }

	   }


	 return modes
}

//根据id查询商品
func (c *commdityServer)Getid(id int)(model.CommdityModel,error)  {
	 modes := new(model.CommdityModel)
	_,err := c.engine.Where("id=?",id).Get(modes)
	if err != nil {
		return *modes,err
	}

	return *modes,nil
}


//修改商品
func (c *commdityServer)PutCommdity(modelcomm model.CommdityModel) error {
	   _,err := c.engine.Id(modelcomm.Id).Cols("commdity_name","commdity_number","commdity_moye").Update(&modelcomm)
	if err != nil {
		return err
	}
	return nil
}

//删除商品
func (c *commdityServer)DelectCommdity(id int)error {
	sql := "delete from commdity_model where id = ?"

	_,err :=c.engine.Exec(sql,id)
	if err != nil {
		return err
	}

	return nil
}