package main

import (
	"bufio"
	"fmt"
	"./pluginmanager"
	"os"
	"strings"
)

func main() {
	pins := pluginmanager.PluginManager{}
	pins.Initialize()

	bio := bufio.NewReader(os.Stdin)

	input := make(chan string)
	output := make(chan string)

	go func(reader *bufio.Reader, in chan string) {
		for {
			line, err := reader.ReadString('\n')
			if err == nil {
				in <- strings.Trim(line, "\n")
			}
		}
	}(bio, input)

	go func() {
		for {
			result, err := pins.Execute(<-input)
			if err != nil {
				output <- fmt.Sprintf("Error: %s", err)
			} else {
				output <- result
			}
		}
	}()

	for {
		fmt.Println(<-output)
	}
}
