package main

import "fmt"

// define a struct
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// modifying tip
func (b *bill) updateTip(num float64) {
	b.tip = num
}

// adding items to our map
func (b *bill) addItems(key string, val float64) {
	b.items[key] = val
}

// method to output total
func (b bill) totalAmount() float64 {
	total := b.tip

	for _, v := range b.items {
		total += v
	}
	return total
}

// prints bill
func (b bill) printBill() {
	fmt.Println("\nBill name \n", b.name)
	fmt.Println("items:")
	for item, price := range b.items {
		fmt.Printf(" - %s: $%.2f\n", item, price)
	}
	fmt.Println("Tip:", b.tip)
	fmt.Println("Total Amount:", b.totalAmount())
}

func main() {
	myBill := bill{name: "Alice's Bill", items: map[string]float64{}}

	myBill.addItems("Eggs", 5.1)
	myBill.addItems("Cake", 10.2)

	myBill.updateTip(20.0)

	myBill.printBill()
}
