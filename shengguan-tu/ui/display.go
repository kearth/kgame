package ui

import (
	"fmt"
	"shengguan-tu/player"
)

// Display 显示工具
type Display struct{}

// NewDisplay 创建新的显示工具
func NewDisplay() *Display {
	return &Display{}
}

// ShowBoard 显示棋盘
func (d *Display) ShowBoard() {
	fmt.Println("====================================")
	fmt.Println("                升官图                ")
	fmt.Println("====================================")
	fmt.Println()
	fmt.Println("1. 太师 (最高官职)")
	fmt.Println("2. 尚书")
	fmt.Println("3. 谏议大夫")
	fmt.Println("4. 侍中")
	fmt.Println("5. 州刺史")
	fmt.Println("6. 县令")
	fmt.Println("7. 牧长")
	fmt.Println("8. 亭长")
	fmt.Println("9. 里正")
	fmt.Println("10. 白丁 (最低官职)")
	fmt.Println()
}

// ShowPlayerStatus 显示玩家状态
func (d *Display) ShowPlayerStatus(players []*player.Player) {
	fmt.Println("====================================")
	fmt.Println("              玩家状态                ")
	fmt.Println("====================================")

	for i, p := range players {
		fmt.Printf("玩家 %d: %s\n", i+1, p.Name)
		fmt.Printf("当前职位: %s (级别: %d)\n", p.CurrentPos.Name, p.CurrentLevel)
		fmt.Printf("总移动次数: %d\n", p.TotalMoves)
		fmt.Println()
	}
}

// ShowDiceResult 显示骰子结果
func (d *Display) ShowDiceResult(result string) {
	fmt.Println("====================================")
	fmt.Printf("            骰子结果: %s            \n", result)
	fmt.Println("====================================")
	fmt.Println()
}

// ShowMoveResult 显示移动结果
func (d *Display) ShowMoveResult(player *player.Player, direction int, diceResult string) {
	var action string
	var step string

	if direction > 0 {
		action = "升官"
		step = fmt.Sprintf("%d", direction)
	} else if direction < 0 {
		action = "贬官"
		step = fmt.Sprintf("%d", -direction)
	} else {
		action = "职位不变"
		step = "0"
	}

	fmt.Printf("%s 掷出了 %s，%s %s 级，成为 %s\n", player.Name, diceResult, action, step, player.CurrentPos.Name)
	fmt.Println()
}

// ShowWinner 显示获胜者
func (d *Display) ShowWinner(winner *player.Player) {
	fmt.Println("====================================")
	fmt.Println("              游戏获胜                ")
	fmt.Println("====================================")
	fmt.Printf("恭喜 %s 成为太师，获得胜利！\n", winner.Name)
	fmt.Println()
}

// ShowSeparator 显示分隔符
func (d *Display) ShowSeparator() {
	fmt.Println("------------------------------------")
}
