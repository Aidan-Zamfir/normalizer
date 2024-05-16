package data

<<<<<<< HEAD
var temp_array = []int{10, 32, 44, 12, 92, 40, 19, 32, 1, 9, 6, 20, 17}
=======


var temp_arr = []int{10, 20, 31, 32, 19, 43, 71, 18, 2, 98, 93, 29, 61, 5}

func Normalize() {
>>>>>>> dev


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

<<<<<<< HEAD

=======
>>>>>>> dev
//normalize:
//for i in row[x]:
// normD = (i-min(row[x]) / (max(row[x]) - min(row[x]))
// LIST.append(normD)
// the func will return float32

//z-score norm:
// normD = (i - mean) / standard deviation

//