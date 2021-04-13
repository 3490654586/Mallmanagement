package database

import (
    "fmt"
    "testing"
    "unsafe"
)

func TestSwaf(t *testing.T)  {
  //  NewEngine()
  //
  //models := &model.OrderModel{
  //    MemberPhone: "15965650030",
  //    Moyer:       100,
  //    SetIntegral: 100,
  //}
  //  order := server.NewOrderXormEng(DataEngine)
  //  order.SetOrder(*models)
  type T1 struct {
      a struct{}
      x int64
  }
  fmt.Println("nihao",unsafe.Sizeof(T1{}))
}
