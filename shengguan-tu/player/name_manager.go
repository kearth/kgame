package player

import "math/rand/v2"

var AncientSurnames = []string{
	"李", "王", "张", "刘", "陈",
	"杨", "赵", "黄", "周", "徐",
	"孙", "胡", "林", "何", "郭",
	"马", "罗", "梁", "宋", "郑",
	"谢", "韩", "唐", "冯", "曹",
	"曾", "彭", "邓", "潘", "苏",
	"卢", "魏", "蒋", "蔡", "贾",
	"丁", "薛", "叶", "杜", "余",
	"程", "汪", "沈", "吕", "姜",
	"范", "方", "石", "尹", "钟",
	"廖", "袁", "钱", "萧", "欧阳",
	"司马", "诸葛", "文", "柳", "陆",
	"崔", "孔", "孟", "颜", "董",
	"熊", "贺", "汤", "毛", "朱",
	"秦", "顾", "白", "康", "史",
	"郝", "施", "傅", "陶", "洪",
	"邵", "万", "黎", "高", "夏侯",
	"温", "燕", "翁", "侯", "倪",
	"滕", "殷", "柯", "尤", "龚",
	"章", "裴", "狄", "阮", "包",
	// 额外经典复姓（科举史上真实出现过）
	"上官", "令狐", "皇甫", "慕容", "尉迟",
	"呼延", "宇文", "鲜于", "公孙", "长孙",
}

// 男性高频字（权重更高）
var MaleChars = []string{
	"文", "武", "德", "明", "贤", "良", "仁", "杰", "宏", "伟",
	"浩", "博", "远", "安", "邦", "辅", "之", "子", "卿", "甫",
	"庭", "玉", "伯", "仲", "叔", "季", "元", "正", "世", "永",
	"光", "辉", "景", "章", "泰", "宗", "时", "中", "大", "天",
	"君", "公", "汝", "斯", "彦", "希", "若", "允", "继", "承",
}

// 女性高频字（柔美、古典）
var FemaleChars = []string{
	"婉", "柔", "静", "萱", "兰", "芙", "蓉", "芷", "茹", "琴",
	"雪", "瑶", "玉", "莲", "蕙", "怡", "慧", "清", "霏", "嫣",
	"娴", "姝", "婵", "婷", "妮", "琳", "璇", "琼", "莹", "芳",
	"慕容", "若", "素", "凝", "昭", "华", "嫕", "娉", "婧", "姣",
	"若兰", "玉环", "翠", "红", "绿", "碧", "秋", "菊", "梅", "竹",
}

// 写一个随机姓名系统
type NameManager struct {
	Random                   *rand.Rand // 随机数生成器
	AncientSurnamesLen       int        // 古代姓氏数组长度
	MaleCharsLen             int        // 男性高频字数组长度
	FemaleCharsLen           int        // 女性高频字数组长度
	MaleSingleOrDoubleRate   int        // 男单字名率（70%）
	FemaleSingleOrDoubleRate int        // 女单字名率（90%）
}

// NewNameManager 创建一个新的姓名管理器
func NewNameManager(random *rand.Rand) *NameManager {
	return &NameManager{
		Random:                   random,
		AncientSurnamesLen:       len(AncientSurnames),
		MaleCharsLen:             len(MaleChars),
		FemaleCharsLen:           len(FemaleChars),
		MaleSingleOrDoubleRate:   70,
		FemaleSingleOrDoubleRate: 90,
	}
}

// GenerateName 随机生成姓名
func (n *NameManager) GenerateName(gender Gender) string {
	switch gender {
	case Male:
		return n.AncientNameMale()
	case Female:
		return n.AncientNameFemale()
	default:
		// 50% 概率生成男性姓名，50% 概率生成女性姓名
		if n.Random.IntN(2) == 0 {
			return n.AncientNameMale()
		}
		return n.AncientNameFemale()
	}
}

// AncientNameMale 随机生成古代男性姓名
func (n *NameManager) AncientNameMale() string {
	var name string
	surname := AncientSurnames[n.Random.IntN(n.AncientSurnamesLen)]
	// 男性：70% 单字，30% 双字（符合古代士人习惯）
	if n.Random.IntN(100) < n.MaleSingleOrDoubleRate {
		name = surname + MaleChars[n.Random.IntN(n.MaleCharsLen)]
	} else {
		c1 := MaleChars[n.Random.IntN(n.MaleCharsLen)]
		c2 := MaleChars[n.Random.IntN(n.MaleCharsLen)]
		name = surname + c1 + c2
	}
	return name
}

// AncientNameFemale 随机生成古代女性姓名
func (n *NameManager) AncientNameFemale() string {
	var name string
	surname := AncientSurnames[n.Random.IntN(n.AncientSurnamesLen)]
	// 女性：90% 双字名，10% 单字名（符合古代女子取名习惯）
	if n.Random.IntN(100) < n.FemaleSingleOrDoubleRate {
		c1 := FemaleChars[n.Random.IntN(n.FemaleCharsLen)]
		c2 := FemaleChars[n.Random.IntN(n.FemaleCharsLen)]
		name = surname + c1 + c2
	} else {
		name = surname + FemaleChars[n.Random.IntN(n.FemaleCharsLen)]
	}
	return name
}
