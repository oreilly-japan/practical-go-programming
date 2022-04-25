package main

// base-function
type Portion int

const (
	Regular    Portion = iota // 普通
	Small                     // 小盛り
	Large                     // 大盛り
	ExtraLarge                // マシマシ
)

// each-func
func NewKakeUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: false,
		ebiten:   0,
	}
}

func NewKitsuneUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: true,
		ebiten:   0,
	}
}

func NewTempuraUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: false,
		ebiten:   3,
	}
}

// each-func

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func main() {
}
