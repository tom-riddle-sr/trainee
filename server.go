package main

import (
	"trainee/fibertrainee3/database/mysql"
	"trainee/fibertrainee3/router"

	_ "github.com/go-sql-driver/mysql"
)

const (
	UserName     string = "root"
	Password     string = "yile1408"
	Addr         string = "127.0.0.1"
	Port         int    = 3306
	Database     string = "test"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

// func main() {
// 	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", UserName, Password, Addr, Port, Database)
// 	// sql.Open 函數來創建一個到資料庫的連接
// 	// return *sql.DB
// 	DB, err := sql.Open("mysql", conn)
// 	if err != nil {
// 			fmt.Println("連線發生錯誤:", err)
// 			return
// 	}
// 	// time.Duration 類型用於表示時間長度，其基礎類型是 int64
// 	DB.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second) // 連接可重用的最大時間,設定0或是小於0就是沒有生命週期
// 	DB.SetMaxOpenConns(MaxOpenConns) // 設定最大連線數,設定0或是小於0就是是無限大
// 	DB.SetMaxIdleConns(MaxIdleConns) // 設定閒置連線量,設定0或是小於0就是無限大

// 	// 確認DB連接是否正常
// 	err = DB.Ping()
// 	if err != nil {
// 			fmt.Println("DB連線異常", err)
// 			return
// 	}

// 	fmt.Println("連線成功")
// }

func main() {
	mysql.New()
	router.Router()
}
