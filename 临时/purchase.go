package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"second_hand/DAO"
	"second_hand/logic"
	"time"
)

// PurchaseCart
// 从购物车中购买二手书
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
func PurchaseCart(c *gin.Context) {
	var carts []DAO.SpCart
	//获取当前用户的用户信息
	curUserInfo := logic.GetUserInfo(c)

	//获取当前用户的余额
	balance := curUserInfo.Balance

	err := c.ShouldBindJSON(&carts)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    5001,
			"message": "数据绑定失败",
		})
		return
	}
	if len(carts) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    4013,
			"message": "购物车为空",
		})
		return
	}

	var total float64 //total为在购物车中选中图书总价

	// 加锁
	// 确保在函数结束时释放锁
	//开启事务
	ts := logic.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			ts.Rollback()
		}
	}()

	for _, v := range carts {
		total += v.Price
	}

	if balance < total {
		c.JSON(http.StatusOK, gin.H{
			"code":    4008,
			"message": "用户余额不足",
		})
		return
	} else {
		for i, v := range carts {
			var book = new(DAO.Books)
			var updateSell = new(DAO.Sell)
			//添加行锁,更新book表
			book = logic.GetBookInfoByIsbnWithLock(carts[i].ISBN, ts)
			err = logic.UpdateBooksWhenSell(book, c)
			if err != nil {
				logic.HandleErrInTransaction(ts, err)
				return
			}

			//利用SpCart中的ISBN,sellerName,price字段查找对应的出售表项
			updateSell, _ = logic.GetSellByISPWithLock(carts[i].ISBN, carts[i].SellerName, carts[i].Price, ts)
			updateSell.Left--
			logic.UpdateSellsLeft(updateSell.Left, updateSell, ts)

			//添加行锁，更新卖家用户信息(增加余额)
			var sellerInfo *DAO.User
			sellerInfo = logic.GetUserInfoByNameWithLock(carts[i].SellerName, ts)
			sellerBalance := sellerInfo.Balance
			sellerBalance += v.Price
			fmt.Println(sellerBalance)
			logic.UpdateUserBalance(sellerBalance, sellerInfo, ts)
		}

		balance -= total
		fmt.Println(balance)
		//更新买家账户余额信息
		logic.UserMutex.Lock()
		logic.UpdateUserBalance(balance, &curUserInfo, ts)
	}

	ts.Commit()
}

func Trading(c *gin.Context) {
	var trade DAO.Trading
	if err := c.ShouldBind(&trade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    5001,
			"message": "5001：数据绑定失败",
		})
		return
	}

}

func DirectPurchase(c *gin.Context) {
	var buy DAO.Sell
	var find = new(DAO.Sell)
	var sellerInfo = new(DAO.User)
	var buyerInfo = new(DAO.User)

	//开启事务
	ts := logic.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			ts.Rollback()
		}
	}()

	//获得买家用户（即当前用户）信息
	buyerInfo = logic.GetUserInfoByNameWithLock(logic.GetUserNameFromToken(c), ts)

	err := c.ShouldBindJSON(&buy)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    5001,
			"message": "数据绑定失败",
		})
		return
	}

	find = logic.GetSellByIdWithLock(buy.ID, ts)
	if find.Left == 0 { //没有剩余图书了
		c.JSON(http.StatusOK, gin.H{
			"code":    4014,
			"message": "无对应书籍",
		})
		logic.HandleErrInTransaction(ts, errors.New("无对应书籍"))
		return
	}
	if find.Price > buyerInfo.Balance { //用户余额不足
		c.JSON(http.StatusOK, gin.H{
			"code":    4008,
			"message": "用户余额不足",
		})
		return
	}
	buyerInfo.Balance -= find.Price
	logic.UpdateUserBalance(buyerInfo.Balance, buyerInfo, ts)
	fmt.Println("UpdateUserBalance success")

	//获得卖家信息,加锁，以便更新卖家用户余额
	sellerInfo = logic.GetUserInfoByNameWithLock(buy.SellerName, ts)

	fmt.Printf("Book: %s -- %.2f\n", buy.BookName, buy.Price)

	//在books表查找相关书籍信息（并添加行锁）并更新信息
	var updateBooks = new(DAO.Books)
	updateBooks = logic.GetBookInfoByIsbnWithLock(buy.ISBN, ts)
	fmt.Println("GetBookInfoByIsbnWithLock success")
	err = logic.UpdateBooksWhenSell(updateBooks, c)
	fmt.Println("UpdateBooksWhenSell success")
	if err != nil {
		logic.HandleErrInTransaction(ts, err)
	}

	//更新sell表中的书籍记录
	find.Left--
	logic.UpdateSellsLeft(find.Left, find, ts)
	fmt.Println("UpdateSellsLeftsuccess")
	//更新买家和卖家的balance字段

	sellerInfo.Balance += find.Price

	logic.UpdateUserBalance(sellerInfo.Balance, sellerInfo, ts)
	fmt.Println("UpdateUserBalance success")

	//更新交易记录
	tradeRecord := DAO.Trading{
		SpCart: DAO.SpCart{
			UserName:   buyerInfo.Uname,
			SellerName: find.SellerName,
			ISBN:       find.ISBN,
			BookName:   find.BookName,
			Price:      find.Price,
		},
		TradeTime: time.Now().Format("2006-01-02 15:04:05"), // 使用当前时间作为交易时间
	}

	// 插入交易记录到数据库
	err = logic.InsertTradingRecord(&tradeRecord, ts)
	if err != nil {
		logic.HandleErrInTransaction(ts, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    5002,
			"message": "交易记录创建失败",
		})
		return
	}
	ts.Commit()
}

func PutBookOnSell(c *gin.Context) {
	var sell DAO.Sell
	if err := c.ShouldBindJSON(&sell); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    5001,
			"message": "5001：数据绑定失败",
		})
		return
	}
	fmt.Println(sell)
	ts := logic.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			ts.Rollback()
			err, _ := r.(error)
			fmt.Println("发生错误:", err)
		}
	}()

	//检查该ISBN书籍是否已存在与books表中
	var book = new(DAO.Books)
	result := logic.DB.Where("isbn=?", sell.ISBN).First(book)
	if result.RowsAffected == 0 { //没有在books表中找到对应书籍，说明这是新书，应该添加进books表中
		newBook := DAO.Books{
			ISBN:     sell.ISBN,
			BookName: sell.BookName,
			Left:     1,
			Author:   sell.Author,
			Count:    0,
			Class:    sell.Class,
		}
		err := logic.InsertIntoBooks(&newBook)
		if err != nil {
			logic.HandleErrInTransaction(ts, err)
		}
	}
	//书籍已经在books表中
	book.Left++
	logic.UpdateBooksByLeft(book.Left, book) //更新book记录库存

	var find = new(DAO.Sell)
	find, err := logic.GetSellByISPWithLock(sell.ISBN, sell.SellerName, sell.Price, ts)
	if err != nil { //在sell中未查找到对应记录
		sell.Left = 1
		err1 := logic.InsertIntoSell(&sell) //插入sell记录

		fmt.Println(sell)
		if err1 != nil {
			logic.HandleErrInTransaction(ts, err1)
		}
	} else { //在sell中查找到了对应记录，更新left字段
		find.Left++
		logic.UpdateSellsLeft(find.Left, find, ts)
	}

	ts.Commit()
}
