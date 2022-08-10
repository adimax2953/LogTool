package LogTool

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	LogType "github.com/adimax2953/log-tool/logType"
)

const dateTimeLayout = "2006-01-02 15:04:05"

func dateTimeNowStr() string {
	return time.Now().Format(dateTimeLayout)
}

func SetOutput(w io.Writer) {
	log.SetOutput(w)
}

// ex: LogTool.LogError("", err)LogType.Error, "Error", err)
func LogBasef(lt LogType.LogType, message string, format string, datas ...any) {

	log.SetFlags(0)

	log.SetPrefix(fmt.Sprintf("%s ", dateTimeNowStr()))

	// 取得呼叫的 func
	_, file1, line1, fileOK1 := runtime.Caller(2)
	if !fileOK1 {
		log.Printf("%s", "[Log Tool Error] Not find call func\n")
		return
	}

	file1 = filepath.Base(file1)
	fileInfo := fmt.Sprintf("%s:%d", file1, line1)

	if file1 == "ServerLog.go" || file1 == "server-log.go" {
		_, file2, line2, fileOK2 := runtime.Caller(3)
		if !fileOK2 {
			log.Printf("%s", "[Log Tool Error] Not find call func\n")
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
			dividerPrint = fmt.Sprintf("%s %s %s", dividerPrint, message, dividerPrint)
		} else {
			dividerPrint = fmt.Sprintf("%s%s", dividerPrint, dividerPrint)
		}
		log.Println(dividerPrint)
	default:
		if len(datas) > 0 {
			log.Printf("[%s] [%s] %s => %s", lt, fileInfo, message, fmt.Sprintf(format, datas...))
		} else {
			log.Printf("[%s] [%s] %s \n", lt, fileInfo, message)
		}
	}

	// Fatal 類型會關閉程式
	if lt == LogType.Fatal {
		os.Exit(1)
	}
}

// ex: LogTool.LogError("", err)LogType.Error, "Error", err)
func LogBase(lt LogType.LogType, message string, datas ...interface{}) {

	log.SetFlags(0)

	log.SetPrefix(fmt.Sprintf("%s ", dateTimeNowStr()))

	// 取得呼叫的 func
	_, file1, line1, fileOK1 := runtime.Caller(2)
	if !fileOK1 {
		log.Printf("%s", "[Log Tool Error] Not find call func\n")
		return
	}

	file1 = filepath.Base(file1)
	fileInfo := fmt.Sprintf("%s:%d", file1, line1)

	if file1 == "ServerLog.go" || file1 == "server-log.go" {
		_, file2, line2, fileOK2 := runtime.Caller(3)
		if !fileOK2 {
			log.Printf("%s", "[Log Tool Error] Not find call func\n")
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
			dividerPrint = fmt.Sprintf("%s %s %s", dividerPrint, message, dividerPrint)
		} else {
			dividerPrint = fmt.Sprintf("%s%s", dividerPrint, dividerPrint)
		}
		log.Println(dividerPrint)
	default:
		if len(datas) > 0 {
			log.Printf("[%s] [%s] %s => %v", lt, fileInfo, message, datas)
		} else {
			log.Printf("[%s] [%s] %s \n", lt, fileInfo, message)
		}
	}

	// Fatal 類型會關閉程式
	if lt == LogType.Fatal {
		os.Exit(1)
	}
}

// ex: LogTool.LogErrorN("", err)LogType.Error, "Error", err)
func LogBaseN(lt LogType.LogType, message string, datas ...interface{}) {

	log.SetFlags(0)

	log.SetPrefix(fmt.Sprint("%s ", dateTimeNowStr()))

	// 取得呼叫的 func
	_, file1, line1, fileOK1 := runtime.Caller(1)
	if !fileOK1 {
		log.Printf("%s", "[Log Tool Error] Not find call func\n")
		return
	}

	file1 = filepath.Base(file1)
	fileInfo := fmt.Sprintf("%s:%d", file1, line1)

	if file1 == "ServerLog.go" || file1 == "server-log.go" {
		_, file2, line2, fileOK2 := runtime.Caller(2)
		if !fileOK2 {
			log.Printf("%s", "[Log Tool Error] Not find call func\n")
			return
		}

		file2 = filepath.Base(file2)
		fileInfo = fmt.Sprintf("%s:%d", file2, line2)
	}

	if len(datas) > 0 {
		log.Printf("[%s] [%s] %s => ", lt, fileInfo, message)
		log.Println(datas...)
	} else {
		log.Printf("[%s] [%s] %s \n", lt, fileInfo, message)
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

func LogFatalN(message string, datas ...interface{}) {
	LogBaseN(LogType.Fatal, message, datas...)
}

func LogErrorN(message string, datas ...interface{}) {
	LogBaseN(LogType.Error, message, datas...)
}

func LogWarningN(message string, datas ...interface{}) {
	LogBaseN(LogType.Warning, message, datas...)
}

func LogInfoN(message string, datas ...interface{}) {
	LogBaseN(LogType.Info, message, datas...)
}

func LogDebugN(message string, datas ...interface{}) {
	LogBaseN(LogType.Debug, message, datas...)
}

func LogDevelopN(message string, datas ...interface{}) {
	LogBaseN(LogType.Develop, message, datas...)
}

func LogSystemN(message string, datas ...interface{}) {
	LogBaseN(LogType.System, message, datas...)
}

func LogCronN(message string, datas ...interface{}) {
	LogBaseN(LogType.Cron, message, datas...)
}

func LogConfigN(message string, datas ...interface{}) {
	LogBaseN(LogType.Config, message, datas...)
}

func LogConnectN(message string, datas ...interface{}) {
	LogBaseN(LogType.Connect, message, datas...)
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

func LogFatalf(message string, format string, datas ...any) {
	LogBasef(LogType.Fatal, message, format, datas...)
}

func LogErrorf(message string, format string, datas ...any) {
	LogBasef(LogType.Error, message, format, datas...)
}

func LogWarningf(message string, format string, datas ...any) {
	LogBasef(LogType.Warning, message, format, datas...)
}

func LogInfof(message string, format string, datas ...any) {
	LogBasef(LogType.Info, message, format, datas...)
}

func LogDebugf(message string, format string, datas ...any) {
	LogBasef(LogType.Debug, message, format, datas...)
}

func LogDevelopf(message string, format string, datas ...any) {
	LogBasef(LogType.Develop, message, format, datas...)
}

func LogSystemf(message string, format string, datas ...any) {
	LogBasef(LogType.System, message, format, datas...)
}

func LogCronf(message string, format string, datas ...any) {
	LogBasef(LogType.Cron, message, format, datas...)
}

func LogConfigf(message string, format string, datas ...any) {
	LogBasef(LogType.Config, message, format, datas...)
}

func LogConnectf(message string, format string, datas ...any) {
	LogBasef(LogType.Connect, message, format, datas...)
}
