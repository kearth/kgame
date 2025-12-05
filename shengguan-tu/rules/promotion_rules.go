package rules

import (
	"shengguan-tu/official"
	"shengguan-tu/player"
)

// CanPromote 检查是否可以升迁
func CanPromote(player *player.Player, direction int) bool {
	newLevel := player.CurrentLevel + direction
	return newLevel < player.CurrentLevel && newLevel >= 1
}

// GetPromotionBonus 获取升迁奖励
func GetPromotionBonus(player *player.Player) string {
	// 根据官职级别返回不同的奖励描述
	switch {
	case player.CurrentLevel == 1:
		return "恭喜！" + player.Name + " 成为太师，位极人臣！"
	case player.CurrentLevel <= 3:
		return player.Name + " 升至高级官职，获得朝廷重用！"
	case player.CurrentLevel <= 6:
		return player.Name + " 官运亨通，仕途顺利！"
	default:
		return player.Name + " 官职晋升，继续努力！"
	}
}

// GetPromotionMessage 获取升迁消息
func GetPromotionMessage(player *player.Player, steps int) string {
	fromOfficial, _ := official.GetOfficialByLevel(player.CurrentLevel + steps)
	return player.Name + " 从 " + fromOfficial.Name + " 晋升为 " + player.CurrentPos.Name + "！"
}
