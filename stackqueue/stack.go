package stackqueue

type Stack []int64

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(num int64) {
	*s = append(*s, num)
}

func (s *Stack) Pop() int64 {
	if s.isEmpty() {
		return 0
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *Stack) Top() int64 {
	return (*s)[len(*s)-1]
}
