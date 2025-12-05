package dice

import (
	"math/rand"
	"time"
)

var (
	DiceFace = []string{"德", "才", "功", "赃"}
)

// Dice 骰子
type Dice struct {
	random  *rand.Rand
	FaceLen int // 骰子面数
}

// NewDice 创建新骰子
func NewDice() *Dice {
	return &Dice{
		random:  rand.New(rand.NewSource(time.Now().UnixNano())),
		FaceLen: len(DiceFace),
	}
}

// Roll 掷骰子
func (d *Dice) Roll() string {
	results := DiceFace
	index := d.random.Intn(d.FaceLen)
	return results[index]
}
