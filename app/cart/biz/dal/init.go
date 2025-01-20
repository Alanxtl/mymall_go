package dal

import (
	"github.com/Alanxtl/mymall_go/app/cart/biz/dal/mysql"
	"github.com/Alanxtl/mymall_go/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
