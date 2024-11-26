package rpc

import (
	"github.com/cloudwego/kitex/client"
	"gomall/app/checkout/conf"
	CheckOutUtils "gomall/app/checkout/utils"
	"gomall/common/clientsuite"
	"gomall/rpc_gen/kitex_gen/cart/cartservice"
	"gomall/rpc_gen/kitex_gen/order/orderservice"
	"gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"sync"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
	err           error
	registryAddr  string
	serviceName   string
	commonSuite   client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Registry.RegistryAddress[0]
		serviceName = conf.GetConf().Kitex.Service
		commonSuite = client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: serviceName,
			RegistryAddr:       registryAddr,
		})
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", commonSuite)
	CheckOutUtils.MustHandleError(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", commonSuite)
	CheckOutUtils.MustHandleError(err)
}

func initPaymentClient() {
	PaymentClient, err = paymentservice.NewClient("payment", commonSuite)
	CheckOutUtils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	CheckOutUtils.MustHandleError(err)
}
