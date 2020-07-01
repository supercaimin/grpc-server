package config

// cmdid 参考cmdid-const-define.h文件

const (
	//define MAC_CMD_AGING_TIME			 	 0x01400005
	//#define MAC_CMD_BLACK_HOLE			     0x02400005
	//#define MAC_CMD_FLAP_AGING_TIME		     0x03400005

	//#define MAC_CMD_FLAP_DETECT			     0x04400005
	//#define MAC_CMD_FLAP_DETECT_LEVEL_LOW    0x05400005
	//#define MAC_CMD_FLAP_DETECT_LEVEL_MIDDLE 0x06400005
	//#define MAC_CMD_FLAP_DETECT_LEVEL_HIGH   0x07400005
	MacCmdAgingTime             = 0x01400005
	MacCmdBlackHole             = 0x02400005
	MacCmdFlapAgingTime         = 0x03400005
	MacCmdFlapDetect            = 0x04400005
	MacCmdFlapDetectLevelLow    = 0x05400005
	MacCmdFlapDetectLevelMiddle = 0x06400005
	MacCmdFlapDetectLevelHigh   = 0x07400005
	MacCmdStatic                = 0x08400005
	MacCmdLearn                 = 0x09400005
	MacCmdLimit                 = 0x10400005
	MacCmdPortBridge            = 0x11400005
	//#define INTERFACE_CMD_SET_TYPE_NUM       0x01020006
	//#define INTERFACE_CMD_LOOKBACK        	 0x02020006
	//#define INTERFACE_CMD_NULL        	  	 0x03020006

	IfCmdSetTypeNum       = 0x01020006
	IfCmdLoopBack         = 0x02020006
	IfCmdNull             = 0x03020006
	IfCmdDescrip          = 0x04020006
	IfCmdClearCfgThis     = 0x05020006
	IfCmdClearCfgIntf     = 0x06020006
	IfCmdProtoUpDelay     = 0x07020006
	IfCmdFlowStatInterval = 0x08020006
	IfCmdShutdown         = 0x09020006
	IfCmdShutdownIntf     = 0x0a020006
	IfCmdShutdownNet      = 0x0b020006
	IfCmdShutdownTransmit = 0x0c020006
	IfCmdShutdownReceive  = 0x0d020006

	IfCmdFecModeBase = 0x0e020006
	IfCmdFecModeNone = 0x0f020006
	IfCmdFecModeRS   = 0x10020006
	IfCmdControlFlap = 0x11020006
	IfCmdIfRange     = 0x12020006
	IfCmdJumboframe  = 0x13020006
	IfCmdLoopback    = 0x14020006

	IfCmdPortModeLan   = 0x15020006
	IfCmdPortModeWan   = 0x16020006
	IfCmdPortMode10g   = 0x17020006
	IfCmdPortModeGe    = 0x18020006
	IfCmdPortMode25ge  = 0x19020006
	IfCmdPortMode100ge = 0x1a020006

	IfCmdPortSwitch      = 0x1b020006
	IfCmdPortSwitchBatch = 0x1c020006

	IfCmdSetUpDelay = 0x1d020006

	IfCmdSpeed10    = 0x1e020006
	IfCmdSpeed100   = 0x1f020006
	IfCmdSpeed1000  = 0x20020006
	IfCmdSpeed10000 = 0x21020006
	IfCmdSpeed40000 = 0x22020006

	IfCmdSpeedAuto10    = 0x23020006
	IfCmdSpeedAuto100   = 0x24020006
	IfCmdSpeedAuto1000  = 0x25020006
	IfCmdSpeedAuto10000 = 0x26020006

	//subinterface statistics
	IfCmdSubIf      = 0x27020006
	IfCmdStatistics = 0x28020006
	IfCmdMtu        = 0x29020006
)
