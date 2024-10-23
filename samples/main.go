package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func update(n *string) string {
	*n = "new name"
	return *n //address the value of the pointer

}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func main() {
	myBill := newBill("old name")

	fmt.Println("bill:", myBill.format())

	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Enter new name: ", reader)

	myBill.updateName(name)
	fmt.Println("bill:", myBill.format())
	data := []byte(myBill.format())

	err := os.WriteFile("bill.txt", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}

}
