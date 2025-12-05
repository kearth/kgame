package main

import (
	"fmt"
	"log"
	"os"
	"shengguan-tu/game"
	"shengguan-tu/ui"
)

func main() {
	fmt.Println("欢迎来到升官图游戏！")
	fmt.Println("===================")

	// 初始化游戏
	g, err := game.NewGame()
	if err != nil {
		log.Fatalf("初始化游戏失败: %v", err)
	}

	// 初始化UI
	cliUI := ui.NewCLI(g)

	// 开始游戏
	if err := cliUI.Start(); err != nil {
		log.Printf("游戏运行失败: %v", err)
	}

	fmt.Println("游戏结束，感谢参与！")
	os.Exit(0)
}
