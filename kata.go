package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateArabian(a, b int, sign string) int {
	switch sign {
	case "+":
		return a + b
	case "-":
		return a - b
	case "/":
		return a / b
	case "*":
		return a * b
	}
	return 0
}

func romanToInt(s string) int {
	know := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	lengthOfString := len(s)
	lastElement := s[len(s)-1 : lengthOfString]
	var result int
	result = know[lastElement]
	for i := len(s) - 1; i > 0; i-- {
		if know[s[i:i+1]] <= know[s[i-1:i]] {
			result += know[s[i-1:i]]
		} else {
			result -= know[s[i-1:i]]
		}
	}
	return result
}

func intToRoman(num int) string {
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}

	return roman
}

func isRoman(a string) int {
	if strings.ContainsAny(a, "IVXLCDM") {
		return 1
	} else {
		return -1
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)
	data := strings.Split(text, " ")
	if len(data) > 3 {
		fmt.Println("Error")
	} else {
		var a int
		var b int
		check := isRoman(data[0]) + isRoman(data[2])
		if check == -2 {
			a, _ = strconv.Atoi(data[0])
			b, _ = strconv.Atoi(data[2])
			fmt.Println(calculateArabian(a, b, data[1]))
		} else {
			if check == 2 {
				a = romanToInt(data[0])
				b = romanToInt(data[2])
				res := calculateArabian(a, b, data[1])
				if res < 0 {
					fmt.Println("Error")
				} else {
					fmt.Println(intToRoman(calculateArabian(a, b, data[1])))
				}
			} else {
				fmt.Println("Error")
			}
		}
	}

}
