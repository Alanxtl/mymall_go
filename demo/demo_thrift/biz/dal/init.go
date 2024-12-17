package dal

import (
	"github.com/cloudwego/biz-demo/tree/main/gomall/tutorial/ch02/demo/demo_thrift/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/tree/main/gomall/tutorial/ch02/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
