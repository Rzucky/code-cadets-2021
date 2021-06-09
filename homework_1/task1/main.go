package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"code-cadets-2021/homework1/task1/fizzbuzz"
)

func main() {

	var start, end int

	// default values are 1-10 unless stated
	flag.IntVar(&start, "start", 1, "Value (inclusive) from which to start counting numbers for FizzBuzz")
	flag.IntVar(&end, "end", 10, "Value (inclusive) to end FizzBuzz game with")

	flag.Parse()

	// Output is in a slice
	fizzBuzzOutputSlice, err := fizzbuzz.FizzBuzz(start, end)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(strings.Join(fizzBuzzOutputSlice, " "))

}
