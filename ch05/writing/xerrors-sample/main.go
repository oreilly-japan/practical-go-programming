package main

// error-main
import (
	"fmt"

	"golang.org/x/xerrors"
)

func main() {
	err := xerrors.New("xerrors error")
	fmt.Printf("%+v\n", err)
}

// error-main
