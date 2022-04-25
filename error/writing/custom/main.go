package main

import (
	"fmt"
	"io"
	"net/http"
)

// custom-error-struct-start
// HTTPError はステータスコードが200以外の場合のエラーを扱う構造体
type HTTPError struct {
	StatusCode int
	URL        string
}

// このメソッドを実装することで HTTPError 構造体のポインターは Error インタフェースを満たしています
func (he *HTTPError) Error() string {
	return fmt.Sprintf("http status code = %d, url = %s", he.StatusCode, he.URL)
}

// custom-error-struct-end

// custom-error-use-start
func ReadContents(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// レスポンスのステータスコードが200ではない場合は *HTTPError としてエラーを返却
	if resp.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			URL:        url,
		}
	}
	return io.ReadAll(resp.Body)
}

// custom-error-use-end
