package main

import (
    "fmt"
    "math"
    "strconv"
)

// Функция для проверки корректности ввода
func isValidInput(input string) bool {
    if _, err := strconv.Atoi(input); err != nil {
        return false
    }
    return true
}

// Функция для выполнения арифметических операций
func calculate(operator string, operand1 int, operand2 int) int {
    switch operator {
    case "+":
        return operand1 + operand2
    case "-":
        return operand1 - operand2
    case "*":
        return operand1 * operand2
    case "/":
        return operand1 / operand2
    default:
        return 0 // Возвращаем ноль, если оператор неверен
    }
}

func main() {
    var input string
    var operand1 int
    var operand2 int
    var operator string

    for {
        fmt.Print("Введите число: ")
        fmt.Scanln(&input)

        if !isValidInput(input) {
            fmt.Println("Некорректный ввод. Попробуйте еще раз.")
            continue
        }

        operand1, _ = strconv.Atoi(input)

        fmt.Print("Введите операцию (+, -, *, /): ")
        fmt.Scanln(&operator)

        if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
            fmt.Println("Неверная операция. Попробуйте еще раз.")
            continue
        }

        fmt.Print("Введите второе число: ")
        fmt.Scanln(&input)

        if !isValidInput(input) {
            fmt.Println("Некорректный ввод. Попробуйте еще раз.")
            continue
        }

        operand2, _ = strconv.Atoi(input)

        result := calculate(operator, operand1, operand2)
        fmt.Println("Результат:", result)

        fmt.Print("Хотите продолжить? (y/n): ")
        var choice string
        fmt.Scanln(&choice)

        if choice == "n" || choice == "N" {
            break
        }
    }
}
