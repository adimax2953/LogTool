# log-tool

Packet Log-Tool implements a way to use Log more easily

## Install

```console
go get -u -v github.com/adimax2953/log-tool
```


```go
package main

import (
    "github.com/adimax2953/log-tool"
)


func main() {
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
```


## TODO

1. [X] Add test cases using "testify".
2. [X] Add log test method.
4. [ ] Improve or remove useless code.
5. [ ] Check code formatting.
