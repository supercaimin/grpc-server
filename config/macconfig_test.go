package config

import (
	"golang.org/x/net/context"
	pb "hstcmler/cml"
	"testing"
)

func TestNewMacConfig(t *testing.T) {
	var testmaconfig = new(MacConfig)
	*testmaconfig = MacConfig{}
	testmaconfig = NewMacConfig()
	if nil != testmaconfig {
		t.Error("TestNewMacConfig error")
	}
}

func TestSetMacGlobalCfg(t *testing.T) {
	var testblackhole = new(pb.Macblackhole)
	*testblackhole = pb.Macblackhole{
		Mac:    "01-02-03-04-05-06",
		Vlanid: 11,
	}

	var teststatic = []*pb.Macstatic{
		{Mac: "01-02-03-04-05-06", Ifname: "11", Vlanid: 11},
		{Mac: "01-02-03-04-05-07", Ifname: "22", Vlanid: 22},
	}
	var testmacflag = new(pb.Macflapcfg)
	*testmacflag = pb.Macflapcfg{
		Flapingagetime:  11,
		Flapdetectlevel: "11",
	}
	var testreq = new(pb.Macglobalcfg)
	*testreq = pb.Macglobalcfg{
		Agingtime:  60,
		Machole:    testblackhole,
		Macflap:    testmacflag,
		Static:     teststatic,
		Updatetime: 32,
	}
	var maccoftest = new(MacConfig)
	*maccoftest = MacConfig{}
	var testerr error
	_, testerr = maccoftest.SetMacGlobalCfg(context.Background(), testreq)
	if nil != testerr {
		t.Error("testSetMacGlobalCfg error")
	}
}

func TestDelMacGlobalCfg(t *testing.T) {
	var testblackhole = new(pb.Macblackhole)
	*testblackhole = pb.Macblackhole{
		Mac:    "01-02-03-04-05-06",
		Vlanid: 11,
	}

	var teststatic = []*pb.Macstatic{
		{Mac: "01-02-03-04-05-06", Ifname: "11", Vlanid: 11},
		{Mac: "01-02-03-04-05-07", Ifname: "22", Vlanid: 22},
	}
	var testmacflag = new(pb.Macflapcfg)
	*testmacflag = pb.Macflapcfg{
		Flapingagetime:  11,
		Flapdetectlevel: "11",
	}
	var testreq = new(pb.Macglobalcfg)
	*testreq = pb.Macglobalcfg{
		Agingtime:  60,
		Machole:    testblackhole,
		Macflap:    testmacflag,
		Static:     teststatic,
		Updatetime: 32,
	}
	var maccoftest = new(MacConfig)
	*maccoftest = MacConfig{}
	var testerr error
	_, testerr = maccoftest.DelMacGlobalCfg(context.Background(), testreq)
	if nil != testerr {
		t.Error("testDelMacGlobalCfg error")
	}
}

func TestShowMacGlobalCfg(t *testing.T) {
	var testmacflag = new(pb.Macflapcfg)
	*testmacflag = pb.Macflapcfg{
		Flapingagetime:  11,
		Flapdetectlevel: "11",
	}
	var testreq = new(pb.Showcfginfo)
	*testreq = pb.Showcfginfo{
		Cmdcode:    60,
		Showoption: "111",
		Inputstr:   "111",
		Regstr:     "111",
	}
	var maccoftest = new(MacConfig)
	*maccoftest = MacConfig{}
	var testerr error
	_, testerr = maccoftest.ShowMacGlobalCfg(context.Background(), testreq)
	if nil != testerr {
		t.Error("testDelMacGlobalCfg error")
	}
}

func TestSetMacIfCfg(t *testing.T) {
	var testmaclearn = new(pb.Maclearndis)
	*testmaclearn = pb.Maclearndis{
		Learndisable: 11,
		Action:       "11",
	}
	var testmaclimit = new(pb.Maclimit)
	*testmaclimit = pb.Maclimit{
		Maxnum: 11,
		Action: "11",
		Alarm:  11,
	}
	var testmacflag = new(pb.Macflapcfg)
	*testmacflag = pb.Macflapcfg{
		Flapingagetime:  11,
		Flapdetectlevel: "11",
	}
	var testreq = new(pb.Macifcfg)
	*testreq = pb.Macifcfg{
		Ifname:     "11",
		Maclearn:   testmaclearn,
		Maclimit:   testmaclimit,
		Portbridge: 11,
		Updatetime: 32,
	}
	var maccoftest = new(MacConfig)
	*maccoftest = MacConfig{}
	var testerr error
	_, testerr = maccoftest.SetMacIfCfg(context.Background(), testreq)
	if nil != testerr {
		t.Error("testDelMacGlobalCfg error")
	}
}

func TestDelMacIfCfg(t *testing.T) {
	var testmaclearn = new(pb.Maclearndis)
	*testmaclearn = pb.Maclearndis{
		Learndisable: 11,
		Action:       "11",
	}
	var testmaclimit = new(pb.Maclimit)
	*testmaclimit = pb.Maclimit{
		Maxnum: 11,
		Action: "11",
		Alarm:  11,
	}
	var testmacflag = new(pb.Macflapcfg)
	*testmacflag = pb.Macflapcfg{
		Flapingagetime:  11,
		Flapdetectlevel: "11",
	}
	var testreq = new(pb.Macifcfg)
	*testreq = pb.Macifcfg{
		Ifname:     "11",
		Maclearn:   testmaclearn,
		Maclimit:   testmaclimit,
		Portbridge: 11,
		Updatetime: 32,
	}
	var maccoftest = new(MacConfig)
	*maccoftest = MacConfig{}
	var testerr error
	_, testerr = maccoftest.DelMacIfCfg(context.Background(), testreq)
	if nil != testerr {
		t.Error("testDelMacGlobalCfg error")
	}
}

func TestShowMacIfCfg(t *testing.T) {
	var testmaclearn = new(pb.Maclearndis)
	*testmaclearn = pb.Maclearndis{
		Learndisable: 11,
		Action:       "11",
	}
	var testmaclimit = new(pb.Maclimit)
	*testmaclimit = pb.Maclimit{
		Maxnum: 11,
		Action: "11",
		Alarm:  11,
	}
	var testmacflag = new(pb.Macflapcfg)
	*testmacflag = pb.Macflapcfg{
		Flapingagetime:  11,
		Flapdetectlevel: "11",
	}
	var testreq = new(pb.Showcfginfo)
	*testreq = pb.Showcfginfo{
		Cmdcode:    11,
		Showoption: "11",
		Inputstr:   "11",
		Regstr:     "11",
	}
	var maccoftest = new(MacConfig)
	*maccoftest = MacConfig{}
	var testerr error
	_, testerr = maccoftest.ShowMacIfCfg(context.Background(), testreq)
	if nil != testerr {
		t.Error("testDelMacGlobalCfg error")
	}
}
