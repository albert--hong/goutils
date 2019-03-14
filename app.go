package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

// AppEnv 是程序运行的全局配置
type AppEnv struct {
	AppName  string
	AppVer   string
	Config   Config
	Logger   *log.Logger
	LogFile  *os.File
	TodayStr string
}

// Config 是程序运行的配置文件
type Config struct {
}

func main() {
	// 初始化运行环境
	env := initialEnv()

	// TODO: 主函数主体

	// 结束运行环境
	destroyEnv(env)
}

// initialEnv 的目标是初始化运行环境
func initialEnv() *AppEnv {
	env := AppEnv{}
	env.AppName = "app"
	env.TodayStr = time.Now().Format("2006-01-02")
	// 1. 解析命令行
	var configPath = "./conf/" + env.AppName + ".yaml"
	flag.StringVar(&configPath, "conf_file", configPath, "Application's configure file!")
	flag.Parse()
	// 2. 配置日志文件。只添加不删除。
	var logPath = "./log/" + env.AppName + "." + env.TodayStr + ".log"
	logWriter, err := os.OpenFile(logPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln("Cannot open or create the log file:", logPath)
	}
	logger := log.New(logWriter, "", log.Lshortfile|log.Ldate|log.Ltime)
	env.Logger = logger
	env.LogFile = logWriter
	// 3. 读取配置文件
	confStr, err := ioutil.ReadFile(configPath)
	if err != nil {
		env.Logger.Fatalln("Open the configure file failed. path:", configPath)
	}
	env.Config = Config{}
	err = yaml.Unmarshal(confStr, &env.Config)
	if err != nil {
		env.Logger.Fatal("Parse the configure file failed. path:", configPath)
	}
	// 4. 其他初始化内容。包括打开文件等

	return &env
}

// destroy 销毁App运行环境
func destroyEnv(env *AppEnv) {
	// 关闭所有文件
	env.LogFile.Sync()
	env.LogFile.Close()
}
