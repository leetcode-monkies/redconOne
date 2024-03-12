package problem0001

func TwoSum(nums []int, target int) []int {
	dict := make(map[int]int, 0)

	for idx, num := range nums {
		if pos, ok := dict[target-num]; ok {
			return []int{pos, idx}
		}
		dict[num] = idx
	}

	return []int{}
}
