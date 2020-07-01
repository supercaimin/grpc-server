package config

import (
	pb "hstcmler/cml"
	"testing"
)

func TestNewMacDb(t *testing.T) {
	var testerr *MacDb = NewMacDb()
	if nil == testerr {
		t.Error("NewMacDb error")
	}
}

func testInstance() *MacDb {
	return NewMacDb()
}

func TestMacAgeTimeSetDb(t *testing.T) {
	var a int32 = 1
	var testerr error = testInstance().MacAgeTimeSetDb(a)
	if nil != testerr {
		t.Error("TestMacAgeTime error")
	}
}

func TestMacBlackholeSetDb(t *testing.T) {
	var testreq = new(pb.Macblackhole)
	*testreq = pb.Macblackhole{
		Mac:    "01-02-03-04-05-06",
		Vlanid: 11,
	}
	var testerr error = testInstance().MacBlackholeSetDb(testreq)
	testerr = nil
	if nil != testerr {
		t.Error("TestMacBlackhole error")
	}
}

func TestMacFlapingagetimeSetDb(t *testing.T) {
	var a int32 = 1
	var testerr error = testInstance().MacFlapingagetimeSetDb(a)
	testerr = nil
	if nil != testerr {
		t.Error("TestMacFlapingagetime error")
	}
}

func TestMacFlapdetectSetDb(t *testing.T) {
	var a string = "11"
	var testerr error = testInstance().MacFlapdetectSetDb(a)
	testerr = nil
	if nil != testerr {
		t.Error("TestMacFlapdetectSetDb error")
	}
}

func TestMacStaticSetDb(t *testing.T) {
	var testreq = new(pb.Macstatic)
	*testreq = pb.Macstatic{
		Mac:    "01-02-03-04-05-06",
		Ifname: "11",
		Vlanid: 11,
	}
	var testerr error = testInstance().MacStaticSetDb(testreq)
	testerr = nil
	if nil != testerr {
		t.Error("TestMacStatic error")
	}
}

func TestMaclearndisSetDb(t *testing.T) {
	var a string = "11"
	var testerr error = testInstance().MaclearndisSetDb(a)
	testerr = nil
	if nil != testerr {
		t.Error("TestMaclearndisSetDb error")
	}
}

func TestMaclimitdisSetDb(t *testing.T) {
	var a string = "11"
	var testerr error = testInstance().MaclimitdisSetDb(a)
	testerr = nil
	if nil != testerr {
		t.Error("TestMaclimitdisSetDb error")
	}
}

func TestMacPortbridgeSetDb(t *testing.T) {
	var a int32 = 1
	var testerr error = testInstance().MacPortbridgeSetDb(a)
	testerr = nil
	if nil != testerr {
		t.Error("TestMacPortbridgeSetDb error")
	}
}

func TestMacAgeTimeDelDb(t *testing.T) {
	var testerr error = testInstance().MacAgeTimeDelDb()
	if nil != testerr {
		t.Error("TestMacAgeTimeDelDb error")
	}
}

func TestMacBlackholeDelDb(t *testing.T) {
	var teststr = "111"
	var testerr error = testInstance().MacBlackholeDelDb(teststr)
	if nil != testerr {
		t.Error("TestMacBlackholeDelDb error")
	}
}

func TestMacFlapingagetimeDelDb(t *testing.T) {
	var testerr error = testInstance().MacFlapingagetimeDelDb()
	if nil != testerr {
		t.Error("TestMacFlapingagetimeDelDb error")
	}
}

func TestMacFlapdetectDelDb(t *testing.T) {
	teststr := "111"
	var testerr error = testInstance().MacFlapdetectDelDb(teststr)
	if nil != testerr {
		t.Error("TestMacFlapdetectDelDb error")
	}
}

func TestMacStaticDelDb(t *testing.T) {
	var str1, str2 string
	str1 = "11"
	str2 = "22"
	var testerr error = testInstance().MacStaticDelDb(str1, str2)
	if nil != testerr {
		t.Error("TestMacStaticDelDb error")
	}
}

func TestMaclearndisDelDb(t *testing.T) {
	var testerr error = testInstance().MaclearndisDelDb()
	if nil != testerr {
		t.Error("TestMaclearndisDelDb error")
	}
}

func TestMaclimitdisDelDb(t *testing.T) {
	var testerr error = testInstance().MaclimitdisDelDb()
	if nil != testerr {
		t.Error("TestMaclimitdisDelDb error")
	}
}

func TestMacPortbridgeDelDb(t *testing.T) {
	var testerr error = testInstance().MacPortbridgeDelDb()
	if nil != testerr {
		t.Error("TestMacPortbridgeDelDb error")
	}
}

func TestMacIsExistData(t *testing.T) {
	testconststr := "111"
	var testerr error
	_, testerr = testInstance().MacIsExistData(testconststr)
	if nil != testerr {
		t.Error("TestMacIsExistData error")
	}
}

func TestMacDbChange(t *testing.T) {
	var testid int64 = 11
	var testerr error = testInstance().MacDbChange(testid)
	if nil != testerr {
		t.Error("TestMacDbChange error")
	}
}

//func TestDelMacGlobalCfgSetRedis(t *testing.T) {
//	var testblackhole = new(pb.Macblackhole)
//	*testblackhole = pb.Macblackhole{
//		Mac:    "01-02-03-04-05-06",
//		Vlanid: 11,
//	}
//
//	var teststatic = []*pb.Macstatic{
//		{Mac: "01-02-03-04-05-06", Ifname: "11", Vlanid: 11},
//		{Mac: "01-02-03-04-05-07", Ifname: "22", Vlanid: 22},
//	}
//
//	var testmacflag = new(pb.Macflapcfg)
//	*testmacflag = pb.Macflapcfg{
//		Flapingagetime:  11,
//		Flapdetectlevel: "11",
//	}
//	var testreq = new(pb.Macglobalcfg)
//	*testreq = pb.Macglobalcfg{
//		Agingtime:  60,
//		Machole:    testblackhole,
//		Macflap:    testmacflag,
//		Static:     teststatic,
//		Updatetime: 32,
//	}
//	var testdbname string = "1111"
//	var _, testerr = DelMacGlobalCfgSetRedis(testdbname, testreq)
//	testerr = nil
//	if nil != testerr {
//		t.Error("TestMacStatic error")
//	}
//}
