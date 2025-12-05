package player

import "shengguan-tu/official"

// Gender 性别定义
type Gender int

const (
	Male    Gender = 0 // 男性
	Female  Gender = 1 // 女性
	Unknown Gender = 2 // 未知性别
)

// Player 玩家定义
type Player struct {
	ID           string             // 玩家ID
	Name         string             // 玩家名称
	Gender       Gender             // 性别
	CurrentLevel int                // 当前官职级别
	CurrentPos   *official.Official // 当前官职
	TotalMoves   int                // 总移动次数
}

// NewPlayer 创建新玩家
func NewPlayer(id, name string, gender Gender) (*Player, error) {
	// 从最低官职开始
	lowest, err := official.GetOfficialByLevel(official.MaxLevel())
	if err != nil {
		return nil, err
	}

	return &Player{
		ID:           id,
		Name:         name,
		CurrentLevel: lowest.Level,
		CurrentPos:   lowest,
		TotalMoves:   0,
	}, nil
}

// Move 移动玩家职位
func (p *Player) Move(direction int) error {
	newLevel := p.CurrentLevel + direction

	// 确保级别在有效范围内
	if newLevel < 1 {
		newLevel = 1
	}
	if newLevel > official.MaxLevel() {
		newLevel = official.MaxLevel()
	}

	// 获取新官职
	newPos, err := official.GetOfficialByLevel(newLevel)
	if err != nil {
		return err
	}

	p.CurrentLevel = newLevel
	p.CurrentPos = newPos
	p.TotalMoves++

	return nil
}

// IsWinner 检查玩家是否获胜
func (p *Player) IsWinner() bool {
	return p.CurrentPos.IsHighest()
}
