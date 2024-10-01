package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/nukestarmaster/Pokedex/internal/pokeapi"
)

func Repl(cfg *pokeapi.Config) {
	myscanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("> ")
		myscanner.Scan()
		mytext := myscanner.Text()
		tokens := cleanInput(mytext)
		if len(tokens) == 0 {
			continue
		}
		commandName := tokens[0]
		availiableCommands := getCommands()
		command, ok := availiableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		err := command.callback(cfg, tokens[1:])
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *pokeapi.Config, args []string) error
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
    	"help": {
        	name:        "help",
        	description: "Displays a help message.",
        	callback:    commandHelp,
    	},
    	"exit": {
        	name:        "exit",
        	description: "Exit the Pokedex.",
        	callback:    commandExit,
    	},
		"map": {
			name:		 "map",
			description: "Prints the first 20 area. Repeated use gives the next 20 areas.",
			callback:	 commandMap,
		},
		"mapb": {
			name:		 "map back",
			description: "Prints the previous list of 20 areas. Returns an error if map has not been called yet",
			callback: 	 commandMapB,
		},
		"explore": {
			name:		 "explore",
			description: "Explores the chosen area and lists the pokemon found there",
			callback: 	 commandExplore,
		},
	}
} 