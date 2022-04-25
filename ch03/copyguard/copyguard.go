package main

import (
	"log"
	"sync"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// copyguard
// ポインターとしてのみ利用する構造体
type NoCopyStruct struct {
	self  *NoCopyStruct
	Value *string
}

// 初期化時にポインターを保持しておく
func NewNoCopyStruct(value string) *NoCopyStruct {
	r := &NoCopyStruct{
		Value: &value,
	}
	r.self = r
	return r
}

// メソッドの中でチェック
func (n *NoCopyStruct) String() string {
	if n != n.self {
		panic("should not copy NoCopyStruct instance without Copy() method")
	}
	return *n.Value
}

// copyguard

func (*NoCopyStruct) Lock() {}

// copymethod
// 明示的なコピー用メソッド
func (n *NoCopyStruct) Copy() *NoCopyStruct {
	str := *n.Value
	p2 := &NoCopyStruct{
		Value: &str,
	}
	p2.self = p2
	return p2
}

// copymethod

func main() {
	s := NewNoCopyStruct("hello")
	s2 := s.Copy()

	log.Println(s.String())
	log.Println(s2.String())

	var s3 NoCopyStruct
	CopyByValue(s3)

	var a sync.Mutex
	b := a    // doesn't detect
	var c = a // doesn't detect
	b = a     // detects
	log.Println(b, c)
}

func CopyByValue(s NoCopyStruct) {
	log.Println(s.String())
}
