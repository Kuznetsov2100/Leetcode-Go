package problem0001

func TwoSum(nums []int, target int) []int {
    record := make(map[int]int)
    for index, value := range nums {
        complement := target - value
        if _, ok := record[complement]; ok {
            return []int{record[complement], index}
        } else {
            record[value] = index
        }
        
    }
    return []int{}
}