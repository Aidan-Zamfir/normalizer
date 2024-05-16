package data



var temp_arr = []int{10, 20, 31, 32, 19, 43, 71, 18, 2, 98, 93, 29, 61, 5}

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

//normalize:
//for i in row[x]:
// normD = (i-min(row[x]) / (max(row[x]) - min(row[x]))
// LIST.append(normD)
// the func will return float32

//z-score norm:
// normD = (i - mean) / standard deviation

//