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
//old logic
func twoSumloop(nums []int, target int) []int {
	for index, value := range nums {
		wanted := target
		if index == len(nums)-1 {
			break
		}

		sum := value + nums[index+1]
		fmt.Println("outside", wanted, index, value, sum)

		if sum == wanted {
			return []int{index, index + 1}
		} else {

			for j := index + 1; j < len(nums); j++ {
				if index == len(nums)-1 {
					break
				}
				sum := value + nums[j]
				fmt.Println("inside loop", wanted, index, value, sum)
				if sum == wanted {
					return []int{index, j}
				}

			}
		}
	}
	return []int{}

}