package main

import "fmt"

type Bill struct {
	name  string
	items map[string]float64
	tips  float64
}

// make new bill
func newBill(name string) Bill {
	b := Bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "coke": 9.99},
		tips:  0,
	}

	return b
}

// format the bill- Receiver functions
func (b Bill) format() string {
	fs := "Bill Breakdown:\n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-10v ...$%v\n", k+":", v) //-25 adds some padding to the left
		total += v
	}
	fs += fmt.Sprintf("%-10v ...$%.2f", "total:", total)

	return fs
}
