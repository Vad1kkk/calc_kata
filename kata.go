package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calc(a, b int, c string) int {
	if c == "+" {
		res := a + b
		return res
	} else if c == "-" {
		res := a - b
		return res
	} else if c == "*" {
		res := a * b
		return res
	} else {
		res := a / b
		return res
	}
}

func main() {

	var o string
	var t, s, z int
	arab := "12345678910"
	roman := "IIIVIIIX"
	op := "+*-/"
	arabNum := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	arabRoman := []string{
		"0", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
		"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
		"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
		"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
		"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C",
	}
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)
	text = strings.ToUpper(text)
	ar := strings.Split(text, " ")
	n1 := []int{}
	n2 := []int{}
	for _, x := range ar {
		z++
		if strings.Contains(arab, x) {
			t++
			x, _ := strconv.Atoi(x)
			n1 = append(n1, x)
		} else if strings.Contains(roman, x) {
			s++
			x, _ := arabNum[x]
			n2 = append(n2, x)
		} else if strings.Contains(op, x) {
			o = x
		} else {
			fmt.Println("Выдача паники, так как число больше 10")
		}
	}
	if z > 3 {
		fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	if z < 3 {
		fmt.Println("Выдача паники, так как строка не является математической операцией.")
	}
	if t == 2 {
		fmt.Println(calc(n1[0], n1[1], o))
	} else if s == 2 {
		res := calc(n2[0], n2[1], o)
		if res < 0 {
			fmt.Println("Выдача паники, так как в римской системе нет отрицательных чисел.")
		} else if res == 0 {
			fmt.Println("Выдача паники, так как в римской системе нет ноля.")
		} else {
			fmt.Println(arabRoman[res])
		}
	} else if t == 1 && s == 1 {
		fmt.Println("Выдача паники, так как используются одновременно разные системы счисления.")
	}

}
