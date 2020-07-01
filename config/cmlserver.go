// cmlServer.go
package config

import (
	"fmt"

	cmodel "hstcmler/cml"

	"hstcmler/errcode"
	"hstcmler/xlog"
	"net"

	"google.golang.org/grpc"
)

type T_CML_SERVER struct {
	myaddr        net.IPAddr
	port          int
	is_init       bool
	log           *xlog.XLogger
	listener      net.Listener
	server        *grpc.Server
	aclService    cmodel.CmlAclServiceServer
	bgpService    cmodel.CmlBgpServiceServer
	sonicService  cmodel.SonicCmdServiceServer
	syslogService cmodel.SyslogConfigServiceServer
	qosService    cmodel.CmlqosserviceServer
	ifService     cmodel.CmlifserviceServer
	macService    cmodel.CmlmacserviceServer
}

var cmlServer *T_CML_SERVER

func init() {
	cmlServer = new(T_CML_SERVER)
}

/*初始化Cml模型层*/
/*func InitCml(gRpcServer *net.IPAddr, gRpcPort int, logServer *net.IPAddr, logPort int) {*/
func InitCml(gRpcServer string, gRpcPort int, logServer string, logPort int) {
	//初始化Log系统，方便后续使用
	if cmlServer.is_init {
		return
	}
	//init syslog client info
	cmlServer.log = xlog.NewXLogger("CML", gRpcServer, 0, logServer, logPort)

	cmlServer.InitRpcServer()
	cmlServer.is_init = true

	InitCmlDbInfo()
	/*从文件读入告警码,需要考虑重启恢复配置的情况 chenwei 20200224*/
	ReadAlarmDefInfo()
}

//初始化RPC Server，将每个配置服务注册到gRPC Server
func (self *T_CML_SERVER) InitRpcServer() errcode.RESULT {
	var err error
	//self.listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", self.myaddr.String(), self.port))

	self.listener, err = net.Listen("tcp", "127.0.0.1:50001")
	if err != nil {
		self.log.Errorf("CmlInit", errcode.RESULT_ERROR_NOT_INITIATED, "Init gRpc Server failed,err: %v", err)
		fmt.Println("Init gRpc Server failed !")
		return errcode.RESULT_ERROR_NOT_INITIATED
	}
	self.server = grpc.NewServer()

	//创建ACL、QoS的对象，这两者都实现了相应的proto文件定义的ACL、QoS的gRPC接口
	self.aclService = NewAclConfig()
	self.sonicService = NewSonicConfig()
	self.syslogService = NewSyslogConfig()
	self.qosService = NewQosConfig()
	self.ifService = NewIfConfig()
	self.macService = NewMacConfig()

	//将gRPC Server接口实现实例绑定到gRPC实例上
	cmodel.RegisterCmlAclServiceServer(self.server, self.aclService)
	cmodel.RegisterSonicCmdServiceServer(self.server, self.sonicService)
	cmodel.RegisterSyslogConfigServiceServer(self.server, self.syslogService)
	cmodel.RegisterCmlqosserviceServer(self.server, self.qosService)
	cmodel.RegisterCmlifserviceServer(self.server, self.ifService)
	cmodel.RegisterCmlmacserviceServer(self.server, self.macService)

	//启动RPC服务，后续调用由gRPC框架直接分发到上述注册的实现类方法中
	go self.server.Serve(self.listener)
	//self.server.Serve(self.listener)

	return errcode.RESULT_SUCCESS
}
