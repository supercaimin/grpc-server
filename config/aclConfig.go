// aclConfig.go
package config

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"	
	"strings"
	pb "hstcmler/cml"
//	"github.com/go-redis/redis"	
)


type AclConfig struct {
}

/*新建一个类型 aclconfig*/
func NewAclConfig() *AclConfig {
	return new(AclConfig)
}

type aclFieldTable struct{
	field string
	value string
}


type aclRuleItem int32

/*CONFIG DB ACL RULE ITEM*/
const(
	ACL_NUMBER aclRuleItem=0
	ACL_STEP   aclRuleItem=1
	ACL_RULE_ID   aclRuleItem=2
	ACL_RULE_ID1   aclRuleItem=3
	ACL_RULE_ID2   aclRuleItem=4	
	ACL_ACTION   aclRuleItem=5
	ACL_FRAGMENT_TYPE aclRuleItem=6
	ACL_TIME_RANGE	aclRuleItem=7
	ACL_SRC_IP_MASK aclRuleItem=8
	ACL_DST_IP_MASK aclRuleItem=10
	ACL_SRC_MAC_MASK aclRuleItem=11
	ACL_DST_MAC_MASK aclRuleItem=12
	ACL_ETHER_TYPE aclRuleItem=13		
	ACL_PACKET_TYPE aclRuleItem=14
	ACL_VLAN_ID aclRuleItem=15
	ACL_DOT1P aclRuleItem=16
	ACL_IN_VLAN_ID aclRuleItem=17
	ACL_IN_DOT1P aclRuleItem=18
	ACL_DTAG aclRuleItem=19	
	ACL_PROTO_TYPE aclRuleItem=20
	ACL_SRC_PORT_MIN aclRuleItem=21
	ACL_SRC_PORT_MAX aclRuleItem=22
	ACL_DST_PORT_MIN aclRuleItem=23
	ACL_DST_PORT_MAX aclRuleItem=24
	ACL_ICMP_TYPE aclRuleItem=25
	ACL_ICMP_CODE aclRuleItem=26
	ACL_IGMP_TYPE aclRuleItem=27
	ACL_TCP_FLAG aclRuleItem=28
	ACL_DSCP aclRuleItem=29
	ACL_PRECEDENCE aclRuleItem=30
	ACL_TOS aclRuleItem=31	
	ACL_TTL_EXPIRED aclRuleItem=32	
	ACL_VPN_INSTANCE aclRuleItem=33
	ACL_CMDLINE 	aclRuleItem=34	
	ACL_LAST_ITEM aclRuleItem=35 
)
//将chip type 类型转换为字符串
func (item aclRuleItem) String() string {
    switch item {
    case ACL_NUMBER:
        return "ACL_NUMBER"
    case ACL_STEP:
        return "ACL_STEP"
    case ACL_RULE_ID:
		return "ACL_RULE_ID"
    case ACL_RULE_ID1:
        return "ACL_RULE_ID1"
    case ACL_RULE_ID2:
        return "ACL_RULE_ID2"
    case ACL_ACTION:
		return "ACL_ACTION"
    case ACL_FRAGMENT_TYPE:
        return "ACL_FRAGMENT_TYPE"
    case ACL_TIME_RANGE:
        return "ACL_TIME_RANGE"
    case ACL_SRC_IP_MASK:
		return "ACL_SRC_IP_MASK"
    case ACL_DST_IP_MASK:
		return "ACL_DST_IP_MASK"
    case ACL_SRC_MAC_MASK:
		return "ACL_SRC_MAC_MASK"
    case ACL_DST_MAC_MASK:
		return "ACL_DST_MAC_MASK"		
    case ACL_ETHER_TYPE:
		return "ACL_ETHER_TYPE"
    case ACL_PACKET_TYPE:
		return "ACL_PACKET_TYPE"
    case ACL_VLAN_ID:
		return "ACL_VLAN_ID"
    case ACL_DOT1P:
		return "ACL_DOT1P"
    case ACL_IN_VLAN_ID:
		return "ACL_VLAN_ID"
    case ACL_IN_DOT1P:
		return "ACL_IN_DOT1P"
   	case ACL_DTAG:
		return "ACL_DTAG"
    case ACL_PROTO_TYPE:
		return "ACL_PROTO_TYPE"			
    case ACL_SRC_PORT_MIN:
		return "ACL_SRC_PORT_MIN"	
    case ACL_SRC_PORT_MAX:
		return "ACL_SRC_PORT_MAX"	
    case ACL_DST_PORT_MIN:
		return "ACL_DST_PORT_MIN"	
    case ACL_DST_PORT_MAX:
		return "ACL_DST_PORT_MAX"		
	case ACL_ICMP_TYPE:
		return "ACL_ICMP_TYPE"	
	case ACL_ICMP_CODE:
		return "ACL_ICMP_CODE"
	case ACL_IGMP_TYPE:
		return "ACL_IGMP_TYPE"		
	case ACL_TCP_FLAG:
		return "ACL_TCP_FLAG"
	case ACL_DSCP:
		return "ACL_DSCP"
	case ACL_TOS:
		return "ACL_TOS"
	case ACL_PRECEDENCE:
		return "ACL_PRECEDENCE"
	case ACL_TTL_EXPIRED:
		return "ACL_TTL_EXPIRED"
    case ACL_VPN_INSTANCE:
		return "ACL_VPN_INSTANCE"
    case ACL_CMDLINE:
		return "ACL_CMDLINE"		
	}
    return "N/A"
}
type aclTableItem int32
const(
	ACL_TABLE_NUMBER aclTableItem=0
	ACL_TYPE   aclTableItem=1
	ACL_TABLE_ACTION   aclTableItem=2
	ACL_TABLE_STEP   aclTableItem=3
	ACL_BIND_TYPE   aclTableItem=4	
	ACL_BIND_TYPE_NUMBER   aclTableItem=5
	ACL_TABLE_LAST_ITEM    aclTableItem=6
)
func (item aclTableItem) String() string {
    switch item {
		case ACL_TABLE_NUMBER:
			return "ACL_NUMBER"
		case ACL_TYPE:
			return "ACL_TYPE"
		case ACL_TABLE_ACTION:
			return "ACL_ACTION"
		case ACL_TABLE_STEP:
			return "ACL_STEP"
		case ACL_BIND_TYPE:
			return "ACL_BIND_TYPE"
		case ACL_BIND_TYPE_NUMBER:
			return "ACL_BIND_TYPE_NUMBER"
	}
	return "N/A"
}

type aclRuleType int32
const(
	ACL_BASIC aclRuleType=0
	ACL_ADVANCE aclRuleType=1
	ACL_LINK aclRuleType=2
	ACL_ARP aclRuleType=3			
	ACL_MIXED aclRuleType=4
)
var dequeueItemNum int32 =30
var RULE_UPDATE string ="RULE_UPDATE:"
var RULE_DELETE string ="RULE_DELETE:"
var TABLE_BIND string ="TABLE_BIND:"
var TABLE_UNBIND string ="TABLE_UNBIND:"
var TABLE_DELETE string ="TABLE_DELETE:"

var msgAclRuleUpdateQueue map[string]string 
var msgAclRuleDeleteQueue map[string]string 
var msgAclTableBindQueue map[string]string //这个map用来传送acl table更新
var msgAclTableUnbindQueue map[string]string
var msgAclTableDeleteQueue map[string]string
var msgTimerStatus int =0;

var msgAclQueueMutex sync.Mutex


var msgTimer *time.Timer
/***************************************************************************
* 函数说明: 将CONFIG PUB的消息增加消息头，
* msgBody:消息体，msgHeader:消息头
* @return  msgBody：新消息
*  创建者  于士超   2019.11.27
****************************************************************************/
func msgAclRuleAddHeader(msgBody string,msgHeader string) string{
	//切片发送
	msgBody =strings.Replace(msgBody,"ACL",msgHeader+"ACL",1)
	
	return msgBody
}
/***************************************************************************
* 函数说明: 将CONFIG PUB的消息增加消息头，
* msgBody:消息体，msgHeader:消息头
* @return  msgBody：新消息
*  创建者  于士超   2019.11.27
****************************************************************************/
func msgAclTableAddHeader(msgBody string,msgHeader string) string{
	msgBody = msgHeader + msgBody
	
	return msgBody
}


func msgAclRuleSend(msgBody string,msgHeader string) int32{
	if msgBody!=""{
		msgBody =msgAclRuleAddHeader(msgBody,msgHeader)
		aclRedisClient.Publish("CONFIG_ACL_CHANNEL",msgBody).Result()
	}
	
	return 0
}
func msgAclTableSend(msgBody string,msgHeader string) int32{
	if msgBody!=""{
		msgBody =msgAclTableAddHeader(msgBody,msgHeader)
		aclRedisClient.Publish("CONFIG_ACL_CHANNEL",msgBody).Result()
	}
		
	return 0
}
/***************************************************************************
* 函数说明: msgTimer回调函数，该函数主要是检查消息队列是否为空，如果为空，则stop定时器，否则发布消息 
*  创建者  于士超   2019.11.27
****************************************************************************/

