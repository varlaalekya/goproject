package greet

import "fmt"

func Hello(name string) string {
	if name == "" {
		name = "there"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
