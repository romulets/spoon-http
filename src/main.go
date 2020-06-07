package main

import (
	"fmt"
	"os"
)

func main() {
	uri, flags, err := ParseParameters(os.Args)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(uri)
	fmt.Println(flags)
}
