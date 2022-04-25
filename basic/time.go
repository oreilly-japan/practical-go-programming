package main

import (
	"fmt"
	"time"

	// tz-bundle
	_ "time/tzdata"
	// tz-bundle
)

func timeSample() {
	// time-sample

	// 現在時刻のtime.Timeインスタンス取得
	now := time.Now()

	// 指定日時のtime.Timeインスタンス取得
	tz, _ := time.LoadLocation("America/Los_Angeles")
	future := time.Date(2015, time.October, 21, 7, 28, 0, 0, tz)

	fmt.Println(now.String())
	fmt.Println(future.Format(time.RFC3339Nano))
	// time-sample
}

func timeZone() {
	// timezone-sample
	// 定義済みのローカルタイムゾーン
	now := time.Date(1985, time.October, 26, 9, 0, 0, 0, time.Local)

	// 定義済みのUTCタイムゾーン
	past := time.Date(1955, time.November, 12, 6, 38, 0, 0, time.UTC)
	// timezone-sample

	fmt.Println(now.String())
	fmt.Println(past.String())
}

func timeDuration() {
	// timeduration-create
	// 5分を作成
	// Nanosecond, Millisecond, Second, Minute, Hourが定義済み
	fiveMinute := 5 * time.Minute

	// intとは型違いで直接演算できないので、即値との計算以外は
	// time.Durationへの明示的なキャストが必要
	// キャストがないと、次のエラーが発生する
	// invalid operation: seconds * time.Second (mismatched types int and time.Duration)
	var seconds int = 10
	tenSeconds := time.Duration(seconds) * time.Second

	// Timeの演算でDuration作成
	past := time.Date(1955, time.November, 12, 6, 38, 0, 0, time.UTC)
	dur := time.Now().Sub(past)
	// timeduration-create
	fmt.Println(dur.String())

	// timeduration-calc
	// 1時間にまとめてバッチで読み込むファイル名を取得
	filepath := time.Now().Truncate(time.Hour).Format("20060102150405.json")

	// 5分後と5分前の時刻
	fiveMinuteAfter := time.Now().Add(fiveMinute)
	fiveMinuteBefore := time.Now().Add(-fiveMinute)

	// timeduration-calc
	fmt.Println(filepath)
	fmt.Println(fiveMinuteAfter, fiveMinuteBefore)

	// timeduration-use
	// 3秒停止
	fmt.Println("3秒スリープスタート")
	time.Sleep(3 * time.Second)
	fmt.Println("3秒スリープ完了")

	// 10秒間待つ
	fmt.Println("10秒停止スタート")
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()
	<-timer.C
	fmt.Println("10秒停止完了")
	// timeduration-use

}

func main() {
	timeDuration()
}
