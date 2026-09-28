package Functions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PrintART(content *os.File, userInput []string) {
	var strn string
	var result string
	for _, art := range userInput {
		if art != "" {
			for line := 0; line < 8; line++ {
				for _, character := range art {
					// Calculate the line number to start reading from.
					startLine := 2 + int(character-32)*9 + line
					content.Seek(0, 0) // Reset file position to the beginning.
					reader := bufio.NewReader(content)
					// Read up to the start line.
					for i := 0; i < startLine; i++ {
						strn, _ = reader.ReadString('\n')
					}
					strn = strings.TrimRight(strn, "\r\n")
					result += strn
				}
				fmt.Println(result)
				result = ""
			}
		} else if len(userInput) != 1 {
			fmt.Println()
		}
	}
}
