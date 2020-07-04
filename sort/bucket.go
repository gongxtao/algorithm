package sort

import (
	"container/list"
	"fmt"
	"math"
)

func BucketSort(arrs []int, bsize int) []int {
	max, min := findMaxMin(arrs)
	irange := float64(max - min + 1) / float64(bsize)
	return mergeBucket(insertBucket(arrs, irange, min), bsize, len(arrs))
}

func findMaxMin(arrs []int) (int, int) {
	max, min := math.MinInt64, math.MaxInt64

	for _, el := range arrs {
		if max < el {
			max = el
		}

		if min > el {
			min = el
		}
	}

	return max, min
}

func insertBucket(arrs []int, irange float64, min int) map[int]*list.List {
	result := make(map[int]*list.List)

	fmt.Printf("irange: %f, min: %d\n", irange, min)
	for _, el := range arrs {
		index := int(math.Floor(float64(float64(el - min) / irange)))

		elist, ok := result[index]
		if !ok {
			elist = list.New()
			result[index] = elist
		}

		fmt.Printf("element: %d, bucket index: %d\n", el, index)
		insertList(elist, el)
	}
	return result
}

func insertList(elist *list.List, eli int) *list.List {
	if elist.Len() == 0 {
		elist.PushBack(eli)
	} else {
		efront := elist.Front()
		for ; efront != nil; efront = efront.Next() {
			ev := efront.Value.(int)
			if ev >= eli { break }
		}
		if efront != nil {
			elist.InsertBefore(eli, efront)
		} else {
			elist.PushBack(eli)
		}
	}
	return elist
}

func mergeBucket(bucket map[int]*list.List, bsize, size int) []int {
	result := make([]int, size)
	index := 0
	for i := 0; i < bsize; i ++ {
		elist, ok := bucket[i]
		if !ok {
			continue
		}

		fmt.Printf("bucket: %d, element: ", i)
		efront := elist.Front()
		for ; efront != nil ; efront = efront.Next() {
			result[index] = efront.Value.(int)
			fmt.Printf("%d ", result[index])
			index ++
		}
		fmt.Println()
	}
	return result
}