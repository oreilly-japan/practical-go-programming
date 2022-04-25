package main


func main() {
   // iはint型
   var i int

   // ErrorCodeはintを拡張した型
   type ErrorCode int
   var e ErrorCode

   i = e // エラー
   e = i // エラー
}