package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arithmetic_expression := bufio.NewScanner(os.Stdin)

	for arithmetic_expression.Scan() {
		text := arithmetic_expression.Text()
		if text != "" {
			num1, oper_type, num2 := parse(strings.ToUpper(text))
			is_arabic := check_num_type(num1, num2)

			if is_arabic {
				arabic_calculator(num1, oper_type, num2)
			} else {
				roman_calculator(num1, oper_type, num2)
			}
		}

	}

}

func parse(arithmetic_expression string) (string, string, string) {
	// Получаем тип операции и проверяем ее корректность
	addition := strings.Contains(arithmetic_expression, "+")
	if addition {
		result := strings.Split(arithmetic_expression, "+")
		if len(result) != 2 || result[0] == "" || result[1] == "" {
			panic("Cтрока не является математической операцией или не соответствуем формату - два операнда и один оператор (+, -, /, *)")
		}
		return strings.TrimSpace(result[0]), "+", strings.TrimSpace(result[1])
	}

	subtraction := strings.Contains(arithmetic_expression, "-")
	if subtraction {
		result := strings.Split(arithmetic_expression, "-")
		if len(result) != 2 || result[0] == "" || result[1] == "" {
			panic("Cтрока не является математической операцией или не соответствуем формату - два операнда и один оператор (+, -, /, *)")
		}
		return strings.TrimSpace(result[0]), "-", strings.TrimSpace(result[1])
	}

	multiplication := strings.Contains(arithmetic_expression, "*")
	if multiplication {
		result := strings.Split(arithmetic_expression, "*")
		if len(result) != 2 || result[0] == "" || result[1] == "" {
			panic("Cтрока не является математической операцией или не соответствуем формату - два операнда и один оператор (+, -, /, *)")
		}
		return strings.TrimSpace(result[0]), "*", strings.TrimSpace(result[1])
	}

	division := strings.Contains(arithmetic_expression, "/")
	if division {
		result := strings.Split(arithmetic_expression, "/")
		if len(result) != 2 || result[0] == "" || result[1] == "" {
			panic("Cтрока не является математической операцией или не соответствуем формату - два операнда и один оператор (+, -, /, *)")
		}
		return strings.TrimSpace(result[0]), "/", strings.TrimSpace(result[1])
	}

	panic("Cтрока не является математической операцией или не соответствуем формату - два операнда и один оператор (+, -, /, *)")

}

func check_num_type(num1 string, num2 string) bool {
	// Проверяем тип чисел
	_, num1_int_err := strconv.Atoi(num1)
	_, num2_int_err := strconv.Atoi(num2)

	if num1_int_err != nil && num2_int_err != nil {
		return false // Если римские
	} else if num1_int_err != nil || num2_int_err != nil {
		panic("Mixing of types") // Если разные типы
	}
	return true // Если арабские

}

func arabic_calculator(num1 string, oper_type string, num2 string) {
	// Калькулятор арабских

	num1_int, num1_int_err := strconv.Atoi(num1)
	num2_int, num2_int_err := strconv.Atoi(num2)

	if num1_int_err != nil || num2_int_err != nil {
		fmt.Println(num1_int_err, num2_int_err)
	}

	if num1_int < 1 || num1_int > 10 || num2_int < 1 || num2_int > 10 {
		panic("Число должно находиться в диапазоне от 1 до 10 (I...X)")
	}

	var arabic_result int

	switch oper_type {
	case "+":
		arabic_result = num1_int + num2_int

	case "-":
		arabic_result = num1_int - num2_int

	case "*":
		arabic_result = num1_int * num2_int

	case "/":
		arabic_result = num1_int / num2_int
	}
	fmt.Println(arabic_result)
}

func roman_calculator(num1 string, oper_type string, num2 string) {
	// Калькулятор римских

	num1_int := rom_map[num1]
	num2_int := rom_map[num2]

	if num1_int > 10 || num2_int > 10 {
		panic("Число должно находиться в диапазоне от 1 до 10 (I...X)")
	} else if num1_int == 0 || num2_int == 0 {
		panic("В римской системе нет одного из введенных чисел")
	}

	var arabic_result int

	switch oper_type {
	case "+":
		arabic_result = num1_int + num2_int
	case "-":
		arabic_result = num1_int - num2_int
	case "*":
		arabic_result = num1_int * num2_int
	case "/":
		arabic_result = num1_int / num2_int
	}
	roman_result := arabic_to_roman(arabic_result)
	fmt.Println(roman_result)
}

func arabic_to_roman(num int) string {
	// Конвертируем арабские в римские
	if num < 0 {
		panic("Результат работы меньше единицы")
	}
	var roman_result string
	for num > 0 {
		for key, value := range rom_map {
			if value == num {
				roman_result = key
				return roman_result
			}
		}

	}

	return roman_result
}
