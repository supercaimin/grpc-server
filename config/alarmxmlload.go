// alarmProcess.go
// 创建者 陈维 2020.2.10
// 告警定义加载处理文件
// 告警定义采用xml文件,每个模块独立一个文件,文件格式为

package config

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-redis/redis"
)

type AlarmData struct {
	XMLName  xml.Name `xml:"ALARMDEF"`
	CODE     int      `xml:"CODE"`
	LEVEL    int      `xml:"LEVEL"`
	POSITION string   `xml:"POSITION"`
}

type alarmabs struct {
	XMLName    xml.Name `xml:"INFO"`
	MODULENAME string   `xml:"MODULENAME"`
	UPDATE     string   `xml:"UPDATE"`
}

type RecurlyAlarm struct {
	XMLName    xml.Name `xml:"MODULEALARM"` //最外层的标签
	Moduleinfo alarmabs
	ALARMDEF   []AlarmData `xml:"ALARMDEF"` //里面循环的标签
}

const (
	alarm_define_fix string = ".alarm.xml" /*告警定义文件后缀,文件名以.alarm.xml结尾*/
)

/*逐个读取目录下各业务模块的告警定义文件,然后写入config-db数据库中*/
/*每次进程启动都遍历读取一次,以保证最新的能在db中*/
/*现在的实现如果db实现了持久化会存在一个问题:即如果原有的告警code删除了,db中无法删除;除非每*/
/*次都先把db中的告警码与级别全删除然后逐个增加。但这会导致不兼容,考虑兼容以前定义的告警不允许删除*/
func ReadAlarmDefInfo() {

	/*读取当前目录下的文件*/
	finfo, _ := ioutil.ReadDir(".")

	for _, file := range finfo {
		if file.IsDir() {
			continue
		} else {
			/*以xxx.alarm.xml结尾的文件为告警码定义文件*/
			if -1 != strings.LastIndex(file.Name(), alarm_define_fix) {
				/*告警定义文件,读取加载告警*/
				ReadAlarmDefFile(file.Name())
			} else {
				continue
			}
		}
	}
	/*加载完以后发一次pub, syslog收到后更新其内存数据*/
	/*syslog启动初始化后也主动读一次db,消除两个进程启动顺序的依赖*/
	CmlDbCfgPubstring("CONFIG_DB", PUB_SYSLOGCHAN, "100")

	return
}

func ReadAlarmDefFile(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("open xml file error")
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("read file stream error")
		return
	}
	//解析xml文件并输出：
	alarm := RecurlyAlarm{}
	err = xml.Unmarshal(data, &alarm)
	if err != nil {
		fmt.Println("format xml data failed")
		return
	}

	/*逐条将告警码信息写入数据库*/
	//Key:Alarm|module
	//Field:code-level::posistr;

	fmt.Printf(" Info:module-name :%s,update:%s\n", alarm.Moduleinfo.MODULENAME, alarm.Moduleinfo.UPDATE)

	for index, data := range alarm.ALARMDEF {
		AlarmcodeSetRedis(alarm.Moduleinfo.MODULENAME, data.POSITION, data.CODE, data.LEVEL)

		fmt.Printf(" %d code:%d\n", index+1, data.CODE)
		fmt.Printf(" %d level:%d\n", index+1, data.LEVEL)
		fmt.Printf(" %d position:%s\n", index+1, data.POSITION)
	}
	return
}

/*初始化告警码写数据库 Key:Alarm|module, code - level:position为field-value对 */
func AlarmcodeSetRedis(modulename, position string, code, level int) error {

	var err error = nil
	var dkey, cfield string
	var dbclient *redis.Client

	err = CmlDbReconnect(CFG_DB)
	if err != nil {
		fmt.Print("db state error!\n")
		return err
	}

	/*模块名统一转换为小写*/
	dbclient = cmldbinfo.cmldbclient[CmlDbNo[CFG_DB]]
	dkey = KEY_ALARM_LEVEL_ORIGINAL + strings.ToLower(modulename)
	cfield = fmt.Sprintf("%d", code)

	/*按DB格式写入数据库 code-level:position 域*/
	err = dbclient.HSet(dkey, cfield, fmt.Sprintf("%d%s%s", level, DATA_SPLIT_STR, position)).Err()
	if err != nil {
		fmt.Printf("Write level-field failed,key:%s code:%d\n", dkey, code)
	}

	return err
}
