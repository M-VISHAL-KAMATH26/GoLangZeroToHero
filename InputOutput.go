package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter your name....!")
	name, _ := reader.ReadString('\n')
	fmt.Println(name)

}
