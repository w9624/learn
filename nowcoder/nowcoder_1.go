package main

import "log"

// 二维数组中的查找
// 技巧型问题：利用左-右，上-下递增，因此从右上角到左下角依此递减规律
func Find(target int, array [][]int) bool {

	if len(array) == 0 || len(array[0]) == 0 {
		return false
	}

	row, column := len(array)-1, len(array[0])-1
	for i := 0; i <= row; i++ {
		for j := column; j >= 0; j-- {
			tmp := array[i][j]

			if target == tmp {
				return true
			} else if target > tmp {
				j--
			} else {
				i++
			}
		}
	}

	return false
}

func main() {
	cases := []struct {
		target   int
		array    [][]int
		expected bool
	}{
		{
			target: 7,
			array: [][]int{
				{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15},
			},
			expected: true,
		},
		{
			target: 7,
			array: [][]int{
				{1},
			},
			expected: false,
		},
		{
			target:   7,
			array:    [][]int{},
			expected: false,
		},
	}

	for _, c := range cases {
		if c.expected != Find(c.target, c.array) {
			log.Fatalf("test fail, c: %#v", c)
		}
	}
}
