package sort

import "testing"

func TestBucketSort(t *testing.T) {
	arrs := []int{5,4,3,2,1,2,3,4,5,6,7,0,0,0,0}
	t.Log(BucketSort(arrs, 10))
}
