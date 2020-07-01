// db_define.go
// 配置CML中各业务db key,field定义文件

package config

/*qos数据库key以及field的名称定义*/
const (
	KEY_QOS_POLICER_GLOBAL     string = "POLIGLOBAL"           /*policy-glo,time,color action*/
	KEY_QOS_POLICER                   = "POLICER|"             /*POLICER|Policer name(string)*/
	FIELD_METERTYPE                   = "METER_TYPE"           /*fix: byte, not support pps*/
	FIELD_MODE                        = "MODE"                 /*SR_TCM   TR_TCM*/
	FIELD_COLOR_SOURCE                = "COLOR_SOURCE"         /*AWARE,BLIND*/
	FIELD_CBS                         = "CBS"                  /*dec string*/
	FIELD_CIR                         = "CIR"                  /*dec string*/
	FIELD_PBS                         = "PBS"                  /*dec string*/
	FIELD_PIR                         = "PIR"                  /*dec string*/
	FIELD_GREEN_PACKET_ACTION         = "GREEN_PACKET_ACTION"  /*DROP, FORWARD*/
	FIELD_YELLOW_PACKET_ACTION        = "YELLOW_PACKET_ACTION" /*DROP, FORWARD*/
	FIELD_RED_PACKET_ACTION           = "RED_PACKET_ACTION"    /*DROP, FORWARD*/
	FIELD_COMMANDLINE                 = "COMMANDLINE"          /*用于CLI模式的配置显示(命令行)*/
	FIELD_UPDATETIME                  = "updatetime"

	KEY_QOS_DIFFSERVER      = "DIFFSERVNAME" /*QOS diffserv name*/
	FIELD_DIFFSERV_IN8021P  = "IN8021P|"     /*in 8021P INFO*/
	FIELD_DIFFSERV_OUT8021P = "OUT8021P|"    /*OUT 8021P INFO*/
	FIELD_DIFFSERV_INDSCP   = "INDSCP"
	FIELD_DIFFSERV_OUTDSCP  = "OUTDSCP"
	FIELD_DIFFSERV_INEXP    = "INEXP"
	FIELD_DIFFSERV_OUTEXP   = "OUTEXP"
	FIELD_DIFFSERV_INMPLS   = "INMPLS"
	FIELD_DIFFSERV_OUTMPLS  = "OUTMPLS"
)

//接口相关的DB KEY与FIELD定义
//port(物理接口属性)信息
const (
	PORT_TABLE_NAME          string = "PORT"
	PORT_ETHERNET_NAME              = "ethernet"
	KEY_PORT_ETHERNET_PREFIX        = PORT_ETHERNET_NAME
	//接口下保存的field信息
	FIELD_PORT_LANES        = "lanes"
	FIELD_PORT_JUMBER       = "jumber"
	FIELD_PORT_SPEED_MODE   = "speedmode"   //速率模式,10G-G,25G-10G, 40G-100G
	FIELD_PORT_SPEED_NEGO   = "speednego"   //速率协商模式
	FIELD_PORT_UP_DELAY     = "updelay"     //物理up延迟上报
	FIELD_PORT_STATUS_DELAY = "statusdelay" //状态延迟上报
	FIELD_PORT_FLAP_INFO    = "flap"        //震荡抑制信息
	FIELD_PORT_SPEED_REAL   = "speed"       //接口当前的工作速率
	FIELD_PORT_LOOP_MODE    = "loopmode"    //环回模式,如内环
	FIELD_PORT_IFG_INFO     = "ifg"         //帧间隙配置
	FIELD_PORT_LAN_WAN      = "lanwan"      //lan wan工作模式
	FIELD_PORT_WAN_INFO     = "waninfo"     //wan模式下的信息格式
	FIELD_PORT_FEC_MODE     = "fecmode"     //fec模式
)

