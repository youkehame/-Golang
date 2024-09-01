package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Карты для преобразования римских чисел в арабские и обратно
var romanToArabicMap = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicToRomanMap = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	20: "XX", 30: "XXX", 40: "XL", 50: "L",
	60: "LX", 70: "LXX", 80: "LXXX", 90: "XC", 100: "C",
}

// Функция для преобразования римских чисел в арабские
func romanToArabic(roman string) (int, error) {
	if val, exists := romanToArabicMap[roman]; exists {
		return val, nil
	}
	return 0, errors.New("некорректное римское число")
}

// Функция для преобразования арабских чисел в римские
func arabicToRoman(number int) (string, error) {
	if number < 1 {
		return "", errors.New("результат меньше I в римских числах")
	}
	result := ""
	for _, value := range []int{100, 90, 50, 40, 10, 9, 5, 4, 1} {
		for number >= value {
			result += arabicToRomanMap[value]
			number -= value
		}
	}
	return result, nil
}

// Основная функция вычисления
func calculate(expression string) (string, error) {
	// Добавляем пробелы вокруг операторов
	expression = strings.ReplaceAll(expression, "+", " + ")
	expression = strings.ReplaceAll(expression, "-", " - ")
	expression = strings.ReplaceAll(expression, "*", " * ")
	expression = strings.ReplaceAll(expression, "/", " / ")

	parts := strings.Fields(expression)
	if len(parts) != 3 {
		return "", errors.New("некорректный формат выражения")
	}

	num1Str, operator, num2Str := parts[0], parts[1], parts[2]

	var num1, num2 int
	var isRoman bool

	// Проверка и преобразование первого числа
	if n, err := strconv.Atoi(num1Str); err == nil {
		num1 = n
	} else if n, err := romanToArabic(num1Str); err == nil {
		num1 = n
		isRoman = true
	} else {
		return "", errors.New("некорректная система счисления")
	}

	// Проверка и преобразование второго числа
	if n, err := strconv.Atoi(num2Str); err == nil {
		num2 = n
	} else if n, err := romanToArabic(num2Str); err == nil {
		num2 = n
		isRoman = true
	} else {
		return "", errors.New("некорректная система счисления")
	}

	// Проверка диапазона чисел
	if (num1 < 1 || num1 > 10) || (num2 < 1 || num2 > 10) {
		return "", errors.New("числа должны быть от 1 до 10")
	}

	// Выполнение операции
	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return "", errors.New("деление на ноль")
		}
		result = num1 / num2
	default:
		return "", errors.New("некорректный оператор")
	}

	// Обработка результата для римских чисел
	if isRoman {
		return arabicToRoman(result)
	}

	return strconv.Itoa(result), nil
}

func main() {
	for {
		fmt.Print("Введите выражение: ")
		var expression string
		fmt.Scanln(&expression)

		if expression == "exit" {
			break
		}

		result, err := calculate(expression)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Результат:", result)
		}
	}
}
