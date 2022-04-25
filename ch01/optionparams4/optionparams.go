package main

// base-function
type Portion int

const (
	Regular    Portion = iota // 普通
	Small                     // 小盛り
	Large                     // 大盛り
	ExtraLarge                // マシマシ
)

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

// fluent-option
type fluentOpt struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon(p Portion) *fluentOpt {
	// デフォルトはコンストラクタ関数で設定
	// 必須オプションはここに付与可能
	return &fluentOpt{
		men:      p,
		aburaage: false,
		ebiten:   1,
	}
}

func (o *fluentOpt) Aburaage() *fluentOpt {
	o.aburaage = true
	return o
}

func (o *fluentOpt) Ebiten(n uint) *fluentOpt {
	o.ebiten = n
	return o
}

func (o *fluentOpt) Order() *Udon {
	return &Udon{
		men:      o.men,
		aburaage: o.aburaage,
		ebiten:   o.ebiten,
	}
}

func useFluentInterface() {
	oomoriKitsune := NewUdon(Large).Aburaage().Order()
}

// fluent-option

func main() {
	useFluentInterface()
}
