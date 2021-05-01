package lib

import "fmt"

// ===============================================
// Types
// ===============================================

type Answers struct {
	Name        string
	Feeling     string
	Temperature string
	Sick        string
	Sore        string
	Contact     string
	Location    string
}

// ===============================================
// Functions
// ===============================================

// makeStrArray() -> creates a str array from struct Answers
func (a Answers) makeStrArray() []string {
	return []string{a.Name, a.Feeling, a.Temperature, a.Sick, a.Sore, a.Contact, a.Location}
}

// iter() -> Iterates thru array with func
func iter(arr []string, f func(string)) {
	for _, val := range arr {
		f(val)
	}
}

// CheckAll -> Checks if all the values are filled
func (a Answers) CheckAll() bool {
	answers := a.makeStrArray()

	retBool := true

	iter(answers, func(s string) {
		if s == "" {
			retBool = false
		}
	})
	return retBool
}

// PrintAll -> Prints all the values
func (a Answers) PrintAll() {
	answers := a.makeStrArray()

	iter(answers, func(s string) {
		fmt.Printf("%s ", s)
	})
	fmt.Println()
}

// Calc() -> Calculates the score for the
func (a Answers) Calc() int8 {
	return int8(8)
}
