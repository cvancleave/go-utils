package utils

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func UserLoop(before, onExit, onEmpty func(), after func(string)) {

	reader := bufio.NewReader(os.Stdin)

	for {
		time.Sleep(250 * time.Millisecond)

		if before != nil {
			before()
		}

		fmt.Printf("user -> ")

		// read user input
		input, _ := reader.ReadString('\n')

		endIndex := len(input) - 2
		if runtime.GOOS != "windows" {
			endIndex = len(input) - 1 // linux
		}

		input = input[0:endIndex]

		if input == "" {
			if onEmpty != nil {
				onEmpty()
			}
			continue
		}

		if strings.EqualFold(input, "exit") {
			if onExit != nil {
				onExit()
			}
			break
		}

		if after != nil {
			after(input)
		}
	}
}
