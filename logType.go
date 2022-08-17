package LogTool

type LogType int

const (
	ReturnLine LogType = iota // 換行
	Divider                   // 分隔線
	Fatal                     // 嚴重錯誤, 會關閉程式
	Error                     // 錯誤, 可能造成系統 Crash
	Warning                   // 警示, 可能會造成資料異常
	Info                      // 顯示資訊
	Debug                     // 除錯使用
	Develop                   // 開發環境才紀錄
	System                    // 系統呼叫
	Cron                      // 定時排程
	Config                    // config 資訊
	Connect                   // 發送其他主機
)

func (logType LogType) String() string {
	result := ""
	switch logType {
	case ReturnLine:
		result = "ReturnLine"
	case Divider:
		result = "Divider"
	case Fatal:
		result = "Fatal"
	case Error:
		result = "Error"
	case Warning:
		result = "Warning"
	case Info:
		result = "Info"
	case Debug:
		result = "Debug"
	case Develop:
		result = "Develop"
	case System:
		result = "System"
	case Cron:
		result = "Cron"
	case Config:
		result = "Config"
	case Connect:
		result = "Connect"
	}
	return result
}
