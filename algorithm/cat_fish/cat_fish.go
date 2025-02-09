package main

import (
	"fmt"
	"golearn/algorithm/queue"
	"golearn/algorithm/stack"
	"math/rand"
	"time"
)

func CatFish(size int) {
	cards := make([]int, 0)
	var mid int
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		cards = append(cards, rand.Intn(100)+1)
	}
	book := make(map[int]int)
	if size%2 == 0 {
		mid = size / 2
	} else {
		mid = size/2 + 1
	}
	aCard := cards[0:mid]
	bCard := cards[mid:]
	fmt.Println(aCard, bCard)

	aQueue := queue.NewQueue[int](0)
	aQueue.Data = aCard
	aQueue.Tail = len(aCard) - 1

	bQueue := queue.NewQueue[int](0)
	bQueue.Data = bCard
	bQueue.Tail = len(bCard) - 1

	tableStack := stack.NewStack[int](0)

	for aQueue.Head < aQueue.Tail && bQueue.Head < bQueue.Tail {
		UserSendCard[int](aQueue, tableStack, book)
		if aQueue.Head == aQueue.Tail {
			fmt.Println("B win")
			fmt.Printf("B cards: %v\n", bQueue.Data[bQueue.Head:])
			if tableStack.Top > 0 {
				fmt.Printf("tableStack.data: %v,top:%v\n", tableStack.Data, tableStack.Top)
			}
			break
		}

		UserSendCard[int](bQueue, tableStack, book)
		if bQueue.Head == bQueue.Tail {
			fmt.Println("A win")
			fmt.Printf("A cards: %v\n", aQueue.Data[aQueue.Head:])
			if tableStack.Top > 0 {
				fmt.Printf("tableStack.data: %v,top:%v\n", tableStack.Data, tableStack.Top)
			}
			break
		}

	}

}

func UserSendCard[T int](userQueue *queue.Queue[T], tableStack *stack.Stack[T], book map[T]T) {
	flag := 0
	t := userQueue.Data[userQueue.Head]

	v, ok := book[t]

	if ok && v > 0 {
		flag = 1
	}

	switch flag {
	case 1:
		userQueue.Head++
		userQueue.Data = append(userQueue.Data, t)
		userQueue.Tail++
		for tableStack.Data[tableStack.Top-1] != t {
			userQueue.Data = append(userQueue.Data, tableStack.Data[tableStack.Top-1])
			userQueue.Tail++
			book[tableStack.Data[tableStack.Top-1]] -= 1
			tableStack.Top--
			tableStack.Data = tableStack.Data[:tableStack.Top]
		}
		userQueue.Data = append(userQueue.Data, tableStack.Data[tableStack.Top-1])
		userQueue.Tail++
		book[tableStack.Data[tableStack.Top-1]] -= 1
		tableStack.Top--
		tableStack.Data = tableStack.Data[:tableStack.Top]
	case 0:
		userQueue.Head++
		tableStack.Push(t)
		book[t] = 1
	}

}

func main() {
	CatFish(100)
}
