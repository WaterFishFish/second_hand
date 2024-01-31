package logic

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"second_hand/DAO"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dbc := "root:123456@(127.0.0.1:3306)/SecondHand?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dbc)
	if err != nil {
		return err
	}

	DB.AutoMigrate(&DAO.User{})
	DB.Table("tradings").AutoMigrate(&DAO.Trading{})
	DB.AutoMigrate(&DAO.Sell{})
	DB.AutoMigrate(&DAO.Books{})
	DB.AutoMigrate(&DAO.SpCart{})
	DB.AutoMigrate(&DAO.ExaggerateInfo{})
	DB.AutoMigrate(&DAO.IsInCart{})

	// 设置锁等待超时时间为 10 秒
	if err := DB.Exec("SET innodb_lock_wait_timeout = 10").Error; err != nil {
		fmt.Println("Failed to set innodb_lock_wait_timeout:", err)
		return err
	}

	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
