package main

import "fmt"

type Student struct {
	Name string
	ID   int
	age  float64
}

func addStudent(students []string, student string) []string {
	return append(students, student)
}

func addStudentID(students []int, student int) []int {
	return append(students, student)
}

func addStudentStruct(students []Student, student Student) []Student {
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

	student2 := []Student{}
	result2 := addStudentStruct(student2, Student{"John", 213, 17.5})
	result2 = addStudentStruct(student2, Student{"James", 111, 18.75})
	result2 = addStudentStruct(student2, Student{"Marsha", 110, 16.25})
	fmt.Println(result2)
}
