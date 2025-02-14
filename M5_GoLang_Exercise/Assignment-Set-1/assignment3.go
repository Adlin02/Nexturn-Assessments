package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Item struct representing an item in the stock
type Item struct {
	ID    int
	Name  string
	Price float64
	Quantity int
}

// Stock slice to hold all items
var stock []Item

// AddItem adds a new item to the stock
func AddItem(id int, name string, price float64, quantity int) error {
	for _, item := range stock {
		if item.ID == id {
			return errors.New("item ID already exists")
		}
	}

	stock = append(stock, Item{ID: id, Name: name, Price: price, Quantity: quantity})
	return nil
}

// UpdateQuantity updates the quantity of a specific item by ID
func UpdateQuantity(id int, newQuantity int) error {
	if newQuantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	for i, item := range stock {
		if item.ID == id {
			stock[i].Quantity = newQuantity
			return nil
		}
	}
	return errors.New("item not found")
}

// FindItem searches for an item by ID or name
func FindItem(query string) (*Item, error) {
	for _, item := range stock {
		if strings.EqualFold(item.Name, query) || fmt.Sprintf("%d", item.ID) == query {
			return &item, nil
		}
	}
	return nil, errors.New("item not found")
}

// DisplayStock displays all items in the stock
func DisplayStock() {
	fmt.Printf("%-5s %-20s %-10s %-10s\n", "ID", "Name", "Price", "Quantity")
	fmt.Println(strings.Repeat("-", 50))
	for _, item := range stock {
		fmt.Printf("%-5d %-20s %-10.2f %-10d\n", item.ID, item.Name, item.Price, item.Quantity)
	}
}

// SortStock sorts items by a specified attribute (price or quantity)
func SortStock(by string) {
	switch by {
	case "price":
		sort.Slice(stock, func(i, j int) bool {
			return stock[i].Price < stock[j].Price
		})
	case "quantity":
		sort.Slice(stock, func(i, j int) bool {
			return stock[i].Quantity < stock[j].Quantity
		})
	}
}

func main() {
	fmt.Println("Welcome to the Stock Management System")

	// Preloaded items
	_ = AddItem(101, "Pens", 5.00, 100)
	_ = AddItem(102, "Notebooks", 50.00, 200)
	_ = AddItem(103, "Markers", 15.00, 80)
	_ = AddItem(104, "Folders", 25.00, 50)

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Add Item")
		fmt.Println("2. Update Quantity")
		fmt.Println("3. Find Item")
		fmt.Println("4. Display Stock")
		fmt.Println("5. Sort Stock by Price")
		fmt.Println("6. Sort Stock by Quantity")
		fmt.Println("7. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id, quantity int
			var name string
			var price float64
			fmt.Print("Enter Item ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Item Name: ")
			fmt.Scan(&name)
			fmt.Print("Enter Item Price: ")
			fmt.Scan(&price)
			fmt.Print("Enter Item Quantity: ")
			fmt.Scan(&quantity)
			if err := AddItem(id, name, price, quantity); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Item added successfully.")
			}
		case 2:
			var id, newQuantity int
			fmt.Print("Enter Item ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter New Quantity: ")
			fmt.Scan(&newQuantity)
			if err := UpdateQuantity(id, newQuantity); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Quantity updated successfully.")
			}
		case 3:
			var query string
			fmt.Print("Enter Item Name or ID: ")
			fmt.Scan(&query)
			if item, err := FindItem(query); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Found: %+v\n", *item)
			}
		case 4:
			DisplayStock()
		case 5:
			SortStock("price")
			fmt.Println("Stock sorted by price.")
			DisplayStock()
		case 6:
			SortStock("quantity")
			fmt.Println("Stock sorted by quantity.")
			DisplayStock()
		case 7:
			fmt.Println("Exiting... Thank you!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}