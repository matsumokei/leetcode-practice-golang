package leetcode

// https://leetcode.com/problems/two-sum/description/

/*
TODO: paste description from leetcode
*/

func twoSum(numbers []int, target int) []int {
    m := make(map[int]int)
    for k, v := range numbers {
        if idx, ok := m[target-v]; ok {
            return []int{idx, k}
        }
        m[v] = k
    }
    return nil
}