//接口的key为 iftype+id
//不建议把ip地址体现在key中因为改变后需要再rename key,而且接口下配置多个地址时不好体现
const (
	INTERFACE_TABLE_NAME      string = "INTERFACE"
	INTERFACE_LOOPBACK_NAME          = "loopback"
	INTERFACE_ETHERNET_NAME          = "ethernet"
	INTERFACE_VLANIF_NAME            = "vlan" //vlan下配置了aggregate命令后就为supervlan特性
	INTERFACE_ETRUNK_IF_NAME         = "eth-trunk"
	INTERFACE_BRIDGE_NAME            = "bridge"
	KEY_INTF_LOOPBAKC_PREFIX         = INTERFACE_LOOPBACK_NAME //Loopback+id
	KEY_INTF_ETHERNET_PREFIX         = INTERFACE_ETHERNET_NAME //Ethernet+id
	KEY_INTF_VLANIF_PREFIX           = INTERFACE_VLANIF_NAME
	KEY_INTF_ETRUNK_IF_PREFIX        = INTERFACE_ETRUNK_IF_NAME
	KEY_INTF_BRIDGE_IF_PREFIX        = INTERFACE_BRIDGE_NAME
	//接口下存储信息的field
	FIELD_INTF_INDEX            = "index"       //接口索引,自动分配,非配置
	FIELD_INTF_ALIAS            = "alias"       //接口别名
	FIELD_INTF_DESCRIP          = "descrip"     //接口描述
	FIELD_INTF_VLAN_LIST        = "vlanlist"    //接口可接入的vlan 列表,对vlanif就只有一个vlan
	FIELD_INTF_ACCESS_TYPE      = "accesstype"  //接口接入类型,access,hybrid,trunk
	FIELD_INTF_PVID             = "pvid"        //接口默认的vlan id
	FIELD_INTF_ETH_TRUNKID      = "eth-trunkid" //接口归属的ETRUNK接口ID
	FIELD_INTF_MTU              = "mtu"         //接口mtu
	FIELD_INTF_JUMBO_ETRUNK     = "jumbo"       //etrunk口的Jumbo信息
	FIELD_INTF_ADMIN_STATUS     = "adminstatus"
	FIELD_INTF_SPEED            = "speed"
	FIELD_INTF_SWITCH_MODE      = "switchmode" //L2 或L3转发模式
	FIELD_INTF_PROTO_UPDELAY    = "protoupdelay"
	FIELD_INTF_STATISTIC_ENABLE = "statistic"
	FIELD_INTF_STATIS_INTERVAL  = "statisinterval"
	FIELD_INTF_DOWN_TRAP_DIS    = "trapdis"
	FIELD_INTF_UNNUMBER_IF      = "unnumberif"
	FIELD_INTF_TTL_SOURCE       = "ttlsource"
	FIELD_INTF_UNREACH_SOURCE   = "unreachsource"
	FIELD_INTF_ARP_AGEINFO      = "arpage"
	FIELD_INTF_V6_ENABLE        = "v6enable"
	FIELD_INTF_IPADDR           = "ipv4addr" //接口地址列表,ip/mask, 多个地址用,分隔
	//v6 接口相关的信息
	FIELD_INTF_V6_MTU         = "v6mtu"
	FIELD_INTF_AUTO_LINKLOCAL = "autolinklocal"
	FIELD_INTF_ANYCAST        = "anycast" //anycast地址
	FIELD_INTF_LINK_LOCAL     = "linklocal"
	FIELD_INTF_EUID_ADDR      = "euidaddr"
	FIELD_INTF_CGA_ADDR       = "cgaaddr"
	FIELD_INTF_CGA_LINKLOCAL  = "cgalinklocal"
	FIELD_INTF_ND_AUTO_CFG    = "ndautocfg"
	FIELD_INTF_ND_AUTO_DETECT = "ndautodetect"
	FIELD_INTF_ND_RA_INTERVAL = "rainterval"
)

//mac 数据库定义
const (
	//mac 全局表信息
	MAC_GLOBAL_TABLE_NAME  string = "MAC_GLOBAL"
	KEY_MAC_GLOBAL                = "MAC_GLOBAL"
	FIELD_MAC_AGTIME              = "AGETIME"
	FIELD_MAC_BLACK_PREFIX        = "BLACK|" //黑洞mac地址,field为 BLACK| + vlanid
	FIELD_MAC_FLAP_INFO           = "FLAP"   //mac漂移信息,探测使能及探测间隔等
	//mac 静态条目表信息
	MAC_STATIC_TABLE_NAME = "MAC_STATIC"
	KEY_MAC_STATIC_PREFIX = "MAC_STATIC|" //KEY为 MAC_STATIC|ifname
	FIELD_MAC_ITEM_PREFIX = "MAC|"        //key为 MAC|macaddr string
	//mac配置信息在接口下的存储,接口表KEY:ifname|ip/mask(接口名系统不重复,所以ifname可定位)
	FIELD_IF_MAC_LEARN         = "MACLEARN" //该值是在接口表下,
	FIELD_IF_MAC_LIMIT         = "MACLIMIT"
	FIELD_IF_MAC_BRIDGE_ENABLE = "BRIDGEENABLE"
)

/*syslog db*/
const (
	KEY_SYSLOG         string = "syslog"
	FIELD_CONSOLE             = "console"
	FIELD_MONITOR             = "monitor"
	FIELD_TRAP                = "trap"
	FIELD_PERSISTENT          = "persistent"
	FIELD_SERVER              = "server"
	FIELD_BUFFERSIZE          = "buffersize"
	FIELD_STORAGESIZE         = "memsize"
	FIELD_SORCEIP             = "sourceip"
	FIELD_SOURCEPORT          = "sourceport"
	FIELD_SERVERIP            = "serverip"
	FIELD_FILEDURATION        = "fileduration"

	KEY_SYSLOG_SEVER     = "logserver"
	FIELD_SYSLOG_SERVER1 = "server1"
	FIELD_SYSLOG_SERVER2 = "server2"
	FIELD_LOGSERVER_VRF  = "vrf"
	FIELD_LOGSERVER_IP   = "ip"
	FIELD_LOGSERVER_PORT = "port"
)

/*alarm level的相关DB KEY及field*/
/*Alarm|modulename为key,code-level为field-value对*/
const (
	KEY_ALARM_LEVEL_ORIGINAL string = "Alarm|"
	KEY_ALARM_LEVEL_NEW             = "AlarmNew|"
	KEY_ALARM_CURRNET_INFO          = "AlarmCur|"
	KEY_ALARM_UPDATEINFO            = "Alarm|updateinfo"
	FIELD_ALARM_UPDATETIME          = "updatetime"
	DATA_SPLIT_STR                  = "::" /*用于分割level与position*/
	KEY_DATA_SPLIT_STR              = "|"  /*用与分割key各元素*/
	DATA_UPDATEINFO_STR             = "updateinfo"
)

/*sonic 命令key*/
const (
	KEY_FRR_CMD_INPUT  string = "FRRCMD"  /*frr命令,各vty通道放一起*/
	KEY_FRR_CMD_OUTPUT        = "FRROUT|" /*frr命令输出,各vty通道分开*/
)

/*app server config change pub channel*/
const (
	SyslogCfgChan string = "syslog.cfg"
	AclCfgChan           = "acl.cfg"
	MacCfgChan           = "mac.cfg"
	IfCfgChan            = "if.cfg"
)
