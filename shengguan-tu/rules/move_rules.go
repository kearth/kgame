package rules

import (
	"shengguan-tu/official"
	"shengguan-tu/player"
	"strconv"
)

// CalculateMove 根据骰子结果计算移动方向
func CalculateMove(player *player.Player, diceResult string) int {
	// 获取当前官职的移动规则
	moveRules := player.CurrentPos.MoveRules

	// 根据骰子结果获取移动步数
	step, exists := moveRules[diceResult]
	if !exists {
		// 如果没有对应的规则，默认不移动
		return 0
	}

	return step
}

// ValidateMove 验证移动是否有效
func ValidateMove(player *player.Player, direction int) int {
	newLevel := player.CurrentLevel + direction

	// 确保级别在有效范围内
	if newLevel < 1 {
		newLevel = 1
	}
	if newLevel > official.MaxLevel() {
		newLevel = official.MaxLevel()
	}

	// 返回实际的移动方向
	return newLevel - player.CurrentLevel
}

// GetMoveDescription 获取移动描述
func GetMoveDescription(player *player.Player, direction int, diceResult string) string {
	if direction > 0 {
		return player.Name + " 掷出了 " + diceResult + "，升官 " + strconv.Itoa(direction) + " 级，成为 " + player.CurrentPos.Name
	} else if direction < 0 {
		return player.Name + " 掷出了 " + diceResult + "，贬官 " + strconv.Itoa(-direction) + " 级，成为 " + player.CurrentPos.Name
	} else {
		return player.Name + " 掷出了 " + diceResult + "，职位不变，仍为 " + player.CurrentPos.Name
	}
}
