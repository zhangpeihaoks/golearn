package queue

import (
	"fmt"
	"log/slog"
	"math/rand"
	"testing"
	"time"
)

func main() {

	rand.New(rand.NewSource(time.Now().UnixNano()))
	queue := NewQueue[int](100)
	for i := 0; i < 10; i++ {
		queue.Data[queue.Tail] = rand.Intn(10)
		queue.Tail++
	}

	slog.Info("queue", "data", queue.Data)

	for queue.Head < queue.Tail {
		// 打印队首,并队首出列
		fmt.Println(queue.Data[queue.Head])
		queue.Head++

		// 先将新队首的数添加到队尾
		queue.Data[queue.Tail] = queue.Data[queue.Head]
		queue.Tail++
		queue.Head++

	}
	fmt.Println(queue.Data)
}

func Test(t *testing.T) {
	main()
}
