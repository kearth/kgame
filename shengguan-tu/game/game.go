package game

import (
	"fmt"
	"math/rand/v2"
	"shengguan-tu/board"
	"shengguan-tu/dice"
	"shengguan-tu/player"
	"time"
)

// Game 游戏核心结构体
type Game struct {
	State         GameState             // 游戏状态
	Board         *board.Board          // 棋盘
	Dice          *dice.Dice            // 骰子
	PlayerManager *player.PlayerManager // 玩家管理器
	NameManager   *player.NameManager   // 玩家姓名管理器
	CurrentTurn   int                   // 当前玩家索引
	TotalTurns    int                   // 总轮数
	Winner        *player.Player        // 赢家
	Random        *rand.Rand            // 随机数生成器
}

// NewGame 创建新游戏
func NewGame() (*Game, error) {
	random := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano())))
	return &Game{
		State:         GameStateInit,
		Board:         board.NewBoard(),
		Dice:          dice.NewDice(),
		PlayerManager: player.NewPlayerManager(),
		NameManager:   player.NewNameManager(random),
		CurrentTurn:   0,
		TotalTurns:    0,
		Winner:        nil,
		Random:        random,
	}, nil
}

// AddPlayer 添加玩家
func (g *Game) AddPlayer(id, name string, gender player.Gender) error {
	p, err := player.NewPlayer(id, name, gender)
	if err != nil {
		return err
	}

	g.PlayerManager.AddPlayer(p)
	return nil
}

// StartGame 开始游戏
func (g *Game) StartGame() error {
	if !g.PlayerManager.PlayersNumIsValid() {
		return fmt.Errorf("玩家数量必须在%d到%d之间", g.PlayerManager.AllowMinPlayers(), g.PlayerManager.AllowMaxPlayers())
	}

	g.State = GameStatePlaying
	return nil
}

// NextTurn 下一轮
func (g *Game) NextTurn() error {
	if g.State != GameStatePlaying {
		return fmt.Errorf("游戏未在进行中")
	}

	// 获取当前玩家
	currentPlayer := g.PlayerManager.GetCurrentPlayer()
	if currentPlayer == nil {
		return fmt.Errorf("当前玩家不存在")
	}

	// 掷骰子
	result := g.Dice.Roll()

	// 计算移动方向
	moveStep := currentPlayer.CurrentPos.MoveRules[result]

	// 移动玩家
	if err := currentPlayer.Move(moveStep); err != nil {
		return err
	}

	// 检查是否获胜
	if currentPlayer.IsWinner() {
		g.Winner = currentPlayer
		g.State = GameStateEnded
		return nil
	}

	// 切换到下一个玩家
	g.PlayerManager.NextPlayer()
	g.TotalTurns++

	return nil
}

// GetGameResult 获取游戏结果
func (g *Game) GetGameResult() *GameResult {
	if g.Winner == nil {
		return nil
	}

	// 简单排序（按级别排序）
	// 实际实现可以更复杂
	return &GameResult{
		Winner:     PlayerID(g.Winner.ID),
		Rankings:   []PlayerID{PlayerID(g.Winner.ID)},
		TotalMoves: g.Winner.TotalMoves,
	}
}
