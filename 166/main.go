package main

import (
	"bytes"
	"fmt"
)

func fractionToDecimal(numerator int, denominator int) string {
	if denominator == 0 {
		return ""
	}
	if numerator == 0 {
		return "0"
	}

	var ret bytes.Buffer
	neg := false
	if numerator < 0 {
		numerator = -numerator
		neg = true
	}
	if denominator < 0 {
		denominator = -denominator
		neg = !neg
	}
	if neg {
		ret.WriteString("-")
	}

	intPart := numerator / denominator
	ret.WriteString(fmt.Sprintf("%d", intPart))

	numerator %= denominator
	if numerator != 0 {
		var fracPart bytes.Buffer
		idx := 0
		m := map[int]int{}
		ret.WriteString(".")
		for numerator != 0 {
			if idx, ok := m[numerator]; ok {
				str := fracPart.String()
				if idx != 0 {
					ret.WriteString(str[:idx])
				}
				ret.WriteString("(")
				ret.WriteString(str[idx:])
				ret.WriteString(")")
				return ret.String()
			}
			m[numerator] = idx
			idx++
			numerator *= 10
			t := numerator / denominator
			fracPart.WriteString(fmt.Sprintf("%d", t))
			numerator %= denominator
		}
		ret.WriteString(fracPart.String())
	}
	return ret.String()
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(2, 3))
	fmt.Println(fractionToDecimal(1000, 999))
}
