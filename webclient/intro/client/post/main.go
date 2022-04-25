package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// client-post-user-start
type User struct {
	Name string
	Addr string
}

// client-post-user-end

func main() {
	// client-post-start
	u := User{
		Name: "O'Reilly Japan",
		Addr: "東京都新宿区四谷坂町",
	}

	payload, err := json.Marshal(u)
	if err != nil {
		// ...
	}

	resp, err := http.Post("http://example.com/", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		// ...
	}
	defer resp.Body.Close()

	// client-post-end

	if resp.StatusCode != http.StatusCreated {
		// ...
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// ...
	}
	fmt.Println(string(body))
}
