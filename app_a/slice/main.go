package main

import "fmt"

func useArray() {
	// use-array
	// 要素数3の整数の配列
	var nums [3]int = [3]int{1, 2, 3}

	// 要素の値を取り出して表示
	fmt.Println(nums[0])

	// 長さを取得
	fmt.Println(len(nums))

	// use-array
	fmt.Println(nums)

	// use-slice
	// スライスの変数宣言
	var nums1 []int

	// 1, 2, 3の要素を持つスライスを作成して代入
	nums2 := []int{1, 2, 3}

	// あるいは既存の配列やスライスからも範囲アクセスでスライス作成
	nums3 := nums[0:2]  // 配列から
	nums4 := nums2[1:3] // スライスから

	// 配列と同じようにブラケットで要素取得可能
	// 範囲外アクセスはパニック
	fmt.Println(nums2[1]) // 2

	// 要素の割り当ても可能
	nums2[0] = 100

	// 長さも取得可能
	fmt.Println(len(nums2)) // 3

	// スライスに要素を追加
	// 再代入が必要
	nums2 = append(nums2, 4)

	// use-slice
	fmt.Println(nums1, nums2, nums3, nums4)

	// use-map
	// 数字がキーで値が文字列のマップ
	// HTTPステータスを格納
	hs := map[int]string{
		200: "OK",
		404: "Not Found",
	}

	// makeで作る
	authors := make(map[string][]string)

	// ブラケットで要素アクセス
	// 代入
	authors["Go"] = []string{"Robert Griesemer", "Rob Pike", "Ken Thompson"}

	// データ取得
	status := hs[200]
	fmt.Println(status)
	// "OK"

	// 存在しない要素にアクセスするとゼロ値
	fmt.Println(hs[0]) // panic

	// あるかどうかの情報も一緒に取得
	status, ok := hs[304]
	// status = ""
	// ok = false
	// use-map

	fmt.Println(hs, ok)

}
