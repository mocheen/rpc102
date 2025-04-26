package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/doutokk/doutok/common/clientsuite"
	"rpc102/app/uid/kitex_gen/uid"
	"rpc102/app/uid/kitex_gen/uid/uidservice"
)

func main() {
	fmt.Println("ttt")
	registryAddr := "127.0.0.1:8500"
	serviceName := "uid"
	commonSuite := client.WithSuite(clientsuite.CommonGrpcClientSuite{
		RegistryAddr:       registryAddr,
		CurrentServiceName: serviceName,
	})
	newClient, err := uidservice.NewClient("uid", commonSuite)
	if err != nil {
		fmt.Println(err)
	}
	r, err := newClient.UidGen(context.Background(), &uid.Empty{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(r.Uid)
}
