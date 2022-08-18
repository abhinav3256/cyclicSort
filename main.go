package main

import "fmt"

func main() {

	arr := []int{4, 3, 1, 5, 2}
	result := sort(arr)
	fmt.Println(result)

}

func sort(arr []int) []int {

	i := 0
	for i < len(arr) {

		if arr[i]-1 == i {
			i++
		} else {
			swap(arr, i, arr[i]-1)
		}
	}
	return arr
}

func swap(arr []int, i, j int) []int {

	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp

	return arr

}

// func missingNumber(nums []int) int {

//     opt := 0
//     act := nums[0]
//     for i:= 1; i<= len(nums); i++ {
//         opt = opt ^ i
//     }
//     for i:= 1; i< len(nums); i++ {
//         act = act ^ nums[i]
//     }

//     fmt.Println(opt, act)
//     ans := opt ^ act
//     return ans
// }
