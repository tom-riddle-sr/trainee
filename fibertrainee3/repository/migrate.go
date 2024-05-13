package repository

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	// GORM 提供的一個基本模型，它包含了幾個常用的欄位
	// 如 ID、CreatedAt、UpdatedAt、DeletedAt
	Code  string
	Price uint
}

// Migrate 遷移 幫助創建Product結構的表
// 如果這個資料表已經存在，AutoMigrate 會修改它，以使其結構與 Product 結構體相匹配
func Migrate(db *gorm.DB, table *Product) error {
	if err := db.AutoMigrate(table); err != nil {
		return err
	}
	return nil
}
