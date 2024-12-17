package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi/echoservice"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"log"

	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")

	if err != nil {
		log.Fatal(err)
	}

	c, err := echoservice.NewClient("demo_proto",
		client.WithResolver(r),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithMiddleware(middleware.Middleware))

	if err != nil {
		log.Fatal(err)
	}

	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_protocol_client")

	echo, err := c.Echo(ctx, &pbapi.Request{Message: "hello"})

	var bizErr *kerrors.BizStatusError

	if err != nil {
		ok := errors.As(err, &bizErr)
		if ok {
			fmt.Println(bizErr)
		}
		log.Fatal(err)
	}

	fmt.Printf("%v", echo)

}
