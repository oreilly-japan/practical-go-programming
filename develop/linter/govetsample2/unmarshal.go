package govetsample2

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	ISBN      string `json:"isbn"`
}

func Unmarshal() {

	// vet-sample-start
	s := `{ "title":"Real World HTTP",
			"author":"渋川よしき",
			"publisher":"オライリー・ジャパン",
			"isbn":"4873119030"}`

	var b Book
	// json.Unmarshal の引数はポインター型として渡す必要があるが、値として渡している
	if err := json.Unmarshal([]byte(s), b); err != nil {
		// ...
	}
	// vet-sample-end
	fmt.Println(b)
}
