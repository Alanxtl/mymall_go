package rpc

import (
	"github.com/Alanxtl/mymall_go/app/checkout/conf"
	"github.com/Alanxtl/mymall_go/common/clientsuite"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/order/orderservice"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/Alanxtl/mymall_go/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"sync"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client

	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]

	once sync.Once
	err  error
)

func Init() {
	once.Do(func() {
		initProductClient()
		initCartClient()
		initPaymentClient()
		initOrderClient()
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
	if err != nil {
		panic(err)
	}
}

func initCartClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}

	CartClient, err = cartservice.NewClient("cart", opts...)
	if err != nil {
		panic(err)
	}
}

func initPaymentClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}

	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	if err != nil {
		panic(err)
	}

}

func initOrderClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}

	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		panic(err)
	}

}
