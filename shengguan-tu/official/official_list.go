package official

import "fmt"

// "白丁":  {"德": 3, "才": 2, "功": 1, "赃": -1},
// "童生": {"德": 4, "才": 3, "功": 1, "赃": -1},
// "案首": {"德": 4, "才": 5, "功": 1, "赃": -1},
// "监生": {"德": 4, "才": 5, "功": 1, "赃": -1},
// "秀才": {"德": 5, "才": 6, "功": 4, "赃": 1},
// "廪生": {"德": 6, "才": 7, "功": 4, "赃": 1},
// "贡生": {"德": 7, "才": 8, "功": 5, "赃": 1},
// "举人": {"德": 8, "才": 9, "功": 6, "赃": 1},
// "解元": {"德": 9, "才": 10, "功": 7, "赃": 1},
// "进士": {"德": 10, "才": 11, "功": 8, "赃": 1},
// "会元": {"德": 11, "才": 12, "功": 9, "赃": 1},
// "传胪": {"德": 12, "才": 13, "功": 10, "赃": 1},
// "探花": {"德": 13, "才": 14, "功": 11, "赃": 1},
// "榜眼": {"德": 14, "才": 14, "功": 12, "赃": 1},
// "状元": {"德": 15, "才": 15, "功": 14, "赃": 1},
// DefaultOfficials 默认官职列表
func DefaultOfficials() []*Official {
	return []*Official{
		{
			ID:          "tai_shi",
			Name:        "太师",
			Level:       1,
			Description: "最高官职，位极人臣",
			MoveRules: map[string]int{
				"德": 0,
				"才": 0,
				"功": 0,
				"赃": -1,
			},
		},
		{
			ID:          "shang_shu",
			Name:        "尚书",
			Level:       2,
			Description: "六部尚书，朝廷重臣",
			MoveRules: map[string]int{
				"德": 1,
				"才": 1,
				"功": 1,
				"赃": -2,
			},
		},
		{
			ID:          "jian_yi_da_fu",
			Name:        "谏议大夫",
			Level:       3,
			Description: "负责谏言的官员",
			MoveRules: map[string]int{
				"德": 1,
				"才": 1,
				"功": 1,
				"赃": -2,
			},
		},
		{
			ID:          "shi_zhong",
			Name:        "侍中",
			Level:       4,
			Description: "皇帝近臣",
			MoveRules: map[string]int{
				"德": 1,
				"才": 1,
				"功": 1,
				"赃": -2,
			},
		},
		{
			ID:          "zhou_zhou_cishi",
			Name:        "州刺史",
			Level:       5,
			Description: "地方行政长官",
			MoveRules: map[string]int{
				"德": 1,
				"才": 1,
				"功": 1,
				"赃": -2,
			},
		},
		{
			ID:          "xian_ling",
			Name:        "县令",
			Level:       6,
			Description: "县级行政长官",
			MoveRules: map[string]int{
				"德": 1,
				"才": 1,
				"功": 1,
				"赃": -2,
			},
		},
		{
			ID:          "mu_zhang",
			Name:        "牧长",
			Level:       7,
			Description: "管理牧场的官员",
			MoveRules: map[string]int{
				"德": 1,
				"才": 1,
				"功": 1,
				"赃": -2,
			},
		},
		{
			ID:          "ting_zhang",
			Name:        "亭长",
			Level:       8,
			Description: "基层小吏",
			MoveRules: map[string]int{
				"德": 1,
				"才": 1,
				"功": 1,
				"赃": -2,
			},
		},
		{
			ID:          "li_zheng",
			Name:        "里正",
			Level:       9,
			Description: "乡里小吏",
			MoveRules: map[string]int{
				"德": 1,
				"才": 1,
				"功": 1,
				"赃": -1,
			},
		},
		{
			ID:          "bai_ding",
			Name:        "白丁",
			Level:       10,
			Description: "平民百姓，无官职",
			MoveRules: map[string]int{
				"德": 1,
				"才": 1,
				"功": 1,
				"赃": 0,
			},
		},
	}
}

// GetOfficialByID 根据ID获取官职
func GetOfficialByID(id string) (*Official, error) {
	for _, official := range DefaultOfficials() {
		if official.ID == id {
			return official, nil
		}
	}
	return nil, fmt.Errorf("官职不存在: %s", id)
}

// GetOfficialByLevel 根据级别获取官职
func GetOfficialByLevel(level int) (*Official, error) {
	if level < 1 || level > maxLevel {
		return nil, fmt.Errorf("无效的官职级别: %d", level)
	}

	for _, official := range DefaultOfficials() {
		if official.Level == level {
			return official, nil
		}
	}
	return nil, fmt.Errorf("级别对应的官职不存在: %d", level)
}
