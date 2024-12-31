//Exercise 2: Bank Transaction System

package main

import (
	"errors"
	"fmt"
)

// Account struct to store account details
type Account struct {
	AccountID       int
	HolderName      string
	Balance         float64
	Transactions    []string
}

// Global slice to store all accounts
var bankAccounts []Account

// Deposit method to add funds to an account
func (acc *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	acc.Balance += amount
	acc.Transactions = append(acc.Transactions, fmt.Sprintf("Deposited: %.2f", amount))
	return nil
}

// Withdraw method to deduct funds from an account
func (acc *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if acc.Balance < amount {
		return errors.New("insufficient funds")
	}
	acc.Balance -= amount
	acc.Transactions = append(acc.Transactions, fmt.Sprintf("Withdrew: %.2f", amount))
	return nil
}

// ViewTransactions method to display account transaction history
func (acc *Account) ViewTransactions() {
	fmt.Println("Transaction History:")
	if len(acc.Transactions) == 0 {
		fmt.Println("No transactions recorded.")
		return
	}
	for index, transaction := range acc.Transactions {
		fmt.Printf("%d. %s\n", index+1, transaction)
	}
}

// FindAccount searches for an account by its ID
func FindAccount(id int) (*Account, error) {
	for i := range bankAccounts {
		if bankAccounts[i].AccountID == id {
			return &bankAccounts[i], nil
		}
	}
	return nil, errors.New("account not found")
}

// DisplayMenu function shows the available options
func DisplayMenu() {
	fmt.Println("\n--- Welcome to Bank Transaction System ---")
	fmt.Println("1. Create Account")
	fmt.Println("2. Deposit Funds")
	fmt.Println("3. Withdraw Funds")
	fmt.Println("4. Check Balance")
	fmt.Println("5. View Transactions")
	fmt.Println("6. Exit")
	fmt.Print("Please select an option: ")
}

func main() {
	// Preloaded accounts
	bankAccounts = append(bankAccounts, Account{AccountID: 101, HolderName: "Alice", Balance: 5000})
	bankAccounts = append(bankAccounts, Account{AccountID: 102, HolderName: "Bob", Balance: 3000})
	bankAccounts = append(bankAccounts, Account{AccountID: 103, HolderName: "Charlie", Balance: 7000})

	for {
		DisplayMenu()
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter Account ID: ")
			var id int
			fmt.Scan(&id)

			fmt.Print("Enter Account Holder Name: ")
			var name string
			fmt.Scan(&name)

			fmt.Print("Enter Initial Deposit: ")
			var initialDeposit float64
			fmt.Scan(&initialDeposit)

			bankAccounts = append(bankAccounts, Account{AccountID: id, HolderName: name, Balance: initialDeposit})
			fmt.Println("Account successfully created.")

		case 2:
			fmt.Print("Enter Account ID: ")
			var id int
			fmt.Scan(&id)

			account, err := FindAccount(id)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Print("Enter Deposit Amount: ")
			var amount float64
			fmt.Scan(&amount)

			if err := account.Deposit(amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful.")
			}

		case 3:
			fmt.Print("Enter Account ID: ")
			var id int
			fmt.Scan(&id)

			account, err := FindAccount(id)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Print("Enter Withdrawal Amount: ")
			var amount float64
			fmt.Scan(&amount)

			if err := account.Withdraw(amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful.")
			}

		case 4:
			fmt.Print("Enter Account ID: ")
			var id int
			fmt.Scan(&id)

			account, err := FindAccount(id)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Printf("Current Balance for %s: %.2f\n", account.HolderName, account.Balance)

		case 5:
			fmt.Print("Enter Account ID: ")
			var id int
			fmt.Scan(&id)

			account, err := FindAccount(id)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			account.ViewTransactions()

		case 6:
			fmt.Println("Thank you for using the Bank Transaction System. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}