package rpc

import (
	"github.com/Alanxtl/mymall_go/app/frontend/conf"
	"github.com/Alanxtl/mymall_go/app/frontend/utils"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/user/userservice"
	client "github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client

	once sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
	})
}

func initUserClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	utils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	UserClient, err = userservice.NewClient("user", opts...)
	utils.MustHandleError(err)
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	utils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	utils.MustHandleError(err)

}
