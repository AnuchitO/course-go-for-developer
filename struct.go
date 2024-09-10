package main

import "fmt"

type Flight struct {
	number      int
	airlineCode string
	departure   string
}

// public Flight() {
// }

// public Flight(String airlineCode, String number) {
// 	this.airlineCode = airlineCode;
// 	this.number = number;
// }

// public String getFlightNumber() {
// 	return this.airlineCode + this.number;
// }

func main() {
	// Flight f = new Flight();
	f := Flight{}
	// var f Flight
	fmt.Printf("%#v\n", f)
	fmt.Println(f.airlineCode)
	fmt.Println(f.number)

	// Flight f = new Flight("AA", "123");
	fl := Flight{123, "AA", "DMK"}
	fmt.Printf("%#v\n", fl)
	ff := Flight{number: 123, airlineCode: "AA"}
	fmt.Printf("%#v\n", ff)
}
