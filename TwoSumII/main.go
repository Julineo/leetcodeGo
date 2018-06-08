package main

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int)
	for j, v := range numbers {
		i, ok := m[target - v]
		if ok {
			return []int{i+1, j+1}
		}
		m[v] = j
	}
	return []int{}
}
