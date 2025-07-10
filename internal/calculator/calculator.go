// 计算器核心逻辑模块
package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Calculator 计算器结构体
type Calculator struct {
	displayValue string    // 当前显示值
	operator     string    // 当前运算符
	operand      float64   // 操作数
	waitingForOperand bool // 是否等待操作数
	history      []string  // 计算历史
}

// HistoryEntry 历史记录条目
type HistoryEntry struct {
	Expression string    `json:"expression"`
	Result     string    `json:"result"`
	Timestamp  time.Time `json:"timestamp"`
}

// New 创建新的计算器实例
func New() *Calculator {
	return &Calculator{
		displayValue:      "0",
		operator:          "",
		operand:           0,
		waitingForOperand: false,
		history:           make([]string, 0),
	}
}

// GetDisplay 获取当前显示值
func (c *Calculator) GetDisplay() string {
	return c.displayValue
}

// InputDigit 输入数字
func (c *Calculator) InputDigit(digit string) {
	if c.waitingForOperand {
		c.displayValue = digit
		c.waitingForOperand = false
	} else {
		if c.displayValue == "0" {
			c.displayValue = digit
		} else {
			c.displayValue += digit
		}
	}
}

// InputDecimal 输入小数点
func (c *Calculator) InputDecimal() {
	if c.waitingForOperand {
		c.displayValue = "0."
		c.waitingForOperand = false
	} else if !strings.Contains(c.displayValue, ".") {
		c.displayValue += "."
	}
}

// InputOperator 输入运算符
func (c *Calculator) InputOperator(nextOperator string) error {
	inputValue, err := strconv.ParseFloat(c.displayValue, 64)
	if err != nil {
		return fmt.Errorf("无效的数字: %v", err)
	}

	if !c.waitingForOperand && c.operator != "" {
		result, err := c.calculate(c.operand, inputValue, c.operator)
		if err != nil {
			return err
		}
		c.displayValue = c.formatResult(result)
		c.operand = result
	} else {
		c.operand = inputValue
	}

	c.waitingForOperand = true
	c.operator = nextOperator
	return nil
}

// Calculate 执行计算
func (c *Calculator) Calculate() error {
	inputValue, err := strconv.ParseFloat(c.displayValue, 64)
	if err != nil {
		return fmt.Errorf("无效的数字: %v", err)
	}

	if c.operator != "" {
		// 记录计算表达式
		expression := fmt.Sprintf("%s %s %s", c.formatResult(c.operand), c.operator, c.displayValue)
		
		result, err := c.calculate(c.operand, inputValue, c.operator)
		if err != nil {
			return err
		}
		
		c.displayValue = c.formatResult(result)
		c.operand = result
		
		// 添加到历史记录
		historyEntry := fmt.Sprintf("%s = %s", expression, c.displayValue)
		c.history = append(c.history, historyEntry)
		
		c.operator = ""
		c.waitingForOperand = true
	}

	return nil
}

// calculate 执行具体的计算操作
func (c *Calculator) calculate(firstOperand, secondOperand float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return firstOperand + secondOperand, nil
	case "-":
		return firstOperand - secondOperand, nil
	case "×":
		return firstOperand * secondOperand, nil
	case "÷":
		if secondOperand == 0 {
			return 0, errors.New("除数不能为零")
		}
		return firstOperand / secondOperand, nil
	case "%":
		if secondOperand == 0 {
			return 0, errors.New("除数不能为零")
		}
		return float64(int(firstOperand) % int(secondOperand)), nil
	default:
		return secondOperand, nil
	}
}

// formatResult 格式化结果显示
func (c *Calculator) formatResult(value float64) string {
	// 如果是整数，不显示小数点
	if value == float64(int64(value)) {
		return fmt.Sprintf("%.0f", value)
	}
	// 否则显示最多6位小数，去除尾随零
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.6f", value), "0"), ".")
}

// Clear 清除所有内容
func (c *Calculator) Clear() {
	c.displayValue = "0"
	c.operator = ""
	c.operand = 0
	c.waitingForOperand = false
}

// Backspace 退格删除
func (c *Calculator) Backspace() {
	if len(c.displayValue) > 1 {
		c.displayValue = c.displayValue[:len(c.displayValue)-1]
	} else {
		c.displayValue = "0"
	}
}

// GetHistory 获取计算历史
func (c *Calculator) GetHistory() []string {
	return c.history
}

// ClearHistory 清除历史记录
func (c *Calculator) ClearHistory() {
	c.history = make([]string, 0)
}

// SetHistory 设置历史记录（用于从配置文件加载）
func (c *Calculator) SetHistory(history []string) {
	c.history = history
}