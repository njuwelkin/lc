package main

import "fmt"

func numberToWords(num int) string {
	unitStr := []string{"", "thousand ", "million ", "billion "}
	numStr := []string{"", "one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "fifteen",
		"sixteen", "seventeen", "eighteen", "nineteen", "twenty"}
	tenStr := []string{"", "ten ", "twenty ", "thirty ", "forty ", "fifty ",
		"sixty ", "seventy ", "eighty ", "ninety "}
	unit := []int{1, 1000, 1000000, 1000000000}

	for i := len(unit) - 1; i >= 0; i-- {
		if num < unit[i] {
			continue
		}
		val := num / unit[i]
		num %= unit[i]
		h := val / 100
		val %= 100
		if h > 0 {
			fmt.Printf("%s handred ", numStr[h])
			if val > 0 {
				fmt.Print("and ")
			}
		}
		if val == 0 {
			continue
		} else if val < 20 {
			fmt.Print(numStr[val], " ")
		} else {
			fmt.Print(tenStr[val/10])
			fmt.Print(numStr[val%10], " ")
		}
		fmt.Print(unitStr[i])
	}

	return ""
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(numberToWords(123))
	fmt.Println(numberToWords(12345))
	fmt.Println(numberToWords(1234567))
	fmt.Println(numberToWords(1234567891))
}
