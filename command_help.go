package main

import "fmt"

func commandHelp() error {
	commands := getCommands()
	fmt.Println("Showing Help Menu")
	for _, c := range commands {
		fmt.Println(c.name, ": ", c.description)
	}
	return nil
}