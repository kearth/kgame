package dice

// DiceResult 骰子结果类型
type DiceResult int

const (
	ResultDe DiceResult = iota
	ResultCai
	ResultGong
	ResultZang
)

// String 将结果转换为字符串
func (dr DiceResult) String() string {
	switch dr {
	case ResultDe:
		return "德"
	case ResultCai:
		return "才"
	case ResultGong:
		return "功"
	case ResultZang:
		return "赃"
	default:
		return "未知"
	}
}

// ParseResult 解析结果字符串为DiceResult类型
func ParseResult(s string) (DiceResult, bool) {
	switch s {
	case "德":
		return ResultDe, true
	case "才":
		return ResultCai, true
	case "功":
		return ResultGong, true
	case "赃":
		return ResultZang, true
	default:
		return 0, false
	}
}
