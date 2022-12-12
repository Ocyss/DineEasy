package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

// 服务器配置
var (
	AppMode  string
	HttpPort string
)

// 数据库配置
var (
	DbHost     string
	DbPost     string
	DbUser     string
	DbPassword string
	DbName     string
)

var Qiniu = map[string]string{
	"AccessKey":  "",
	"SecretKey":  "",
	"Bucket":     "",
	"QiniuSever": "",
}

func init() {
	file, err := ini.Load("conf/config.ini")
	if err != nil {
		panic(fmt.Sprintf("配置文件读取出错：%v", err))
	}
	LoadServer(file)
	LoadDatabase(file)
	LoadQiniu(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

func LoadDatabase(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("DbHost")
	DbPost = file.Section("database").Key("DbPost").MustString("DbPost")
	DbUser = file.Section("database").Key("DbUser").MustString("DbUser")
	DbPassword = file.Section("database").Key("DbPassword").MustString("DbPassword")
	DbName = file.Section("database").Key("DbName").MustString("DbName")
}

func LoadQiniu(file *ini.File) {
	Qiniu["AccessKey"] = file.Section("qiniu").Key("AccessKey").String()
	Qiniu["SecretKey"] = file.Section("qiniu").Key("SecretKey").String()
	Qiniu["Bucket"] = file.Section("qiniu").Key("Bucket").String()
	Qiniu["QiniuSever"] = file.Section("qiniu").Key("QiniuSever").String()
}
