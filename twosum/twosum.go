package twosum

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
func TwoSum(nums []int, target int) []int {
	out := []int{}
	visited := map[int]int{} // Space O(n)
	for i, j := range nums { // Time O(n)
		diff := target - j
		if v, ok := visited[diff]; ok {
			out = append(out, i)
			out = append(out, v)
			break
		}
		visited[j] = i
	}
	return out
}

func TwoSumBruteForce(nums []int, target int) []int {
    out := []int{}
    // range over num slice
    for i, j := range nums {
        for x, y := range nums {
            // check if output slice is less than 2
            if len(out) < 2{
                // check if same element
                if i == x {
                    continue
                }
                // check if sum equals target
                sum := j + y
                if sum == target {
                    out = append(out, i)
                    out = append(out, x)
                    break
                }
            }
        }
    }
    return out
}