package main

import (
	"fmt"
	"os"
)

type bill struct {
	bill_name        string
	items_and_prices map[string]float64
	tip              float64
}

func createBill(name string) bill {
	init_bill := bill{
		bill_name:        name,
		items_and_prices: map[string]float64{},
		tip:              0,
	}

	return init_bill
}

func (bill_to_format *bill) format_bill() string {
	formated_bill := "Bill Details\n"
	var total float64 = 0

	// bill name
	formated_bill += fmt.Sprintf("Bill name: %v\n", bill_to_format.bill_name)

	// List map items
	for key, value := range bill_to_format.items_and_prices {
		formated_bill += fmt.Sprintf("%-25v ... $%v\n", key+":", value)
		total += value
	}

	// tip
	formated_bill += fmt.Sprintf("%-25v ... $%v\n", "Tip:", bill_to_format.tip)

	// total
	formated_bill += fmt.Sprintf("%-25v ... $%0.2f\n", "Total:", total+bill_to_format.tip)

	return formated_bill
}

func (bill_to_add_tip *bill) add_tip_to_bill(tip float64) {
	bill_to_add_tip.tip = tip
}

func (bill_to_add_item *bill) add_item_to_bill(name string, price float64) {
	bill_to_add_item.items_and_prices[name] = price
}

func (bill_to_save bill) saveBill() {
	data := []byte(bill_to_save.format_bill())

	err := os.WriteFile("E:/archive_root/CSE/Go/protfolio_go/projects_go/bill_project_go/"+bill_to_save.bill_name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill is saved to file.")
}
