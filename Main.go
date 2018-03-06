package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ricardoaat/bioschemas-govalid/config"
	"github.com/ricardoaat/bioschemas-govalid/validator"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

var (
	version   string
	buildDate string
)

func logInit() {

	//now := time.Now()
	//logfile := config.Conf.Path.LogPath + fmt.Sprintf("govalid_%s.log", now.Format("20060102T150405"))
	logfile := config.Conf.Path.LogPath + "govalid.log"
	fmt.Println("Loging to " + logfile)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{})
	pathMap := lfshook.PathMap{
		log.DebugLevel: logfile,
		log.InfoLevel:  logfile,
		log.ErrorLevel: logfile,
		log.WarnLevel:  logfile,
		log.PanicLevel: logfile,
	}
	log.AddHook(lfshook.NewHook(
		pathMap,
		&log.JSONFormatter{},
	))
}

func main() {

	err := config.LoadConfig("config.toml")
	if err != nil {
		fmt.Println("Couldn't load config.toml ", err)
	}
	logInit()

	log.Info("--------------Init program--------------")
	log.Info(fmt.Sprintf("Version: %s Build Date: %s", version, buildDate))
	log.Info("Loaded configuration " + fmt.Sprint(config.Conf))
	v := flag.Bool("v", false, "Returns the binary version and built date info")
	flag.Parse()
	if !*v {
		validator.Validate()
	}
}
