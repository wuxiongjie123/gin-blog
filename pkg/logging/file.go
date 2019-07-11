package logging

import (
	"fmt"
	"gin-blog/pkg/file"
	"gin-blog/pkg/setting"
	"os"
	"time"
)

//var (
//	LogSavePath = "runtime/logs/"
//	LogSaveName = "log"
//	LogFileExt = "log"
//	TimeFormat = "20060102"
//)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt, )
}

func openLogFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := file.CheckPermission(src)
	if perm == true {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("fail to OpenFile :%v", err)
	}
	return f, nil
}

//func getLogFileFullPath() string {
//	prefixPath := getLogFilePath()
//	suffixPath := fmt.Sprintf("%s%s.%s",
//		setting.AppSetting.LogSaveName,
//		time.Now().Format(setting.AppSetting.TimeFormat),
//		setting.AppSetting.LogFileExt)
//	return fmt.Sprintf("%s%s",prefixPath,suffixPath)
//}

//func openLogFile(filePath string) *os.File {
//	_,err := os.Stat(filePath)
//	switch  {
//	case os.IsNotExist(err):  //文件或者目录不存在 返回一个bool
//		mkDir()
//	case os.IsPermission(err):  //反回一个bool 得知权限是否满足
//		log.Fatalf("Permission :%v",err)
//	}
//	// os.OpenFile调用文件,支持传入文件名称,指定的模式调用文件,文件权限
//	handle,err:=os.OpenFile(filePath,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//	if err!=nil{
//		log.Fatalf("Fail to OpenFile :%v",err)
//	}
//	return handle
//}
//
//func mkDir()  {
//	dir,_:=os.Getwd()  //返回与当前目录对应的根路径名
//	//os.MkdirAll创建对应的目录以及所需的子目录,若成功则返回nil,否则返回error
//	//os.ModePerm const定义ModePerm FileMode = 0777
//	err:=os.MkdirAll(dir+"/"+getLogFilePath(),os.ModePerm)
//	if err!=nil {
//		panic(err)
//	}
//}
