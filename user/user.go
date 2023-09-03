package user

import (
	"fmt"
)

// INTERFACE (implicit implementation)
// everything implements empty interface{} cause it doesn`t contain methods
// could be composition of interfaces
type Student interface {
	LeaveUni()
}

type Age uint

func (age Age) IsAdult() bool {
	return age >= 18
}

type User struct {
	name   string
	Aage    Age
	score  rune
	insane bool
	salary float64
}

func (user User) LeaveUni() {
	fmt.Println(user.name + " leaves University")
}

// value receiver (obj copied)
func (user User) PrintUser() {
	fmt.Println("Name:", user.name, "\nAge:", user.Aage, "\nScore:", user.score, "\nInsanity:", user.insane, "\nSalary:", user.salary)
}

// pointer receiver (obj copied)
func (user *User) UpdateUser(newAge Age) {
	user.Aage = newAge
}

func UserConstructor(name string, age uint8, score rune, insane bool, salary float64) User {
	return User{
		name:   name,
		Aage:    Age(age),
		score:  score,
		salary: salary,
		insane: !insane,
	}
}

