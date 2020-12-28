package main

import (
	"log"
	"math"
)

// 把字符串转成整数
// 串 -> char&askll
func StrToInt(str string) int {
	if len(str) == 0 {
		return 0
	}

	op := 1
	integer := 0
	chars := []rune(str)
	for i := 0; i < len(chars); i++ {
		char := chars[i]
		if char >= '0' && char <= '9' {
			// 判断数据溢出
			if op == 1 && int(char-'0') > math.MaxInt32-integer {
				return 0
			} else if op == -1 && int(char-'0') < integer-math.MinInt32 {
				return 0
			}
			integer = integer*10 + int(char-'0')
		} else if char == '-' && i == 0 {
			op = -1
		} else if char == '+' && i == 0 {
		} else {
			return 0
		}

		//log.Println(integer)
	}

	return op * integer
}

func main() {
	log.Println(math.MinInt32)
	log.Println(StrToInt("+214748364700000"))
	log.Println(StrToInt("1a33"))
}
