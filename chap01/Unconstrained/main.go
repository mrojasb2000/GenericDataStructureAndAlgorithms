package main

import "fmt"

type Student struct {
	Name string
	ID   int
	Age  float32
}

func addStudent[T any](students []T, student T) []T {
	return append(students, student)
}

func main() {
	students := []string{}
	result := addStudent(students, "Michale")
	result = append(result, "Jennifer")
	result = append(result, "Elaine")
	fmt.Println(result)
}
