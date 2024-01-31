package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CollectRoutes 控制路由执行哪个函数;调用models的具体操作方法
func CollectRoutes(r *gin.Engine) {

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "reg.html", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "loginpage.html", nil)
	})
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})
	r.GET("/home/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.html", nil)
	})
	r.GET("/home/detail", func(c *gin.Context) {
		c.HTML(http.StatusOK, "detail.html", nil)
	})
	r.GET("/home/cart", func(c *gin.Context) {
		c.HTML(http.StatusOK, "cart.html", nil)
	})
	r.GET("/home/DP", func(c *gin.Context) {
		c.HTML(http.StatusOK, "buy.html", nil)
	})
	r.GET("/home/DirectPurchase", SendRenderInfo)
	r.GET("/home/getList", getCatList)
	r.GET("/home/bookList", GetGoodsList)
	r.GET("/home/bookList/detail", GetDetailInfo)
	//r.GET("/home/DirectPurchase", DirectPurchase)
	r.GET("/home/DirectPurchase/:id", DirectPurchase)

	r.POST("/register", Register)
	r.POST("/login", Login)
	//r.POST("/home/bookList", ReadGoodsList)
	r.POST("/home/ShoppingCarts/add", AddToSC)
	//r.POST("/home/DirectPurchase", DirectPurchase)

	r.DELETE("/home/ShoppingCarts/remove", RemoveFromSC)

	r.PUT("/home/PutBookOnSell/:id", PutBookOnSell)

	//测试区

}

func getCatList(c *gin.Context) {
	// 一个返回给前端的列表数据
	list := [...]string{
		"童书",
		"教育考试",
		"文学小说",
		"人文社科",
		"科技IT",
		"经管励志",
		"艺术",
		"生活",
		"原版",
	}

	// 返回 JSON 数据
	c.JSON(http.StatusOK, gin.H{
		"list": list,
	})
}
