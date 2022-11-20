package main

import (
	"fmt"
	"sort"
)

type Ordered interface {
	~int | ~float64 | ~string
}

type Student struct {
	Name string
	ID   int
	Age  float64
}

func addStudent[T any](students []T, student T) []T {
	return append(students, student)
}

type OrderedSlice[T Ordered] []T

func (s OrderedSlice[T]) Len() int {
	return len(s)
}

func (s OrderedSlice[T]) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s OrderedSlice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type SortType[T any] struct {
	slice   []T
	compare func(T, T) bool
}

func (s SortType[T]) Len() int {
	return len(s.slice)
}

func (s SortType[T]) Less(i, j int) bool {
	return s.compare(s.slice[i], s.slice[j])
}

func (s SortType[T]) Swap(i, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}

func PerformSort[T any](slice []T, compare func(T, T) bool) {
	sort.Sort(SortType[T]{slice, compare})
}

func main() {
	students := []string{}
	result := addStudent(students, "Michale")
	result = addStudent(result, "Jennifer")
	result = addStudent(result, "Elaine")
	sort.Sort(OrderedSlice[string](result))
	fmt.Println(result)
}
