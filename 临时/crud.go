package logic

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
	"second_hand/DAO"
	"sync"
)

var (
	SellMutex   sync.Mutex
	BookMutex   sync.Mutex
	UserMutex   sync.Mutex
	SpCartMutex sync.Mutex
)

//var Mutex sync.Mutex
//var Mutex sync.Mutex

type Claims struct {
	UserName int64 `json:"uname"`
	jwt.RegisteredClaims
}

func GetUserBalance(userName string) (float64, error) {
	var userInfo DAO.User
	err := DB.Where("userName = ?", userName).First(&userInfo).Error
	return userInfo.Balance, err
}

// GetUserInfoByName 根据用户名获取用户信息（不在事务中使用）
func GetUserInfoByName(userName string) DAO.User {
	var user DAO.User
	DB.Where("uname=?", userName).First(&user)
	return user
}

// GetUserInfoWithLock 获取当前用户信息（在事务中使用）
func GetUserInfoWithLock(c *gin.Context, ts *gorm.DB) *DAO.User {
	return GetUserInfoByNameWithLock(GetUserNameFromToken(c), ts)
}

// GetUserInfo 获取当前用户信息（不在事务中使用）
func GetUserInfo(c *gin.Context) DAO.User {
	return GetUserInfoByName(GetUserNameFromToken(c))
}

// GetUserInfoByNameWithLock 根据用户名获得用户信息(在事务中使用)
func GetUserInfoByNameWithLock(userName string, ts *gorm.DB) *DAO.User {
	var user = new(DAO.User)
	//if err := ts.Set("gorm:query_option", "FOR UPDATE").
	//	Where("uname=?", userName).First(&user).Error; err != nil {
	//	HandleErrInTransaction(ts, err)
	//}
	UserMutex.Lock()
	if err := ts.Where("uname=?", userName).First(&user).Error; err != nil {
		HandleErrInTransaction(ts, err)
	}
	return user
}

// GetBookInfoByIsbnWithLock 根据书籍的ISBN编码获取书籍信息
func GetBookInfoByIsbnWithLock(isbn string, ts *gorm.DB) *DAO.Books {
	book := new(DAO.Books)
	//if err := ts.Set("gorm:query_option", "FOR UPDATE").
	//	Where("isbn = ?", isbn).First(book).Error; err != nil { //查找书籍，并将结果存在find中
	//	HandleErrInTransaction(ts, err)
	//}
	BookMutex.Lock()
	if err := ts.Where("isbn = ?", isbn).First(book).Error; err != nil { //查找书籍，并将结果存在find中
		HandleErrInTransaction(ts, err)
	}

	return book

}

// GetSellByIdWithLock 根据ID号寻找对应的Sell表表项
func GetSellByIdWithLock(SId int64, ts *gorm.DB) *DAO.Sell {
	sell := new(DAO.Sell)
	//if err := ts.Set("gorm:query_option", "FOR UPDATE").Debug().
	//	Where("id = ?", SId).First(sell).Error; err != nil { //查找买家选中的书籍，并将结果存在find中
	//	HandleErrInTransaction(ts, err)
	//}
	SellMutex.Lock()
	if err := ts.Where("id = ?", SId).First(sell).Error; err != nil { //查找买家选中的书籍，并将结果存在find中
		HandleErrInTransaction(ts, err)
	}
	return sell
}

// GetSellByISPWithLock 根据ISBN，SellerName，price查找唯一对用的出售表项。I：ISBN，S：sellerName，P：price
func GetSellByISPWithLock(isbn string, sellerName string, price float64, ts *gorm.DB) (*DAO.Sell, error) {
	sell := new(DAO.Sell)
	var err error
	//tmp := ts.Set("gorm:query_option", "FOR UPDATE").
	//	Where("isbn = ? AND Seller_Name = ? AND price = ?", isbn, sellerName, price).First(sell)
	//if tmp.Error != nil { //查找买家选中的书籍，并将结果存在find中
	//	HandleErrInTransaction(ts, tmp.Error)
	//}
	SellMutex.Lock()
	tmp := ts.Where("isbn = ? AND Seller_Name = ? AND price = ?", isbn, sellerName, price).First(sell)
	if tmp.Error != nil { //查找买家选中的书籍，并将结果存在find中
		HandleErrInTransaction(ts, tmp.Error)
	}
	if tmp.RowsAffected == 0 { //未查找到对应记录
		err = errors.New("出售列表中无对应记录")
	}
	return sell, err
}

func GetInfoByUserName(userName string) *DAO.ExaggerateInfo {
	var info = new(DAO.ExaggerateInfo)
	DB.Where("user_name = ?", userName).First(&info)
	return info
}

// UpdateSellsLeft 方法更新sell表中的left字段
func UpdateSellsLeft(newLeft int64, s *DAO.Sell, ts *gorm.DB) {
	if newLeft == 0 { //用户出售的某种图书剩余库存为0，在数据库中删除这条记录
		err := DB.Where("id = ?", s.ID).Unscoped().Delete(&s).Error
		if err != nil {
			HandleErrInTransaction(ts, err)
		}
	} else { //不为0则正常更新left字段即可
		err := DB.Model(s).Debug().Update("left", newLeft).Error
		if err != nil {
			HandleErrInTransaction(ts, err)
		}
	}
	SellMutex.Unlock()
}

// UpdateUserBalance 更新用户余额
func UpdateUserBalance(newBalance float64, user *DAO.User, ts *gorm.DB) {
	err1 := DB.Model(&user).Update("balance", newBalance).Error
	if err1 != nil {
		HandleErrInTransaction(ts, err1)
	}
	UserMutex.Unlock()
}

// UpdateBooksWhenSell 在交易时更新books表
func UpdateBooksWhenSell(book *DAO.Books, c *gin.Context) error {
	if book.Left == 0 {
		//c.JSON(http.StatusOK, gin.H{
		//	"code":    4014,
		//	"message": "无对应书籍",
		//})
		return errors.New("无对应书籍")
	}

	book.Count++ //销量+1
	book.Left--  //余量-1

	//更新books表中对应表项的count和left字段
	DB.Model(&book).Debug().Updates(map[string]interface{}{"count": book.Count, "left": book.Left})
	BookMutex.Unlock()
	return nil
}

// UpdateBooksByLeft 指定left更新books的left字段
func UpdateBooksByLeft(left int64, book *DAO.Books) {
	book.Left = left
	//更新books表中对应表项的left字段
	DB.Model(&book).Updates(map[string]interface{}{"left": book.Left})
}

// InsertIntoBooks 插入记录到books表中
func InsertIntoBooks(book *DAO.Books) error {
	err := DB.Create(book).Error
	return err
}

// InsertIntoSell 插入记录到books表中
func InsertIntoSell(sell *DAO.Sell) error {
	err := DB.Create(sell).Error
	return err
}

func SaveInfos(info *DAO.ExaggerateInfo) error {
	err := DB.Debug().Where("username = ?", info.UserName).Save(info).Error
	return err
}
func InsertTradingRecord(trade *DAO.Trading, ts *gorm.DB) error {
	// 插入交易记录到数据库
	if err := ts.Create(trade).Error; err != nil {
		// 如果创建记录时发生错误，返回错误信息

		return err
	}
	// 没有错误，返回nil表示成功
	return nil
}
