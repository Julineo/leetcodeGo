package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	topN := 5

	ss := []int{134,1,43,54,4,23,5456,23,787,3,453,7,2,4,86,9,0}

	h := &IntHeap{}
	counter := 0
	for _, v := range ss {
		// Insert top 5 to the heap
		if counter <  topN {
			if counter == 0 {
				heap.Init(h)
				heap.Push(h, v)
				counter++
				continue
			}
			heap.Push(h, v)
			counter++
			continue
		}

		if v > (*h)[0] {
			(*h)[0] = v
			heap.Fix(h, 0)
		}
	}

	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
