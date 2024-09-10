package main

import "fmt"

func NewFlight(airlineCode string, number int, destination string) Flight {
	f := Flight{
		number:      number,
		airlineCode: airlineCode,
		destination: destination,
	}

	return f
}

type Flight struct {
	number      int
	airlineCode string
	destination string
}

func (f Flight) getFlightNumber() string {
	code := fmt.Sprintf("%s %d", f.airlineCode, f.number)
	return code
}

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

	code := fl.getFlightNumber()
	fmt.Println("code in main: ", code)
}
