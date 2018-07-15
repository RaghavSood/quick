package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("HISTFILE:", os.Getenv("HISTFILE"))
	fmt.Println("SHELL:", os.Getenv("SHELL"))
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