func msgTimerHandle(){
	again:	
	msg,num:=msgRuleDeleteDequeueItem()
	if num != 0{
		msgAclRuleSend(msg,RULE_DELETE)		
	}else if num==dequeueItemNum{
		goto again
	}
	again_1:
	msg,num=msgRuleUpdateDequeueItem()
	if num != 0{
		msgAclRuleSend(msg,RULE_UPDATE)		
	}else if num==dequeueItemNum{
		goto again_1
	}
	again_2:
	msg,num=msgTableDeleteDequeueItem()
	if num != 0{
		msgAclTableSend(msg,TABLE_DELETE)		
	}else if num==dequeueItemNum{
		goto again_2
	}
	again_3:
	msg,num=msgTableBindDequeueItem()
	if num != 0{
		msgAclTableSend(msg,TABLE_BIND)		
	}else if num==dequeueItemNum{
		goto again_3
	}
	again_4:	
	msg,num=msgTableUnbindDequeueItem()
	if num != 0{
		msgAclTableSend(msg,TABLE_UNBIND)		
	}else if num==dequeueItemNum{
		goto again_4
	}
	msgTimer.Stop()
	msgTimerStatus =1;

}
/***************************************************************************
* 函数说明: 从msgAclRuleDeleteQueue中取出条目，一次最多取出30个条目
* @rturn msg:消息内容，num:条目个数
*  创建者  于士超   2019.11.27
****************************************************************************/
func msgRuleDeleteDequeueItem()(string,int32){
	/*如果定时器没有创建，那么创建定时器，定时广播消息*/
	var msg string
	var num int32
	msgAclQueueMutex.Lock()	

	for index,_:=range msgAclRuleDeleteQueue{
		msg = msg + "," + index
		delete(msgAclRuleDeleteQueue,index)	
		num++;
		if	num>=dequeueItemNum{
			break;
		}
	}	
	msgAclQueueMutex.Unlock()
	msg=strings.Trim(msg,",")
	return msg,num
}
/***************************************************************************
* 函数说明: 从msgAclRuleUpdateQueue中取出条目，一次最多取出30个条目
* @rturn msg:消息内容，num:条目个数
*  创建者  于士超   2019.11.27
****************************************************************************/
func msgRuleUpdateDequeueItem()(string,int32){
	/*如果定时器没有创建，那么创建定时器，定时广播消息*/
	var msg string
	var num int32
	msgAclQueueMutex.Lock()	

	for index,_:=range msgAclRuleUpdateQueue{
		msg = msg + "," + index
		delete(msgAclRuleUpdateQueue,index)	
		num++;
		if	num>=dequeueItemNum{
			break;
		}
	}	
	msgAclQueueMutex.Unlock()
	msg=strings.Trim(msg,",")	
	return msg,num
}
/***************************************************************************
* 函数说明: 从msgAclTableDeleteQueue中取出条目，一次最多取出30个条目
* @rturn msg:消息内容，num:条目个数
*  创建者  于士超   2019.11.27
****************************************************************************/
func msgTableDeleteDequeueItem()(string,int32){
	/*如果定时器没有创建，那么创建定时器，定时广播消息*/
	var msg string
	var num int32
	msgAclQueueMutex.Lock()	

	for index,_:=range msgAclTableDeleteQueue{
		msg = msg + "," + index
		delete(msgAclTableDeleteQueue,index)	
		num++;
		if	num>=dequeueItemNum{
			break;
		}
	}		
	msgAclQueueMutex.Unlock()
	msg=strings.Trim(msg,",")	
	return msg,num
}
/***************************************************************************
* 函数说明: 从msgAclTableBindQueue取出条目，一次最多取出30个条目
* @rturn msg:消息内容，num:条目个数
*  创建者  于士超   2019.11.27
****************************************************************************/
func msgTableBindDequeueItem()(string,int32){
	/*如果定时器没有创建，那么创建定时器，定时广播消息*/
	var msg string
	var num int32
	msgAclQueueMutex.Lock()	

	for index,_:=range msgAclTableBindQueue{
		msg = msg + "," + index
		delete(msgAclTableBindQueue,index)	
		num++;
		if	num>=dequeueItemNum{
			break;
		}
	}		
	msgAclQueueMutex.Unlock()
	msg=strings.Trim(msg,",")	
	return msg,num
}
/***************************************************************************
* 函数说明: 从msgAclTableUnbindQueue取出条目，一次最多取出30个条目
* @rturn msg:消息内容，num:条目个数
*  创建者  于士超   2019.11.27
****************************************************************************/
func msgTableUnbindDequeueItem()(string,int32){
	/*如果定时器没有创建，那么创建定时器，定时广播消息*/
	var msg string
	var num int32
	msgAclQueueMutex.Lock()	

	for index,_:=range msgAclTableUnbindQueue{
		msg = msg + "," + index
		delete(msgAclTableUnbindQueue,index)	
		num++;
		if	num>=dequeueItemNum{
			break;
		}
	}	
	msgAclQueueMutex.Unlock()
	msg=strings.Trim(msg,",")	
	return msg,num
}
/***************************************************************************
* 函数说明: 向msgAclRuleDeleteQueue中插入条目
*  创建者  于士超   2019.11.27
****************************************************************************/

func msgRuleDeleteQueueInsert(key string,value string) {
	/*如果定时器没有创建，那么创建定时器，定时广播消息*/
	if	msgTimer == nil{
		msgTimer = time.AfterFunc(time.Second * 5, func(){
			msgTimerHandle()			
		})
	} else if msgTimerStatus ==1{
		msgTimer.Reset(time.Second * 5)
		msgTimerStatus =0;
	}

	msgAclQueueMutex.Lock()	
	if msgAclRuleDeleteQueue == nil{
		msgAclRuleDeleteQueue=make(map[string]string)
	}
	msgAclRuleDeleteQueue[key]=value
	msgAclQueueMutex.Unlock()
}
/***************************************************************************
* 函数说明: 向msgAclRuleUpdateQueue中插入条目
*  创建者  于士超   2019.11.27
****************************************************************************/
func msgRuleUpdateQueueInsert(key string,value string) {
	/*如果定时器没有创建，那么创建定时器，定时广播消息*/
	if	msgTimer == nil{
		msgTimer = time.AfterFunc(time.Second * 5, func(){
			msgTimerHandle()			
		})
	}else if msgTimerStatus ==1{
		msgTimer.Reset(time.Second * 5)
		msgTimerStatus =0;
	}

	msgAclQueueMutex.Lock()	
	if msgAclRuleUpdateQueue == nil{
		msgAclRuleUpdateQueue=make(map[string]string)
	}
	msgAclRuleUpdateQueue[key]=value
	msgAclQueueMutex.Unlock()
}
/***************************************************************************
* 函数说明: 向msgAclTableDeleteQueue中插入条目
*  创建者  于士超   2019.11.27
****************************************************************************/
func msgTableDeleteQueueInsert(key string,value string) {
	/*如果定时器没有创建，那么创建定时器，定时广播消息*/
	if	msgTimer == nil{
		msgTimer = time.AfterFunc(time.Second * 5, func(){
			msgTimerHandle()			
		})
	}else if msgTimerStatus ==1{
		msgTimer.Reset(time.Second * 5)
		msgTimerStatus =0;
	}

	msgAclQueueMutex.Lock()	
	if msgAclTableDeleteQueue == nil{
		msgAclTableDeleteQueue=make(map[string]string)
	}
	msgAclTableDeleteQueue[key]=value
	msgAclQueueMutex.Unlock()
}
/***************************************************************************
* 函数说明: 向msgAclTableBindQueue中插入条目
*  创建者  于士超   2019.11.27
****************************************************************************/
func msgTableBindQueueInsert(key string,value string) {
	/*如果定时器没有创建，那么创建定时器，定时广播消息*/
	if	msgTimer == nil{
		msgTimer = time.AfterFunc(time.Second * 5, func(){
			msgTimerHandle()			
		})
	}else if msgTimerStatus ==1{
		msgTimer.Reset(time.Second * 5)
		msgTimerStatus =0;
	}

	msgAclQueueMutex.Lock()	
	if msgAclTableBindQueue == nil{
		msgAclTableBindQueue=make(map[string]string)
	}
	msgAclTableBindQueue[key]=value
	msgAclQueueMutex.Unlock()
}
/***************************************************************************
* 函数说明: 向msgAclTableUnbindQueue中插入条目
*  创建者  于士超   2019.11.27
****************************************************************************/
func msgTableUnbindQueueInsert(key string,value string) {
	/*如果定时器没有创建，那么创建定时器，定时广播消息*/
	if	msgTimer == nil{
		msgTimer = time.AfterFunc(time.Second * 5, func(){
			msgTimerHandle()			
		})
	}else if msgTimerStatus ==1{
		msgTimer.Reset(time.Second * 5)
		msgTimerStatus =0;
	}

	msgAclQueueMutex.Lock()	
	if msgAclTableUnbindQueue == nil{
		msgAclTableUnbindQueue=make(map[string]string)
	}
	msgAclTableUnbindQueue[key]=value
	msgAclQueueMutex.Unlock()
}

