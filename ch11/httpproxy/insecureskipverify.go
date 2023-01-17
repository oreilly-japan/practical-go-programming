package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// go-insecure-start
	hc := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			Proxy: http.ProxyFromEnvironment, // 要設定
		}}
	// go-insecure-end

	resp, _ := hc.Get("https://www.oreilly.co.jp/index.shtml")
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
