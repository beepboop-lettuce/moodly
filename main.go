package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("[your name]")
		return
	}

	userName := os.Args[1]

	mood := [6]string{
		"feels good",
		"feels bad",
		"feels sad",
		"feels happy",
		"feels awesome",
		"feels terrible",
	}

	rand.Seed(time.Now().UnixNano())
	myMood := mood[rand.Intn(len(mood))]

	fmt.Printf("%s %s\n", userName, myMood)
}
