package official

// Rank 官职等级
type Rank struct {
	Name  string
	Level int
	Yamen Yamen
}

var (
	ranks = []string{
		"正一品",
		"从一品",
		"正二品",
		"从二品",
		"正三品",
		"从三品",
		"正四品",
		"从四品",
		"正五品",
		"从五品",
		"正六品",
		"从六品",
		"正七品",
		"从七品",
		"正八品",
		"从八品",
		"正九品",
		"从九品",
		"未入流",
	}
)
