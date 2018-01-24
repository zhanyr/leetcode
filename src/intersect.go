package main

//import "fmt"
//
//func main() {
//    num1 := []int{1,2,2,1,3,4, 5,6,3}
//    num2 := []int{2,2,4,3,4,2,3,6}
//    ret := intersect(num1, num2)
//    fmt.Println(ret)
//}
/**
350. Intersection of Two Arrays II

Given two arrays, write a function to compute their intersection.

Example:
Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2, 2].

Note:
Each element in the result should appear as many times as it shows in both arrays.
The result can be in any order.
Follow up:
What if the given array is already sorted? How would you optimize your algorithm?
What if nums1’s size is small compared to nums2’s size? Which algorithm is better?
What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?

 */

func intersect(nums1 []int, nums2 []int) []int {
    ret := make([]int, 0)
    //如果有一个为空，返回空数组
    if len(nums1) == 0 || len(nums2) == 0 {
        return ret
    }
    //为节省内存，找出长度小的num存map
    //map的结构是：key-num,value-num出现的次数
    if len(nums1) > len(nums2) {
        num := nums1
        nums1 = nums2
        nums2 = num
    }
    tmpNum := make(map[int]int, len(nums1))
    for _, num := range nums1 {
        tmpNum[num] = tmpNum[num] + 1
    }

    //遍历nums2 找出map中存在的value
    for _, num := range nums2 {
        if value, ok := tmpNum[num]; ok {
            if value != 0 {
                ret = append(ret, num)
                tmpNum[num] = value - 1
            }
        }
    }
    return ret
}
