package console

import (
	"bufio"
	"fmt"
	"os"
)

func StartConsole() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		if text == "exit\n" {
			break
		}
		// Handle commands
		fmt.Println("Command received:", text)
	}
}
