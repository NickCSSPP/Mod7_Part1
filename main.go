package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("What is your favorite food?")
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println(s + "! Nice choice!")

	fmt.Println("How do you like it cooked?")
	cooked, _ := in.ReadString('\n')
	cooked = strings.TrimSpace(cooked)
	cooked = strings.ToLower(cooked)
	fmt.Println("You like it cooked", cooked+"?, I'm more of a medium.")
}
