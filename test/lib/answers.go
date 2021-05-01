package lib

import (
	"fmt"
	"strconv"
	"strings"
)

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

	total := 0

	feeling, _ := strconv.ParseInt(a.Feeling, 10, 64)

	// Six elements are because feeling is not indexed at zero
	trueFeeling := [6]int{0, 5, 4, 3, 2, 1}

	total += trueFeeling[feeling]

	temperature, _ := strconv.ParseFloat(a.Temperature, 32)

	if temperature < 97.0 || temperature > 99.0 {
		total += 1
	}

	radioQuestions := [4]string{a.Sick, a.Sore, a.Contact, a.Location}

	for _, val := range radioQuestions {
		lowercase := strings.ToLower(val)

		if lowercase == "yes" {
			total += 1
		}
	}

	fmt.Println("Total=", total)

	if total <= 3 {
		return 3
	} else if total <= 6 {
		return 2
	} else {
		return 1
	}

}
