package breadth_first_search

import (
	"golearn/algorithm/queue"
	"math/rand"
	"time"
)

func BuildMap(row, col int, obstacleCount int) [][]int {
	tmpMap := make([][]int, row)
	for i := range tmpMap {
		tmpMap[i] = make([]int, col)
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < obstacleCount; i++ {
		for {
			obX := rand.Intn(len(tmpMap))
			obY := rand.Intn(len(tmpMap[obX]))
			if tmpMap[obX][obY] == 0 {
				tmpMap[obX][obY] = 1
				break
			}
		}
	}
	return tmpMap
}

func SetTarget(theMap [][]int) (int, int) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		targetX := rand.Intn(len(theMap))
		targetY := rand.Intn(len(theMap[targetX]))
		if theMap[targetX][targetY] == 0 {
			theMap[targetX][targetY] = 2
			return targetX, targetY
		}
	}
}

var next = [4][2]int{
	{0, 1}, {0, -1}, {1, 0}, {-1, 0},
}

type Note struct {
	X int
	Y int
	F int
	S int
}

func BfsSaveHa(theMap [][]int, book [][]int, startX, startY int) int {
	list := queue.NewQueue[*Note](0)

	list.Data = append(list.Data, &Note{
		X: startX,
		Y: startY,
		F: 0,
		S: 0,
	})

	book[startX][startY] = 1
	var flag int
	var tx, ty int
	for list.Head <= list.Tail {
		for i := 0; i < 4; i++ {
			tx = list.Data[list.Head].X + next[i][0]
			ty = list.Data[list.Head].Y + next[i][1]
			if tx < 0 || tx >= len(theMap) || ty < 0 || ty >= len(theMap[tx]) {
				continue
			}

			if theMap[tx][ty] != 1 && book[tx][ty] == 0 {
				book[tx][ty] = 1
				list.Data = append(list.Data, &Note{
					X: tx,
					Y: ty,
					F: list.Head,
					S: list.Data[list.Head].S + 1,
				})
				list.Tail++
			}

			if theMap[tx][ty] == 2 {
				flag = 1
				break
			}
		}
		if flag == 1 {
			break
		}
		list.Head++
	}
	//if flag == 0 {
	//	return -1
	//}
	return list.Data[list.Tail].S
}
