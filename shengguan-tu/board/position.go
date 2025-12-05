package board

import "shengguan-tu/official"

// Position 职位节点定义
type Position struct {
	Official  *official.Official
	UpPaths   []*Path
	DownPaths []*Path
}

// NewPosition 创建新职位节点
func NewPosition(off *official.Official) *Position {
	return &Position{
		Official:  off,
		UpPaths:   []*Path{},
		DownPaths: []*Path{},
	}
}

// AddUpPath 添加向上路径
func (p *Position) AddUpPath(path *Path) {
	p.UpPaths = append(p.UpPaths, path)
}

// AddDownPath 添加向下路径
func (p *Position) AddDownPath(path *Path) {
	p.DownPaths = append(p.DownPaths, path)
}

// GetPathByType 根据路径类型获取路径
func (p *Position) GetPathByType(pathType string) *Path {
	// 先检查向上路径
	for _, path := range p.UpPaths {
		if path.Type == pathType {
			return path
		}
	}

	// 再检查向下路径
	for _, path := range p.DownPaths {
		if path.Type == pathType {
			return path
		}
	}

	return nil
}
