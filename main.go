package main

import (
	"flag"
	"github.com/unbyte/beeep"
	"log"
	"os"
)

var (
	i string
)

func init() {
	flag.StringVar(&i, "i", "", "指定配置文件路径，默认是在程序同目录下的 config.yaml")
}

func main() {
	beeep.AppID = "Teemo"
	flag.Parse()
	if len(i) == 0 {
		i = getDefaultConfigPath()
	}
	tasks, err := parseConfig(i)
	if err != nil {
		log.Printf("解析配置文件时出错: %s", err.Error())
		os.Exit(1)
		return
	}
	for _, t := range tasks {
		if err := t.Start(); err != nil {
			os.Exit(1)
			return
		}
	}
	select {}
}
