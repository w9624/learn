package main

func main() {

}

func BF(s, sub string) bool {
	if len(s) == 0 || len(sub) == 0 {
		return false
	}

	i, j := 0, 0
	for i < len(s) && j < len(sub) {
		if s[i:i+1] == sub[j:j+1] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}

	if j == len(sub) {
		return true
	}

	return false
}

func RK(s, sub string) bool {

	return false
}
