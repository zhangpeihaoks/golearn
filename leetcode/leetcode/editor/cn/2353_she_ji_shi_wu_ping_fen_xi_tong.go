package leetcode

import (
	"sort"
)

//设计一个支持下述操作的食物评分系统：
//
//
// 修改 系统中列出的某种食物的评分。
// 返回系统中某一类烹饪方式下评分最高的食物。
//
//
// 实现 FoodRatings 类：
//
//
// FoodRatings(String[] foods, String[] cuisines, int[] ratings) 初始化系统。食物由
//foods、cuisines 和 ratings 描述，长度均为 n 。
//
//
//
// foods[i] 是第 i 种食物的名字。
// cuisines[i] 是第 i 种食物的烹饪方式。
// ratings[i] 是第 i 种食物的最初评分。
//
//
// void changeRating(String food, int newRating) 修改名字为 food 的食物的评分。
// String highestRated(String cuisine) 返回指定烹饪方式 cuisine 下评分最高的食物的名字。如果存在并列，返回 字典
//序较小 的名字。
//
//
// 注意，字符串 x 的字典序比字符串 y 更小的前提是：x 在字典中出现的位置在 y 之前，也就是说，要么 x 是 y 的前缀，或者在满足 x[i] !=
//y[i] 的第一个位置 i 处，x[i] 在字母表中出现的位置在 y[i] 之前。
//
//
//
// 示例：
//
// 输入
//["FoodRatings", "highestRated", "highestRated", "changeRating",
//"highestRated", "changeRating", "highestRated"]
//[[["kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"], ["korean",
//"japanese", "japanese", "greek", "japanese", "korean"], [9, 12, 8, 15, 14, 7]], [
//"korean"], ["japanese"], ["sushi", 16], ["japanese"], ["ramen", 16], ["japanese"]]
//输出
//[null, "kimchi", "ramen", null, "sushi", null, "ramen"]
//
//解释
//FoodRatings foodRatings = new FoodRatings(["kimchi", "miso", "sushi",
//"moussaka", "ramen", "bulgogi"], ["korean", "japanese", "japanese", "greek", "japanese",
// "korean"], [9, 12, 8, 15, 14, 7]);
//foodRatings.highestRated("korean"); // 返回 "kimchi"
//                                    // "kimchi" 是分数最高的韩式料理，评分为 9 。
//foodRatings.highestRated("japanese"); // 返回 "ramen"
//                                      // "ramen" 是分数最高的日式料理，评分为 14 。
//foodRatings.changeRating("sushi", 16); // "sushi" 现在评分变更为 16 。
//foodRatings.highestRated("japanese"); // 返回 "sushi"
//                                      // "sushi" 是分数最高的日式料理，评分为 16 。
//foodRatings.changeRating("ramen", 16); // "ramen" 现在评分变更为 16 。
//foodRatings.highestRated("japanese"); // 返回 "ramen"
//                                      // "sushi" 和 "ramen" 的评分都是 16 。
//                                      // 但是，"ramen" 的字典序比 "sushi" 更小。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 2 * 10⁴
// n == foods.length == cuisines.length == ratings.length
// 1 <= foods[i].length, cuisines[i].length <= 10
// foods[i]、cuisines[i] 由小写英文字母组成
// 1 <= ratings[i] <= 10⁸
// foods 中的所有字符串 互不相同
// 在对 changeRating 的所有调用中，food 是系统中食物的名字。
// 在对 highestRated 的所有调用中，cuisine 是系统中 至少一种 食物的烹饪方式。
// 最多调用 changeRating 和 highestRated 总计 2 * 10⁴ 次
//
//
// Related Topics 设计 数组 哈希表 字符串 有序集合 堆（优先队列） 👍 60 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
type FoodRatings struct {
	FoodsMap            map[string]*Food
	HighestRatedFoodMap map[string]*Food
	CuisineToFoods      map[string][]*Food
}

type Food struct {
	Name    string
	Cuisine string
	Rating  int
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	fr := FoodRatings{
		FoodsMap:            make(map[string]*Food),
		HighestRatedFoodMap: make(map[string]*Food),
		CuisineToFoods:      make(map[string][]*Food),
	}
	for i := range foods {
		food := &Food{
			Name:    foods[i],
			Cuisine: cuisines[i],
			Rating:  ratings[i],
		}
		fr.FoodsMap[foods[i]] = food
		fr.CuisineToFoods[cuisines[i]] = append(fr.CuisineToFoods[cuisines[i]], food)
	}

	for cuisine, cuisineFoods := range fr.CuisineToFoods {
		sort.Slice(cuisineFoods, func(i, j int) bool {
			if cuisineFoods[i].Rating != cuisineFoods[j].Rating {
				return cuisineFoods[i].Rating > cuisineFoods[j].Rating
			}
			return cuisineFoods[i].Name < cuisineFoods[j].Name
		})
		fr.HighestRatedFoodMap[cuisine] = cuisineFoods[0]
	}

	return fr
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	currentFood := this.FoodsMap[food]
	currentFood.Rating = newRating
	cuisine := currentFood.Cuisine
	currentHighest := this.HighestRatedFoodMap[cuisine]

	if currentHighest.Name == currentFood.Name {
		var newHighest *Food
		for _, f := range this.CuisineToFoods[cuisine] {
			if newHighest == nil {
				newHighest = f
				continue
			}
			if f.Rating > newHighest.Rating || (f.Rating == newHighest.Rating && f.Name < newHighest.Name) {
				newHighest = f
			}
		}
		this.HighestRatedFoodMap[cuisine] = newHighest
	} else {
		if newRating > currentHighest.Rating || (newRating == currentHighest.Rating && currentFood.Name < currentHighest.Name) {
			this.HighestRatedFoodMap[cuisine] = currentFood
		}
	}
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	return this.HighestRatedFoodMap[cuisine].Name
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
//leetcode submit region end(Prohibit modification and deletion)
