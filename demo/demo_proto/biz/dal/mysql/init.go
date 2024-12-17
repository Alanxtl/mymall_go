package mysql

import (
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/model"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	//fmt.Println(v)
}