/***************************************************************************
* 函数说明: 删除一个ACL规则条目，并向消息队列中插入一个ACL规则删除的条目
*  创建者  于士超   2019.11.27
****************************************************************************/
func redisDelRuleKey(value string){
	hlen := aclRedisClient.HLen(value).Val()
	if hlen !=0{
		m,_ := aclRedisClient.HGetAll(value).Result()

		for index,_ := range m{
			aclRedisClient.HDel(value,index)	
		}
		msgRuleDeleteQueueInsert(value,RULE_DELETE)	/*更新消息删除*/
	}	
}
/***************************************************************************
* 函数说明: 删除一个ACL表项，并向消息队列中插入一个ACL表项删除的条目
*  创建者  于士超   2019.11.27
****************************************************************************/
func redisDelTableKey(value string){
	hlen := aclRedisClient.HLen(value).Val()
	if hlen !=0{
		m,_ := aclRedisClient.HGetAll(value).Result()

		for index,_ := range m{
			aclRedisClient.HDel(value,index)	
		}
		msgTableDeleteQueueInsert(value,TABLE_DELETE)	/*更新消息删除*/
	}	
}
/***************************************************************************
* 函数说明: 从acl表中，查找step值
*  创建者  于士超   2019.11.27
****************************************************************************/
func aclNumberStepQuery(key string,field string) int{

	var step int =5
	/*首先从数据库获取该number对应的步长*/
	val :=aclRedisClient.HGet(key,field).Val()
	if val != "" {
		step, _ = strconv.Atoi(val)/*转化为整形*/
	}
	return step	
}
/***************************************************************************
* 函数说明: 从数据库中查找一张表中一个域的值
*  创建者  于士超   2019.11.27
****************************************************************************/
func redisGetKeyFiled(key string,field string) string{

	/*首先从数据库获取该number对应的步长*/
	val :=aclRedisClient.HGet(key,field).Val()
	return val	
}
/***************************************************************************
* 函数说明: 设置一张表中一个域的值
*  创建者  于士超   2019.11.27
****************************************************************************/
func redisSetKeyFiledValue(key string,field string,value string){

	/*首先从数据库获取该number对应的步长*/
	aclRedisClient.HSet(key,field,value)	
}
func aclTableMsgCopy(outMsg *[ACL_TABLE_LAST_ITEM]aclFieldTable,acl *pb.ACL_TABLE){

	outMsg[ACL_TABLE_NUMBER].field 			= fmt.Sprint(ACL_NUMBER)
	outMsg[ACL_TABLE_NUMBER].value 			= fmt.Sprint(acl.AclNumber)
	outMsg[ACL_TYPE].field 					= fmt.Sprint(ACL_TYPE)
	outMsg[ACL_TYPE].value 					= fmt.Sprint(acl.AclType)		
	outMsg[ACL_TABLE_ACTION].field 			= fmt.Sprint(ACL_ACTION)
	outMsg[ACL_TABLE_ACTION].value 			= fmt.Sprint(acl.Action)
	outMsg[ACL_TABLE_STEP].field 			= fmt.Sprint(ACL_STEP)
	outMsg[ACL_TABLE_STEP].value 			= fmt.Sprint(acl.Step)	
	outMsg[ACL_BIND_TYPE].field 			= fmt.Sprint(ACL_BIND_TYPE)
	outMsg[ACL_BIND_TYPE].value 			= fmt.Sprint(acl.BindType)
	outMsg[ACL_BIND_TYPE_NUMBER].field 		= fmt.Sprint(ACL_BIND_TYPE_NUMBER)
	outMsg[ACL_BIND_TYPE_NUMBER].value 		= fmt.Sprint(acl.BindTypeNumber)

}

func basicMsgCopy(outMsg *[ACL_LAST_ITEM]aclFieldTable,acl *pb.BASIC_ACLRULE){
			
	outMsg[ACL_NUMBER].field 				= fmt.Sprint(ACL_NUMBER)
	outMsg[ACL_NUMBER].value 				= fmt.Sprint(acl.AclNumber)
	outMsg[ACL_STEP].field 					= fmt.Sprint(ACL_STEP)
	outMsg[ACL_STEP].value 					= fmt.Sprint(acl.Step)	
	outMsg[ACL_RULE_ID].field	 			= fmt.Sprint(ACL_RULE_ID)
	outMsg[ACL_RULE_ID].value 				= fmt.Sprint(acl.RuleId)
	outMsg[ACL_RULE_ID1].field 				= fmt.Sprint(ACL_RULE_ID1)
	outMsg[ACL_RULE_ID1].value 				= fmt.Sprint(acl.RuleId1)
	outMsg[ACL_RULE_ID2].field 				= fmt.Sprint(ACL_RULE_ID2)
	outMsg[ACL_RULE_ID2].value 				= fmt.Sprint(acl.RuleId2)		
	outMsg[ACL_ACTION].field 				= fmt.Sprint(ACL_ACTION)
	outMsg[ACL_ACTION].value 				= fmt.Sprint(acl.Action)
	outMsg[ACL_CMDLINE].field 				= fmt.Sprint(ACL_CMDLINE)
	outMsg[ACL_CMDLINE].value 				= fmt.Sprint(acl.Cmdline)	
	outMsg[ACL_FRAGMENT_TYPE].field 		= fmt.Sprint(ACL_FRAGMENT_TYPE)
	outMsg[ACL_FRAGMENT_TYPE].value 		= fmt.Sprint(acl.Match.GetFragmentType())
	outMsg[ACL_TIME_RANGE].field 			= fmt.Sprint(ACL_TIME_RANGE)
	outMsg[ACL_TIME_RANGE].value 			= fmt.Sprint(acl.Match.GetTimeRange())
	outMsg[ACL_SRC_IP_MASK].field 			= fmt.Sprint(ACL_SRC_IP_MASK)
	outMsg[ACL_SRC_IP_MASK].value 			= fmt.Sprint(acl.Match.GetSrcIpMask())
	outMsg[ACL_VPN_INSTANCE].field 			= fmt.Sprint(ACL_VPN_INSTANCE)
	outMsg[ACL_VPN_INSTANCE].value 			= fmt.Sprint(acl.Match.GetVpnInst())

}

