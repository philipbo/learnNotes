package main

import "log"

func f() (i int) {
	defer func() {
		i++
	}()

	return 0
}

func f1() int {
	i := 0
	defer func() {
		i++
	}()

	return i
}

func f2() int {
	i := 5
	defer func() {
		i = i + 5
	}()

	return i
}

func f3() (i int) {
	i = 5
	defer func() {
		i = i + 5
	}()

	return i
}

func f4() (i int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f5() (i int) {
	defer func(i int) {
		i = i + 5
	}(i)

	return 1
}

func f6() int {
	i := 0
	defer log.Printf("in defer i %d", i) //output 0
	i++
	return i
}

func main() {
	log.Printf("f %d", f())
	log.Printf("f1 %d", f1())
	log.Printf("f2 %d", f2())
	log.Printf("f3 %d", f3())
	log.Printf("f4 %d", f4())
	log.Printf("f5 %d", f5())
	log.Printf("f6 %d", f6())
}
