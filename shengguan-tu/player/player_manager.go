package player

// PlayerManager 玩家管理器
type PlayerManager struct {
	Players    []*Player // 玩家列表
	CurrentIdx int       // 当前玩家索引
	TotalTurns int       // 总轮数
	MaxPlayers int       // 最大玩家数量
	MinPlayers int       // 最小玩家数量
}

// NewPlayerManager 创建玩家管理器
func NewPlayerManager() *PlayerManager {
	return &PlayerManager{
		Players:    []*Player{},
		CurrentIdx: 0,
		TotalTurns: 0,
		MaxPlayers: 4,
		MinPlayers: 2,
	}
}

// AddPlayer 添加玩家
func (pm *PlayerManager) AddPlayer(player *Player) {
	pm.Players = append(pm.Players, player)
}

// GetCurrentPlayer 获取当前回合玩家
func (pm *PlayerManager) GetCurrentPlayer() *Player {
	if len(pm.Players) == 0 {
		return nil
	}
	return pm.Players[pm.CurrentIdx]
}

// NextPlayer 切换到下一个玩家
func (pm *PlayerManager) NextPlayer() *Player {
	if len(pm.Players) == 0 {
		return nil
	}

	pm.CurrentIdx = (pm.CurrentIdx + 1) % len(pm.Players)
	pm.TotalTurns++

	return pm.Players[pm.CurrentIdx]
}

// GetPlayerCount 获取玩家数量
func (pm *PlayerManager) GetPlayerCount() int {
	return len(pm.Players)
}

// GetWinner 获取获胜者
func (pm *PlayerManager) GetWinner() *Player {
	for _, player := range pm.Players {
		if player.IsWinner() {
			return player
		}
	}
	return nil
}

// HasWinner 检查是否有获胜者
func (pm *PlayerManager) HasWinner() bool {
	return pm.GetWinner() != nil
}

// AllowMaxPlayers 允许的最大玩家数量
func (pm *PlayerManager) AllowMaxPlayers() int {
	return pm.MaxPlayers
}

// AllowMinPlayers 允许的最小玩家数量
func (pm *PlayerManager) AllowMinPlayers() int {
	return pm.MinPlayers
}

// PlayersNumIsValid 检查玩家数量是否有效
func (pm *PlayerManager) PlayersNumIsValid() bool {
	return pm.GetPlayerCount() >= pm.MinPlayers && pm.GetPlayerCount() <= pm.MaxPlayers
}
