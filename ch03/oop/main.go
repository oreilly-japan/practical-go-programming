package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"strconv"
)

// ok-override
type I interface {
	m1(i I)
	m2()
}

type Parent struct{}

func (p Parent) m1(i I) {
	i.m2()
}

func (p Parent) m2() {
	fmt.Println("Parent")
}

type Child struct {
	Parent
}

func (c Child) m2() {
	fmt.Println("Child")
}

// ok-override

func overrideTest() {
	// override-test
	p := Parent{}
	p.m1(p) // Parent

	c := Child{}
	c.m1(c) // Child
	// override-test
}

func reverseProxy() {
	// reverse-proxy
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = "localhost:9001"
	}
	modifier := func(res *http.Response) error {
		body := make(map[string]interface{})
		dec := json.NewDecoder(res.Body)
		dec.Decode(&body)
		body["fortune"] = "大吉"
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.Encode(&body)
		res.Body = io.NopCloser(&buf)
		res.Header.Set("Content-Length", strconv.Itoa(buf.Len()))
		return nil
	}
	rp := &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifier,
	}
	http.ListenAndServe(":9000", rp)
	// reverse-proxy
}

func reverseProxyTest() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"fortune": "小吉"}`)
	}))
	defer server.Close()
	u, _ := url.Parse(server.URL)
	director := func(req *http.Request) {
		req.URL = u
	}
	modifier := func(res *http.Response) error {
		body := make(map[string]interface{})
		dec := json.NewDecoder(res.Body)
		dec.Decode(&body)
		body["fortune"] = "大吉"
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.Encode(&body)
		res.Body = io.NopCloser(&buf)
		res.Header.Set("Content-Length", strconv.Itoa(buf.Len()))
		return nil
	}
	rp := &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifier,
	}
	http.ListenAndServe(":9000", rp)
}

func main() {
	overrideTest()
	//overrideTest()
	//reverseProxyTest()
}
