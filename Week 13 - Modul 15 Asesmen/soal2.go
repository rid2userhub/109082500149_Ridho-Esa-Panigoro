package main

import "fmt"

type Pemain struct {
	firstName string
	lastName  string
	gol       int
	assist    int
}

func main() {
	var n int

	fmt.Println("Masukkan Data Input :")
	fmt.Scan(&n)

	var arr [1005]Pemain

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i].firstName, &arr[i].lastName, &arr[i].gol, &arr[i].assist)
	}

	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {

			if arr[j].gol > arr[maxIdx].gol {
				maxIdx = j
			} else if arr[j].gol == arr[maxIdx].gol {

				if arr[j].assist > arr[maxIdx].assist {
					maxIdx = j
				}
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}

	fmt.Println("\nHasil Sorting :")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %d %d\n", arr[i].firstName, arr[i].lastName, arr[i].gol, arr[i].assist)
	}
}
