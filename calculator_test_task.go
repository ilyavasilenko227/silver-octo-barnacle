package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calc(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
	}

	return result
}

func toarabicmap(a string) int {
	convert := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	return convert[a]
}

func torimmap(a int) string {
	convert := map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VII",
		9:   "IX",
		10:  "X",
		11:  "XI",
		12:  "XII",
		13:  "XIII",
		14:  "XIV",
		15:  "XV",
		16:  "XVI",
		17:  "XVII",
		18:  "XVIII",
		19:  "XIX",
		20:  "XX",
		21:  "XXI",
		24:  "XXIV",
		25:  "XXV",
		27:  "XXVII",
		28:  "XXVIII",
		30:  "XXX",
		32:  "XXXII",
		35:  "XXXV",
		36:  "XXXVI",
		40:  "XL",
		42:  "XLII",
		45:  "XLV",
		48:  "XLVIII",
		49:  "XLIX",
		50:  "L",
		54:  "LIV",
		56:  "LVI",
		60:  "LX",
		63:  "LXIII",
		64:  "LXVI",
		70:  "LXX",
		72:  "LXXII",
		80:  "LXXX",
		81:  "LXXXI",
		90:  "XC",
		100: "C"}
	return convert[a]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	operands := strings.Fields(expression)
	if len(operands) == 1 {
		fmt.Println("Строка не являестя математической операцией")
		return
	} else if len(operands) != 3 {
		fmt.Println("Формат математической операции не удовлетворяет заданию - два операнда и один оператор (+, -, /, *)")
		return
	}
	a := operands[0]
	b := operands[2]
	op := operands[1]

	num1, err1 := strconv.Atoi(a)
	num2, err2 := strconv.Atoi(b)

	if err1 == nil && err2 == nil {
		if num1 > 0 && num1 <= 10 && num2 > 0 && num2 <= 10 {
			result := calc(num1, num2, op)
			fmt.Println(result)
		} else {
			fmt.Println("Введите два числа от одного до десяти включительно")
		}
	} else if err1 != nil && err2 != nil {
		var rimA int = toarabicmap(a)
		var rimB int = toarabicmap(b)
		if rimA > 0 && rimA <= 10 && rimB > 0 && rimB <= 10 {
			result := calc(rimA, rimB, op)
			if result > 0 {
				println(torimmap(result))
			} else {
				println("В римской системе счисления нет отрицательных чисел")
			}
		} else {
			println("Введите два числа от одного до десяти включительно")
		}
	} else {
		println("Используются одновреммено разные системы счисления")
	}

}
