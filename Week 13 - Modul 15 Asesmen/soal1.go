package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func SelectionSort(T *arrInt, n int) {

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if T[j] < T[minIdx] {
				minIdx = j
			}
		}
		// Tukar data
		T[i], T[minIdx] = T[minIdx], T[i]
	}
}

func median(T arrInt, n int) float64 {

	if n == 0 {
		return 0
	}
	if n%2 != 0 {
		return float64(T[n/2])
	} else {
		return float64(T[n/2-1]+T[n/2]) / 2.0
	}
}

func main() {
	var A arrInt
	var x int
	var n int = 0

	fmt.Println("Input data masukan :")

	fmt.Scan(&x)
	for x != -5313541 && n < NMAX {
		if x == 0 {
			SelectionSort(&A, n)
			med := median(A, n)

			fmt.Println("Median :")
			fmt.Printf("%g\n", med)
		} else {
			A[n] = x
			n++
		}
		fmt.Scan(&x)
	}
}
