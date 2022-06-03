package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(promt string, reader *bufio.Reader) (string, error) {
	fmt.Print(promt)
	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createNewBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Enter the bill name: ", reader)

	new_bill := createBill(name)

	fmt.Println("Bill", new_bill.bill_name, "is Created.")

	return new_bill
}

func promtOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	options, _ := getInput("Choose an option (a - add item, t - add tip, s - save bill): ", reader)

	switch options {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		price_value, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number.")
			promtOptions(b)
		}
		b.add_item_to_bill(name, price_value)

		fmt.Println("Item added - ", name, price_value)
		promtOptions(b)
	case "t":
		tip, _ := getInput("Tip amount ($): ", reader)

		tip_value, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number.")
			promtOptions(b)
		}
		b.add_tip_to_bill(tip_value)

		fmt.Println("Tip added - ", tip_value)
		promtOptions(b)
	case "s":
		b.saveBill()
		fmt.Println("File saved - ", b.bill_name)
	default:
		fmt.Println("that was not a valid option...")
		promtOptions(b)
	}
}

func main() {
	new_bill := createNewBill()
	promtOptions(new_bill)
}
