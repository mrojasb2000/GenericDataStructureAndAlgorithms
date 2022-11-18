package main

import "fmt"

func addStudent(students []string, student string) []string {
	return append(students, student)
}

func main() {
	students := []string{}
	result := addStudent(students, "Michael")
	result = addStudent(result, "Jennifer")
	result = append(result, "Elaine")
	fmt.Println(result)
}
