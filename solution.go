package solution

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func GetMessage() {
	pizzaMessage := emoji.Sprint("Hello :world_map:!")
	fmt.Println(pizzaMessage)
}
