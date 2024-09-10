package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func main() {

	{
		var result int
		result = sum(1, 2)
		fmt.Println(result)
		{
			fmt.Println(result)
		}
	}

	fmt.Println("Hello, World!")
}

/* Java

import System.out;

public class Hello {
	public static void main(String[] args) {
		Integer number = 2;
		out.println("Hello, World!");
	}
}
*/
