package logic

import "github.com/jinzhu/gorm"

func HandleErrInTransaction(tx *gorm.DB, err error) {
	tx.Rollback()
	panic(err)
}
