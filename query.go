package main

import (
	"database/sql"
	"fmt"
)

func Query(DB *sql.DB) {
	var memberId string
	var status string
	var createTime string
	var updateTime string



// 單筆資料
// row := DB.QueryRow("select MemberId,Status,CreateTime,UpdateTime from Member where MemberId=?", 00000001)
// 	//Scan對應的欄位與select語法的欄位順序一致
// if err := row.Scan(&memberId, &status,&createTime,&updateTime );err != nil {
// 	fmt.Printf("scan failed, err:%v\n", err)
// 	return
// }
// fmt.Printf("memberId:%s,status:%s,createTime:%s,updateTime:%s\n", memberId, status, createTime, updateTime)

// 多筆資料
rows, err := DB.Query("select MemberId,Status,CreateTime,UpdateTime from Member")
	//記得要close掉連線，不然會一直卡連線
defer func() {
		rows.Close()
}()
if err != nil {
	fmt.Printf("Query failed,err:%v\n", err)
	return
}

var customers []map[string]string

// 一筆一筆讀取
for rows.Next() {
	err = rows.Scan(&memberId, &status, &updateTime, &createTime)
	if err != nil {
		fmt.Printf("Scan failed,err:%v\n", err)
		return
	}
	customer := map[string]string{
		"memberId":   memberId,
		"status":     status,
		"createTime": createTime,
		"updateTime": updateTime,
}
customers = append(customers, customer)
}

for _, customer := range customers {
	fmt.Printf("customer:%+v\n", customer)
}
}
