package utlls

import "github.com/gin-gonic/gin"

var (
	SESSE = 0
	FILY  =1
)

func ResponseJson(c *gin.Context,v interface{}){
	c.JSON(200,map[string]interface{}{
		"message":SESSE,
		"data":v,
		"messages":"成功",
	})
}

func ResponseJsonShi(c *gin.Context,v interface{}){
	c.JSON(200,map[string]interface{}{
		"message":FILY,
		"data":v,
	})
}
