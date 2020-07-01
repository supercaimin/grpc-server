// cmlmain.go
package main

import (
	"fmt"
	"net"

	"hstcmler/config"

	"google.golang.org/grpc"
)

var gRpcServerstr string = "127.0.0.1"
var logServerstr string = "127.0.0.1"

func main() {

	var gRpcserver, logServer net.IPAddr
	var rpcport, logport int
	var ch chan int

	gRpcserver.IP = net.ParseIP(gRpcServerstr)
	logServer.IP = net.ParseIP(logServerstr)
	rpcport = 50001
	logport = 30001

	config.InitCml(gRpcServerstr, rpcport, logServerstr, logport)

	fmt.Println("hello,test!")
	fmt.Println(grpc.Version)

	/*wait until rpc call exit*/
	<-ch
}
