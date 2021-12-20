package maintidx

import "fmt"

func f1(arg string) string { // want "Function name: f1, Cyclomatic Complexity: 3, Halstead Volume: 81.41, Maintainability Index: 64"
	var test = "aaa"
	if false {
		if false {
			return arg
		}
		return test
	}
	return ""
}

func f2() { // want "Function name: f2, Cyclomatic Complexity: 1, Halstead Volume: 18.09, Maintainability Index: 80"
	print("Hello, World")
}

func f3() { // want "Function name: f3, Cyclomatic Complexity: 1, Halstead Volume: 101.58, Maintainability Index: 67"
	a := 2
	b := 1
	c := 3
	avg := (a + b + c) / 3
	println(avg)
}

func f4() { // want "Function name: f4, Cyclomatic Complexity: 2, Halstead Volume: 25.27, Maintainability Index: 71"
	if false {

	} else {

	}
}

func f5() { // want "Function name: f5, Cyclomatic Complexity: 8, Halstead Volume: 144.00, Maintainability Index: 55"
	for true {
		if false {

		} else if false {

		} else if false {

		} else if false {
			n := 0
			switch n {
			case 0:
			case 1:
			default:
			}
		} else {

		}
	}
}

type t1 struct {
}

func (t *t1) receive() { // want "Function name: receive, Cyclomatic Complexity: 1, Halstead Volume: 22.46, Maintainability Index: 83"
}

func callIncr() { // want "Function name: callIncr, Cyclomatic Complexity: 1, Halstead Volume: 38.04, Maintainability Index: 73"
	var a int
	a++
	print(a)
}

func callDefer() { // want "Function name: callDefer, Cyclomatic Complexity: 1, Halstead Volume: 41.21, Maintainability Index: 75"
	defer fmt.Println("world")
	fmt.Println("hello")
}

func callGo() { // want "Function name: callGo, Cyclomatic Complexity: 1, Halstead Volume: 41.21, Maintainability Index: 75"
	go fmt.Println("hello")
	fmt.Println("world")
}

func f10() { // want "Function name: f10, Cyclomatic Complexity: 1, Halstead Volume: 101.58, Maintainability Index: 67"
	a := make(chan string)
	go func() { a <- "send" }()

	b := <-a
	fmt.Println(b)
}

func f11() { // want "Function name: f11, Cyclomatic Complexity: 1, Halstead Volume: 27.00, Maintainability Index: 76"
	fmt.Println("hello, world")
	return
}

func f12() { // want "Function name: f12, Cyclomatic Complexity: 3, Halstead Volume: 92.00, Maintainability Index: 63"
	var a int
	for a < 5 {
		if a < 3 {
			continue
		} else {
			break
		}
		a++
	}
}

func f13() { // want "Function name: f13, Cyclomatic Complexity: 3, Halstead Volume: 136.16, Maintainability Index: 62"
	c1 := make(chan string)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		}
	}
}

func f14() { // want "Function name: f14, Cyclomatic Complexity: 2, Halstead Volume: 82.04, Maintainability Index: 69"
	a := []int{0, 1, 2}
	for b := range a {
		fmt.Println(b)
	}
}
