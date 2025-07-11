package main

import (
	"fmt"
	"strings"
)
func average(input map[string]float64) float64 {
    var total float64
    for _, grade := range input {
        total += grade
    }
    if len(input) == 0 {
        return 0
    }
    return total / float64(len(input))
}

func main() {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scan(&name)
    name = strings.TrimSpace(name)

    var subjects int
    fmt.Print("Enter the number of subjects: ")
    _, err := fmt.Scan(&subjects)
    if err != nil || subjects <= 0 {
        fmt.Println("Invalid input. Please enter a positive number for subjects.")
		return
    }

    inputs := make(map[string]float64)
    var subject string
    var grade float64
    
	i := 0
    for i < subjects {
        fmt.Print("Enter the subject name and the grade (e.g., Math 90): ")
        _, err := fmt.Scan(&subject, &grade)
        if err != nil {
            fmt.Println("Invalid input. Try again with a subject name followed by a numeric grade.")
            var discard string
        	fmt.Scanln(&discard)
			continue
        }
        if grade < 0 || grade > 100 {
            fmt.Println("Invalid range, Grade must be between 0 and 100.")
            continue
        }
        inputs[strings.TrimSpace(subject)] = grade
        i++
    }

    fmt.Printf("\nGrade Report of %s\n", name)
	fmt.Println("====================")
    fmt.Printf("%s | %s\n", "Subject", "Grade")
    for sub, grade := range inputs {
        fmt.Printf("%s | %.2f\n", sub, grade)
    }
    fmt.Println("====================")
    avg := average(inputs)
    fmt.Printf("%s | %.2f\n", "Average", avg)
}
