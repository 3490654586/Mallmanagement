package utlls

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)




var Jwtkey = []byte("89js82js72")


type Myclaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//生成token
func SetToken(username string)string{
          SetClaime := &Myclaims{
			  Username:       username,
			  StandardClaims: jwt.StandardClaims{
				  ExpiresAt:  time.Now().Add(7 * 24 *time.Hour).Unix(), //到期时间
				  Issuer:    "zhang",
			  },
		  }

		  reqclaims :=  jwt.NewWithClaims(jwt.SigningMethodHS256,SetClaime)
		  token ,_ := reqclaims.SignedString(Jwtkey)
		  return token
}

//验证token
func CheckToken (token string) (*Myclaims,string) {
	    var claims Myclaims

	setToken,err := jwt.ParseWithClaims(token,&claims, func(token *jwt.Token) (interface{}, error) {
			 return Jwtkey,nil
		 })


	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil,"令牌格式不正确"
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, "令牌过期"
			} else {
				return nil, "令牌错误"
			}
		}
	}

	if setToken != nil {
		if key, ok := setToken.Claims.(*Myclaims); ok && setToken.Valid {
			return key, "验证成功"
		} else {
			return nil, "验证失败"
		}
	}
	return nil, "验证失败"
}

//jwt中间件
// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
        //从请求头获取token
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			ResponseJson(c,"TOKEN不存在")
			c.Abort()
			return
		}

		//去除空格,判断token的长度
		//checkToken := strings.Split(tokenHeader, " ")
		////长度为0返回
		//if len(checkToken) == 0 {
		//	ResponseJson(c,"TOKEN格式错误")
		//	c.Abort()
		//	return
		//}
		//
        //fmt.Println("checkToken=",checkToken)
		//
		//
		////如果token的长度不等于并且开头不等于Bearer返回
		//if len(checkToken) != 2 && checkToken[0] != "Bearer" {
		//	ResponseJson(c,"TOKEN格式错误")
		//	c.Abort()
		//	return
		//}

		key, tCode := CheckToken(tokenHeader)
		if tCode !="验证成功" {
			ResponseJson(c,tCode)
			c.Abort()
			return
		}
		c.Set("username", key)
		c.Next()
	}
}



