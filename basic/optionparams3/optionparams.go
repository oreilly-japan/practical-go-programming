package main

import "time"

// base-function
type Portion int

const (
	Regular    Portion = iota // 普通
	Small                     // 小盛り
	Large                     // 大盛り
	ExtraLarge                // マシマシ
)

// base-function

// struct-option
type Option struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon(opt Option) *Udon {
	// ゼロ値に対するデフォルト値処理は関数/メソッド内部で行う
	// 朝食時間は海老天1本無料
	if opt.ebiten == 0 && time.Now().Hour() < 10 {
		opt.ebiten = 1
	}
	return &Udon{
		men:      opt.men,
		aburaage: opt.aburaage,
		ebiten:   opt.ebiten,
	}
}

// struct-option

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func main() {
}
