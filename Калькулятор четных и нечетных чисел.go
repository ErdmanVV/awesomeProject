package main

import "fmt"

func main() {
	// Создаем массив последовательных чисел от 1 до 40
	numbers := make([]int, 40)
	for i := 1; i <= 40; i++ {
		numbers[i-1] = i
	}

	// Инициализируем переменные для сумм четных и нечетных чисел
	evenSum := 0
	oddSum := 0

	// Суммируем четные и нечетные числа из массива
	for _, num := range numbers {
		if num%2 == 0 {
			evenSum += num
		} else {
			oddSum += num
		}
	}

	// Выводим результаты
	fmt.Println("Сумма четных чисел:", evenSum)
	fmt.Println("Сумма нечетных чисел:", oddSum)
}
