package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"second_hand/DAO"
	"time"
)

var jwtSecret = []byte("4dv6s46sdv4s6d9vfsf")

// CreateToken
// @Summary 生成Token
// @Description 使用jwt包生成Token
// @Tags 登录
// @Accept json
// @Produce json

//func CreateToken(c *gin.Context) (string, error) {
//	var customerClaims Claims
//	err := c.ShouldBind(&customerClaims)
//	if err != nil {
//		fmt.Println(err)
//		return "", err
//	}
//
//	now := time.Now()
//	expireTime := jwt.NumericDate{
//		now.Add(5 * time.Hour),
//	}
//
//	//设置过期时间
//	customerClaims.RegisteredClaims.ExpiresAt = &expireTime
//
//	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, customerClaims)
//	token, err := tokenClaims.SignedString(jwtSecret) //对token进行签名
//	return token, err
//}
//
//// ValidateToken 验证token
//func ValidateToken(tokenString string) (*Claims, error) {
//	token, err := jwt.ParseWithClaims(tokenString, &Claims{},
//		func(token *jwt.Token) (interface{}, error) {
//			return jwtSecret, nil
//		})
//
//	if err != nil {
//		return nil, err
//	}
//
//	//检查令牌是否有效
//	if !token.Valid {
//		return nil, fmt.Errorf("Invalid tokenString")
//	}
//
//	customerClaims, ok := token.Claims.(*Claims)
//
//	if !ok {
//		return nil, fmt.Errorf("Invalid claims")
//	}
//	return customerClaims, nil
//}

func GetUserNameFromToken(c *gin.Context) (userName string) {
	//在请求头中获取JWTToken，从而解析出用户名
	tokenString := c.GetHeader("Authorization")
	fmt.Println("Headers&Method:", c.Request.Header.Get("Authorization"), c.Request.Method)
	fmt.Println("path is:", c.Request.RequestURI)
	fmt.Println("token is:", tokenString)
	fmt.Println()

	//解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		fmt.Println("Error parsing token:", err)
		return ""
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Error parsing token:.", err)
		return ""
	}

	//获得用户id和用户名
	userName = claims["userName"].(string)
	return userName
}

func GetUserNameFromTokenString(tokenString string) (userName string) {
	fmt.Println("token is:", tokenString)

	//解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		fmt.Println("Error parsing token:", err)
		return ""
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Error parsing token:.", err)
		return ""
	}

	//获得用户id和用户名
	userName = claims["userName"].(string)
	return userName
}

func CreateUserToken(requestUser *DAO.User, c *gin.Context) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   requestUser.ID,
		"userName": requestUser.Uname,
		"exp":      expirationTime,
	})

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(signedToken)
	return signedToken, err

}
