package models

import (
	"fmt"

	"samh_common_lib/base"
	"samh_common_lib/utils"
	"samh_common_lib/utils/config"
	"samh_common_lib/utils/log"

	"github.com/davecgh/go-spew/spew"
	"github.com/fsnotify/fsnotify"
	raven "github.com/getsentry/raven-go"
	"github.com/go-redis/redis"
	"github.com/xormplus/xorm"
)

var (
	Env                string
	Config             *config.Config
	CommerceDB         *xorm.Engine
	RedisClient        *redis.Client
	RedisClusterClient *redis.ClusterClient
)

func Init() {
	filePath := fmt.Sprintf("config/%s.toml", Env)
	config.Init(filePath)
	spew.Printf("config init succ,filepath:%s\n", filePath)

	Config = config.ConfigInstance()

	log.Init(Config.Log_info_item)

	CommerceDB = utils.InitDB(Config.DB_arr["Commerce"])

	// RedisClient = utils.InitRedisClient(Config.Redis_item)
	// RedisClusterClient = utils.InitRedisCluster(Config.Redis_cluster_item)

	err := raven.SetDSN(Config.Sentry_dsn)
	if err != nil {
		log.Panic(err)
	}

	// initCronDealPayfException()
}

func NewConfigWatcher() {
	log.Info(base.NowFunc())
	defer log.Info(base.NowFunc() + " end")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Panic(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					log.Error(base.NowFuncError())
					return
				}
				if event.Name == fmt.Sprintf("config/%s.toml", Env) && event.Op == fsnotify.Write {
					log.Info("modified file:", event.Name)
					Init()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					log.Error(base.NowFuncError())
					return
				}
				log.Error(err)
			}
		}
	}()

	err = watcher.Add("config")
	if err != nil {
		log.Panic(err)
	}
	<-done
}

// func initCronDealPayfException() {
// log.Info(base.NowFunc())
// defer log.Info(base.NowFunc() + " end")

// c := CronDealPayfException
// if c != nil {
// c.Stop()
// }
// c = cron.New()
// cronStr := Config.Cron_arr["deal_pay_exception"].Cron_str
// log.Infof("cronStr::%s", cronStr)
// err := c.AddFunc(cronStr, func() {
// timingDealPayException()
// })
// if err != nil {
// log.Panic(err)
// }
// c.Start()
// }

// func timingDealPayException() {
// log.Info(base.NowFunc())
// defer log.Info(base.NowFunc() + " end")

// // 将配置的n s以上的未支付的订单状态更为超时失败
// nowTime := time.Now().Unix()
// _, err := VipDB.Sql("UPDATE vip_upgrade_order"+
// " SET status=?,update_time=?"+
// " WHERE update_time-?<? and status=?",
// UpgradeDataStatusCode_TimeoutTimingDealToFailed, nowTime,
// nowTime, Config.Cron_arr["deal_pay_exception"].Time_out, UpgradeDataStatusCode_Ordering,
// ).Execute()
// if err != nil {
// log.Error(err)
// return
// }

// return
// }
