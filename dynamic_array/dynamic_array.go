package dynamic_array

type DynamicArray struct {
	arr  []int
	size int
	cap  int
}

func NewDynamicArray(cap int) *DynamicArray {
	arr := make([]int, 0, cap)
	return &DynamicArray{arr: arr, size: 0, cap: cap}
}
