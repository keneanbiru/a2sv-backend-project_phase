package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type student struct {
	name    string
	subject map[string]float64
}

func grade_calculator(grades map[string]float64) (float64, string, error) {
	if len(grades) == 0 {
		return 0, "No Grades", fmt.Errorf("no grades provided")
	}

	var total float64
	var count int
	for _, grade := range grades {
		total += grade
		count++
	}

	average := total / float64(count)
	var finalGrade string
	if average >= 90 {
		finalGrade = "A"
	} else if average >= 80 {
		finalGrade = "B"
	} else if average >= 70 {
		finalGrade = "C"
	} else if average >= 60 {
		finalGrade = "D"
	} else {
		finalGrade = "F"
	}

	return average, finalGrade, nil
}

func main() {
	var nosub int
	var sub string
	var grade float64
	var err error

	student1 := &student{
		subject: make(map[string]float64),
	}

	fmt.Println("Hey, What is your name?")
	reader := bufio.NewReader(os.Stdin)
	for {
		student1.name, err = reader.ReadString('\n')
		if err == nil {
			student1.name = strings.TrimSpace(student1.name)
			break
		}
		fmt.Println("Error reading name:", err)
	}

	fmt.Println("please, enter the number of subjects you take?")
	for {
		_, err = fmt.Scanln(&nosub)
		if err == nil {
			break
		}
		fmt.Println("Error reading number of subjects:", err)
	}

	for i := 0; i < nosub; i++ {
		for {
			fmt.Println("Enter the " + strconv.Itoa(i+1) + " subject name")
			sub, err = reader.ReadString('\n')
			if err == nil && strings.TrimSpace(sub) != "" {
				break
			}
			if err != nil {
				fmt.Println("Error reading subject name:", err)
			} else {
				fmt.Println("Subject name cannot be empty. Please try again.")
			}
		}

		for {
			fmt.Println("Enter the " + strings.TrimSpace(sub) + " subject grade")
			_, err = fmt.Scanln(&grade)
			if err == nil && grade >= 0 && grade <= 100 {
				student1.subject[strings.TrimSpace(sub)] = grade
				break
			}
			if err != nil {
				fmt.Println("Error reading subject grade:", err)
			} else {
				fmt.Println("Invalid grade value. Please try again.")
			}
		}
	}

	res, ans, err := grade_calculator(student1.subject)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Student: %s\n", student1.name)
	fmt.Printf("Subjects and grades:\n")
	for subject, grade := range student1.subject {
		fmt.Printf("  %s: %.2f\n", subject, grade)
	}
	fmt.Printf("Average score: %.2f, Final grade: %s\n", res, ans)
}
