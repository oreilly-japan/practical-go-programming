package main

// wrap
import (
	"bufio"
	"io"
	"os"
	"strings"

	"golang.org/x/text/unicode/norm"
)

// コアとなる基本実装
func Normalize(w io.Writer, r io.Reader) error {
	br := bufio.NewReader(r)
	for {
		s, err := br.ReadString('\n')
		if s != "" {
			io.WriteString(w, norm.NFKC.String(s))
		}
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
	}
}

// ファイルとの読み書きがしやすいようにファイル名を受け取るようにしたラッパー
func NormalizeFile(input, output string) error {
	r, err := os.Open(input)
	if err != nil {
		return err
	}
	w, err := os.Create(output)
	if err != nil {
		return err
	}
	return Normalize(r, w)
}

// ユニットテストや他のコードから呼び出しやすいように文字列を渡すだけで利用できるラッパー
func NormalizeString(i string) (string, error) {
	r := strings.NewReader(i)
	var w strings.Builder
	err := Normalize(&w, r)
	if err != nil {
		return "", err
	}
	return w.String(), nil
}

// wrap

func assersion() {
	// assersion

	// assersion

}
