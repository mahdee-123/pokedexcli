package main

import (
	"bufio"
	"fmt"
	"os"
)

func initRepl() {

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		cleanText := cleanInput(input)


		if len(cleanText) == 0 {
			continue
		}

		arg := ""
		if len(cleanText) > 1 {
			arg = cleanText[1]
		}

		if cmd, ok := commands[cleanText[0]]; ok {
			if err := cmd.callback(arg); err != nil {
				fmt.Println(err)
			}
			} else {
				fmt.Println("Unknown command")
			}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
