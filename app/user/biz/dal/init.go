package dal

import (
	"github.com/Alanxtl/mymall_go/app/user/biz/dal/mysql"
	"github.com/Alanxtl/mymall_go/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
