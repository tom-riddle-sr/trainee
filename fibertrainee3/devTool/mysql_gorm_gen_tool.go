package main

import (
	"fmt"
	"trainee/fibertrainee3/config"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Gorm版

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "../model/entity/mysql/fibertrainee",                               // outputㄉ位置
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // 生成模式ㄉ設置
		ModelPkgPath: "fibertrainee",                                                     // 生成ㄉ包
	})

	db, _ := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		config.UserName, config.Password, config.Addr, config.Port, config.Database)))
	g.UseDB(db) // 設置 GORM Gen 生成器使用的 DB 實例

	// 生成model
	g.GenerateModel("accountData")

	// //忽略address欄位,指定id欄位為int64
	// g.GenerateModel("users", gen.FieldIgnore("address"), gen.FieldType("id", "int64")),

	// // tagsㄉJSON欄位
	// g.GenerateModel("customer", gen.FieldType("tags", "datatypes.JSON")),
	// g.ApplyBasic(
	// 	//根據db table生成所有struct
	// 	g.GenerateAllTables(),
	// )
	// 執行生成
	g.Execute()
}
