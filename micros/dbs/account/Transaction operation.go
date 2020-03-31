package account

import "github.com/jinzhu/gorm"

func TransactionOperation(dbs ...*gorm.DB) bool {

	tx := DB.Begin()

	for _,db := range dbs{
		tx = db
		if tx.Error != nil {
			tx.Rollback()
			return false
		}
	}

	tx.Commit()
	return true
}