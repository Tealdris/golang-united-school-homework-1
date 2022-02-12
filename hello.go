package main

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func main() {
	pizzaMessage := emoji.Sprint("Hello :world_map:!")
	fmt.Println(pizzaMessage)
}
