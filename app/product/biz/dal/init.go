package dal

import (
	"github.com/Alanxtl/mymall_go/app/product/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
