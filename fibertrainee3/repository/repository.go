package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type IRepo interface {
	Create(db *gorm.DB, value interface{}) error
	Query(db *gorm.DB, query string, model interface{}, values ...interface{}) error
	Update(db *gorm.DB, condition string, value interface{}, model interface{}) error
	Delete(db *gorm.DB, condition string, value interface{}, model interface{}) error
}

type Repo struct{}

func NewRepo() IRepo {
	return &Repo{}
}

func (repo *Repo) Create(db *gorm.DB, value interface{}) error {
	fmt.Print(value)
	result := db.Create(value)
	if result.Error != nil {
		return fmt.Errorf("create失敗: %w", result.Error)
	}
	fmt.Println("create成功")
	return nil
}

// db.Select("Name", "Age", "CreatedAt").Create(&user)
// 新增 Name, Age, CreatedAt欄位
// db.Omit("Name", "Age", "CreatedAt").Create(&user)
// 新增除了Name, Age, CreatedAt欄位之外的欄位
// db.CreateInBatches(users, 100)
// 新增多筆資料，每次新增100筆

func (repo *Repo) Query(db *gorm.DB, query string, model interface{}, values ...interface{}) error {
	if err := db.Model(&model).Where(query, values...).First(model).Error; err != nil {
		return err
	}
	return nil
}

//  db.First(&queryAccountData, 1)
// query accountData中primary key為1的資料
//  db.First(&queryAccountData, account="?",value.Account)
// query accountData中primary key為1的資料

// db.First(&user): return第一條記錄。如果表中沒有任何記錄，return一個錯誤。
// db.Take(&user): return第一條記錄。如果表中沒有任何記錄，return一個空的 user 對象。
// db.Last(&user): return最後一條記錄。如果表中沒有任何記錄，return一個錯誤。
// db.Find(&users) return所有記錄
// db.Limit(3).Find(&users) return前三條記錄
// db.Offset(5).Find(&users) return從第五條記錄開始的所有記錄
// db.Model(&User{}).Select("users.name, emails.email")
// .Joins("left join emails on emails.user_id = users.id").Scan(&result{})
// return users表中的name和emails表中的email

func (repo *Repo) Update(db *gorm.DB, condition string, value interface{}, model interface{}) error {
	if result := db.Model(model).Where(condition, value).Updates(model); result.Error != nil {
		return result.Error
	}

	return nil
}

// Model:要更新的table結構
// Where:條件
// Updates:要更新的值

func (repo *Repo) Delete(db *gorm.DB, condition string, value interface{}, model interface{}) error {
	if result := db.Where(condition, value).Delete(&model); result.Error != nil {
		return result.Error
	}
	return nil
}
