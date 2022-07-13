package LogTool

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const dateTimeLayout = "2006-01-02 15:04:05"

func dateTimeNowStr() string {
	return time.Now().Format(dateTimeLayout)
}

// ex: LogTool.LogError("", err)LogType.Error, "Error", err)
func LogBase(lt LogType.LogType, message string, datas ...interface{}) {

	nowTime := dateTimeNowStr()

	// 取得呼叫的 func
	_, file1, line1, fileOK1 := runtime.Caller(1)
	if !fileOK1 {
		log.Printf("%s %s", nowTime, "[Log Tool Error] Not find call func\n")
		return
	}

	file1 = filepath.Base(file1)
	fileInfo := fmt.Sprintf("%s:%d", file1, line1)

	if file1 == "ServerLog.go" || file1 == "server-log.go" {
		_, file2, line2, fileOK2 := runtime.Caller(2)
		if !fileOK2 {
			log.Printf("%s %s", nowTime, "[Log Tool Error] Not find call func\n")
			return
		}

		file2 = filepath.Base(file2)
		fileInfo = fmt.Sprintf("%s:%d", file2, line2)
	}

	switch lt {
	case LogType.ReturnLine: // 換行
		log.Println("")
	case LogType.Divider: // 分隔線
		dividerLen := 10
		dividerStr := "="
		dividerPrint := ""
		for index := 0; index < dividerLen; index++ {
			dividerPrint = dividerPrint + dividerStr
		}
		if message != "" {
			dividerPrint = fmt.Sprintf("%s %s %s %s", nowTime, dividerPrint, message, dividerPrint)
		} else {
			dividerPrint = fmt.Sprintf("%s %s%s", nowTime, dividerPrint, dividerPrint)
		}
		fmt.Println(dividerPrint)
	default:
		if len(datas) > 0 {
			log.Printf("%s [%s] [%s] %s => ", nowTime, lt, fileInfo, message)
			log.Println(datas...)
		} else {
			log.Printf("%s [%s] [%s] %s \n", nowTime, lt, fileInfo, message)
		}
	}

	// Fatal 類型會關閉程式
	if lt == LogType.Fatal {
		os.Exit(1)
	}
}

func LogReturnLine() {
	LogBase(LogType.ReturnLine, "")
}

func LogDivider(message string) {
	LogBase(LogType.Divider, message)
}

func LogFatal(message string, datas ...interface{}) {
	LogBase(LogType.Fatal, message, datas...)
}

func LogError(message string, datas ...interface{}) {
	LogBase(LogType.Error, message, datas...)
}

func LogWarning(message string, datas ...interface{}) {
	LogBase(LogType.Warning, message, datas...)
}

func LogInfo(message string, datas ...interface{}) {
	LogBase(LogType.Info, message, datas...)
}

func LogDebug(message string, datas ...interface{}) {
	LogBase(LogType.Debug, message, datas...)
}

func LogDevelop(message string, datas ...interface{}) {
	LogBase(LogType.Develop, message, datas...)
}

func LogSystem(message string, datas ...interface{}) {
	LogBase(LogType.System, message, datas...)
}

func LogCron(message string, datas ...interface{}) {
	LogBase(LogType.Cron, message, datas...)
}

func LogConfig(message string, datas ...interface{}) {
	LogBase(LogType.Config, message, datas...)
}

func LogConnect(message string, datas ...interface{}) {
	LogBase(LogType.Connect, message, datas...)
}
