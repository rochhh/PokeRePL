package main

import (
	"fmt";
	"os";
	"bufio";
	"strings"
)


func startRepl(){
	
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(" >")

		scanner.Scan()
		text := scanner.Text()
		cleanWord := cleanInput(text)

		if len(cleanWord) == 0 {
			continue  // when enter , the shell does not quit like normal 
		}

		fmt.Println( "echoing ->" , cleanWord )
	}
}

func cleanInput(str string ) []string {
	lower := strings.ToLower(str)
	words := strings.Fields(lower)
	return words
}
