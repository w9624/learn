package main

func KMP(s, sub string) bool {
	if len(s) == 0 || len(sub) == 0 {
		return false
	}

	next := func() []int {
		next := []int{-1}
		i, k := 0, -1
		for i < len(sub) {
			if k == -1 || sub[i:i+1] == sub[k:k+1] {
				i++
				k++
				next = append(next, k)
			} else {
				k = next[k]
			}
		}
		return next
	}()

	i, j := 0, 0
	for i < len(s) && j < len(sub) {
		if j == -1 || s[i:i+1] == sub[j:j+1] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == len(sub) {
		return true
	}

	return false
}
