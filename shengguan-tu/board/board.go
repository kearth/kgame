package board

import "shengguan-tu/official"

// Board 棋盘定义
type Board struct {
	Positions []*official.Official // 所有职位
}

// NewBoard 创建新棋盘
func NewBoard() *Board {
	return &Board{
		Positions: official.DefaultOfficials(),
	}
}

// GetPosition 获取指定ID的职位
func (b *Board) GetPosition(id string) (*official.Official, error) {
	return official.GetOfficialByID(id)
}

// GetPositionByLevel 获取指定级别的职位
func (b *Board) GetPositionByLevel(level int) (*official.Official, error) {
	return official.GetOfficialByLevel(level)
}

// GetNextPosition 获取下一个职位
func (b *Board) GetNextPosition(current *official.Official, direction int) (*official.Official, error) {
	newLevel := current.Level + direction
	return b.GetPositionByLevel(newLevel)
}

// GetTotalPositions 获取总职位数
func (b *Board) GetTotalPositions() int {
	return len(b.Positions)
}
