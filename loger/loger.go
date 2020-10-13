package loger

import (
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var(
	Loger        *logrus.Logger
	fileRootPath string
)

func Setup(savePath string){
	fileRootPath = savePath

	if _, err := os.Stat(fileRootPath); os.IsNotExist(err) {
		err := os.MkdirAll(fileRootPath, os.ModePerm)
		if err != nil {
			fmt.Errorf("Loger Setup Error ###", err.Error())
		}
	}

	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Errorf("Loger File Open Error ###", err.Error())
	}

	Loger = logrus.New()
	Loger.SetReportCaller(true)

	//输出到文件
	Loger.Out = src

	logerFormatter := new(LogerFormatter)
	Loger.SetFormatter(logerFormatter)

	Loger.AddHook(newLocalFileLogHook(logrus.ErrorLevel, logerFormatter))
	Loger.AddHook(newLocalFileLogHook(logrus.InfoLevel, logerFormatter))

	Loger.SetOutput(os.Stdout)

	Info("# Loger SetUp Success.")
}

/**
	写入本地日志文件， 按日期、日志级别分割为不同的文件
*/
func newLocalFileLogHook(level logrus.Level, formatter logrus.Formatter) logrus.Hook {

	fileName := filepath.Join(fileRootPath, level.String() + "_%Y%m%d.log")

	//文件分割
	writer, err := rotatelogs.New(
		fileName,
		// 最大保存时间(30天)
		rotatelogs.WithMaxAge(30*24*time.Hour),
		// 日志分割间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	if err != nil {
		fmt.Errorf("config local file system for Loger error: %v", err)
	}

	return lfshook.NewHook(lfshook.WriterMap{
		level: writer,
	}, formatter)

}

func Infof(format string, args ...interface{}){
	setPrefix("Infof")
	Loger.Infof(format, args)
}

func Info(args ...interface{}){
	Loger.Info(setPrefix("Info"), args)
}

func Error(args ...interface{}){
	Loger.Error(setPrefix("Error"), args)
}


// setPrefix set the prefix of the log output
func setPrefix(level string) string{

	pc, file, line, ok := runtime.Caller(2)
	if ok {
		loc, _ := time.LoadLocation("Asia/Shanghai")
		funcName := runtime.FuncForPC(pc).Name()
		funcName = strings.TrimPrefix(filepath.Ext(funcName), ".")
		timestamp := time.Now().In(loc).Format("2006-01-02 15:04:05")

		//path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		//path = path[1:] + "/"

		return fmt.Sprintf("[%s][%s][%s:%d:%s]", strings.ToUpper(level), timestamp, filepath.Base(file), line, funcName)
	}
	return ""
}

/*
	日志输出格式
*/
type LogerFormatter struct{}

func (s *LogerFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	msg := fmt.Sprintf("%s \n", entry.Message)
	return []byte(msg), nil
}