package main

import "fmt"

func main() {
	numList := make([]int, 10)
	book := make([]int, 10)
	i, total, sum := 0, 0, 0
	for numList[1] = 1; numList[1] < 10; numList[1]++ {
		for numList[2] = 1; numList[2] < 10; numList[2]++ {
			for numList[3] = 1; numList[3] < 10; numList[3]++ {
				for numList[4] = 1; numList[4] < 10; numList[4]++ {
					for numList[5] = 1; numList[5] < 10; numList[5]++ {
						for numList[6] = 1; numList[6] < 10; numList[6]++ {
							for numList[7] = 1; numList[7] < 10; numList[7]++ {
								for numList[8] = 1; numList[8] < 10; numList[8]++ {
									for numList[9] = 1; numList[9] < 10; numList[9]++ {

										for i = 1; i <= 9; i++ {
											book[i] = 0
										}

										for i = 1; i < 9; i++ {
											book[numList[i]] = 1
										}

										sum = 0

										for i = 1; i <= 9; i++ {
											sum += book[numList[i]]
										}

										if sum == 9 && numList[1]*100+numList[2]*10+numList[3]+numList[4]*100+numList[5]*10+numList[6] == numList[7]*100+numList[8]*10+numList[9] {
											total++
											fmt.Printf("%d%d%d+%d%d%d=%d%d%d\t", numList[1], numList[2], numList[3], numList[4], numList[5], numList[6], numList[7], numList[8], numList[9])
										}

									}
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("总可能性:", total)
}
