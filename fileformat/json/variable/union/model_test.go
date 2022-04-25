package union

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

func Test(t *testing.T) {
	jsonStr, err := os.ReadFile("../sensor.json")
	if err != nil {
		log.Fatal(err)
	}

	var r Response
	if err := json.Unmarshal(jsonStr, &r); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", r)
}
