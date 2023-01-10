package main

import (
	"bufio"
	"fmt"
	"os"
)

var krot mole

func main() {
	krot := Newmole()
	krot.Start()
}
func command() string {
	fmt.Println("Enter command:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Println("You entered:", input)
	return input
}