func advanceMsgCopy(temp *[ACL_LAST_ITEM]aclFieldTable,acl *pb.ADVANCE_ACLRULE){
			
	temp[ACL_NUMBER].field 					= fmt.Sprint(ACL_NUMBER)
	temp[ACL_NUMBER].value 					= fmt.Sprint(acl.AclNumber)
	temp[ACL_STEP].field 					= fmt.Sprint(ACL_STEP)
	temp[ACL_STEP].value 					= fmt.Sprint(acl.Step)	
	temp[ACL_RULE_ID].field 				= fmt.Sprint(ACL_RULE_ID)
	temp[ACL_RULE_ID].value 				= fmt.Sprint(acl.RuleId)
	temp[ACL_RULE_ID1].field 				= fmt.Sprint(ACL_RULE_ID1)
	temp[ACL_RULE_ID1].value 				= fmt.Sprint(acl.RuleId1)
	temp[ACL_RULE_ID2].field 				= fmt.Sprint(ACL_RULE_ID2)
	temp[ACL_RULE_ID2].value 				= fmt.Sprint(acl.RuleId2)	
	temp[ACL_ACTION].field 					= fmt.Sprint(ACL_ACTION)
	temp[ACL_ACTION].value 					= fmt.Sprint(acl.Action)
	temp[ACL_CMDLINE].field 				= fmt.Sprint(ACL_CMDLINE)
	temp[ACL_CMDLINE].value 				= fmt.Sprint(acl.Cmdline)	
	temp[ACL_FRAGMENT_TYPE].field 			= fmt.Sprint(ACL_FRAGMENT_TYPE)
	temp[ACL_FRAGMENT_TYPE].value 			= fmt.Sprint(acl.Match.GetFragmentType())
	temp[ACL_TIME_RANGE].field 				= fmt.Sprint(ACL_TIME_RANGE)
	temp[ACL_TIME_RANGE].value 				= fmt.Sprint(acl.Match.GetTimeRange())
	temp[ACL_SRC_IP_MASK].field 			= fmt.Sprint(ACL_SRC_IP_MASK)
	temp[ACL_SRC_IP_MASK].value 			= fmt.Sprint(acl.Match.GetSrcIpMask())
	temp[ACL_DST_IP_MASK].field 			= fmt.Sprint(ACL_DST_IP_MASK)
	temp[ACL_DST_IP_MASK].value 			= fmt.Sprint(acl.Match.GetDstIpMask())
	temp[ACL_PROTO_TYPE].field 				= fmt.Sprint(ACL_PROTO_TYPE)
	temp[ACL_PROTO_TYPE].value 				= fmt.Sprint(acl.Match.GetProtoType())
	temp[ACL_SRC_PORT_MIN].field 			= fmt.Sprint(ACL_SRC_PORT_MIN)
	temp[ACL_SRC_PORT_MIN].value 			= fmt.Sprint(acl.Match.GetSrcPortMin())	
	temp[ACL_SRC_PORT_MAX].field 			= fmt.Sprint(ACL_SRC_PORT_MAX)
	temp[ACL_SRC_PORT_MAX].value 			= fmt.Sprint(acl.Match.GetSrcPortMax())
	temp[ACL_DST_PORT_MIN].field 			= fmt.Sprint(ACL_DST_PORT_MIN)
	temp[ACL_DST_PORT_MIN].value 			= fmt.Sprint(acl.Match.GetDstPortMin())	
	temp[ACL_DST_PORT_MAX].field 			= fmt.Sprint(ACL_DST_PORT_MAX)
	temp[ACL_DST_PORT_MAX].value 			= fmt.Sprint(acl.Match.GetDstPortMax())
	temp[ACL_ICMP_TYPE].field 				= fmt.Sprint(ACL_ICMP_TYPE)
	temp[ACL_ICMP_TYPE].value 				= fmt.Sprint(acl.Match.GetIcmpType())
	temp[ACL_ICMP_CODE].field 				= fmt.Sprint(ACL_ICMP_CODE)
	temp[ACL_ICMP_CODE].value 				= fmt.Sprint(acl.Match.GetIcmpCode())
	temp[ACL_IGMP_TYPE].field 				= fmt.Sprint(ACL_IGMP_TYPE)
	temp[ACL_IGMP_TYPE].value 				= fmt.Sprint(acl.Match.GetIgmpType())
	temp[ACL_TCP_FLAG].field 				= fmt.Sprint(ACL_TCP_FLAG)
	temp[ACL_TCP_FLAG].value 				= fmt.Sprint(acl.Match.GetTcpFlag())
	temp[ACL_DSCP].field 					= fmt.Sprint(ACL_DSCP)
	temp[ACL_DSCP].value 					= fmt.Sprint(acl.Match.GetDscp())
	temp[ACL_TOS].field 					= fmt.Sprint(ACL_TOS)
	temp[ACL_TOS].value 					= fmt.Sprint(acl.Match.GetTos())
	temp[ACL_PRECEDENCE].field 				= fmt.Sprint(ACL_PRECEDENCE)
	temp[ACL_PRECEDENCE].value 				= fmt.Sprint(acl.Match.GetPrecedence())
	temp[ACL_TTL_EXPIRED].field 			= fmt.Sprint(ACL_TTL_EXPIRED)
	temp[ACL_TTL_EXPIRED].value 			= fmt.Sprint(acl.Match.GetTtlExpired())
	temp[ACL_VPN_INSTANCE].field 			= fmt.Sprint(ACL_VPN_INSTANCE)
	temp[ACL_VPN_INSTANCE].value 			= fmt.Sprint(acl.Match.GetVpnInst())

}
func l2MsgCopy(temp *[ACL_LAST_ITEM]aclFieldTable,acl *pb.L2_ACLRULE){
			
	temp[ACL_NUMBER].field 					= fmt.Sprint(ACL_NUMBER)
	temp[ACL_NUMBER].value 					= fmt.Sprint(acl.AclNumber)
	temp[ACL_STEP].field 					= fmt.Sprint(ACL_STEP)
	temp[ACL_STEP].value 					= fmt.Sprint(acl.Step)	
	temp[ACL_RULE_ID].field 				= fmt.Sprint(ACL_RULE_ID)
	temp[ACL_RULE_ID].value 				= fmt.Sprint(acl.RuleId)
	temp[ACL_RULE_ID1].field 				= fmt.Sprint(ACL_RULE_ID1)
	temp[ACL_RULE_ID1].value 				= fmt.Sprint(acl.RuleId1)
	temp[ACL_RULE_ID2].field 				= fmt.Sprint(ACL_RULE_ID2)
	temp[ACL_RULE_ID2].value 				= fmt.Sprint(acl.RuleId2)	
	temp[ACL_ACTION].field 					= fmt.Sprint(ACL_ACTION)
	temp[ACL_ACTION].value 					= fmt.Sprint(acl.Action)
	temp[ACL_CMDLINE].field 				= fmt.Sprint(ACL_CMDLINE)
	temp[ACL_CMDLINE].value 				= fmt.Sprint(acl.Cmdline)	
	temp[ACL_TIME_RANGE].field 				= fmt.Sprint(ACL_TIME_RANGE)
	temp[ACL_TIME_RANGE].value 				= fmt.Sprint(acl.Match.GetTimeRange())
	temp[ACL_SRC_MAC_MASK].field 			= fmt.Sprint(ACL_SRC_MAC_MASK)
	temp[ACL_SRC_MAC_MASK].value 			= fmt.Sprint(acl.Match.GetSrcMacMask())
	temp[ACL_DST_MAC_MASK].field 			= fmt.Sprint(ACL_DST_MAC_MASK)
	temp[ACL_DST_MAC_MASK].value 			= fmt.Sprint(acl.Match.GetDstMacMask())
	temp[ACL_ETHER_TYPE].field 				= fmt.Sprint(ACL_ETHER_TYPE)
	temp[ACL_ETHER_TYPE].value 				= fmt.Sprint(acl.Match.GetEthertype())
	temp[ACL_PACKET_TYPE].field 			= fmt.Sprint(ACL_PACKET_TYPE)
	temp[ACL_PACKET_TYPE].value 			= fmt.Sprint(acl.Match.GetPacketType())
	temp[ACL_VLAN_ID].field 				= fmt.Sprint(ACL_VLAN_ID)
	temp[ACL_VLAN_ID].value 				= fmt.Sprint(acl.Match.GetVlanId())
	temp[ACL_DOT1P].field 					= fmt.Sprint(ACL_DOT1P)
	temp[ACL_DOT1P].value 					= fmt.Sprint(acl.Match.GetDot1P())
	temp[ACL_IN_VLAN_ID].field 				= fmt.Sprint(ACL_IN_VLAN_ID)
	temp[ACL_IN_VLAN_ID].value 				= fmt.Sprint(acl.Match.GetInvlan())
	temp[ACL_IN_DOT1P].field 				= fmt.Sprint(ACL_IN_DOT1P)
	temp[ACL_IN_DOT1P].value 				= fmt.Sprint(acl.Match.GetIndot1P())
	temp[ACL_DTAG].field 					= fmt.Sprint(ACL_DTAG)
	temp[ACL_DTAG].value			 		= fmt.Sprint(acl.Match.GetDtag())
}
func arpMsgCopy(temp *[ACL_LAST_ITEM]aclFieldTable,acl *pb.ARP_ACLRULE){

	temp[ACL_NUMBER].field 					= fmt.Sprint(ACL_NUMBER)
	temp[ACL_NUMBER].value 					= fmt.Sprint(acl.AclNumber)
	temp[ACL_STEP].field 					= fmt.Sprint(ACL_STEP)
	temp[ACL_STEP].value 					= fmt.Sprint(acl.Step)	
	temp[ACL_RULE_ID].field 				= fmt.Sprint(ACL_RULE_ID)
	temp[ACL_RULE_ID].value 				= fmt.Sprint(acl.RuleId)
	temp[ACL_RULE_ID1].field 				= fmt.Sprint(ACL_RULE_ID1)
	temp[ACL_RULE_ID1].value 				= fmt.Sprint(acl.RuleId1)
	temp[ACL_RULE_ID2].field 				= fmt.Sprint(ACL_RULE_ID2)
	temp[ACL_RULE_ID2].value 				= fmt.Sprint(acl.RuleId2)		
	temp[ACL_ACTION].field 					= fmt.Sprint(ACL_ACTION)
	temp[ACL_ACTION].value 					= fmt.Sprint(acl.Action)
	temp[ACL_CMDLINE].field 				= fmt.Sprint(ACL_CMDLINE)
	temp[ACL_CMDLINE].value 				= fmt.Sprint(acl.Cmdline)	
	temp[ACL_SRC_MAC_MASK].field 			= fmt.Sprint(ACL_SRC_MAC_MASK)
	temp[ACL_SRC_MAC_MASK].value 			= fmt.Sprint(acl.Match.GetSrcMacMask())
	temp[ACL_DST_MAC_MASK].field 			= fmt.Sprint(ACL_DST_MAC_MASK)
	temp[ACL_DST_MAC_MASK].value 			= fmt.Sprint(acl.Match.GetDstMacMask())
	temp[ACL_SRC_IP_MASK].field 			= fmt.Sprint(ACL_SRC_IP_MASK)
	temp[ACL_SRC_IP_MASK].value 			= fmt.Sprint(acl.Match.GetSrcIpMask())
	temp[ACL_DST_IP_MASK].field 			= fmt.Sprint(ACL_DST_IP_MASK)
	temp[ACL_DST_IP_MASK].value 			= fmt.Sprint(acl.Match.GetDstIpMask())
	temp[ACL_TIME_RANGE].field 				= fmt.Sprint(ACL_TIME_RANGE)
	temp[ACL_TIME_RANGE].value 				= fmt.Sprint(acl.Match.GetTimeRange())	

}
func mixedMsgCopy(temp *[ACL_LAST_ITEM]aclFieldTable,acl *pb.MIXED_ACLRULE){

	/*RULE 条目接收到本地*/
	temp[ACL_NUMBER].field 					= fmt.Sprint(ACL_NUMBER)
	temp[ACL_NUMBER].value 					= fmt.Sprint(acl.AclNumber)
	temp[ACL_STEP].field 					= fmt.Sprint(ACL_STEP)
	temp[ACL_STEP].value 					= fmt.Sprint(acl.Step)	
	temp[ACL_RULE_ID].field 				= fmt.Sprint(ACL_RULE_ID)
	temp[ACL_RULE_ID].value 				= fmt.Sprint(acl.RuleId)
	temp[ACL_RULE_ID1].field 				= fmt.Sprint(ACL_RULE_ID1)
	temp[ACL_RULE_ID1].value 				= fmt.Sprint(acl.RuleId1)
	temp[ACL_RULE_ID2].field 				= fmt.Sprint(ACL_RULE_ID2)
	temp[ACL_RULE_ID2].value 				= fmt.Sprint(acl.RuleId2)	
	temp[ACL_ACTION].field 					= fmt.Sprint(ACL_ACTION)
	temp[ACL_ACTION].value 					= fmt.Sprint(acl.Action)
	temp[ACL_CMDLINE].field 				= fmt.Sprint(ACL_CMDLINE)
	temp[ACL_CMDLINE].value 				= fmt.Sprint(acl.Cmdline)	
	temp[ACL_TIME_RANGE].field 				= fmt.Sprint(ACL_TIME_RANGE)
	temp[ACL_TIME_RANGE].value 				= fmt.Sprint(acl.Match.GetTimeRange())
	temp[ACL_SRC_MAC_MASK].field 			= fmt.Sprint(ACL_SRC_MAC_MASK)
	temp[ACL_SRC_MAC_MASK].value 			= fmt.Sprint(acl.Match.GetSrcMacMask())
	temp[ACL_DST_MAC_MASK].field 			= fmt.Sprint(ACL_DST_MAC_MASK)
	temp[ACL_DST_MAC_MASK].value 			= fmt.Sprint(acl.Match.GetDstMacMask())
	temp[ACL_ETHER_TYPE].field 				= fmt.Sprint(ACL_ETHER_TYPE)
	temp[ACL_ETHER_TYPE].value 				= fmt.Sprint(acl.Match.GetEthertype())
	temp[ACL_PACKET_TYPE].field 			= fmt.Sprint(ACL_PACKET_TYPE)
	temp[ACL_PACKET_TYPE].value 			= fmt.Sprint(acl.Match.GetPacketType())
	temp[ACL_VLAN_ID].field 				= fmt.Sprint(ACL_VLAN_ID)
	temp[ACL_VLAN_ID].value 				= fmt.Sprint(acl.Match.GetVlanId())
	temp[ACL_DOT1P].field 					= fmt.Sprint(ACL_DOT1P)
	temp[ACL_DOT1P].value 					= fmt.Sprint(acl.Match.GetDot1P())
	temp[ACL_IN_VLAN_ID].field 				= fmt.Sprint(ACL_IN_VLAN_ID)
	temp[ACL_IN_VLAN_ID].value 				= fmt.Sprint(acl.Match.GetInvlan())
	temp[ACL_IN_DOT1P].field 				= fmt.Sprint(ACL_IN_DOT1P)
	temp[ACL_IN_DOT1P].value 				= fmt.Sprint(acl.Match.GetIndot1P())
	temp[ACL_DTAG].field 					= fmt.Sprint(ACL_DTAG)
	temp[ACL_DTAG].value 					= fmt.Sprint(acl.Match.GetDtag())	
	temp[ACL_SRC_IP_MASK].field 			= fmt.Sprint(ACL_SRC_IP_MASK)
	temp[ACL_SRC_IP_MASK].value 			= fmt.Sprint(acl.Match.GetSrcIpMask())
	temp[ACL_DST_IP_MASK].field 			= fmt.Sprint(ACL_DST_IP_MASK)
	temp[ACL_DST_IP_MASK].value 			= fmt.Sprint(acl.Match.GetDstIpMask())
	temp[ACL_PROTO_TYPE].field 				= fmt.Sprint(ACL_PROTO_TYPE)
	temp[ACL_PROTO_TYPE].value 				= fmt.Sprint(acl.Match.GetProtoType())
	temp[ACL_SRC_PORT_MIN].field 			= fmt.Sprint(ACL_SRC_PORT_MIN)
	temp[ACL_SRC_PORT_MIN].value 			= fmt.Sprint(acl.Match.GetSrcPortMin())	
	temp[ACL_SRC_PORT_MAX].field 			= fmt.Sprint(ACL_SRC_PORT_MAX)
	temp[ACL_SRC_PORT_MAX].value 			= fmt.Sprint(acl.Match.GetSrcPortMax())
	temp[ACL_DST_PORT_MIN].field 			= fmt.Sprint(ACL_DST_PORT_MIN)
	temp[ACL_DST_PORT_MIN].value 			= fmt.Sprint(acl.Match.GetDstPortMin())	
	temp[ACL_DST_PORT_MAX].field 			= fmt.Sprint(ACL_DST_PORT_MAX)
	temp[ACL_DST_PORT_MAX].value 			= fmt.Sprint(acl.Match.GetDstPortMax())
	temp[ACL_ICMP_TYPE].field 				= fmt.Sprint(ACL_ICMP_TYPE)
	temp[ACL_ICMP_TYPE].value 				= fmt.Sprint(acl.Match.GetIcmpType())
	temp[ACL_ICMP_CODE].field 				= fmt.Sprint(ACL_ICMP_CODE)
	temp[ACL_ICMP_CODE].value 				= fmt.Sprint(acl.Match.GetIcmpCode())
	temp[ACL_IGMP_TYPE].field 				= fmt.Sprint(ACL_IGMP_TYPE)
	temp[ACL_IGMP_TYPE].value 				= fmt.Sprint(acl.Match.GetIgmpType())
	temp[ACL_TCP_FLAG].field 				= fmt.Sprint(ACL_TCP_FLAG)
	temp[ACL_TCP_FLAG].value 				= fmt.Sprint(acl.Match.GetTcpFlag())
	temp[ACL_DSCP].field 					= fmt.Sprint(ACL_DSCP)
	temp[ACL_DSCP].value 					= fmt.Sprint(acl.Match.GetDscp())
	temp[ACL_TOS].field 					= fmt.Sprint(ACL_TOS)
	temp[ACL_TOS].value 					= fmt.Sprint(acl.Match.GetTos())
	temp[ACL_PRECEDENCE].field 				= fmt.Sprint(ACL_PRECEDENCE)
	temp[ACL_PRECEDENCE].value 				= fmt.Sprint(acl.Match.GetPrecedence())
	temp[ACL_TTL_EXPIRED].field 			= fmt.Sprint(ACL_TTL_EXPIRED)
	temp[ACL_TTL_EXPIRED].value 			= fmt.Sprint(acl.Match.GetTtlExpired())

}
/***************************************************************************
* 函数说明: 创建基本acl rule
*  创建者  于士超   2019.11.27
****************************************************************************/

