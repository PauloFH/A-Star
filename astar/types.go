package astar

type Item struct {
	City     string
	Priority int
	G        int
	Index    int
	Parent   *Item
}

type PriorityQueue []*Item

func (pq *PriorityQueue) Len() int { return len(*pq) }

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].Priority < (*pq)[j].Priority
}

func (pq *PriorityQueue) Swap(i, j int) {
	q := *pq
	q[i], q[j] = q[j], q[i]
	q[i].Index = i
	q[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.Index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}
