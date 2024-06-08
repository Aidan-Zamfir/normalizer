package data

import (
	"math"
	"time"
)

func MinMax(arr []float64, c chan []float64) {
	mn := minimum(arr)
	mx := maximum(arr)
	n := []float64{}

	for i := 0; i < len(arr); i++ {
		ni := (arr[i] - mn) / (mx - mn)
		n = append(n, ni)
	}

	c <- n
	time.Sleep(time.Millisecond * 10)
	close(c)
}

func Standardise(arr []float64, c chan []float64) {
	mean := mean(arr)
	x := standardDeviation(arr)
	n := []float64{}

	for i := 0; i < len(arr); i++ {
		s := (arr[i] - mean) / float64(x)
		n = append(n, s)
	}

	c <- n
	time.Sleep(time.Millisecond * 10)
	close(c)
}

func minimum(arr []float64) float64 { //return single value -> fl64
	min := arr[0]
	for _, i := range arr {
		if i < min {
			min = i
		}
	}
	return min
}

func maximum(arr []float64) float64 { //return single value -> fl64
	mmax := arr[0]
	for _, i := range arr {
		if i > mmax {
			mmax = i
		}
	}
	return mmax
}

func mean(arr []float64) float64 {
	var t float64 = 0
	for _, i := range arr {
		t += i
	}

	var mean = t / float64((len(arr)))

	return mean
}

func standardDeviation(arr []float64) float64 { //works
	var standardDev float64 = 0
	var powNums float64 = 0
	var sum = float64(len(arr))

	xm := float64(mean(arr))

	for _, v := range arr {
		xi := float64(v) //the current number in data set
		powNums += math.Pow((xi - xm), 2)
	}

	standardDev = math.Sqrt(powNums / (sum - 1))
	return standardDev
}
