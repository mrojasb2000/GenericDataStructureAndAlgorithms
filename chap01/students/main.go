package main

import "fmt"

func addStudent(students []string, student string) []string {
	return append(students, student)
}

func addStudentID(students []int, student int) []int {
	return append(students, student)
}

func main() {
	students := []string{}
	result := addStudent(students, "Michael")
	result = addStudent(result, "Jennifer")
	result = append(result, "Elaine")
	fmt.Println(result)

	students1 := []int{}
	result1 := addStudentID(students1, 155)
	result1 = addStudentID(result1, 112)
	result1 = addStudentID(result1, 120)
	fmt.Println(result1)
}
