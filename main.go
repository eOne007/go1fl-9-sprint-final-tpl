package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{} // returns an empty slice if size <= 0.
	}
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1000)
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {

	if len(data) == 0 {
		return 0
	}

	if len(data) == 1 {
		return data[0]
	}

	maxValue := data[0]
	for _, v := range data[1:] {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) < CHUNKS {
		return maximum(data)
	}
	allMaxes := make([]int, CHUNKS) // create new slice for all maxValues.
	var wg sync.WaitGroup
	chunkSize := (len(data) + CHUNKS - 1) / CHUNKS // for partitioning all elements of a slice into new slices.

	for i := 0; i < CHUNKS; i++ {

		index := i
		startOfSlice := i * chunkSize
		endOfSlice := min(startOfSlice+chunkSize, len(data))
		newSlice := data[startOfSlice:endOfSlice]

		wg.Add(1)
		go func() {
			defer wg.Done()
			allMaxes[index] = maximum(newSlice)
		}()
	}

	wg.Wait()
	return maximum(allMaxes)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
