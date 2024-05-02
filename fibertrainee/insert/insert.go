package insert

import (
	"database/sql"
	"fmt"
	"trainee/fibertrainee/fiber"
)

const (
	UserName = "root"
	Password = "yile1408"
	Addr     = "127.0.0.1"
	Port     = 3003
	Database = "fibertrainee"
)

const (
	dataQuantity   = 10
	stringQuanity  = 6
	numberQuantity = 6
)

func Insert() {
	DB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", UserName, Password, Addr, Port, Database))
	if err != nil {
		fmt.Println("連接到mysql 失敗:", err)
		return
	}
	fmt.Println("連接到mysql 成功")

	// sqlValues := []string{}

	// for i := 0; i < 10; i++ {
	// 		sqlValue := fmt.Sprintf(`('%s','%s')`, common.HashString(common.GetRand(stringQuanity, numberQuantity)), common.HashString(common.GetRand(stringQuanity, numberQuantity)))
	// 		sqlValues = append(sqlValues, sqlValue)
	// }
	// sqlStatement := fmt.Sprintf(`INSERT INTO accountdata (account, password) VALUES %s`, strings.Join(sqlValues, ", "))

	// step2.🍤做一些假資料🍤
	// 	sqlValues:=fmt.Sprintf(`('%s','%s')`,common.HashString("test"),common.HashString("test"))
	// 	sqlStatement := fmt.Sprintf(`INSERT INTO accountdata (account, password) VALUES %s`, sqlValues)

	// if _,err:=DB.Exec(sqlStatement);err!=nil{
	// 		fmt.Println("插入資料失敗",err)
	// 		return
	// }
	// fmt.Println("插入資料成功")

	fiber.Fiber(DB)
}