//ACL ADVANCE Acl Rule相关配置
func (self *AclConfig) CreateBasicACLRule(ctx context.Context, acl *pb.BASIC_ACLRULE) (*pb.CommonRespHdr, error) {

	var temp [ACL_LAST_ITEM]aclFieldTable
	var rtnCode pb.CommonRespHdr
	var tableKey string
	var entryField string
	var ruleKey string
	rtnCode.RetCode = 0
	
	basicMsgCopy(&temp,acl)	
	tableKey =fmt.Sprintf("acl_table|acl_basic_%s",temp[ACL_NUMBER].value)
	entryField = "ACL_ENTRIES"

	
	/*如果命令不带有RULE ID，那么根据默认步长自定义RULE ID*/
	if temp[ACL_RULE_ID].value =="-1"{			
		step :=aclNumberStepQuery(tableKey,temp[ACL_STEP].field)			
		for i :=step;i<10000;i=(i+step) {//10000作为测试使用，后续更改
			hlen := aclRedisClient.HLen(fmt.Sprintf("ACL_BASIC_%s|RULE_%d",temp[ACL_NUMBER].value,i)).Val()
			if hlen ==0 {
				temp[ACL_RULE_ID].value = fmt.Sprint(i)	
				break;				
			}
		}		
	}
	ruleKey =fmt.Sprintf("ACL_BASIC_%s|RULE_%s",temp[ACL_NUMBER].value,temp[ACL_RULE_ID].value)

	redisDelRuleKey(ruleKey)				
	hlen := aclRedisClient.HLen(ruleKey).Val()
	if hlen ==0{
		for i :=3;i<len(temp);i++{
			if temp[i].value !="" && temp[i].value !="-1"{
				aclRedisClient.HSet(ruleKey, temp[i].field, temp[i].value).Err()
			}
		}			
	}
	entries:=redisGetKeyFiled(tableKey,entryField)
	if entries != ""{
		r := strings.Contains(entries,ruleKey)
		if r== false{
			entries = entries + "," + ruleKey
			entries = strings.Trim(entries,",")
			aclRedisClient.HSet(tableKey,entryField,entries)				
		}
	}else{
		aclRedisClient.HSet(tableKey,entryField,ruleKey)				
	}
	msgRuleUpdateQueueInsert(ruleKey,RULE_UPDATE)		

	return &rtnCode, nil
}

