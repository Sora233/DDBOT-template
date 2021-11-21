package main

import (
	"github.com/Sora233/DDBOT"
	//_ "github.com/Sora233/DDBOT/_example/concern"
	// 别忘记在这里引入刚刚编写的插件
)

func main() {
	// 使用默认的日志格式配置
	DDBOT.SetUpLog()
	// 启动bot，会自动阻塞
	DDBOT.Run()
}
