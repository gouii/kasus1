package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/** ini berfungsi untuk menagkap inputan dari terminal */
func ReadString() string {
	input := bufio.NewReader(os.Stdin)
	str, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("error input", err)
	}

	strNew := strings.Split(str, "\n")
	return strNew[0]
}
