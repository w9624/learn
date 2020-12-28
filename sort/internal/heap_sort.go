package main

import "log"

func HeadSort(array []int) {
	build(array)
	for i := len(array) - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		rebuild(array, 0, i-1)
	}
}

func build(array []int) {
	for i := len(array)/2 - 1; i >= 0; i-- {
		rebuild(array, i, len(array)-2)
	}
}

func rebuild(array []int, head, tail int) {
	root := array[head]
	i := head
	j := 2*i + 1

	for j <= tail {
		if j < tail && array[j] < array[j+1] {
			j++
		}

		if root > array[j] {
			break
		}

		array[i] = array[j]
		i = j
		j = 2*i + 1
	}

	array[i] = root
}

func main() {
	array := []int{3, 2, 2, 3, 9, 4, 6, 5, 1, 4}
	HeadSort(array)
	log.Println(array)
}
