package algorithms

// Task from leetcode.com.
// LEVEL: MEDIUM

// Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.
// The integer division should truncate toward zero, which means losing its fractional part. For example, 8.345 would be truncated to 8, and -2.7335 would be truncated to -2.
// Return the quotient after dividing dividend by divisor.
// Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−2^31, 2^31 − 1]. 
// For this problem, if the quotient is strictly greater than 2^31 - 1, then return 2^31 - 1, and if the quotient is strictly less than -2^31, then return -2^31.

func Divide(dividend int, divisor int) int {
	if divisor == 0 { return 0 } 			// на ноль делить нельзя
	var quotient int 						// результат, он же частное
	var signed bool							// нужно для установки знака полученного результата

	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) { // если один из параметров отрицательный, а другой положительный,
		signed = true													// то результат будет отрицательным
	}

	if dividend < 0 {dividend = -dividend} 	// убираем знак для супер-упрощения цикла
	if divisor < 0 {divisor = -divisor}		// убираем знак для супер-упрощения цикла
    
    for dividend >= divisor {				// выполняем деление, пока делимое больше или равно делителю. в остальных случаях результат останется 0
        dividend = dividend - divisor	 	// вычитаем делитель из делимого и обновляем делимое
        quotient++							// увеличиваем результат на 1
    }
    
    if signed && quotient != 0 {			// меняем знак у результата, если это необходимо
    	quotient = -quotient
    }

    return quotient
}