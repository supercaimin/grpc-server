// cmlmain.go
package main

import (
	context "context"
	"fmt"

	"hstcmler/cml"

	"google.golang.org/grpc"
)

type cmlClient struct {
	gRpcServerstr string
	logServerstr  string
	gRpcServPort  string
	gLogServPort  string
	gSendMsgNum   uint64
	gLastMsg      string //最后发的消息的结构信息
}

var gCmlClient cmlClient

func gClientlobalVarInit() {

	//cmlClient = new(cmlClient)

	gCmlClient.gRpcServerstr = "127.0.0.1"
	gCmlClient.logServerstr = "0.0.0.0"
	gCmlClient.gRpcServPort = "8000"
	gCmlClient.gLogServPort = "510"
	gCmlClient.gSendMsgNum = 0
	gCmlClient.gLastMsg = "unkown"

}

func main() {
	var i uint32
	var client cml.SonicCmdServiceClient
	var soniccmdinfo cml.SONICCmdInputProfile
	//var rtninfo cml.SONICCmdRtnProfile

	gClientlobalVarInit()

	//连接cml server;
	conn, err := grpc.Dial((gCmlClient.gRpcServerstr + ":" + gCmlClient.gRpcServPort), grpc.WithInsecure())
	//conn, err := grpc.Dial("127.0.0.1:520", grpc.WithInsecure())
	if err != nil {
		fmt.Println("connect server failed !\n")
	}
	defer conn.Close()
	client = cml.NewSonicCmdServiceClient(conn)

	for i = 0; i < 5; i++ {
		soniccmdinfo.CmdId = i
		//fmt.Printf(soniccmdinfo.CmdinputLine, "This is Client test %d \n", i)
		soniccmdinfo.CmdinputLine = "This is Client test !\n"
		rtninfo, _ := client.ExecSonicCfgProfile(context.Background(), &soniccmdinfo)
		if nil != rtninfo {
			fmt.Println("cmdid %h,rtn str:%s", rtninfo.CmdId, rtninfo.CmdrtnStr)
		} else {
			fmt.Println("rtninfo is null !")
		}
	}

	fmt.Println("hello,client test finished !")
	fmt.Println(grpc.Version)

}