/***************************************************************************
* 函数说明: 删除基本acl规则
*  创建者  于士超   2019.11.27
****************************************************************************/
func (self *AclConfig) DeleteBasicACLRule(ctx context.Context, acl *pb.BASIC_ACLRULE) (*pb.CommonRespHdr, error) {

	var temp [ACL_LAST_ITEM]aclFieldTable
	var rtnCode pb.CommonRespHdr
	var tableKey string
	var tableField string
	rtnCode.RetCode = 0

	basicMsgCopy(&temp,acl)	
	tableKey = fmt.Sprintf("acl_table|acl_basic_%s",temp[ACL_NUMBER].value)
	tableField = "ACL_ENTRIES"

	
	entries:=redisGetKeyFiled(tableKey,tableField)
	temp1 := strings.Split(entries,",")
	if acl.RuleId2 == -1{
		redisDelRuleKey(fmt.Sprintf("ACL_BASIC_%s|RULE_%s",temp[ACL_NUMBER].value,temp[ACL_RULE_ID1].value))
		if entries != ""{
			for i,value :=range temp1{
				if value ==fmt.Sprintf("ACL_BASIC_%s|RULE_%s",temp[ACL_NUMBER].value,temp[ACL_RULE_ID1].value){
					temp1 = append(temp1[:i], temp1[i+1:]...)
				} 					
			}
		}			
	}else{

		for i :=acl.RuleId2;i>=acl.RuleId1;i--{
			redisDelRuleKey(fmt.Sprintf("ACL_BASIC_%s|RULE_%d",temp[ACL_NUMBER].value,i))
			if entries != ""{
				for j,value :=range temp1{
					if value ==fmt.Sprintf("ACL_BASIC_%s|RULE_%d",temp[ACL_NUMBER].value,i){
						temp1 = append(temp1[:j], temp1[j+1:]...)
					} 					
				}
			}
		}					
	}
	newEntries :=strings.Join(temp1,",")
	aclRedisClient.HSet(tableKey,tableField,newEntries)		
	return &rtnCode, nil
}
/***************************************************************************
* 函数说明: 创建高级acl规则
*  创建者  于士超   2019.11.27
****************************************************************************/
func (self *AclConfig) CreateExtendACLRule(ctx context.Context, acl *pb.ADVANCE_ACLRULE) (*pb.CommonRespHdr, error) {

	var temp [ACL_LAST_ITEM]aclFieldTable
	var rtnCode pb.CommonRespHdr
	var tableKey string
	var ruleKey string
	var entryField string
	rtnCode.RetCode = 0

	/*RULE 条目接收到本地*/
	advanceMsgCopy(&temp,acl)
	tableKey = 	fmt.Sprintf("acl_table|acl_advance_%s",temp[ACL_NUMBER].value)
	entryField ="ACL_ENTRIES"


	
	if temp[ACL_RULE_ID].value =="-1"{			
		/*首先从数据库获取该number对应的步长*/
		step :=aclNumberStepQuery(tableKey,temp[ACL_STEP].field)			
		//10000作为测试使用，后续更改
		for i :=step;i<10000;i=(i+step) {
			hlen := aclRedisClient.HLen(fmt.Sprintf("ACL_ADVANCE_%s|RULE_%d",temp[ACL_NUMBER].value,i)).Val()
			if hlen ==0 {
				temp[ACL_RULE_ID].value = fmt.Sprint(i)	
				break;				
			}
		}
		
	}
	ruleKey = fmt.Sprintf("ACL_ADVANCE_%s|RULE_%s",temp[ACL_NUMBER].value,temp[ACL_RULE_ID].value)
	/*首先判断是否存在该条目，如果存在，那么先删除，后下发；如果不存在，那么直接下发配置*/
	redisDelRuleKey(ruleKey)
	hlen := aclRedisClient.HLen(ruleKey).Val()
	if hlen ==0{/*下发配置到CONFIG DB*/
		for i :=3;i<len(temp);i++{
			if temp[i].value !="" && temp[i].value !="-1"{
				aclRedisClient.HSet(ruleKey, temp[i].field, temp[i].value).Err()
			}
		}
	}

	entries:=redisGetKeyFiled(tableKey,entryField)
	if entries != ""{
		r := strings.Contains(entries,ruleKey)
		if r== false{
			entries = entries + ","+ ruleKey
			entries =strings.Trim(entries,",")
			aclRedisClient.HSet(tableKey,entryField,entries)				
		}
	}else{
		aclRedisClient.HSet(tableKey,entryField,ruleKey)				
	}

	msgRuleUpdateQueueInsert(ruleKey,RULE_UPDATE)					


	return &rtnCode, nil	
}
/***************************************************************************
* 函数说明: 删除高级acl规则
*  创建者  于士超   2019.11.27
****************************************************************************/
func (self *AclConfig) DeleteExtendACLRule(ctx context.Context, acl *pb.ADVANCE_ACLRULE) (*pb.CommonRespHdr, error) {

	var temp [ACL_LAST_ITEM]aclFieldTable
	var rtnCode pb.CommonRespHdr
	var tableKey string
	var tableField string
	rtnCode.RetCode = 0

	/*RULE 条目接收到本地*/
	advanceMsgCopy(&temp,acl)
	tableKey = fmt.Sprintf("acl_table|acl_advance_%s",temp[ACL_NUMBER].value)
	tableField = "ACL_ENTRIES"

	fmt.Println(acl.AclNumber,acl.Step,acl.RuleId,acl.RuleId1,acl.RuleId2,acl.Action)


	entries:=redisGetKeyFiled(tableKey,tableField)
	temp1 := strings.Split(entries,",")		
	if acl.RuleId2 == -1{
		redisDelRuleKey(fmt.Sprintf("ACL_ADVANCE_%s|RULE_%s",temp[ACL_NUMBER].value,temp[ACL_RULE_ID1].value))
		if entries != ""{
			for i,value :=range temp1{
				if value ==fmt.Sprintf("ACL_ADVANCE_%s|RULE_%s",temp[ACL_NUMBER].value,temp[ACL_RULE_ID1].value){
					temp1 = append(temp1[:i], temp1[i+1:]...)
				} 					
			}
		}			
	}else{
		for i :=acl.RuleId2;i>=acl.RuleId1;i--{	
			redisDelRuleKey(fmt.Sprintf("ACL_ADVANCE_%s|RULE_%d",temp[ACL_NUMBER].value,i))	

			if entries != ""{
				for j,value :=range temp1{
					if value ==fmt.Sprintf("ACL_ADVANCE_%s|RULE_%d",temp[ACL_NUMBER].value,i){
						temp1 = append(temp1[:j], temp1[j+1:]...)
					} 					
				}
			}	

		}
	}
	newEntries :=strings.Join(temp1,",")		
	aclRedisClient.HSet(tableKey,tableField,newEntries)				


	return &rtnCode, nil
}

/***************************************************************************
* 函数说明: 创建L2ACL 规则
*  创建者  于士超   2019.11.27
****************************************************************************/

func (self *AclConfig) CreateL2ACLRule(ctx context.Context, acl *pb.L2_ACLRULE) (*pb.CommonRespHdr, error) {

	var temp [ACL_LAST_ITEM]aclFieldTable
	var rtnCode pb.CommonRespHdr
	var tableKey string
	var entryField string
	var ruleKey string
	rtnCode.RetCode = 0


	/*RULE 条目接收到本地*/
	l2MsgCopy(&temp,acl)
	tableKey =	fmt.Sprintf("acl_table|acl_link_%s",temp[ACL_NUMBER].value)
	entryField = "ACL_ENTRIES"
		
	/*如果命令不带有RULE ID，那么根据默认步长自定义RULE ID*/
	if temp[ACL_RULE_ID].value =="-1"{			
		/*首先从数据库获取该number对应的步长*/
		step :=aclNumberStepQuery(tableKey,temp[ACL_STEP].field)
		//10000作为测试使用，后续更改
		for i :=step;i<10000;i=(i+step) {
			hlen := aclRedisClient.HLen(fmt.Sprintf("ACL_LINK_%s|RULE_%d",temp[ACL_NUMBER].value,i)).Val()
			if hlen ==0 {
				temp[ACL_RULE_ID].value = fmt.Sprint(i)	
				break;				
			}
		}

	}
	ruleKey = fmt.Sprintf("ACL_LINK_%s|RULE_%s",temp[ACL_NUMBER].value,temp[ACL_RULE_ID].value)		
	/*首先判断是否存在该条目，如果存在，那么先删除，后下发；如果不存在，那么直接下发配置*/
	redisDelRuleKey(ruleKey)		
	hlen := aclRedisClient.HLen(ruleKey).Val()
	if hlen ==0{/*下发配置到CONFIG DB*/
		for i :=3;i<len(temp);i++{
			if temp[i].value !="" && temp[i].value !="-1"{
				aclRedisClient.HSet(ruleKey, temp[i].field, temp[i].value).Err()
			}
		}
	}
	entries:=redisGetKeyFiled(tableKey,entryField)
	if entries != ""{
		r := strings.Contains(entries,ruleKey)
		if r== false{
			entries = entries + "," + ruleKey
			entries = strings.Trim(entries,",")
			aclRedisClient.HSet(tableKey,entryField,entries)				
		}
	}else{
		aclRedisClient.HSet(tableKey,entryField,ruleKey)				
	}		
	msgRuleUpdateQueueInsert(ruleKey,RULE_UPDATE)			

	/*发布消息*/

	return &rtnCode, nil
}

/***************************************************************************
* 函数说明: 删除L2ACL 规则
*  创建者  于士超   2019.11.27
****************************************************************************/
func (self *AclConfig) DeleteL2ACLRule(ctx context.Context, acl *pb.L2_ACLRULE) (*pb.CommonRespHdr, error) {

	var temp [ACL_LAST_ITEM]aclFieldTable
	var rtnCode pb.CommonRespHdr
	rtnCode.RetCode = 0
	var tableKey string
	var tableField string	

	l2MsgCopy(&temp,acl)
	fmt.Println(acl.AclNumber,acl.Step,acl.RuleId,acl.RuleId1,acl.RuleId2,acl.Action)
	tableKey = fmt.Sprintf("acl_table|acl_link_%s",temp[ACL_NUMBER].value)
	tableField = "ACL_ENTRIES"

	entries:=redisGetKeyFiled(tableKey,tableField)
	temp1 := strings.Split(entries,",")					
	if acl.RuleId2 == -1{
		redisDelRuleKey(fmt.Sprintf("ACL_LINK_%s|RULE_%s",temp[ACL_NUMBER].value,temp[ACL_RULE_ID1].value))
		if entries != ""{
			for i,value :=range temp1{
				if value ==fmt.Sprintf("ACL_LINK_%s|RULE_%s",temp[ACL_NUMBER].value,temp[ACL_RULE_ID1].value){
					temp1 = append(temp1[:i], temp1[i+1:]...)
				} 					
			}
		}
	}else{
		for i :=acl.RuleId2;i>=acl.RuleId1;i--{	
			redisDelRuleKey(fmt.Sprintf("ACL_LINK_%s|RULE_%d",temp[ACL_NUMBER].value,i))
			if entries != ""{
				for j,value :=range temp1{
					if value ==fmt.Sprintf("ACL_LINK_%s|RULE_%d",temp[ACL_NUMBER].value,i){
						temp1 = append(temp1[:j], temp1[j+1:]...)
					} 					
				}
			}	
		}
	}
	newEntries :=strings.Join(temp1,",")			
	aclRedisClient.HSet(tableKey,tableField,newEntries)

	return &rtnCode, nil
}
/***************************************************************************
* 函数说明: 创建ARPACL 规则
*  创建者  于士超   2019.11.27
****************************************************************************/

