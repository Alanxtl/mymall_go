package rpc

import (
	"github.com/Alanxtl/mymall_go/app/cart/conf"
	"github.com/Alanxtl/mymall_go/app/cart/utils"
	"github.com/Alanxtl/mymall_go/common/clientsuite"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	ProductClient productcatalogservice.Client
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]

	err  error
	once sync.Once
)

func Init() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	utils.MustHandleError(err)

}
