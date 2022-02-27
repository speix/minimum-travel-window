package main

import "fmt"

type Result struct {
	days  int
	path  []string
	start int
	stop  int
}

func main() {
	A := []string{"New York", "Paris", "New York", "Paris", "Tokyo", "Paris", "Athens", "Tokyo"}
	fmt.Println(solution(A).path)
}

func solution(A []string) *Result {
	places := make(map[string]int)

	for _, value := range A {
		if places[value] == 0 {
			places[value] = 1 // map[Athens:1 New York:1 Paris:1 Tokyo:1]
		}
	}

	remaining := len(places) // 4

	// Create a solution window: wStart->wStop to save the final solution
	start, wStart, wStop := 0, 0, 0
	for stop, value := range A {

		if places[value] > 0 {
			remaining -= 1
		}
		places[value] -= 1 // 1st iter: map[Athens:1 New York:0 Paris:1 Tokyo:1]

		if remaining == 0 { // If all places are visited at least once

			for start < stop && places[A[start]] < 0 { // While there are still duplicate records
				places[A[start]] += 1 // Increase the map counter for current city
				start += 1            // Adjust the start pointer to check the next city
			}

			if stop-start <= wStop-wStart || wStop == 0 { // Adjust the solution window
				wStart, wStop = start, stop

				if len(places) == wStop-wStart+1 { // Break early in case we have an early optimal solution
					return &Result{
						days:  wStop - wStart + 1,
						path:  A[wStart : wStop+1],
						start: wStart,
						stop:  wStop,
					}
				}
			}
		}
	}

	return &Result{
		days:  wStop - wStart + 1,
		path:  A[wStart : wStop+1],
		start: wStart,
		stop:  wStop,
	}
}
