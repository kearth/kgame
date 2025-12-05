package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Input 输入处理工具
type Input struct {
	Scanner *bufio.Scanner
}

// NewInput 创建新的输入处理工具
func NewInput() *Input {
	return &Input{
		Scanner: bufio.NewScanner(os.Stdin),
	}
}

// GetString 获取字符串输入
func (i *Input) GetString(prompt string) string {
	fmt.Print(prompt)
	i.Scanner.Scan()
	return strings.TrimSpace(i.Scanner.Text())
}

// GetInt 获取整数输入
func (i *Input) GetInt(prompt string) (int, error) {
	for {
		input := i.GetString(prompt)
		value, err := strconv.Atoi(input)
		if err == nil {
			return value, nil
		}
		fmt.Println("请输入有效的整数")
	}
}

// GetChoice 获取选择输入
func (i *Input) GetChoice(prompt string, options []string) (int, error) {
	fmt.Println(prompt)
	for idx, option := range options {
		fmt.Printf("%d. %s\n", idx+1, option)
	}

	for {
		choice, err := i.GetInt("请选择: ")
		if err != nil {
			return 0, err
		}

		if choice >= 1 && choice <= len(options) {
			return choice, nil
		}

		fmt.Printf("请输入1-%d之间的数字\n", len(options))
	}
}

// Confirm 获取确认输入
func (i *Input) Confirm(prompt string) bool {
	for {
		input := strings.ToLower(i.GetString(prompt + " (y/n): "))
		switch input {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println("请输入 y 或 n")
		}
	}
}

// WaitForEnter 等待回车键
func (i *Input) WaitForEnter(prompt string) {
	if prompt != "" {
		fmt.Print(prompt)
	}
	i.Scanner.Scan()
}
