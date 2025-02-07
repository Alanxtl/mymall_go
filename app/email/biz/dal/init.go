package dal

import (
	"github.com/Alanxtl/mymall_go/app/email/biz/dal/mysql"
	"github.com/Alanxtl/mymall_go/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
