package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"models"
	"samh_common_lib/utils/log"
	"server/router"
)

func main() {
	flag.StringVar(&models.Env, "env", "local", "env:local,dev,test,official")
	flag.Parse()

	models.Init()
	go models.NewConfigWatcher()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%v", models.Config.Base_info.Port),
		Handler:        router.InitRouter(),
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   2 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
