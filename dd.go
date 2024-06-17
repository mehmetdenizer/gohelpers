package gohelpers

import (
	"fmt"
	"os"
)

func DD(data interface{}) {
	fmt.Printf("%+v\n", data)
	os.Exit(1)
}
