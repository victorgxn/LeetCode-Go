func getConcatenation(nums []int) []int {
    exerciseArray := append(nums, nums...) 
    return exerciseArray
}
