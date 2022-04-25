package main

// base-function
type Portion int

const (
	Regular Portion = iota // 普通
	Small                  // 小盛り
	Large                  // 大盛り
)

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

// 麺の量、油揚げ、海老天の有無でインスタンス作成
func NewUdon(p Portion, aburaage bool, ebiten uint) *Udon {
	return &Udon{
		men:      p,
		aburaage: aburaage,
		ebiten:   ebiten,
	}
}

// 海老天2本入りの大盛り
var tempuraUdon = NewUdon(Large, false, 2)

// base-function

func main() {
	u := New Udon{}
}
