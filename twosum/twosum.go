package twosum

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
