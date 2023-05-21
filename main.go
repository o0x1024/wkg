package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"wkg/conf"
	db2 "wkg/pkg/db"
	"wkg/pkg/libs/env"
	"wkg/pkg/libs/wkgLogger"
	"wkg/routers"
	"wkg/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var addr = ""

func init() {
	configfile := flag.String("c", "conf/dev/config.yaml", "the config fileOp path")
	confAddr := flag.String("addr", "0.0.0.0:7788", "the listen address")
	flag.Parse()
	addr = *confAddr
	cfg, err := conf.LoadConfig(*configfile)
	if err != nil {
		panic(err)
	}
	conf.Cfg = cfg
}

func main() {
	db2.InitDB()
	env.EnvCheck()
	wkgLogger.InitLog()
	go services.InitService()

	f, err := os.OpenFile("log/access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()

	r.Static("img", "upload/img")

	fmt.Println("[+] wkg running")
	routers.InitRouter(r)
	if err := r.Run(addr); err != nil {
		zap.S().Errorf(err.Error())
		os.Exit(1)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("closing database connection")

	defer func() {
		sqldb, err := db2.Orm.DB()
		if err != nil {
			zap.S().Errorf("[!] main.go get sqldb failed. line:29  . [", err, "]")
			os.Exit(1)
		}
		sqldb.Close()
	}()

}
