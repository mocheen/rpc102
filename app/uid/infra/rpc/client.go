package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/doutokk/doutok/common/clientsuite"
	"rpc102/app/user/kitex_gen/user/userservice"
)

var Client userservice.Client
var err error

func Init() {
	registryAddr := "127.0.0.1:8500"
	serviceName := "user"
	commonSuite := client.WithSuite(clientsuite.CommonGrpcClientSuite{
		RegistryAddr:       registryAddr,
		CurrentServiceName: serviceName,
	})

	Client, err = userservice.NewClient("user", commonSuite)
	if err != nil {
		panic(err)
	}

}
