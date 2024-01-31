// Package DAO:Database access Object:数据操作对象
package DAO

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// SpCart :购物车表单
type SpCart struct {
	gorm.Model
	UserName   string  `json:"userName" form:"userName"`
	SellerName string  `json:"sellerName" form:"sellerName" `
	ISBN       string  `json:"isbn" form:"isbn" binding:"required"`
	BookName   string  `json:"bookName" form:"bookName"`
	Price      float64 `json:"price" form:"price"`
}

func (s *SpCart) TableName() string {
	return "SpCart"
}

func (ei *ExaggerateInfo) TableName() string {
	return "ExaggerateInfo"
}

type User struct {
	gorm.Model
	Uname     string  `gorm:"varchar(20);primary_key" column:"user_name" json:"uname" form:"uname"`
	Upassword string  `json:"upassword" gorm:"size:255;not null" form:"upassword"`
	Nickname  string  `json:"nickName" gorm:"varchar(50);not null" form:"nickName"`
	Balance   float64 `json:"balance" gorm:"double;default:1000;comment:'余额' "  form:"balance"`
}

// Books 保存所有正在出售的图书，比如有不同的用户出售同一本书
type Books struct {
	ID       int64  `json:"id" form:"id" gorm:"primary_key"`
	ISBN     string `json:"isbn" gorm:"primary_key" form:"isbn"`
	BookName string `json:"bookName"`
	Left     int64  `json:"left" form:"left"` //剩余图书
	Author   string `json:"author" form:"author"`
	Count    int64  `json:"count" form:"count"` //总销量，排序时用
	Class    string `json:"class" form:"class"`
}

type Trading struct {
	SpCart
	TradeTime   string `json:"tradeTime"`
	Destination string `json:"destination"`
}

// Sell 保存某个用户出售的图书
type Sell struct {
	gorm.Model
	ISBN        string  `json:"isbn" gorm:"primary_key"`
	Description string  `json:"description"`
	SellerName  string  `json:"sellerName"  gorm:"primary_key"`
	Author      string  `json:"author" form:"author"`
	BookName    string  `json:"bookName"`
	Left        int64   `json:"left"`
	Price       float64 `json:"price"  gorm:"primary_key"`
	Src         string  `json:"src"`
	Class       string  `json:"class" form:"class"`
}

type ExaggerateInfo struct {
	UserName string    `json:"userName" gorm:"primary_key" column:"user_name"`
	Current  int       `json:"current"`
	PageSize int       `json:"pagesize"`
	Search   string    `json:"search"`
	Total    int       `json:"total"`
	Class    string    `json:"class"`
	Time     time.Time `json:"time"`
	SortType string    `json:"sortType"`
	AD       string    `json:"AD"`
}

type IsInCart struct {
	SID   int `json:"sid" gorm:"primary_key"`
	Exist int `json:"exist"`
}
