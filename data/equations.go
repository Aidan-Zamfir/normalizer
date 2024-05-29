package data

import (
	"math"
)


func minimum(arr []float64) float64 { //return single value -> fl
	min := arr[0]
	for _, i := range arr {
		if (i < min) {
			min = i
		}
	}
	return min
}

func maximum(arr []float64) float64 { //return single value -> fl
	max := arr[0]
	for _, i := range arr {
		if (i > max) {
			max = i
		}
	}
	return max
}

func mean(arr []float64) float64{
	var t float64 = 0
	for _, i := range arr {
		t += i
	}
	
	var mean float64 = t/float64((len(arr)))

	return mean 
}

func standardDeviation(arr []float64) float64 { 
	var standardDev float64 = 0
	var powNums float64 = 0
	var sum float64 = 0

	xm := float64(mean(arr))

	//subtract (x -x1)^2
	for _, v := range arr {
		xi := float64(v)
		powNums = math.Pow((xi-xm), 2)
	}
	for _, i := range arr {
		sum += i
	}

	standardDev = powNums/(sum-1)
	return standardDev
}

func MinMax(arr []float64) []float64 { 
	min := minimum(arr)
	max := maximum(arr)
	n := []float64{} //stores minmax normalization data -> put in table (?)

	for i := 0; i < len(arr); i++ { 
		ni := (arr[i]-min)/(max-min)
        n = append(n, ni)
	}

	return n
}


func Standardise(arr []float64) []float64{ //currently return array
	mean := mean(arr)
	x := standardDeviation(arr)
	n := []float64{}

	for i := 0; i < len(arr); i++ { 
		s := (arr[i]-mean)/float64(x)
        n = append(n, s)
	}
	
	return n
}


//build standard scalar function