package main

import (
	"fmt"
)

func slice() []int {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{-1, -2, -3}
	return append(append(s1[:1], s2...), s1[1:]...)
}

func slice1() []int {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{-1, -2, -3, -4}
	return append(append(s1[:1], s2...), s1[1:]...)
}

func main() {
	// output?
	fmt.Printf("func sl=%+v\n", slice())
	fmt.Printf("func sl1=%+v\n", slice1())

	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("s len - %v, cap - %v, pointer -  %p, val - %v\n", len(s), cap(s), s, s)
	s1 := s[1:3]
	fmt.Printf("before append s1 len - %v cap - %v pointer - %p val - %v\n", len(s1), cap(s1), s1, s1)
	s1 = append(s1, 10, 11)
	fmt.Printf("afert 1 append s1 len - %v cap - %v pointer - %p val - %v\n", len(s1), cap(s1), s1, s1)

	s1 = append(s1, 12, 13, 14)
	fmt.Printf("afert 2 append s1 len - %v cap - %v pointer - %p val - %v\n", len(s1), cap(s1), s1, s1)

	fmt.Println("修改前")
	sl := []int{1, 2, 3, 4}
	fmt.Println("sl: ", sl)
	sl1 := sl[1:3]
	fmt.Println("sl1: ", sl1)
	sl1[1] = 10
	fmt.Println("修改后")
	fmt.Println("sl: ", sl)
	fmt.Println("sl1: ", sl1)

}
