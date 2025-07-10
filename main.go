// 主程序入口
package main

import (
	"calculator-go/internal/ui"
	"log"
)

func main() {
	// 创建并运行计算器UI
	calcUI := ui.New()
	
	// 设置程序退出时保存配置
	defer func() {
		calcUI.SaveConfig()
		if r := recover(); r != nil {
			log.Printf("程序异常退出: %v", r)
		}
	}()
	
	// 运行应用
	calcUI.Run()
}