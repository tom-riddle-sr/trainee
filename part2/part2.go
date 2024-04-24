package part2

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

func Create(DB *sql.DB){
	createTable := `create table if not exists MemberInfo(
		MemberId varchar(50) primary key not null,
		Level int not null default 1,
		VipLevel int not null default 1); `

		if _,err:=DB.Exec(createTable); err!=nil{
		fmt.Println("創建新table失敗",err)
		return
	}
	fmt.Println("創建新table成功")
}

func Alter(DB *sql.DB){
	if _, err := DB.Exec(`ALTER TABLE MemberInfo
	 ADD CONSTRAINT fk_memberId FOREIGN KEY (MemberId) REFERENCES Member(MemberId)`); err != nil {
			fmt.Println("建立foreign-key失敗", err)
			return
	}
	fmt.Println("建立foreign-key成功")
	}

func Insert(DB *sql.DB) {
    tx, txErr := DB.Begin()
    if txErr != nil {
        fmt.Println("txErr:", txErr)
        return
    }
    if _, execMemberErr := tx.Exec(
        `
        INSERT INTO Member(MemberId,Status)
        VALUES(?,?)
        `, "00000001", "Y",
    ); execMemberErr != nil {
        tx.Rollback()
        fmt.Println("execErr:", execMemberErr)
        return
    }

if _,execMemberInfoErr:=tx.Exec(`
INSERT INTO MemberInfo(MemberId)
VALUES(?)
`,"00000001",
);execMemberInfoErr!=nil{
	fmt.Println("execMemberInfoErr:",execMemberInfoErr)
}

    if commitErr := tx.Commit(); commitErr != nil {
        fmt.Println("commitErr:", commitErr)
        return
    }
		fmt.Println("insert 成功")
}

type Member struct{
	MemberId   string
	Status   	 string
	CreateTime string
	UpdateTime string
}
type MemberInfo struct{
	MemberId string
	Level    string
	VipLevel string
}

type Data struct{
	MemberId string
	Status   	 string
	CreateTime string
	UpdateTime string
	Level    string
	VipLevel string
}

func InsertMultiple(DB *sql.DB) {
	tx, txErr := DB.Begin()
	if txErr != nil {
			fmt.Println("txErr:", txErr)
			return
	}

	memberList:=[]string{}
	memberInfoList:=[]string{}

for i:=3;i<=9;i++{
member:= Member{
	MemberId: "0000000"+strconv.Itoa(i),
	Status:   "Y",
}
memberStr:=fmt.Sprintf("('%s','%s')",member.MemberId,member.Status)
memberList = append(memberList,memberStr)

memberInfo := MemberInfo{
	MemberId: "0000000"+strconv.Itoa(i),
}

memberInfoStr:=fmt.Sprintf("('%s')",memberInfo.MemberId)
memberInfoList = append(memberInfoList,memberInfoStr)
}

memberQuery := fmt.Sprintf(`INSERT INTO Member (MemberId, Status)
VALUES %s`, strings.Join(memberList, ", "))
if _, execMemberErr := tx.Exec(memberQuery); execMemberErr != nil {
	tx.Rollback()
	fmt.Println("execErr:", execMemberErr)
	return
}

memberInfoQuery := fmt.Sprintf(`INSERT INTO MemberInfo (MemberId)
VALUES %s`, strings.Join(memberInfoList, ", "))
if _,execMemberInfoErr:=tx.Exec(memberInfoQuery);execMemberInfoErr!=nil{
	tx.Rollback()
	fmt.Println("execMemberInfoErr:",execMemberInfoErr)
	return
	}

if commitErr := tx.Commit(); commitErr != nil {
	fmt.Println("commitErr:", commitErr)
	return
}
fmt.Println(" insert 成功")
}

func Query(DB *sql.DB) {
	// tx, txErr := DB.Begin()
	// if txErr != nil {
	// 	fmt.Println("txErr:", txErr)
	// 	return
	// }

	// member := Member{
	// 	MemberId: "00000001",
	// }

	// row := tx.QueryRow("SELECT * FROM Member WHERE MemberId=?", member.MemberId)
	//  queryMember:=Member{}
	// if memberQueryErr := row.Scan(&queryMember.MemberId, &queryMember.Status,&queryMember.CreateTime,&queryMember.UpdateTime); memberQueryErr != nil {
	// 	fmt.Println("memberQueryErr:", memberQueryErr)
	// 	tx.Rollback()
	// 	return
	// }

	// memberInfo:=MemberInfo{
	// 	MemberId: "00000001",
	// }

	// infoRow:=tx.QueryRow("SELECT * FROM MemberInfo WHERE MemberId=?", memberInfo.MemberId)
	// queryMemberMemberInfo:=MemberInfo{}
	// if memberQueryErr := infoRow.Scan(&queryMemberMemberInfo.MemberId, &queryMemberMemberInfo.Level,&queryMemberMemberInfo.VipLevel); memberQueryErr != nil {
	// 	fmt.Println("memberQueryErr:", memberQueryErr)
	// 	tx.Rollback()
	// 	return
	// }
	// if commitErr := tx.Commit(); commitErr != nil {
	// 	fmt.Println("commitErr:", commitErr)
	// 	return
	// }
	// fmt.Printf(`Query成功:
	// MemberId:%s,Status:%s,CreateTime:%s,UpdateTime:%s
	// Level:%s,VipLevel:%s` ,queryMember.MemberId,queryMember.Status,queryMember.CreateTime,queryMember.UpdateTime,
	// queryMemberMemberInfo.Level,queryMemberMemberInfo.VipLevel)

row:=DB.QueryRow(`
SELECT Member.MemberId,Member.Status,Member.CreateTime,Member.UpdateTime,
MemberInfo.Level,MemberInfo.VipLevel
FROM Member
INNER JOIN MemberInfo ON Member.MemberId = MemberInfo.MemberId
WHERE Member.MemberId = ?;`,00000001)


var queryMemberData=Data{}
if err := row.Scan(
	&queryMemberData.MemberId,
	&queryMemberData.Status,
	&queryMemberData.CreateTime,
	&queryMemberData.UpdateTime,
	&queryMemberData.Level,
	&queryMemberData.VipLevel); err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Print(queryMemberData)

}