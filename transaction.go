package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)




func transaction(DB *sql.DB) {
// var i int = 1
//Itoa = "integer to ASCII"
// fmt.Println("00000"+ strconv.Itoa(i))

	tx, txErr := DB.Begin()
	if txErr != nil {
		fmt.Println("txErr:", txErr)
		return
	}
	for i := 15; i <= 50; i++ {
		res, execErr := tx.Exec("insert INTO Member(MemberId,Status,CreateTime,UpdateTime) values(?,?,?,?)", "000000"+strconv.Itoa(i), "Y",time.Now().Format("2006-01-02 15:04:05"),time.Now().Format("2006-01-02 15:04:05"))
		if execErr != nil {
			tx.Rollback()
			fmt.Println("execErr:", execErr)
			return
		}
		rowsAffected, rowsErr := res.RowsAffected()
		if rowsErr != nil || rowsAffected != 1 {
			tx.Rollback()
			fmt.Println("execErr:", execErr)
			return
		}
	}
	tx.Commit()
	fmt.Println("transaction success")
}