package Log_Eio

import (
	"os"
	"strings"
	"time"
)

// 系统参数
var SysCanShu int = 0
var ServerURl string = ""

var FilePort = "00"

//var FilePort = "5001"

// BTest == true 测试环境
//var BTest = true

// 日志函数
func Log11(data string, data1 ...string) {
	var datatmp string

	datatmp = data
	// 循环取值
	for _, data1 := range data1 {
		datatmp = datatmp + data1
	}
	//====================================================
	var path string
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		path = "\\"
	} else {
		path = "/"
	}
	dir, _ := os.Getwd()                     // 获取当前的程序路径
	os.MkdirAll(dir+path+"log", os.ModePerm) //生成多级目录
	//====================================================

	//创建日志文件
	t := time.Now()
	filepath := "./log/access_run_" + FilePort + "_" + t.Format("2006-01-02") + ".txt"
	_, err := os.Stat(filepath)
	var file *os.File
	var sTmp string
	if err != nil {
		file, err = os.Create(filepath)
		defer file.Close()
	} else {
		file, err = os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		defer file.Close()
	}
	sTmp = strings.Replace(t.String()[:19], ":", ":", 3) + ":  " + datatmp + "\r\n"
	file.WriteString(sTmp)
	//file.Close()
}

// 访问日志
func Access_runlog(data string, data1 ...string) {
	var datatmp string
	datatmp = data + " "
	// 循环取值
	for _, data1 := range data1 {
		datatmp = datatmp + data1 + " "
	}

	t := time.Now()
	fileName := "./access_run.txt"
	_, err := os.Stat(fileName)
	var sTmp string
	var file *os.File
	if err != nil {
		file, err = os.Create(fileName)
		sTmp = "create file" + fileName + "\n"
	} else {
		file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		sTmp = "Open file" + fileName + "\n"
	}

	sTmp = strings.Replace(t.String()[:19], ":", ":", 3) + ":  " + datatmp + "\n"
	file.WriteString(sTmp + "\n")
	file.Close()
}

// 访问错误日志
func Access_errlog(data string, data1 ...string) {

	var datatmp string
	datatmp = data + " "
	// 循环取值
	for _, data1 := range data1 {
		datatmp = datatmp + data1 + " "
	}

	t := time.Now()
	fileName := "./access_err.txt"
	_, err := os.Stat(fileName)
	var file *os.File
	var sTmp string
	if err != nil {
		file, err = os.Create(fileName)
		sTmp = "create file" + fileName + "\n"
	} else {
		file, err = os.OpenFile(fileName, os.O_APPEND, os.ModeAppend)
		sTmp = "Open file" + fileName + "\n"
	}

	sTmp = strings.Replace(t.String()[:19], ":", "_", 3) + ":  " + datatmp + "\n"
	file.WriteString(sTmp + "\n")
	file.Close()
}

// 统计日志 要用脚本进行解析
// strOutContent := strUserName + " " + strGameName " " + strTime + " " strStatus
func Statistics(data string) {
	strFileName := "../log/ statistics.txt"
	_, err := os.Stat(strFileName)
	var file *os.File
	if err != nil {
		//fmt.Printf("found old log file %s, now remove it\n", logFilename)
		file, _ = os.OpenFile(strFileName, os.O_APPEND, os.ModeAppend)
	} else {
		file, err = os.Create(strFileName)
	}
	file.WriteString(data)
	file.Close()
}
