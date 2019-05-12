package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	c := len(os.Args) - 1
	if c < 1 {
		fmt.Fprint(os.Stderr, "[usage] %s choice1 choice2...", os.Args[0])
		os.Exit(1)
	}
	
	fmt.Println(os.Args[rand.Intn(c) + 1])
}