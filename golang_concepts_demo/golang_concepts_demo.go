package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

//////////////////////
// 1) STRUCT + EMBEDDING
//////////////////////

// Base struct (Person)
type Person struct {
	Name string
	Age  int
}

// Employee embeds Person (inherits fields)
type Employee struct {
	Person
	Role string
}

//////////////////////
// 2) INTERFACE
//////////////////////

type LoginSystem interface {
	Login(username string, password string) error
}

//////////////////////
// 3) CUSTOM ERROR
//////////////////////

type LoginError string

func (e LoginError) Error() string {
	return string(e)
}

//////////////////////
// 4) SYSTEM USING MAP, POINTERS & CLOSURE
//////////////////////

func loginCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

type Auth struct {
	users map[string]string // username → password
}

func (a *Auth) Login(username, password string) error {

	// Simulate delay
	time.Sleep(1 * time.Second)

	// Check if user exists
	if pwd, ok := a.users[username]; ok {
		if pwd == password {
			return nil
		}
		return LoginError("❌ Wrong password")
	}
	return LoginError("❌ User does not exist")
}

//////////////////////
// 5) PANIC + RECOVER + DEFER EXAMPLE
//////////////////////

func riskyOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("⚠️ Recovered from panic:", r)
		}
	}()

	panic("🔥 Something went seriously wrong in system!")
}

//////////////////////
// 6) MAIN FUNCTION
//////////////////////

func main() {

	log.Println("🟢 System Starting...")

	// Create employee using struct & embedding
	emp := Employee{
		Person: Person{"Rahul", 28},
		Role:   "Developer",
	}

	fmt.Println("Employee:", emp.Name, "-", emp.Role)

	// Pointer Example (change age)
	agePtr := &emp.Age
	*agePtr = 30
	fmt.Println("Updated Age using Pointer:", emp.Age)

	// Create auth system using map
	auth := Auth{
		users: map[string]string{
			"rahul": "1234",
		},
	}

	// Closure login counter
	countLogin := loginCounter()

	for i := 1; i <= 3; i++ {
		fmt.Printf("\n🔁 Login Attempt %d\n", countLogin())

		err := auth.Login("rahul", "wrongpass")

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("✅ Login Success!")
		break
	}

	// Normal error example
	_, err := divide(10, 0)
	if err != nil {
		fmt.Println("Divide Error:", err)
	}

	// Panic Example
	riskyOperation()

	// Sleep + log final message
	log.Println("⏳ Closing system in 2 seconds...")
	time.Sleep(2 * time.Second)

	fmt.Println("👋 Program finished safely.")
}

//////////////////////
// 7) NORMAL ERROR HANDLING EXAMPLE
//////////////////////

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("❌ Cannot divide by zero")
	}
	return a / b, nil
}
