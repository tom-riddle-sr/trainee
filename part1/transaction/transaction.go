package transaction

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"trainee/part1/query"
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
	for i := 51; i <= 60; i++ {
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



//insert 多筆資料 一次完成而非一筆一筆insert
func transaction1(DB *sql.DB) {

	members := []string{}

for i := 71; i <= 80; i++ {
    member := query.Member{
        MemberId:   "000000" + strconv.Itoa(i),
        Status:     "Y",
        CreateTime: time.Now().Format("2006-01-02 15:04:05"),
        UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
    }
    memberStr := fmt.Sprintf("('%s', '%s', '%s', '%s')", member.MemberId, member.Status, member.CreateTime, member.UpdateTime)
    members = append(members, memberStr)
}

query := fmt.Sprintf("INSERT INTO Member (MemberId, Status, CreateTime, UpdateTime) VALUES %s", strings.Join(members, ","))
	res, err := DB.Exec(query)
	if err != nil {
			fmt.Println("Exec error:", err)
			return
	}
fmt.Println("transaction1 success",res)



}