// 用户界面模块
package ui

import (
	"calculator-go/internal/calculator"
	"calculator-go/internal/config"
	"fmt"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	fyne "fyne.io/fyne/v2"
)

// UI 计算器UI结构体
type UI struct {
	app        fyne.App
	window     fyne.Window
	calculator *calculator.Calculator
	display    *widget.Entry
	config     *config.Config
	historyList *widget.List
}

// New 创建新的UI实例
func New() *UI {
	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		cfg = config.DefaultConfig()
	}

	// 创建应用和窗口
	myApp := app.New()
	myApp.SetIcon(nil) // 可以在这里设置应用图标
	myWindow := myApp.NewWindow("计算器")
	myWindow.Resize(fyne.NewSize(float32(cfg.WindowWidth), float32(cfg.WindowHeight)))
	myWindow.SetFixedSize(true)

	// 创建计算器实例
	calc := calculator.New()
	calc.SetHistory(cfg.History)

	ui := &UI{
		app:        myApp,
		window:     myWindow,
		calculator: calc,
		config:     cfg,
	}

	// 构建UI
	ui.buildUI()

	return ui
}

// buildUI 构建用户界面
func (ui *UI) buildUI() {
	// 创建显示屏
	ui.display = widget.NewEntry()
	ui.display.SetText(ui.calculator.GetDisplay())
	ui.display.Disable() // 禁用编辑

	// 创建按钮网格
	buttonGrid := ui.createButtonGrid()

	// 创建计算器标签页
	calcTab := container.NewVBox(
		ui.display,
		buttonGrid,
	)

	// 创建历史记录标签页
	historyTab := ui.createHistoryTab()

	// 创建标签页容器
	tabs := container.NewAppTabs(
		container.NewTabItem("计算器", calcTab),
		container.NewTabItem("历史", historyTab),
	)

	// 创建菜单
	menu := ui.createMenu()
	ui.window.SetMainMenu(menu)

	// 设置窗口内容
	ui.window.SetContent(tabs)
}

// createButtonGrid 创建按钮网格
func (ui *UI) createButtonGrid() *container.GridWithColumns {
	// 定义按钮布局
	buttonLayout := [][]string{
		{"C", "←", "%", "÷"},
		{"7", "8", "9", "×"},
		{"4", "5", "6", "-"},
		{"1", "2", "3", "+"},
		{"0", ".", "=", ""},
	}

	var buttons []fyne.CanvasObject

	for _, row := range buttonLayout {
		for _, label := range row {
			if label == "" {
				// 空按钮位置
				buttons = append(buttons, widget.NewLabel(""))
				continue
			}

			button := ui.createButton(label)
			buttons = append(buttons, button)
		}
	}

	return container.NewGridWithColumns(4, buttons...)
}

// createButton 创建单个按钮
func (ui *UI) createButton(label string) *widget.Button {
	button := widget.NewButton(label, func() {
		ui.handleButtonClick(label)
	})

	// 设置按钮样式
	switch label {
	case "=":
		button.Importance = widget.HighImportance
	case "C", "←":
		button.Importance = widget.MediumImportance
	case "+", "-", "×", "÷", "%":
		button.Importance = widget.LowImportance
	}

	return button
}

// createHistoryTab 创建历史记录标签页
func (ui *UI) createHistoryTab() fyne.CanvasObject {
	ui.historyList = widget.NewList(
		func() int {
			return len(ui.calculator.GetHistory())
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			history := ui.calculator.GetHistory()
			if id < len(history) {
				obj.(*widget.Label).SetText(history[len(history)-1-id]) // 倒序显示
			}
		},
	)

	clearHistoryBtn := widget.NewButton("清除历史", func() {
		ui.calculator.ClearHistory()
		ui.historyList.Refresh()
	})
	clearHistoryBtn.Importance = widget.DangerImportance

	return container.NewVBox(
		widget.NewLabel("计算历史:"),
		ui.historyList,
		clearHistoryBtn,
	)
}

// createMenu 创建菜单
func (ui *UI) createMenu() *fyne.MainMenu {
	// 文件菜单
	fileMenu := fyne.NewMenu("文件",
		fyne.NewMenuItem("退出", func() {
			ui.app.Quit()
		}),
	)

	// 编辑菜单
	editMenu := fyne.NewMenu("编辑",
		fyne.NewMenuItem("清除", func() {
			ui.calculator.Clear()
			ui.updateDisplay()
		}),
		fyne.NewMenuItem("清除历史", func() {
			ui.calculator.ClearHistory()
			ui.historyList.Refresh()
		}),
	)

	// 帮助菜单
	helpMenu := fyne.NewMenu("帮助",
		fyne.NewMenuItem("关于", func() {
			ui.showAboutDialog()
		}),
	)

	return fyne.NewMainMenu(fileMenu, editMenu, helpMenu)
}

// handleButtonClick 处理按钮点击事件
func (ui *UI) handleButtonClick(label string) {
	switch label {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		ui.calculator.InputDigit(label)
	case ".":
		ui.calculator.InputDecimal()
	case "+", "-", "×", "÷", "%":
		err := ui.calculator.InputOperator(label)
		if err != nil {
			ui.showErrorDialog(err.Error())
			return
		}
	case "=":
		err := ui.calculator.Calculate()
		if err != nil {
			ui.showErrorDialog(err.Error())
			return
		}
		ui.historyList.Refresh() // 刷新历史列表
	case "C":
		ui.calculator.Clear()
	case "←":
		ui.calculator.Backspace()
	}

	ui.updateDisplay()
}

// updateDisplay 更新显示
func (ui *UI) updateDisplay() {
	ui.display.SetText(ui.calculator.GetDisplay())
}

// showErrorDialog 显示错误对话框
func (ui *UI) showErrorDialog(message string) {
	dialog.ShowError(fmt.Errorf(message), ui.window)
}

// showAboutDialog 显示关于对话框
func (ui *UI) showAboutDialog() {
	aboutText := `Go Fyne 计算器 v1.0.0

一个使用 Go 语言和 Fyne 框架构建的现代化计算器应用程序。

功能特性:
• 基本算术运算 (+, -, ×, ÷, %)
• 小数运算支持
• 计算历史记录
• 现代化图形界面
• 跨平台支持

开发者: Calculator Team
框架: Fyne v2.6.1
语言: Go 1.21+`

	dialog.ShowInformation("关于计算器", aboutText, ui.window)
}

// Run 运行应用
func (ui *UI) Run() {
	ui.window.ShowAndRun()
}

// SaveConfig 保存配置
func (ui *UI) SaveConfig() {
	// 更新配置中的历史记录
	ui.config.History = ui.calculator.GetHistory()
	
	// 保存配置
	if err := ui.config.Save(); err != nil {
		fmt.Printf("保存配置失败: %v\n", err)
	}
}