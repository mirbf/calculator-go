package calculator

import (
	"testing"
)

func TestNew(t *testing.T) {
	calc := New()
	if calc.GetDisplay() != "0" {
		t.Errorf("Expected display to be '0', got '%s'", calc.GetDisplay())
	}
}

func TestInputDigit(t *testing.T) {
	calc := New()
	calc.InputDigit("5")
	if calc.GetDisplay() != "5" {
		t.Errorf("Expected display to be '5', got '%s'", calc.GetDisplay())
	}
	
	calc.InputDigit("3")
	if calc.GetDisplay() != "53" {
		t.Errorf("Expected display to be '53', got '%s'", calc.GetDisplay())
	}
}

func TestInputDecimal(t *testing.T) {
	calc := New()
	calc.InputDigit("5")
	calc.InputDecimal()
	calc.InputDigit("2")
	if calc.GetDisplay() != "5.2" {
		t.Errorf("Expected display to be '5.2', got '%s'", calc.GetDisplay())
	}
	
	// 测试重复输入小数点
	calc.InputDecimal()
	if calc.GetDisplay() != "5.2" {
		t.Errorf("Expected display to remain '5.2', got '%s'", calc.GetDisplay())
	}
}

func TestBasicArithmetic(t *testing.T) {
	calc := New()
	
	// 测试加法: 5 + 3 = 8
	calc.InputDigit("5")
	calc.InputOperator("+")
	calc.InputDigit("3")
	err := calc.Calculate()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if calc.GetDisplay() != "8" {
		t.Errorf("Expected display to be '8', got '%s'", calc.GetDisplay())
	}
	
	// 测试减法: 10 - 4 = 6
	calc.Clear()
	calc.InputDigit("1")
	calc.InputDigit("0")
	calc.InputOperator("-")
	calc.InputDigit("4")
	err = calc.Calculate()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if calc.GetDisplay() != "6" {
		t.Errorf("Expected display to be '6', got '%s'", calc.GetDisplay())
	}
	
	// 测试乘法: 6 × 7 = 42
	calc.Clear()
	calc.InputDigit("6")
	calc.InputOperator("×")
	calc.InputDigit("7")
	err = calc.Calculate()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if calc.GetDisplay() != "42" {
		t.Errorf("Expected display to be '42', got '%s'", calc.GetDisplay())
	}
	
	// 测试除法: 15 ÷ 3 = 5
	calc.Clear()
	calc.InputDigit("1")
	calc.InputDigit("5")
	calc.InputOperator("÷")
	calc.InputDigit("3")
	err = calc.Calculate()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if calc.GetDisplay() != "5" {
		t.Errorf("Expected display to be '5', got '%s'", calc.GetDisplay())
	}
}

func TestDivisionByZero(t *testing.T) {
	calc := New()
	calc.InputDigit("5")
	calc.InputOperator("÷")
	calc.InputDigit("0")
	err := calc.Calculate()
	if err == nil {
		t.Error("Expected error for division by zero")
	}
}

func TestClear(t *testing.T) {
	calc := New()
	calc.InputDigit("1")
	calc.InputDigit("2")
	calc.InputDigit("3")
	calc.Clear()
	if calc.GetDisplay() != "0" {
		t.Errorf("Expected display to be '0' after clear, got '%s'", calc.GetDisplay())
	}
}

func TestBackspace(t *testing.T) {
	calc := New()
	calc.InputDigit("1")
	calc.InputDigit("2")
	calc.InputDigit("3")
	calc.Backspace()
	if calc.GetDisplay() != "12" {
		t.Errorf("Expected display to be '12' after backspace, got '%s'", calc.GetDisplay())
	}
	
	// 测试删除到只剩一位数字
	calc.Backspace()
	calc.Backspace()
	if calc.GetDisplay() != "0" {
		t.Errorf("Expected display to be '0' after backspacing all digits, got '%s'", calc.GetDisplay())
	}
}

func TestHistory(t *testing.T) {
	calc := New()
	
	// 执行一个计算
	calc.InputDigit("5")
	calc.InputOperator("+")
	calc.InputDigit("3")
	calc.Calculate()
	
	history := calc.GetHistory()
	if len(history) != 1 {
		t.Errorf("Expected history length to be 1, got %d", len(history))
	}
	
	if history[0] != "5 + 3 = 8" {
		t.Errorf("Expected history entry to be '5 + 3 = 8', got '%s'", history[0])
	}
	
	// 清除历史
	calc.ClearHistory()
	history = calc.GetHistory()
	if len(history) != 0 {
		t.Errorf("Expected history length to be 0 after clear, got %d", len(history))
	}
}

func TestContinuousCalculation(t *testing.T) {
	calc := New()
	
	// 5 + 3 = 8, 然后 + 2 = 10
	calc.InputDigit("5")
	calc.InputOperator("+")
	calc.InputDigit("3")
	calc.InputOperator("+") // 这应该计算 5+3 并继续
	if calc.GetDisplay() != "8" {
		t.Errorf("Expected display to be '8', got '%s'", calc.GetDisplay())
	}
	
	calc.InputDigit("2")
	calc.Calculate()
	if calc.GetDisplay() != "10" {
		t.Errorf("Expected display to be '10', got '%s'", calc.GetDisplay())
	}
}