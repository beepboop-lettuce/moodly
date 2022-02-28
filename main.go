package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	username := os.Args[1]

	if len(os.Args) != 2 {
		fmt.Println("[your name]")
	}

	mood := [6]string{
		"feels good",
		"feels bad",
		"feels sad",
		"feels happy",
		"feels awesome",
		"feels terrible",
	}

	rand.Seed(time.Now().UnixNano())
}
