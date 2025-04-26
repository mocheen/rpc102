package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/doutokk/doutok/common/clientsuite"
	"rpc102/app/uid/kitex_gen/uid/uidservice"
)

var Client uidservice.Client
var err error

func Init() {
	registryAddr := "127.0.0.1:8500"
	serviceName := "uid"
	commonSuite := client.WithSuite(clientsuite.CommonGrpcClientSuite{
		RegistryAddr:       registryAddr,
		CurrentServiceName: serviceName,
	})

	Client, err = uidservice.NewClient("uid", commonSuite)
	if err != nil {
		panic(err)
	}

}
