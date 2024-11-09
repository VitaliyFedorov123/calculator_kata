package main

import (
	"errors"
	"fmt"
	"strings"
)

var romanToArabicMap = map[string]int{
	"I": 1,
	"II": 2,
	"III": 3,
	"IV": 4,
	"V": 5,
	"VI": 6,
	"VII": 7,
	"VIII": 8,
	"IX": 9,
	"X": 10,
}

var arabicToRomanMap = []struct {
	Value  int
	Symbol string
}{
	{10, "X"},
	{9, "IX"},
	{8, "VIII"},
	{7, "VII"},
	{6, "VI"},
	{5, "V"},
	{4, "IV"},
	{3, "III"},
	{2,"II"},
	{1, "I"},
}

func romanToArabic(roman string) (int, error) {
	roman = strings.ToUpper(roman) 
	total := 0
	prevValue := 0

	
	for i := len(roman) - 1; i >= 0; i-- {
		char := string(roman[i])
		value, exists := romanToArabicMap[char]
		if !exists {
			return 0, errors.New("неверое число")
		}

	
		if value < prevValue {
			total -= value
		} else {
			total += value
		}
		prevValue = value
	}

	return total, nil
}

func arabicToRoman(num int) (string, error) {
	if num <= 0 {
		return "", errors.New("неверно:число должно быть больше нуля")
	}

	result := ""
	for _, entry := range arabicToRomanMap {
		for num >= entry.Value {
			result += entry.Symbol
			num -= entry.Value
		}
	}
	return result, nil
}

func arabicCalculator(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль невозможно")
		}
		return a / b, nil
	default:
		return 0, errors.New("оператор недопустим")
	}
}

func romanCalculator(a, b, operator string) (string, error) {
	numA, err := romanToArabic(a)
	if err != nil {
		return "", err
	}

	numB, err := romanToArabic(b)
	if err != nil {
		return "", err
	}

	arabicResult, err := arabicCalculator(numA, numB, operator)
	if err != nil {
		return "", err
	}

	return arabicToRoman(arabicResult)
}

func main() {
	var mode, operator, input1, input2 string

	fmt.Println("Введите режим: 'arabic' или 'roman':")
	fmt.Scanln(&mode)

	if mode == "arabic" {
		var num1, num2 int
		fmt.Println("Введите первое число:")
		fmt.Scanln(&num1)
		fmt.Println("Введите оператор (+, -, *, /):")
		fmt.Scanln(&operator)
		fmt.Println("Введите второе число:")
		fmt.Scanln(&num2)

		result, err := arabicCalculator(num1, num2, operator)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		fmt.Printf("Результат: %d\n", result)

	} else if mode == "roman" {
		fmt.Println("Введите первое римское число:")
		fmt.Scanln(&input1)
		fmt.Println("Введите оператор (+, -, *, /):")
		fmt.Scanln(&operator)
		fmt.Println("Введите второе римское число:")
		fmt.Scanln(&input2)

		result, err := romanCalculator(input1, input2, operator)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		fmt.Printf("Результат в римских цифрах: %s\n", result)

	} else {
		fmt.Println("Неверный режим. Введите 'arabic' или 'roman'.")
	}
}