func (self *AclConfig) CreateArpACLRule(ctx context.Context, acl *pb.ARP_ACLRULE) (*pb.CommonRespHdr, error) {

	var temp [ACL_LAST_ITEM]aclFieldTable	
	var rtnCode pb.CommonRespHdr
	var tableKey string
	var ruleKey string	
	var entryField string
	rtnCode.RetCode = 0

	arpMsgCopy(&temp,acl)
	tableKey = fmt.Sprintf("acl_table|acl_arp_%s",temp[ACL_NUMBER].value)
	entryField = "ACL_ENTRIES"
	
	/*如果命令不带有RULE ID，那么根据默认步长自定义RULE ID*/
	if temp[ACL_RULE_ID].value =="-1"{			
		/*首先从数据库获取该number对应的步长*/
		step :=aclNumberStepQuery(tableKey,temp[ACL_STEP].field)			
		//10000作为测试使用，后续更改
		for i :=step;i<10000;i=(i+step) {
			hlen := aclRedisClient.HLen(fmt.Sprintf("ACL_ARP_%s|RULE_%d",temp[ACL_NUMBER].value,i)).Val()
			if hlen ==0 {
				temp[ACL_RULE_ID].value = fmt.Sprint(i)	
				break;				
			}
		}		
	}
	ruleKey = fmt.Sprintf("ACL_ARP_%s|RULE_%s",temp[ACL_NUMBER].value,temp[ACL_RULE_ID].value)
	/*首先判断是否存在该条目，如果存在，那么先删除，后下发；如果不存在，那么直接下发配置*/
	redisDelRuleKey(ruleKey)
	hlen := aclRedisClient.HLen(ruleKey).Val()
	if hlen ==0{/*下发配置到CONFIG DB*/
		for i :=3;i<len(temp);i++{
			if temp[i].value !="" && temp[i].value !="-1"{
				aclRedisClient.HSet(ruleKey, temp[i].field, temp[i].value).Err()
			}
		}
	}
	entries:=redisGetKeyFiled(tableKey,entryField)
	if entries != ""{
		r := strings.Contains(entries,ruleKey)
		if r== false{
			entries = entries + "," + ruleKey
			entries = strings.Trim(entries,",")
			aclRedisClient.HSet(tableKey,entryField,entries)				
		}
	}else{
		aclRedisClient.HSet(tableKey,entryField,ruleKey)				
	}		
	msgRuleUpdateQueueInsert(ruleKey,RULE_UPDATE)		


	return &rtnCode, nil
}
/***************************************************************************
* 函数说明: 删除ARPACL 规则
*  创建者  于士超   2019.11.27
****************************************************************************/
func (self *AclConfig) DeleteArpACLRule(ctx context.Context, acl *pb.ARP_ACLRULE) (*pb.CommonRespHdr, error) {

	var temp [ACL_LAST_ITEM]aclFieldTable
	var rtnCode pb.CommonRespHdr
	rtnCode.RetCode = 0
	var tableKey string
	var tableField string

	arpMsgCopy(&temp,acl)
	tableKey = fmt.Sprintf("acl_table|acl_arp_%s",temp[ACL_NUMBER].value)
	tableField = "ACL_ENTRIES"
	
	entries:=redisGetKeyFiled(tableKey,tableField)
	temp1 := strings.Split(entries,",")				
	if acl.RuleId2 == -1{
		redisDelRuleKey(fmt.Sprintf("ACL_ARP_%s|RULE_%s",temp[ACL_NUMBER].value,temp[ACL_RULE_ID1].value))	
		if entries != ""{
			for i,value :=range temp1{
				if value ==fmt.Sprintf("ACL_ARP_%s|RULE_%s",temp[ACL_NUMBER].value,temp[ACL_RULE_ID1].value){
					temp1 = append(temp1[:i], temp1[i+1:]...)
				} 					
			}
		}
	}else{
		for i :=acl.RuleId2;i>=acl.RuleId1;i--{
			redisDelRuleKey(fmt.Sprintf("ACL_ARP_%s|RULE_%d",temp[ACL_NUMBER].value,i))	
			if entries != ""{
				for j,value :=range temp1{
					if value ==fmt.Sprintf("ACL_ARP_%s|RULE_%d",temp[ACL_NUMBER].value,i){
						temp1 = append(temp1[:j], temp1[j+1:]...)
					} 					
				}								
			}
		}
	}
	newEntries :=strings.Join(temp1,",")			
	aclRedisClient.HSet(tableKey,tableField,newEntries)				

	return &rtnCode, nil
}
/***************************************************************************
* 函数说明: 创建混合ACL 规则
*  创建者  于士超   2019.11.27
****************************************************************************/
func (self *AclConfig) CreateMixedACLRule(ctx context.Context, acl *pb.MIXED_ACLRULE) (*pb.CommonRespHdr, error) {

	var temp [ACL_LAST_ITEM]aclFieldTable	
	var rtnCode pb.CommonRespHdr
	var tableKey string
	var entryField string
	var ruleKey string
	rtnCode.RetCode = 0

	mixedMsgCopy(&temp,acl)
	tableKey = fmt.Sprintf("acl_table|acl_mixed_%s",temp[ACL_NUMBER].value)
	entryField ="ACL_ENTRIES"
		
	/*如果命令不带有RULE ID，那么根据默认步长自定义RULE ID*/
	if temp[ACL_RULE_ID].value =="-1"{			
		/*首先从数据库获取该number对应的步长*/
		step :=aclNumberStepQuery(tableKey,temp[ACL_STEP].field)			
		//10000作为测试使用，后续更改
		for i :=step;i<10000;i=(i+step) {
			hlen := aclRedisClient.HLen(fmt.Sprintf("ACL_MIXED_%s|RULE_%d",temp[ACL_NUMBER].value,i)).Val()
			if hlen ==0 {
				temp[ACL_RULE_ID].value = fmt.Sprint(i)	
				break;				
			}
		}			
	}
	ruleKey = fmt.Sprintf("ACL_MIXED_%s|RULE_%s",temp[ACL_NUMBER].value,temp[ACL_RULE_ID].value)
	redisDelRuleKey(ruleKey)		
	hlen := aclRedisClient.HLen(ruleKey).Val()
	if hlen ==0{/*下发配置到CONFIG DB*/
		for i :=3;i<len(temp);i++{
			if temp[i].value !="" && temp[i].value !="-1"{
				aclRedisClient.HSet(ruleKey, temp[i].field, temp[i].value).Err()
			}
		}
	}
	entries:=redisGetKeyFiled(tableKey,entryField)
	if entries != ""{
		r := strings.Contains(entries,ruleKey)
		if r== false{
			entries = entries + "," + ruleKey
			entries = strings.Trim(entries,",")
			aclRedisClient.HSet(tableKey,entryField,entries)				
		}
	}else{
		aclRedisClient.HSet(tableKey,entryField,ruleKey)				
	}			
	msgRuleUpdateQueueInsert(ruleKey,RULE_UPDATE)			

	
	return &rtnCode, nil
}
/***************************************************************************
* 函数说明: 删除混合ACL 规则
*  创建者  于士超   2019.11.27
****************************************************************************/
func (self *AclConfig) DeleteMixedACLRule(ctx context.Context, acl *pb.MIXED_ACLRULE) (*pb.CommonRespHdr, error) {

	var temp [ACL_LAST_ITEM]aclFieldTable
	var rtnCode pb.CommonRespHdr
	var tableKey string
	var tableField string
	rtnCode.RetCode = 0

	mixedMsgCopy(&temp,acl)
	tableKey = fmt.Sprintf("acl_table|acl_mixed_%s",temp[ACL_NUMBER].value)
	tableField = "ACL_ENTRIES"

	
	entries:=redisGetKeyFiled(tableKey,tableField)	
	temp1 := strings.Split(entries,",")					
	if acl.RuleId2 == -1{
		redisDelRuleKey(fmt.Sprintf("ACL_MIXED_%s|RULE_%s",temp[ACL_NUMBER].value,temp[ACL_RULE_ID1].value))
		if entries != ""{
			for i,value :=range temp1{
				if value ==fmt.Sprintf("ACL_MIXED_%s|RULE_%s",temp[ACL_NUMBER].value,temp[ACL_RULE_ID1].value){
					temp1 = append(temp1[:i], temp1[i+1:]...)
				} 					
			}	
		}
	}else{
		for i :=acl.RuleId2;i>=acl.RuleId1;i--{
			redisDelRuleKey(fmt.Sprintf("ACL_MIXED_%s|RULE_%d",temp[ACL_NUMBER].value,i))	
			if entries != ""{
				for j,value :=range temp1{
					if value ==fmt.Sprintf("ACL_MIXED_%s|RULE_%d",temp[ACL_NUMBER].value,i){
						temp1 = append(temp1[:j], temp1[j+1:]...)
					} 					
				}	
			}									
		}
	}
	newEntries :=strings.Join(temp1,",")			
	aclRedisClient.HSet(tableKey,tableField,newEntries)				

	return &rtnCode, nil
}

