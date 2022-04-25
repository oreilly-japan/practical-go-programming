package optionparams

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

// functional-options
type OptFunc func(r *Udon)

func NewUdon(opts ...OptFunc) *Udon {
	r := &Udon{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func OptMen(p Portion) OptFunc {
	return func(r *Udon) { r.men = p }
}

func OptAburaage() OptFunc {
	return func(r *Udon) { r.aburaage = true }
}

func OptEbiten(n uint) OptFunc {
	return func(r *Udon) { r.ebiten = n }
}

func useFuncOption() {
	tokuseiUdon := NewUdon(OptAburaage(), OptEbiten(3))
}

// functional-options

func main() {
	useFuncOption()
}
