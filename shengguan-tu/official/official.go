package official

// Official 官职定义
type Official struct {
	ID          string // 官职ID
	Name        string // 官职名称
	Level       int    // 官职级别（1-最高，数字越大级别越低）
	Description string // 官职描述
	// 移动规则：德、才、功、赃对应的移动步数
	MoveRules map[string]int // key: "德", "才", "功", "赃"; value: 移动步数（正数为升，负数为降）
}

// IsHighest 是否为最高官职
func (o *Official) IsHighest() bool {
	return o.Level == 1
}

// IsLowest 是否为最低官职
func (o *Official) IsLowest() bool {
	return o.Level == maxLevel
}

// MaxLevel 获取最高级别数
func MaxLevel() int {
	return maxLevel
}

// 内部常量
var maxLevel = 10 // 最高级别数
