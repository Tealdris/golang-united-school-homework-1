package solution

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	pizzaMessage := emoji.Sprint("Hello :world_map:!")
	fmt.Println(pizzaMessage)
	return pizzaMessage
}
