package query

import (
	"database/sql"
	"fmt"
)

type Member struct {
	MemberId   string
	Status     string
	CreateTime string
	UpdateTime string
}

func Query(DB *sql.DB) {
	var mList Member

	// 單筆資料
// row := DB.QueryRow("select MemberId,Status,CreateTime,UpdateTime from Member where MemberId=?", 00000001)
// 	//Scan對應的欄位與select語法的欄位順序一致
// if err := row.Scan(&mList.memberId, &mList.status,&mList.createTime,&mList.updateTime );err != nil {
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

var memberList =[]Member{}

// 一筆一筆讀取
for rows.Next() {
	err = rows.Scan(&mList.MemberId, &mList.Status, &mList.UpdateTime, &mList.CreateTime)
	if err != nil {
		fmt.Printf("Scan failed,err:%v\n", err)
		return
	}
	member := Member{
    MemberId:   mList.MemberId,
    Status:     mList.Status,
    CreateTime: mList.CreateTime,
    UpdateTime: mList.UpdateTime,
}
memberList = append(memberList, member)
}

for _, customer := range memberList {
	fmt.Printf("customer:%+v\n", customer)
}
}
