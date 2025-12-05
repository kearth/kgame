package ui

import (
	"bufio"
	"fmt"
	"os"
	"shengguan-tu/game"
	"shengguan-tu/player"
	"strconv"
	"strings"
)

// CLI 命令行界面
type CLI struct {
	Game    *game.Game
	Scanner *bufio.Scanner
}

// NewCLI 创建新的CLI界面
func NewCLI(g *game.Game) *CLI {
	return &CLI{
		Game:    g,
		Scanner: bufio.NewScanner(os.Stdin),
	}
}

// Start 启动CLI界面
func (cli *CLI) Start() error {
	// 欢迎信息
	cli.showWelcome()

	// 设置玩家
	if err := cli.setupPlayers(); err != nil {
		return err
	}

	// 开始游戏
	if err := cli.Game.StartGame(); err != nil {
		return err
	}

	// 游戏主循环
	for cli.Game.State == game.GameStatePlaying {
		// 显示当前游戏状态
		cli.showGameState()

		// 等待玩家输入
		if err := cli.waitForInput(); err != nil {
			return err
		}

		// 执行回合
		if err := cli.Game.NextTurn(); err != nil {
			return err
		}

		// 检查游戏是否结束
		if cli.Game.State == game.GameStateEnded {
			cli.showGameResult()
			break
		}
	}

	return nil
}

// showWelcome 显示欢迎信息
func (cli *CLI) showWelcome() {
	fmt.Println("====================================")
	fmt.Println("          欢迎来到升官图游戏          ")
	fmt.Println("====================================")
	fmt.Println()
	fmt.Println("游戏规则：")
	fmt.Println("1. 玩家从'白丁'开始，通过掷骰子决定官职升降")
	fmt.Println("2. 骰子有四种结果：德、才、功、赃")
	fmt.Println("3. 先升至最高官职'太师'者获胜")
	fmt.Println()
}

// setupPlayers 设置玩家
func (cli *CLI) setupPlayers() error {
	fmt.Printf("请输入玩家数量（%d-%d人）：\n", cli.Game.PlayerManager.AllowMinPlayers(), cli.Game.PlayerManager.AllowMaxPlayers())

	for {
		cli.Scanner.Scan()
		input := strings.TrimSpace(cli.Scanner.Text())

		playerCount, err := strconv.Atoi(input)
		if isValueValid := playerCount >= cli.Game.PlayerManager.AllowMinPlayers() && playerCount <= cli.Game.PlayerManager.AllowMaxPlayers(); !isValueValid {
			fmt.Printf("请输入有效的玩家数量（%d-%d人）：\n", cli.Game.PlayerManager.AllowMinPlayers(), cli.Game.PlayerManager.AllowMaxPlayers())
			continue
		}
		if err != nil {
			fmt.Printf("请输入有效的玩家数量（%d-%d人）：\n", cli.Game.PlayerManager.AllowMinPlayers(), cli.Game.PlayerManager.AllowMaxPlayers())
			continue
		}

		// 添加玩家
		for i := 1; i <= playerCount; i++ {
			// 升官图都是男的，默认性别为男
			fmt.Printf("请输入第 %d 位玩家的姓名：\n", i)
			cli.Scanner.Scan()
			name := strings.TrimSpace(cli.Scanner.Text())
			if name == "" {
				name = cli.Game.NameManager.GenerateName(player.Male)
			}

			if err := cli.Game.AddPlayer(fmt.Sprintf("player_%d", i), name, player.Male); err != nil {
				return err
			}
		}

		break
	}

	fmt.Println()
	fmt.Println("玩家设置完成，游戏即将开始！")
	fmt.Println()

	return nil
}

// showGameState 显示当前游戏状态
func (cli *CLI) showGameState() {
	fmt.Println("====================================")
	fmt.Println("            当前游戏状态              ")
	fmt.Println("====================================")

	// 显示玩家信息
	for i, player := range cli.Game.PlayerManager.Players {
		fmt.Printf("玩家 %d: %s - 当前职位: %s (级别: %d)\n", i+1, player.Name, player.CurrentPos.Name, player.CurrentLevel)
	}

	fmt.Println()
	fmt.Printf("当前回合: 玩家 %s\n", cli.Game.PlayerManager.Players[cli.Game.CurrentTurn].Name)
	fmt.Println()
}

// waitForInput 等待玩家输入
func (cli *CLI) waitForInput() error {
	fmt.Println("按回车键继续（输入 'exit' 退出游戏）：")

	cli.Scanner.Scan()
	input := strings.TrimSpace(cli.Scanner.Text())

	if input == "exit" {
		return fmt.Errorf("玩家退出游戏")
	}

	return nil
}

// showGameResult 显示游戏结果
func (cli *CLI) showGameResult() {
	fmt.Println("====================================")
	fmt.Println("              游戏结束                ")
	fmt.Println("====================================")

	if cli.Game.Winner != nil {
		fmt.Printf("恭喜 %s 成为太师，获得胜利！\n", cli.Game.Winner.Name)
		fmt.Printf("总回合数: %d\n", cli.Game.TotalTurns)
	}

	fmt.Println()
	fmt.Println("最终排名：")

	// 简单排序（按官职级别排序）
	// 实际实现可以更复杂
	for i, player := range cli.Game.PlayerManager.Players {
		fmt.Printf("第 %d 名: %s - %s\n", i+1, player.Name, player.CurrentPos.Name)
	}

	fmt.Println()
	fmt.Println("感谢参与！")
	fmt.Println()
}
