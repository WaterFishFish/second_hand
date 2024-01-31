package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"second_hand/DAO"
	"second_hand/logic"
	"strconv"
)

// AddToSC
// @Summary 添加书籍到购物车
// @Tags 购物
// @Accept json
// @Produce json
// @Param ISBN query string true "书籍的ISBN编号"
// @Param SellerName query string true "出售者用户名"
// @Param userName query string true "当前用户名"
// @Success 2009 {string} string "添加成功"
// @Failure 4009 {string} string "添加失败"
// @Failure 5002 {string} string "数据绑定失败"
// @Router /home/ShoppingCarts/add [post]
func AddToSC(c *gin.Context) {

	token := c.GetHeader("authorization")
	sid := c.PostForm("bookId")
	pSID, _ := strconv.Atoi(sid)
	fmt.Println("*********************psid*************** = :", pSID)
	userName := logic.GetUserNameFromTokenString(token)
	var insertItem DAO.SpCart
	//var sell *DAO.Sell
	var isExist DAO.IsInCart

	sell := logic.GetSellById(int64(pSID))
	insertItem = DAO.SpCart{
		UserName:   userName,
		SellerName: sell.SellerName,
		ISBN:       sell.ISBN,
		BookName:   sell.BookName,
		Price:      sell.Price,
	}

	isExist.SID = pSID
	insertItem.UserName = userName

	//sell, err = logic.GetSellByISP(insertItem.ISBN, insertItem.SellerName, insertItem.Price)
	//if err != nil {
	//	fmt.Println("error in AddToSC:", err)
	//	return
	//}

	_, err := logic.GetCartByID(pSID)

	if err != nil { //购物车中没有这本书的记录,
		isExist.Exist = 0 //存在位置为0
		err = logic.DB.Create(&insertItem).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    4006,
				"message": "加入购物车失败",
				"item":    insertItem,
			})
			return
			// 返回插入的物品信息给客户端
		}
	} else {
		fmt.Println("购物车中已存在该书")
		isExist.Exist = 1
	}
	fmt.Println("更新数据库中")
	err = logic.DB.Debug().Save(&insertItem).Error
	if err != nil {
		fmt.Println(err)
		return
	}

}

// RemoveFromSC
// @Summary 将书籍从购物车清除
// @Tags 购物
// @Accept json
// @Produce json
// @Param ISBN query string true "书籍的ISBN编号"
// @Param SellerName query string true "出售者用户名"
// @Success 2010 {string} string "删除成功"
// @Failure 5001 {string} string "非法图书ID"
// @Failure 4010 {string} string "从购物车删除失败"
// @Router /home/ShoppingCarts/remove [delete]
func RemoveFromSC(c *gin.Context) {
	BookID := c.Param("id")
	var del DAO.SpCart
	id, err := strconv.Atoi(BookID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    5001,
			"message": "Invalid item ID",
		})
		return
	}

	db := logic.DB.Where("id = ?", id).Unscoped().Delete(&del)
	if db.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    4010,
			"message": "删除失败:" + err.Error(),
		})
		return
	}
	if db.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    4009,
			"message": "删除失败：没有对应记录",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    2010,
		"message": "删除成功",
	})
}
