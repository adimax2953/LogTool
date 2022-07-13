package LogTool_test

import (
	"testing"

	LogTool "github.com/adimax2953/log-tool"
)

func Test_redis(t *testing.T) {
	LogTool.LogDebug("testLogDebug")
	LogTool.LogDivider("testLogDivider")
	LogTool.LogInfo("testLogInfo")
	LogTool.LogError("testLogError")
	LogTool.LogWarning("testLogWarning")
	LogTool.LogReturnLine()
	LogTool.LogSystem("testLogSystem")
	LogTool.LogDevelop("testLogDevelop")
	LogTool.LogConfig("testLogConfig")
	LogTool.LogCron("testLogCron")
	LogTool.LogFatal("testLogFatal")

}
