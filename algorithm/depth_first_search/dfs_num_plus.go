package depth_first_search

import "fmt"

func NumPlusDfs(step int, numList []int, book []int, total int) int {
	if step == 10 {
		if numList[1] == 1 && numList[2] == 1 && numList[3] == 1 && numList[4] == 1 && numList[5] == 1 && numList[6] == 1 && numList[7] == 2 && numList[8] == 2 && numList[9] == 2 {
			fmt.Println(numList)
		}
		if numList[1]*100+numList[2]*10+numList[3]+numList[4]*100+numList[5]*10+numList[6] == numList[7]*100+numList[8]*10+numList[9] {
			total++
			fmt.Printf("%d%d%d+%d%d%d=%d%d%d\t", numList[1], numList[2], numList[3], numList[4], numList[5], numList[6], numList[7], numList[8], numList[9])
		}
		return total
	}

	for i := 1; i <= 9; i++ {
		if book[i] == 0 {
			numList[step] = i
			book[i] = 1

			total = NumPlusDfs(step+1, numList, book, total)

			book[i] = 0
		}
	}
	return total
}
