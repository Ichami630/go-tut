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
		items: map[string]float64{},
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
	fs += fmt.Sprintf("%-10v ...$%v\n", "tips:", b.tips)
	fs += fmt.Sprintf("%-10v ...$%.2f", "total:", total+b.tips)

	return fs
}

//receiver functions with pointers
//note: For structs, pointers are automatically dereferenced

// update the tips
func (b *Bill) updateTips(newTip float64) {
	b.tips = newTip
}

// add item to the bill
func (b *Bill) addItem(key string, value float64) {
	b.items[key] = value
}
