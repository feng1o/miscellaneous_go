package main

import "fmt"

func Bubble_sort(arr []int)  {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr) - i; j++ {
			if arr[j] < arr[j+1] {
				Swap(arr, j+1, j)
			}
		}
	}
}

func Bubble_sort_2(arr []int) {
	flag := true
	for flag {
		flag = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] < arr[i+1] {
				Swap(arr, i, i+1)
				flag = true
			}
		}
	}
}

func Swap(arr []int, i int, j int)  {
	var tmp int
	tmp = arr[i]
	arr[i] = arr[j]
	arr[j] =tmp
}

func Insert_sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		index := i
		for j := i-1; j > 0 && arr[j] > tmp; j-- {
			arr[j+1] = arr[j]
			index = j
		}
		arr[index] = tmp
	}
}

func Select_sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		var tmp int
		index := 0
		tmp = arr[0]
		for j := 0; j < len(arr) - i; j++ {
			if arr[j] > tmp {
				index = j
				tmp = arr[index]
			}
		}
		Swap(arr, index, len(arr) - i - 1)
	}
}

func MergeSort(arr []int) {

}

func main()  {
	fmt.Println("----001.buuble sort-----")
	var arr []int
	arr = []int{1,2,25,2,6236,23,2}
	fmt.Println("unsort :", arr)
	//Bubble_sort_2(arr)
	Insert_sort(arr)
	fmt.Println("sort :", arr)

}
