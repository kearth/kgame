package rules

import (
	"shengguan-tu/official"
	"shengguan-tu/player"
)

// CanDemote 检查是否可以贬谪
func CanDemote(player *player.Player, direction int) bool {
	newLevel := player.CurrentLevel + direction
	return newLevel > player.CurrentLevel && newLevel <= official.MaxLevel()
}

// GetDemotionPenalty 获取贬谪惩罚
func GetDemotionPenalty(player *player.Player) string {
	// 根据官职级别返回不同的惩罚描述
	switch {
	case player.CurrentLevel >= official.MaxLevel()-2:
		return player.Name + " 被贬至低位，需更加谨慎！"
	case player.CurrentLevel >= official.MaxLevel()-5:
		return player.Name + " 仕途受挫，需重新振作！"
	default:
		return player.Name + " 官职被贬，切勿灰心！"
	}
}

// GetDemotionMessage 获取贬谪消息
func GetDemotionMessage(player *player.Player, steps int) string {
	fromOfficial, _ := official.GetOfficialByLevel(player.CurrentLevel - steps)
	return player.Name + " 从 " + fromOfficial.Name + " 被贬为 " + player.CurrentPos.Name + "！"
}