/***************************************************************************
* 函数说明: acl table绑定接口
*  创建者  于士超   2019.11.27
****************************************************************************/
func AclTableBind(acl *pb.ACL_TABLE)(*pb.CommonRespHdr,error){

	var temp [ACL_TABLE_LAST_ITEM]aclFieldTable
	var rtnCode pb.CommonRespHdr
	var value string		
	rtnCode.RetCode = 0

	aclTableMsgCopy(&temp,acl)	
	key:= "acl_table|acl_"+ acl.AclType + "_" + acl.AclNumber
	field:= "BIND"+"_" + acl.BindType
	value = aclRedisClient.HGet(key,field).Val()

	temp1 := strings.Split(value,",")
	for _,j:=range temp1{
		if j == acl.BindTypeNumber{
			return &rtnCode, nil		
		}
	}

	value += ","+ acl.BindTypeNumber
	value = strings.Trim(value,",")
	aclRedisClient.HSet(key,field,value)
	msgTableBindQueueInsert(key,TABLE_BIND)
	return &rtnCode, nil
}
/***************************************************************************
* 函数说明: acl table解除绑定接口
*  创建者  于士超   2019.11.27
****************************************************************************/
func AclTableUnbind(acl *pb.ACL_TABLE)(*pb.CommonRespHdr,error){

	var temp [ACL_TABLE_LAST_ITEM]aclFieldTable
	var rtnCode pb.CommonRespHdr
	var value string
	aclTableMsgCopy(&temp,acl)		
	rtnCode.RetCode = 0

	key:= "acl_table|acl_"+ acl.AclType + "_" + acl.AclNumber
	field:= "BIND"+"_" + acl.BindType
	value = aclRedisClient.HGet(key,field).Val()

	temp1 := strings.Split(value,",")
	for i,j:=range temp1{
		if j == acl.BindTypeNumber{
			temp1 = append(temp1[:i], temp1[i+1:]...)
			value = strings.Join(temp1,",")
			aclRedisClient.HSet(key,field,value)
			msgTableUnbindQueueInsert(key,TABLE_UNBIND)
			return 	&rtnCode, nil					
		}
	}
	return &rtnCode, nil
}
/***************************************************************************
* 函数说明: acl table 步长配置
*  创建者  于士超   2019.11.27
****************************************************************************/
func AclTableStepSetup(acl *pb.ACL_TABLE)(*pb.CommonRespHdr,error){

	var temp [ACL_TABLE_LAST_ITEM]aclFieldTable
	var rtnCode pb.CommonRespHdr
	aclTableMsgCopy(&temp,acl)		
	rtnCode.RetCode = 0

	key:= "acl_table|acl_"+ acl.AclType + "_" + acl.AclNumber
	field:= "ACL_STEP"
	aclRedisClient.HSet(key,field,temp[ACL_TABLE_STEP].value)

	return &rtnCode, nil
}
/***************************************************************************
* 函数说明: acl table 恢复为默认步长
*  创建者  于士超   2019.11.27
****************************************************************************/
func AclTableNostepSetup(acl *pb.ACL_TABLE)(*pb.CommonRespHdr,error){

	var temp [ACL_TABLE_LAST_ITEM]aclFieldTable
	var rtnCode pb.CommonRespHdr
	aclTableMsgCopy(&temp,acl)		
	rtnCode.RetCode = 0

	key:= "acl_table|acl_"+ acl.AclType + "_" + acl.AclNumber
	field:= "ACL_STEP"
	aclRedisClient.HSet(key,field,"5")

	return &rtnCode, nil
}
/***************************************************************************
* 函数说明: acl table创建
*  创建者  于士超   2019.11.27
****************************************************************************/
func AclTableCreateSetup(acl *pb.ACL_TABLE)(*pb.CommonRespHdr,error){

	var temp [ACL_TABLE_LAST_ITEM]aclFieldTable
	var rtnCode pb.CommonRespHdr
	aclTableMsgCopy(&temp,acl)		
	rtnCode.RetCode = 0

	key:= "acl_table|acl_"+ acl.AclType + "_" + acl.AclNumber
	field:= "ACL_STEP"
	len := aclRedisClient.HLen(key).Val()	
	if len ==0{
		aclRedisClient.HSet(key,field,"5")
	}

	return &rtnCode, nil
}

/***************************************************************************
* 函数说明: 更新acl table 主要包括绑定接口、解除绑定接口以及步长配置
*  创建者  于士超   2019.11.27
****************************************************************************/

func (self *AclConfig) UpdateACLTable(ctx context.Context, acl *pb.ACL_TABLE) (*pb.CommonRespHdr, error) {

	var rtnCode pb.CommonRespHdr	
	rtnCode.RetCode = 0
	if acl.Action ==0{//bind
		return AclTableBind(acl)
	}else if acl.Action ==1{//unbind
		return AclTableUnbind(acl)
	}else if acl.Action ==2{//step
		return AclTableStepSetup(acl)		
	}else if acl.Action ==3{
		return AclTableNostepSetup(acl)		
	}else if acl.Action ==5{
		return AclTableCreateSetup(acl)		
	}

	return &rtnCode, nil
}
/***************************************************************************
* 函数说明: 删除Acl table
*  创建者  于士超   2019.11.27
****************************************************************************/
func (self *AclConfig) DeleteACLTable(ctx context.Context, acl *pb.ACL_TABLE) (*pb.CommonRespHdr, error) {

	var temp [ACL_TABLE_LAST_ITEM]aclFieldTable
	var rtnCode pb.CommonRespHdr
	var value string
	rtnCode.RetCode = 0

	aclTableMsgCopy(&temp,acl)

	key := "acl_table|acl_"+ acl.AclType + "_" + acl.AclNumber	
	field := "ACL_ENTRIES"
	value = aclRedisClient.HGet(key,field).Val()

	temp1 :=strings.Split(value,",")

	for _,entry :=range temp1{
		redisDelRuleKey(entry)		
	}

	redisDelTableKey(key)

	return &rtnCode, nil
}
/***************************************************************************
* 函数说明: 显示某条acl entry
*  创建者  于士超   2019.11.27
****************************************************************************/
func showAclEntry(acl *pb.ACL_SHOW_MSG)(*pb.ALL_SHOW_RESPONSE, error){

	var aclReponse pb.ALL_SHOW_RESPONSE

	key := "acl_table|acl_"+ acl.AclType + "_" + acl.AclName	
	field := "ACL_ENTRIES"
	portfield := "BIND_PORT"
	vlanfield := "BIND_VLAN"
	stepfield := "ACL_STEP"
	interfacefield := "BIND_INTERFACE"
	entries := aclRedisClient.HGet(key,field).Val()

	aclReponse.Entries = make([]*pb.ACL_SHOW_RESPONSE,1)
	aclReponse.Entries[0]=new(pb.ACL_SHOW_RESPONSE)

	aclReponse.Entries[0].Step = aclRedisClient.HGet(key,stepfield).Val()
	aclReponse.Entries[0].BindPort = aclRedisClient.HGet(key,portfield).Val()
	aclReponse.Entries[0].BindVlan = aclRedisClient.HGet(key,vlanfield).Val()
	aclReponse.Entries[0].BindInterface = aclRedisClient.HGet(key,interfacefield).Val()	

	temp1 :=strings.Split(entries,",")
	len :=strings.Count(entries,",")
	aclReponse.Entries[0].Cmdline = make([]string,len+1)
	
	for i,entry :=range temp1{
		aclReponse.Entries[0].Cmdline[i] =aclRedisClient.HGet(entry,"ACL_CMDLINE").Val()		
	}


	return &aclReponse, nil

}
/***************************************************************************
* 函数说明: 显示所有acl entries
*  创建者  于士超   2019.11.27
****************************************************************************/
func showAclAllEntries(acl *pb.ACL_SHOW_MSG)(*pb.ALL_SHOW_RESPONSE, error){
	var aclReponse pb.ALL_SHOW_RESPONSE
	
	field := "ACL_ENTRIES"
	portfield := "BIND_PORT"
	vlanfield := "BIND_VLAN"
	stepfield := "ACL_STEP"
	interfacefield := "BIND_INTERFACE"

	tableInfo:=aclRedisClient.Keys("acl_table*").Val()
	aclReponse.Entries = make([]*pb.ACL_SHOW_RESPONSE,len(tableInfo))	
	for index,key := range tableInfo{
		aclReponse.Entries[index]=new(pb.ACL_SHOW_RESPONSE)
		aclReponse.Entries[index].AclTableName = key	
		aclReponse.Entries[index].Step = aclRedisClient.HGet(key,stepfield).Val()
		aclReponse.Entries[index].BindPort = aclRedisClient.HGet(key,portfield).Val()
		aclReponse.Entries[index].BindVlan = aclRedisClient.HGet(key,vlanfield).Val()
		aclReponse.Entries[index].BindInterface = aclRedisClient.HGet(key,interfacefield).Val()	
		entries := aclRedisClient.HGet(key,field).Val()	
		temp1 :=strings.Split(entries,",")
		len :=strings.Count(entries,",")
		aclReponse.Entries[index].Cmdline = make([]string,len+1)
		
		for i,entry :=range temp1{
			aclReponse.Entries[index].Cmdline[i] =aclRedisClient.HGet(entry,"ACL_CMDLINE").Val()		
		}

	}

	return &aclReponse, nil
}
/***************************************************************************
* 函数说明: acl 命令回显
*  创建者  于士超   2019.11.27
****************************************************************************/

func (self *AclConfig) ShowACLRule(ctx context.Context, acl *pb.ACL_SHOW_MSG) (*pb.ALL_SHOW_RESPONSE, error) {

	if acl.Action ==1{
		return showAclEntry(acl)
	}else{
		return showAclAllEntries(acl)
	}

}