package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"second_hand/DAO"
	"second_hand/logic"
)

var jwtSecret = []byte("secretString")

// Register
// 注册账号接口
// @Summary 注册新用户
// @Description 注册一个新用户账号
// @Tags 登录
// @Accept json
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param confirmPassword formData string true "确认密码"
// @Param email formData string true "电子邮件"
// @Success 200 {string} string "注册成功"
// @Failure 400 {string} string "该用户名已被使用"
// @Failure 422 {string} string "用户名不能为空"
// @Failure 4003 {string} string "重复密码与第一次输入的密码不一致，请重新输入"
// @Failure 500 {string} string "密码加密错误"
// @Router /register [post]
func Register(c *gin.Context) {

	name := c.PostForm("uname")
	password := c.PostForm("upassword")
	nickName := c.PostForm("nickName")
	repeatPassword := c.PostForm("rpassword")

	if len(name) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    4003,
			"message": "用户名不能为空",
		})
		return
	}

	var user DAO.User
	logic.DB.Where("uname=?", name).First(&user)
	if user.Uname == name {
		c.JSON(http.StatusOK, gin.H{
			"code":    4004,
			"message": "用户名已存在",
		})
		return
	}

	if repeatPassword != password {
		c.JSON(http.StatusOK, gin.H{
			"code":    4005,
			"message": "重复密码与第一次输入的密码不一致，请重新输入",
		})
		return
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "密码加密错误",
		})
	}

	newUser := DAO.User{
		Uname:     name,
		Upassword: string(hasedPassword),
		Nickname:  nickName,
		Balance:   1000,
	}
	logic.DB.Create(&newUser)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "注册成功",
	})

}

// Login
// 登录账号接口
// @Tags 登录
// @Accept json
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param email formData string false "电子邮件"
// @Success 2000 {string} string "登录成功"
// @Failure 4001 {string} string "该用户不存在"
// @Failure 4002 {string} string "密码错误"
// @Router /login [post]
func Login(c *gin.Context) {
	var requestUser DAO.User
	var queryUser DAO.User

	err := c.ShouldBind(&requestUser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    5001,
			"message": "数据绑定失败",
		})
		return
	}
	name := requestUser.Uname
	password := requestUser.Upassword

	db := logic.DB
	db.Where("uname=?", name).First(&queryUser)
	if queryUser.Uname != name {
		c.JSON(http.StatusOK, gin.H{
			"code":    4001,
			"message": "该用户不存在",
		})
		fmt.Println("该用户不存在")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(queryUser.Upassword), []byte(password)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    4002,
			"message": "密码错误",
		})
		fmt.Println("密码错误")

		return
	}

	//生成用户Token,标识会话
	signedToken, err := logic.CreateUserToken(&requestUser, c)
	if err != nil { //生成用户Token失败，直接返回
		fmt.Println("error:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    2000,
		"message": "登录成功",
		"token":   signedToken,
	})
}
