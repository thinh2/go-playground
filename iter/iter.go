package iter

/*
 [[1,2,3],
  [4,5],
  [6],
  [],
  [7,8,9]]
*/
func convert2d(nums [][]int) []int {
	res := []int{}
	isEnd := true
	for col := 0; ; col++ {
		for row := 0; row < len(nums); row++ {
			if col < len(nums[row]) {
				res = append(res, nums[row][col])
				isEnd = false
			}
		}
		if isEnd {
			break
		}
		isEnd = true
	}
	return res
}
