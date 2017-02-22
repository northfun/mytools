package tools

import "fmt"

type RQueueDataItfc interface {
	Equal(d RQueueDataItfc) bool
}

type RoundQueue struct {
	slc         []RQueueDataItfc
	front, tail int // index
}

func (l *RoundQueue) Init(size int) bool {
	if size <= 0 {
		return false
	}
	l.slc = make([]RQueueDataItfc, size+1)
	l.front = 0
	l.tail = 0
	return true
}

func (l *RoundQueue) Reset() {
	l.front = 0
	l.tail = 0
}

func (l *RoundQueue) Len() int    { return (l.tail + len(l.slc) - l.front) % len(l.slc) }
func (l *RoundQueue) Full() bool  { return (l.tail+1)%len(l.slc) == l.front }
func (l *RoundQueue) Empty() bool { return l.front == l.tail }

func (l *RoundQueue) Out() (RQueueDataItfc, bool) {
	if l.Empty() {
		return nil, false
	}
	v := l.slc[l.front]
	l.front = (l.front + 1) % len(l.slc)
	return v, true
}

func (l *RoundQueue) Push(v RQueueDataItfc) {
	if l.Full() {
		l.Out()
	}
	l.slc[l.tail] = v
	l.tail = (l.tail + 1) % len(l.slc)
}

func (l *RoundQueue) PushUnique(v RQueueDataItfc) {
	if d, ok := l.Top(); ok && d == v {
		return
	}
	l.Del(v)
	l.Push(v)
}

func (l *RoundQueue) Pop() (RQueueDataItfc, bool) {
	if l.Empty() {
		return nil, false
	}
	v := l.slc[l.tail]
	l.tail = (l.tail + len(l.slc) - 1) % len(l.slc)
	return v, true
}

func (l *RoundQueue) Top() (RQueueDataItfc, bool) {
	if l.Empty() {
		return nil, false
	}
	return l.slc[(l.tail+len(l.slc)-1)%len(l.slc)], true
}

func (l *RoundQueue) Del(v RQueueDataItfc) {
	if l.Empty() {
		return
	}
	var del bool
	for i := l.front; i != l.tail; i = (i + 1) % len(l.slc) {
		if del || l.slc[i].Equal(v) {
			l.slc[i] = l.slc[(i+1)%len(l.slc)]
			del = true
		}
	}
	if del {
		l.tail = (l.tail + len(l.slc) - 1) % len(l.slc)
	}
}

// 取第0,1,2,....值
func (l *RoundQueue) Get(i int) (RQueueDataItfc, bool) {
	if i < 0 || i >= l.Len() {
		return nil, false
	}
	return l.slc[(l.front+i)%len(l.slc)], true
}

func (l *RoundQueue) Print() {
	for i := l.front; i != l.tail; i = (i + 1) % len(l.slc) {
		fmt.Printf("-%v", l.slc[i])
	}
	fmt.Println("\n", l.front, ",", l.tail, ",len:", l.Len(), ",size:", len(l.slc))
}
