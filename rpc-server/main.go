package main

import (
	"log"
	"os"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	rpc "github.com/ngquyduc/assignment_demo_2023/rpc-server/kitex_gen/rpc/imservice"
)

func main() {
	var link string
	if os.Getenv("ENV") == "PROD" {
		link = "etcd:2379"
	} else {
		link = "127.0.0.1:2379"
	}
	log.Println("link:", link)
	r, err := etcd.NewEtcdRegistry([]string{link}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	svr := rpc.NewServer(new(IMServiceImpl), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "demo.rpc.server",
	}))

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
