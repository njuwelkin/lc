package main

import "fmt"

/*
Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

Example:

n = 15,

Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
*/
func fizzBuzz(n int) []string {
	ret := make([]string, n)
	for i := 3; i <= n; i += 3 {
		ret[i-1] = "Fizz"
	}
	for i := 5; i <= n; i += 5 {
		ret[i-1] += "Buzz"
	}
	for i := 1; i <= n; i++ {
		if ret[i-1] == "" {
			ret[i-1] = fmt.Sprintf("%d", i)
		}
	}
	return ret
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(fizzBuzz(15))
}
