package main

import (
	"flag"
	"time"

	. "models"
	"process/api"
	"samh_common_lib/base"
	"samh_common_lib/utils/log"
)

func init() {
	Env = *flag.String("env", "local", "env:local,dev,test,official")
	flag.Parse()
	Init()
	go NewConfigWatcher()
}

func main() {
	// PrivilegeInfoApiTest()
	// InternalUserPrivilegeInfoApiTest()
}

func PrivilegeInfoApiTest() {
	log.Debug(base.NowFunc())
	start := time.Now()
	request := &PrivilegeInfoRequest{
		SamhBaseRequest: base.SamhBaseRequest{Uid: 1, DeviceId: "1"},
	}
	log.Debug(request)
	rsp, retCode := api.PrivilegeInfoApi(request)
	log.Debug(retCode, rsp)
	cost := time.Since(start)
	log.Debug(cost)
}

func InternalUserPrivilegeInfoApiTest() {
	log.Debug(base.NowFunc())
	start := time.Now()
	request := &InternalUserPrivilegeInfoRequest{
		SamhBaseRequest: base.SamhBaseRequest{Uid: 1, DeviceId: "1"},
	}
	log.Debug(request)
	rsp, retCode := api.InternalUserPrivilegeInfoApi(request)
	log.Debug(retCode, rsp)
	cost := time.Since(start)
	log.Debug(cost)
}
