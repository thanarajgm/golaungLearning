package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// Account struct maps to SQL table columns
type Account struct {
	ID    int
	Name  string
	Email string
	Phone string
}

func main() {

	// 1. CONNECT TO MYSQL
	db, err := sql.Open("mysql", "root:November@2025@tcp(localhost:3306)/testdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Connected to MySQL!")

	// 2. CREATE TABLE IF NOT EXISTS
	createAccountTable(db)

	// 3. CRUD Operations

	// INSERT (CREATE)
	newID := insertAccount(db, "Dhanaraj", "dhanaraj@gmail.com", "9876543210")
	fmt.Println("Inserted ID:", newID)

	// UPDATE
	updateAccount(db, newID, "Dhanaraj GM", "gm@gmail.com", "9999999999")

	// GET ONE (READ)
	acc := getAccount(db, newID)
	fmt.Println("Fetched Account:", acc)

	// GET ALL (READ ALL)
	allAcc := getAllAccounts(db)
	fmt.Println("All Accounts:", allAcc)

	// DELETE
	deleteAccount(db, newID)
	fmt.Println("Deleted Account ID:", newID)
}

// ========================================================================================
// CREATE ACCOUNT TABLE
// ========================================================================================

func createAccountTable(db *sql.DB) {

	query := `
	CREATE TABLE IF NOT EXISTS account (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(150),
		phone VARCHAR(20)
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Account table created (if not exists)")
}

// ========================================================================================
// INSERT ACCOUNT (CREATE)
// ========================================================================================

func insertAccount(db *sql.DB, name, email, phone string) int {

	query := "INSERT INTO account (name, email, phone) VALUES (?, ?, ?)"

	res, err := db.Exec(query, name, email, phone)
	if err != nil {
		panic(err)
	}

	lastID, _ := res.LastInsertId()

	return int(lastID)
}

// ========================================================================================
// UPDATE ACCOUNT
// ========================================================================================

func updateAccount(db *sql.DB, id int, name, email, phone string) {

	query := "UPDATE account SET name=?, email=?, phone=? WHERE id=?"

	_, err := db.Exec(query, name, email, phone, id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Account updated:", id)
}

// ========================================================================================
// GET ONE ACCOUNT (READ)
// ========================================================================================

func getAccount(db *sql.DB, id int) Account {

	query := "SELECT id, name, email, phone FROM account WHERE id=?"

	var acc Account

	err := db.QueryRow(query, id).Scan(&acc.ID, &acc.Name, &acc.Email, &acc.Phone)
	if err != nil {
		panic(err)
	}

	return acc
}

// ========================================================================================
// GET ALL ACCOUNTS (READ ALL)
// ========================================================================================

func getAllAccounts(db *sql.DB) []Account {

	query := "SELECT id, name, email, phone FROM account"

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var accounts []Account

	for rows.Next() {
		var acc Account
		rows.Scan(&acc.ID, &acc.Name, &acc.Email, &acc.Phone)
		accounts = append(accounts, acc)
	}

	return accounts
}

// ========================================================================================
// DELETE ACCOUNT (DELETE)
// ========================================================================================

func deleteAccount(db *sql.DB, id int) {

	query := "DELETE FROM account WHERE id=?"

	_, err := db.Exec(query, id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Account deleted:", id)
}
