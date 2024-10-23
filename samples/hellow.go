package main

import "fmt"

func hello(s string) (string, string) {
	return fmt.Sprintf("Hello, %s!", s), "Hi, world"
}
