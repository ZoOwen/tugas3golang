package main

import "fmt"

func main() {
	status, word, ok := searchString("sublime")
	if !ok {
		fmt.Println(status, word)
	} else {
		fmt.Println(status, word)
	}
}

func searchString(word string) (string, string, bool) {
	s := [...]string{"vscode", "sublime", "node", "net beans"}
	wordslice := s[0:]

	for i := 0; i < len(s); i++ {
		if wordslice[i] == word {
			return "found", word, true
		}

	}
	return "not found ", word, false
}
