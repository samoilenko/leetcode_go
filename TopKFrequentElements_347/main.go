package main

type Element struct {
	digit int
	times int
}

type MaxHeap struct {
	elements []*Element
}

func (h *MaxHeap) Add(item *Element) {
	h.elements = append(h.elements, item)
	if len(h.elements) == 1 {
		return
	}
	var parentIndex int
	queue := []int{len(h.elements) - 1}
	for len(queue) > 0 {
		itemIndex := queue[0]
		queue = queue[1:]
		parentIndex = (itemIndex - 1) / 2
		if h.elements[itemIndex].times > h.elements[parentIndex].times {
			queue = append(queue, parentIndex)
			tmp := h.elements[itemIndex]
			h.elements[itemIndex] = h.elements[parentIndex]
			h.elements[parentIndex] = tmp
		}
	}
}

func (h *MaxHeap) GetRoot() int {
	if len(h.elements) == 0 {
		return 0
	}
	res := h.elements[0]
	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	queue := []int{0}
	for len(queue) > 0 {
		itemIndex := queue[0]
		queue = queue[1:]
		leftChildIndex := (itemIndex * 2) + 1
		rightChildIndex := (itemIndex * 2) + 2

		var newIndex int
		if len(h.elements) > leftChildIndex && len(h.elements) <= rightChildIndex {
			newIndex = leftChildIndex
		} else if len(h.elements) > rightChildIndex && h.elements[leftChildIndex].times > h.elements[rightChildIndex].times {
			newIndex = leftChildIndex
		} else {
			newIndex = rightChildIndex
		}

		if (len(h.elements)) > newIndex && h.elements[newIndex].times > h.elements[itemIndex].times {
			tmp := h.elements[itemIndex]
			h.elements[itemIndex] = h.elements[newIndex]
			h.elements[newIndex] = tmp
			queue = append(queue, newIndex)
		}
	}

	return res.digit
}

func topKFrequent(nums []int, k int) []int {
	counters := make(map[int]*Element)
	for _, n := range nums {
		if _, ok := counters[n]; !ok {
			counters[n] = &Element{digit: n, times: 1}
		} else {
			counters[n].times += 1
		}
	}

	heap := &MaxHeap{}
	for _, el := range counters {
		heap.Add(el)
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.GetRoot()
	}

	return res
}
