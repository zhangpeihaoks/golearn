package depth_first_search

import "fmt"

/*
*
  - void dfs(int step){
    判断边界
    尝试每种可能 for(i=1;i<=n;i++){
    继续下一步
    }
    返回
    }

*
*/
func dfs(step int, boxCount int, boxData []int, book []int) {
	if step == boxCount+1 {
		for i := 1; i <= boxCount; i++ {
			fmt.Printf("%d", boxData[i])
		}
		fmt.Println()
		return
	}

	for i := 1; i <= boxCount; i++ {
		if book[i] == 0 {
			boxData[step] = i
			book[i] = 1
			dfs(step+1, boxCount, boxData, book)
			book[i] = 0
		}
	}
	return
}
