package main

import (
	"container/list"
	"fmt"
	"log"
	"net/url"
)

// http-status
type HTTPStatus int

const (
	StatusOK              HTTPStatus = 200
	StatusUnauthorized    HTTPStatus = 401
	StatusPaymentRequired HTTPStatus = 402
	StatusForbidden       HTTPStatus = 403
)

func (s HTTPStatus) String() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusUnauthorized:
		return "Unauthorized"
	case StatusPaymentRequired:
		return "Payment Required"
	case StatusForbidden:
		return "Forbidden"
	default:
		return fmt.Sprintf("HTTPStatus(%d)", s)
	}
}

// http-status

// national-route
type NationalRoute int

const (
	NagasakiKaido   NationalRoute = 200
	AizuNumataKaido NationalRoute = 401
	HokurikuDo      NationalRoute = 402
	KurinokiBypass  NationalRoute = 403
)

func (n NationalRoute) String() string {
	switch n {
	case NagasakiKaido:
		return "長崎街道"
	case AizuNumataKaido:
		return "会津沼田街道"
	case HokurikuDo:
		return "北陸道"
	case KurinokiBypass:
		return "栗ノ木バイパス"
	default:
		return fmt.Sprintf("国道%d号線", n)
	}
}

// national-route

func urlValues() {
	// init-url-values
	// mapと同じく初期化が必要
	// mapと同じく2通りの初期化方法がある
	v1 := url.Values{}
	v2 := make(url.Values)
	// init-url-values
	log.Println(v1, v2)
}

func useUrlValues() {
	// use-url-values
	vs := url.Values{}
	vs.Add("key1", "value1")
	vs.Add("key2", "value2")
	for k, v := range vs {
		fmt.Printf("%s: %v\n", k, v)
	}
	// use-url-values
}

func getUrlMultiValues() {
	vs := url.Values{}
	// get-url-multi-values
	vs.Add("key1", "value1")
	vs.Add("key1", "value2")                 // 同じキー
	fmt.Printf("key1: %v\n", vs.Get("key1")) // Getメソッドだと最初の要素のみ
	for i, v := range vs["key1"] {           // インデックスアクセスすると全値取得可能
		fmt.Printf("key1[%d]: %v\n", i, v)
	}
	// get-url-multi-values
}

func uselist() {
	// use-list
	// container/listを使う
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	// container/listはNext()で返り値がnilでない間はループ
	for ele := l.Front(); ele != nil; ele = ele.Next() {
		fmt.Println(ele.Value)
	}
	// use-list
}

func main() {
	var s HTTPStatus

	fmt.Println(s)
	//fmt.Printf("%s\n", StatusTeapot.String())
	urlValues()
	useUrlValues()
	uselist()
	getUrlMultiValues()
}
