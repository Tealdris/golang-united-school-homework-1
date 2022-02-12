package solution

import (
	"fmt"

	"github.com/kyokomi/emoji"
)

func Message() {
	pizzaMessage := emoji.Sprint("Hello :world_map:!")
	fmt.Println(pizzaMessage)
}
