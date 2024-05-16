package data

func Normalize() {

}

func Minimum(arr []int) int {
	min := arr[0]
	for _, i := range arr {
		if (i < min) {
			min = i
		}
	}
	return min
}

func Maximum(arr []int) int{
	max := arr[0]
	for _, i := range arr {
		if (i > max) {
			max = i
		}
	}
	return max
}