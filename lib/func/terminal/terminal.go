package terminal

import "fmt"

func ClearTerminal() {
	fmt.Print("\033c")
}
