package game

import "shengguan-tu/player"

// 游戏规则常量定义
const (
	MinPlayers = 2
	MaxPlayers = 4
	StartLevel = 10 // 从最低官职开始
)

// CheckValidMove 检查移动是否有效
func CheckValidMove(currentLevel, direction int, maxLevel int) int {
	newLevel := currentLevel + direction

	// 确保级别在有效范围内
	if newLevel < 1 {
		newLevel = 1
	}
	if newLevel > maxLevel {
		newLevel = maxLevel
	}

	return newLevel
}

// DetermineWinner 确定获胜者
func DetermineWinner(players []*player.Player) *player.Player {
	for _, p := range players {
		if p.IsWinner() {
			return p
		}
	}
	return nil
}
