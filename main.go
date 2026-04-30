package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	name := "Stranger"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	hour := time.Now().Hour()
	greeting := "Hello"
	if hour < 12 {
		greeting = "Good morning"
	} else if hour < 17 {
		greeting = "Good afternoon"
	} else {
		greeting = "Good evening"
	}
	fmt.Printf("%s, %s! Welcome to TurnUpCampus 🔥\n", greeting, name)
	fmt.Printf("Current time: %s\n", time.Now().Format("15:04:05"))
	fmt.Println("Campus: Thika Road | Status: Active")
}