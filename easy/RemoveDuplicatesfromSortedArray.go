func removeDuplicates(nums []int) int {
	var expectedNums = make([]int, len(nums))
	currentNumber := -101
	expectedNumsLength := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != currentNumber {
			currentNumber = nums[i]
			expectedNums[expectedNumsLength] = currentNumber
			expectedNumsLength++
		}
	}
	return expectedNumsLength
}