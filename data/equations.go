package data

import "fmt"

var temp_array = []int{10, 32, 44, 12, 92, 40, 19, 32, 5, 9, 6, 20, 17}

func Caller() {
	minmax(temp_array)
}

func minimum(arr []int) float32 {
	min := arr[0]
	for _, i := range arr {
		if (i < min) {
			min = i
		}
	}
	return float32(min)
}

func maximum(arr []int) float32{
	max := arr[0]
	for _, i := range arr {
		if (i > max) {
			max = i
		}
	}
	return float32(max)
}

func minmax(arr []int) {
	var min float32 = minimum(arr)
	var max float32= maximum(arr)
	n := []float32{} //stores minmax normalization data

	for i := 0; i < len(arr); i++ { 
		ni := (float32(arr[i])-min)/(max-min)
        n = append(n, ni)
	}

	fmt.Println(n, "<-- this is nrm data")

}


//normalize:
//for i in row[x]:
// normD = (i-min(row[x]) / (max(row[x]) - min(row[x]))
// LIST.append(normD)
// the func will return float32

//z-score norm:
// normD = (i - mean) / standard deviation

//