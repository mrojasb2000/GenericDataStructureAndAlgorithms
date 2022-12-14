package main

import "fmt"

type Stringer = interface {
	String() string
}

type Integer int

func (i Integer) String() string {
	return fmt.Sprintf("%d", i)
}

type String string

func (s String) String() string {
	return string(s)
}

type Student struct {
	Name string
	ID   int
	Age  float64
}

func (s Student) String() string {
	return fmt.Sprintf("%s %d %0.2f", s.Name, s.ID, s.Age)
}

func addStudent[T Stringer](students []T, student T) []T {
	return append(students, student)
}

func main() {
	students := []String{}
	result := addStudent(students, "Michael")
	result = addStudent(result, "Jennifer")
	result = append(result, "Elaine")
	fmt.Println(result)

	students1 := []Integer{}
	result1 := addStudent(students1, 155)
	result1 = addStudent(result1, 112)
	result1 = addStudent(result1, 120)
	fmt.Println(result1)

	student2 := []Student{}
	result2 := addStudent(student2, Student{"John", 213, 17.5})
	result2 = addStudent(result2, Student{"James", 111, 18.75})
	result2 = addStudent(result2, Student{"Marsha", 110, 16.25})
	fmt.Println(result2)
}
