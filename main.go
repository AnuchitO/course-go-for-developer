package main

import "fmt"

// Signature of the function
// func(int, int) int
func sum(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}

// TODO: create a function calculate that takes a function as an argument (func(int, int) int)

type myFunc func(int, int) int

func calculate(fn myFunc, a int, b int) int {
	return fn(a, b)
}

func main() {
	var fn = func(i1, i2 int) int {
		return i1 + i2
	}

	fmt.Println(fn(1, 2))

	r1 := calculate(sum, 1, 2)
	fmt.Println(r1) // 3

	r2 := calculate(minus, 1, 2)
	fmt.Println(r2) // -1

	fmt.Println("Hello, World!")
}

/* Java

import System.out;

public class Hello {
	public static void main(String[] args) {
		Integer number = 2;
		out.println("Hello, World!");
	}

	// lambda predicate method sum
	public static Integer calculate(predicate<Integer, Integer> sum, Integer a, Integer b) {
	}

	// signature of the function is Integer, Integer -> Integer
	public static Integer sum(Integer a, Integer b) {
		return a + b;
	}
}
*/
