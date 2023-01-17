package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 証明書を読み込む
	cert, err := os.ReadFile("ca.crt")
	if err != nil {
		panic(err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)
	cfg := &tls.Config{
		RootCAs: certPool,
	}
	cfg.BuildNameToCertificate()

	// クライアントを作成
	hc := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: cfg,
			Proxy:           http.ProxyFromEnvironment, // 要設定
		},
	}

	resp, _ := hc.Get("https://www.oreilly.co.jp/index.shtml")
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

}
