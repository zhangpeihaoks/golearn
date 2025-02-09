package queue

type Queue[T any] struct {
	Head int // 队列头索引
	Tail int // 队列尾索引
	Data []T // 队列
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		Head: 0,
		Tail: 0,
		Data: make([]T, size),
	}
}
