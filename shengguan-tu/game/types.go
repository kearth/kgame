package game

// GameState 游戏状态
type GameState int

const (
	GameStateInit GameState = iota
	GameStatePlaying
	GameStatePaused
	GameStateEnded
)

// DiceResult 骰子结果
type DiceResult int

const (
	DiceResultDe   DiceResult = iota // 德
	DiceResultCai                    // 才
	DiceResultGong                   // 功
	DiceResultZang                   // 赃
)

// PlayerID 玩家ID
type PlayerID string

// PositionID 职位ID
type PositionID string

// MoveResult 移动结果
type MoveResult struct {
	Success      bool
	FromPosition PositionID
	ToPosition   PositionID
	Message      string
}

// GameResult 游戏结果
type GameResult struct {
	Winner     PlayerID
	Rankings   []PlayerID
	TotalMoves int
}
