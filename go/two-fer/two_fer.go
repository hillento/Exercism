// Package twofer https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// ShareWith will take user input and add it to the sentence "One for {name}, one for me." If no user input is given it will automatically assign the name to be 'you'
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
