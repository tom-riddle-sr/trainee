package main

import (
	"database/sql"
	"fmt"
	"time"
)
func Add(DB *sql.DB) {
	// DB.Exec()用於執行 SQL 查詢
	// 新增一筆
	result, err := DB.Exec("insert INTO Member(MemberId, Status, CreateTime, UpdateTime)values(?,?,?,?)", "00000008", "Y",time.Now().Format("2006-01-02 15:04:05"),time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
			fmt.Printf("Insert data failed,err:%v", err)
			return
	}
			//LastInsertId()return 資料庫AUTO_INCREMENT的整數
			lastInsertID, err := result.LastInsertId()
			if err != nil {
					fmt.Printf("Get insert id failed,err:%v", err)
					return
			}
			fmt.Println("Insert data id:", lastInsertID)

			//RowsAffected() 影響的資料筆數
			rowsaffected, err := result.RowsAffected()
			if err != nil {
					fmt.Printf("Get RowsAffected failed,err:%v", err)
					return
			}
			fmt.Println("Affected rows:", rowsaffected)
	}

// 新增多筆
// func Add(DB *sql.DB) {
// 	var memberList =[]struct{
// 		 MemberId string
// 		 Status string
// 		  CreateTime string
// 			 UpdateTime string }{
// 				 {"000000011","Y",time.Now().Format("2006-01-02 15:04:05"),time.Now().Format("2006-01-02 15:04:05")},
// 				 {"000000012","Y",time.Now().Format("2006-01-02 15:04:05"),time.Now().Format("2006-01-02 15:04:05")},
// 				 {"00000013","Y",time.Now().Format("2006-01-02 15:04:05"),time.Now().Format("2006-01-02 15:04:05")},
// 				 {"00000014","Y",time.Now().Format("2006-01-02 15:04:05"),time.Now().Format("2006-01-02 15:04:05")},
// 				}

// 	for _, member := range memberList {
// 			result, err := DB.Exec("INSERT INTO Member(MemberId, Status, CreateTime, UpdateTime) VALUES (?, ?, ?, ?)",
// 					member.MemberId, member.Status, member.CreateTime, member.UpdateTime)
// 			if err != nil {
// 					fmt.Printf("Insert data failed,err:%v", err)
// 					continue
// 			}

// 			//LastInsertId()return 資料庫AUTO_INCREMENT的整數
// 			lastInsertID, err := result.LastInsertId()
// 			if err != nil {
// 					fmt.Printf("Get insert id failed,err:%v", err)
// 					continue
// 			}
// 			fmt.Println("Insert data id:", lastInsertID)

// 			//RowsAffected() 影響的資料筆數
// 			rowsaffected, err := result.RowsAffected()
// 			if err != nil {
// 					fmt.Printf("Get RowsAffected failed,err:%v", err)
// 					continue
// 			}
// 			fmt.Println("Affected rows:", rowsaffected)
// 	}
// }