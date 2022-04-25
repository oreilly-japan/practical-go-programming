package main

// define-sync-pool
import "sync"

// 巨大な構造体
type BigStruct struct {
	Member string
}

// Poolは初期化関数をNewフィールドに設定して作成する
var pool = &sync.Pool{
	New: func() interface{} {
		return &BigStruct{}
	},
}

// define-sync-pool

func main() {
	// use-sync-pool
	// インスタンスはGet()メソッドで取得
	// 在庫があればそれを、なければNew()を呼び出す
	b := pool.Get().(*BigStruct)

	// 使い終わったらPut()でPoolに戻して次回に備える
	pool.Put(b)
	// use-sync-pool
}

// hide-pool
// BigStructのインスタンスを作成するファクトリー関数
// 内部でプールを利用
func NewBigStruct() *BigStruct {
	b := pool.Get().(*BigStruct)
	return b
}

// 自分自身を返却するメソッド
func (b *BigStruct) Release() {
	// 初期化してから格納
	b.Member = ""
	pool.Put(b)
}

// hide-pool
