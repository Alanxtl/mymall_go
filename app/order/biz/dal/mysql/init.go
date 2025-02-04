package mysql

import (
	"fmt"
	"github.com/Alanxtl/mymall_go/app/order/biz/model"
	"github.com/Alanxtl/mymall_go/app/order/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)

	if os.Getenv("GO_ENV") != "online" {
		err := DB.AutoMigrate(
			&model.Order{},
			&model.OrderItem{},
		)
		if err != nil {
			klog.Error(err)
		}
	}
}
