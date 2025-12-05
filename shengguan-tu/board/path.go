package board

// Path 职位间连接路径定义
type Path struct {
	Type     string // 路径类型：德、才、功、赃
	From     *Position
	To       *Position
	Distance int // 距离（步数）
}

// NewPath 创建新路径
func NewPath(pathType string, from, to *Position, distance int) *Path {
	return &Path{
		Type:     pathType,
		From:     from,
		To:       to,
		Distance: distance,
	}
}

// IsUpward 检查路径是否向上
func (p *Path) IsUpward() bool {
	return p.Distance > 0
}

// IsDownward 检查路径是否向下
func (p *Path) IsDownward() bool {
	return p.Distance < 0
}
