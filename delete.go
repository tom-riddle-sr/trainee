package main

import (
	"database/sql"
	"fmt"
)

func Delete(DB *sql.DB) {
	result, err := DB.Exec("delete from Member where MemberId=?", 00000002)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	fmt.Println("delete data successd:", result)

	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("delete RowsAffected failed,err:%v\n", err)
		return
	}
	fmt.Println("delete Affected rows:", rowsaffected)
}