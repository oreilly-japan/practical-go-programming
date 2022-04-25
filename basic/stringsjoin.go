package main

import (
	"log"
	"strings"
)

func main() {
	// plus_op
	src := []string{"Back", "To", "The", "Future", "Part", "III"}
	var title string
	for i, word := range src {
		if i != 0 {
			title += " "
		}
		title += word
	}
	log.Println(title)
	// plus_op

	// join-expression-start
	displayTitle := "1990年7月6日公開 - " + title + " - ロバート・ゼメキス"
	log.Println(displayTitle)
	// join-expression-end

	// use_builder
	var builder strings.Builder
	builder.Grow(100) // 最大100文字以下と仮定できる場合
	for i, word := range src {
		if i != 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(word)
	}
	log.Println(builder.String())
	// use_builder
}
