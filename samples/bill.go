package main

import "fmt"

type bill struct {
	id    int
	items map[string]float64
	name  string
}

func newBill(name string) bill {
	b := bill{
		id:    1,
		items: map[string]float64{"pie": 5.99, "cake": 7.99},
		name:  name,
	}
	return b
}

func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0.0

	fs += fmt.Sprintf("%-15v  %5v\n \n", "Name:", b.name)
	for k, v := range b.items {
		fs += fmt.Sprintf("%-15v  ...$%5v \n", k+":", v)

		total += v
	}
	fs += fmt.Sprintf("%-15v  ...$%3.2f", "Total:", total)

	return fs
}

func (b *bill) updateName(name string) {
	// (*b).name = name //dereference
	b.name = name // no need in stuct
}
