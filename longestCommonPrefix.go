package algorithms

// Task from leetcode.com.
// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string ""

func longestCommonPrefix(strs []string) string {
	var prefix string												// работаем с одной переменной

	if len(strs[0]) > 0 {											// если первая строка не пустая, то выполняем следующий код
		prefix = strs[0]											// присваиваем переменной prefix первую строку среза strs

		for idx := 1; idx < len(strs); idx++ {						// начинаем обход строк strs, начиная со второй, idx хранит индекс строки в срезе strs
			for symbol := 0; symbol < len(prefix); symbol++ {		// проход по каждому символу только первой строки
				if len(strs[idx]) > symbol && strs[idx][symbol] == prefix[symbol] { 	// сравниваем symbol-ый символ первой строки с symbol-ым символом strs[idx] строки
					symbol++										// если равны то переходим к следующему символу
					continue					 					// и продолжаем цикл
				}
				prefix = prefix[:symbol]							// если не равны, то обновляем prefix на совпадающий префикс
				break;			 									// и переходим к вледующей строке 
			}
			if len(prefix) == 0 { break } 							// если общий префикс сравниваемых строк не найден, то выходим из цикла, т.к. дальнейший обход по strs не имеет смысла
		}
	}

	return prefix
}


// Вариант 2

func longestCommonPrefix_2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}