package data

import (
	"fmt"
	"math"
)

var temp_array = []int{10, 32, 44, 12, 92, 40, 19, 32, 5, 9, 6, 20, 17}

func Caller() {
	minmax(temp_array)
	standardise(temp_array)
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

func maximum(arr []int) float32 {
	max := arr[0]
	for _, i := range arr {
		if (i > max) {
			max = i
		}
	}
	return float32(max)
}

func mean(arr []int) float32{
	t := 0
	for _, i := range arr {
		t += i
	}
	
	var mean float32 = float32(t)/float32((len(arr)))

	return mean //return float32(mean) work??
}

func standardDeviation(arr []int) float64 { 
	var standardDev float64 = 0
	var powNums float64 = 0
	var sum int = 0

	xm := float64(mean(arr))

	//subtract (x -x1)^2
	for _, v := range arr {
		xi := float64(v)
		powNums = math.Pow((xi-xm), 2)
	}
	for _, i := range arr {
		sum += i
	}

	standardDev = powNums/(float64(sum)-1)
	return standardDev
}

func minmax(arr []int) { //currently return array
	var min float32 = minimum(arr)
	var max float32= maximum(arr)
	n := []float32{} //stores minmax normalization data -> put in table

	for i := 0; i < len(arr); i++ { 
		ni := (float32(arr[i])-min)/(max-min)
        n = append(n, ni)
	}

	fmt.Println(n, "<-- this is mm data")
}


//normalize:
//for i in row[x]:
// normD = (i-min(row[x]) / (max(row[x]) - min(row[x]))
// LIST.append(normD)
// the func will return float32


func standardise(arr []int) { //currently return array
	mean := mean(arr)
	x := standardDeviation(arr)
	n := []float32{}

	for i := 0; i < len(arr); i++ { 
		s := (float32(arr[i])-mean)/float32(x)
        n = append(n, s)
	}
	
	fmt.Println(n, "<-- this is st data")
}


//build standard scalar function