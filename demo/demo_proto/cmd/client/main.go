package main

import (
	"context"
	"demo_proto/kitex_gen/pbapi"
	"demo_proto/kitex_gen/pbapi/echo"
	"fmt"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Println(1)
		log.Fatal(err)
	}
	c, err := echo.NewClient("demo_proto", client.WithResolver(r))
	if err != nil {
		log.Println(2)
		log.Fatal(err)
	}
	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto_client")
	res, err := c.Echo(ctx, &pbapi.Request{
		Message: "hello",
	})
	if err != nil {
		log.Println(3)
		log.Fatal(err)
	}

	fmt.Println(res)
}
