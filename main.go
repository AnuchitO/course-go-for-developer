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

func main() {
	var fn func(int, int) int
	fn = sum
	result := fn(1, 2)
	fmt.Println(result) // 3

	result = calculate(sum, 1, 2)
	fmt.Println(result) // 3

	result = calculate(minus, 1, 2)
	fmt.Println(result) // 3

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
