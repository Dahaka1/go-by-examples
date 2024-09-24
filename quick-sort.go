package main

import (
	"fmt"
	"golang.org/x/exp/rand"
	"time"
)

func sort(arr []int, start, end int) {
	if start >= end {
		return
	}
	pivot := arr[(start+end)/2]

	i, j := start, end

	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i <= j {
			tmp := arr[j]
			arr[j] = arr[i]
			arr[i] = tmp
			i++
			j--
		}
	}
	sort(arr, start, j)
	sort(arr, i, end)
}

func main() {
	arr := []int{}
	n := 1_000_000
	for i := 0; i < n; i++ {
		rand.Seed(uint64(time.Now().Nanosecond()) << 5)
		arr = append(arr, rand.Int())
	}
	x := 10
	times := []time.Duration{}
	for i := 0; i < x; i++ {
		t := time.Now()
		sort(arr, 0, len(arr)-1)
		times = append(times, time.Since(t))
	}
	var max time.Duration
	var min = times[0]
	for i := 0; i < len(times); i++ {
		if times[i] < min {
			min = times[i]
		}
		if times[i] > max {
			max = times[i]
		}
	}
	fmt.Printf("Minimum sorting time: %s\nMaximum sorting time: %s\n", min, max)
}
