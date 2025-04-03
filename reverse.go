package algorithms

// Task from leetcode.com.
// LEVEL: MEDIUM

// Given a signed 32-bit integer x, return x with its digits reversed.
// If reversing x causes the value to go outside the signed 32-bit integer range [-2^31, 2^31 - 1], then return 0.
// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

func reverse(x int) int {
	var remainder, target int // выделяем 2 переменные

	if (-1<<31) <= x && x <= (1<<31-1) { // если число является 32 битным знаковым числом(по условиям задачи исключены 64-битные числа), то выполняем алгоритм

		for x%10 == 0 { // убираем лишние нули в конце
			x /= 10
		}

		for x != 0 { // реверсирование в 4 строчки кода
			remainder = x % 10             // получаем последнее число из X
			x = int(x / 10)                // убираем у X последнее число
			target = 10*target + remainder // умножаем каждый раз target на 10, тем самым сдвигая его на 1 шаг влево и прибавляем последнее число из X
		}
	}
	return target
}
