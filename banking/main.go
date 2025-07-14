package main

import (
    "fmt"
)

type Account struct {
    ID      int
    Name    string
    Email   string
    Balance float64
}

var accounts []Account
var nextID = 1

func createAccount() Account {
    var name, email string
    fmt.Print("Enter name: ")
    fmt.Scanln(&name)
    fmt.Print("Enter email: ")
    fmt.Scanln(&email)
    acc := Account{ID: nextID, Name: name, Email: email, Balance: 0}
    nextID++
    fmt.Println("Account created:", acc)
    return acc
}

func deposit() {
    var id int
    var amount float64
    fmt.Print("Enter Account ID: ")
    fmt.Scanln(&id)
    fmt.Print("Enter deposit amount: ")
    fmt.Scanln(&amount)
    for i := range accounts {
        if accounts[i].ID == id {
            accounts[i].Balance += amount
            fmt.Println("Deposit successful. New balance:", accounts[i].Balance)
            return
        }
    }
    fmt.Println("Account not found.")
}

func withdraw() {
    var id int
    var amount float64
    fmt.Print("Enter Account ID: ")
    fmt.Scanln(&id)
    fmt.Print("Enter withdrawal amount: ")
    fmt.Scanln(&amount)
    for i := range accounts {
        if accounts[i].ID == id {
            if accounts[i].Balance >= amount {
                accounts[i].Balance -= amount
                fmt.Println("Withdrawal successful. New balance:", accounts[i].Balance)
            } else {
                fmt.Println("Insufficient balance.")
            }
            return
        }
    }
    fmt.Println("Account not found.")
}

func listAccounts() {
    for _, acc := range accounts {
        fmt.Printf("ID: %d, Name: %s, Email: %s, Balance: %.2f\n", acc.ID, acc.Name, acc.Email, acc.Balance)
    }
}

func main() {
    for {
        fmt.Println("\n1. Create Account")
        fmt.Println("2. Deposit")
        fmt.Println("3. Withdraw")
        fmt.Println("4. Show Accounts")
        fmt.Println("5. Exit")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            acc := createAccount()
            accounts = append(accounts, acc)
        case 2:
            deposit()
        case 3:
            withdraw()
        case 4:
            listAccounts()
        case 5:
            return
        default:
            fmt.Println("Invalid option")
        }
    }
}
