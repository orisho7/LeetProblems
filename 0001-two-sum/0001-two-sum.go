
func twoSum(nums []int, target int) []int {
	for index, value := range nums {
		wanted := target
		if index == len(nums)-1 {
			break
		}

		sum := value + nums[index+1]

		if sum == wanted {
			return []int{index, index + 1}
		} else {

			for j := index + 1; j < len(nums); j++ {
				if index == len(nums)-1 {
					break
				}
				sum := value + nums[j]
				if sum == wanted {
					return []int{index, j}
				}

			}
		}
	}
	return []int{}

}