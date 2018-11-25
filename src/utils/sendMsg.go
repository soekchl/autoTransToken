package utils

import (
	"autoTransToken/src/config"
	"fmt"

	"coding.net/soekchl/smsutil"
)

var (
	msgUserid = ""
	msgPwd    = ""
	msgPhone  = ""
	msgInfo   = ""
)

func init() {
	config.InitConfig("utils")
	loadConfig()
}

func loadConfig() {
	config.ReLoadConfig()
	msgUserid = config.GetConfigString("msgUserid")
	msgPwd = config.GetConfigString("msgPwd")
	msgPhone = config.GetConfigString("msgPhone")
	msgInfo = config.GetConfigString("msgInfo")
}

func SendMsg(code int) bool {
	if code < 100000 || code > 999999 {
		return false
	}
	msg := fmt.Sprintf(msgInfo, code)
	sendobj := smsutil.NewSingleSend(msgUserid, msgPwd, msgPhone, msg)
	return smsutil.SendAndRecvOnce(sendobj)
}
