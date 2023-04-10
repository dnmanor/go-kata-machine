package main

func BubbleSort(arr []int) {
	count := len(arr)

	for i := 0; i < count; i++ {
		for j := 0; j < count-1-i; j++ {

			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
}
