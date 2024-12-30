package dal

import (
	"github.com/Alanxtl/mymall_go/app/frontend/biz/dal/mysql"
	"github.com/Alanxtl/mymall_go/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
