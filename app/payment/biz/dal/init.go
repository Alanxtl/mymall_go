package dal

import (
	"github.com/Alanxtl/mymall_go/app/payment/biz/dal/mysql"
	"github.com/Alanxtl/mymall_go/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
