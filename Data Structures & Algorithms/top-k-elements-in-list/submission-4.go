type Pair struct{
	Val int
	Freq int
}

type minHeap []Pair

func (h minHeap) parent(i int) int { return (i - 1) / 2 }
func (h minHeap) left(i int) int { return i * 2 + 1 }
func (h minHeap) right(i int) int { return i * 2 + 2 }


func (h minHeap) siftUp(i int) {
	for i > 0 {
		if h[h.parent(i)].Freq <= h[i].Freq {
			break
		}
		if h[h.parent(i)].Freq > h[i].Freq {
			h[i], h[h.parent(i)] = h[h.parent(i)], h[i]
			i = h.parent(i)
		}

	}
}

func (h minHeap) siftDown(i int) {
	n := len(h)
	for {
		smallest := i
		l ,r := h.left(i), h.right(i)
		if l < n && h[l].Freq < h[smallest].Freq {smallest  = l}
		if r < n && h[r].Freq < h[smallest].Freq {smallest = r}
		if smallest == i {break}
		h[i], h[smallest] = h[smallest], h[i]
		i = smallest
	}
}

func (h *minHeap) push(p Pair) {
	*h = append(*h, p)
	h.siftUp(len(*h)-1)
}


func (h *minHeap) pop() Pair {
	min := (*h)[0]
	last := len(*h) - 1
	(*h)[0] = (*h)[last]
	*h = (*h)[:last]
	if len(*h) > 1 {
		h.siftDown(0)
	}
	return min
}


func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, val := range nums {
		freq[val]++
	} 

	heap := minHeap{}

	for num, count := range freq {
		heap.push(Pair{
			Val : num,
			Freq : count,
		})

		if len(heap) > k {
			heap.pop()
		}
	}

	result := []int{}
	for len(heap) > 0 {
		result = append(result, heap.pop().Val)
	}
	return result
}
