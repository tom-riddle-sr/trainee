package main

import (
	"database/sql"
	"fmt"
)

// func Update(DB *sql.DB) {
// 	result, err := DB.Exec("UPDATE Member set Status=? where MemberId=?", "D", 00000001)
// 	if err != nil {
// 		fmt.Printf("UPDATE failed,err:%v\n", err)
// 		return
// 	}
// 	fmt.Println("update data successd:", result)

// 	rowsaffected, err := result.RowsAffected()
// 	if err != nil {
// 		fmt.Printf("UPDATE RowsAffected failed,err:%v\n", err)
// 		return
// 	}
// 	fmt.Println("UPDATE Affected rows:", rowsaffected)
// }

// 原子性:整個操作應該被視為一個不可分割的單元
// 在資料庫中執行一個操作（例如，插入、更新或刪除資料）時，你可能希望確保所有的變更都成功，或者如果有任何一個變更失敗，則所有的變更都不會發生
// 關心原子性或有多個操作,用tx
func Update(DB *sql.DB) {
	// tx:來操作該事務的 *sql.Tx 對象
	// 事務是一種保證資料庫操作原子性的機制
	tx, err := DB.Begin()
	if err != nil {
			fmt.Printf("Begin transaction failed,err:%v\n", err)
			return
	}

	result, err := tx.Exec("UPDATE Member set Status=? where MemberId=?", "D", 00000001)
	if err != nil {
			fmt.Printf("UPDATE failed,err:%v\n", err)
			tx.Rollback()
			return
	}
	fmt.Println("update data successd:", result)

	rowsaffected, err := result.RowsAffected()
	if err != nil {
			fmt.Printf("UPDATE RowsAffected failed,err:%v\n", err)
			tx.Rollback()
			return
	}
	fmt.Println("UPDATE Affected rows:", rowsaffected)

	// 有操作都成功，則可以提交（commit）事務,所有操作的結果將一起反映在資料庫中
	// 事務中的任何操作失敗，則可以回滾（rollback）事務,這將撤銷事務中已執行的所有操作

	err = tx.Commit()
	if err != nil {
			fmt.Printf("Commit transaction failed,err:%v\n", err)
			tx.Rollback()
			return
	}
}
