package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"second_hand/DAO"
)

func FindBooks(c *gin.Context) *[]DAO.Sell {
	var book DAO.Books
	err := c.ShouldBind(&book)
	if err != nil {
		fmt.Println("Bind failed: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    5001,
			"message": "数据绑定失败",
			"item":    book,
		})
		return nil
	}
	isbn := book.ISBN

	var bookOnSell *[]DAO.Sell

	//在出售清单中搜索书籍
	DB.Where("isbn = ?", isbn).Find(bookOnSell)
	if len(*bookOnSell) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    4007,
			"message": "未找到正在出售的商品",
		})
		return nil
	}
	bookOnSell = SortByPrice(bookOnSell)
	return bookOnSell
}
