package LogTool_test

import (
	"testing"

	LogTool "github.com/adimax2953/log-tool"
)

func Test_log(t *testing.T) {

	LogTool.LogDivider("testLogDivider")
	LogTool.LogReturnLine()
	LogTool.LogDebug("testLogDebug", 3)
	LogTool.LogInfo("testLogInfo", "hello")
	LogTool.LogError("testLogError", "hello")
	LogTool.LogWarning("testLogWarning", "hello")
	LogTool.LogSystem("testLogSystem", "hello")
	LogTool.LogDevelop("testLogDevelop", "hello")
	LogTool.LogConfig("testLogConfig", "hello")
	LogTool.LogCron("testLogCron", "hello")
	//LogTool.LogFatal("testLogFatal", "hello")
}
func Test_logN(t *testing.T) {

	LogTool.LogDebugN("testLogDebugN", 3)
	LogTool.LogInfoN("testLogInfoN", "hello")
	LogTool.LogErrorN("testLogErrorN", "hello")
	LogTool.LogWarningN("testLogWarningN", "hello")
	LogTool.LogSystemN("testLogSystemN", "hello")
	LogTool.LogDevelopN("testLogDevelopN", "hello")
	LogTool.LogConfigN("testLogConfigN", "hello")
	LogTool.LogCronN("testLogCronN", "hello")
	LogTool.LogFatalN("testLogFatalN", "hello")
}
