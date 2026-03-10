func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for index, value := range nums {
		needed := target - value
		val, ok := seen[needed]
		if ok {
			return []int{index, val}
		} else {
			seen[value] = index
		}
	}
	return []int{}
}