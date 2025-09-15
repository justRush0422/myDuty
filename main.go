package main

import (
	"sort"
	"strconv"
)

// 返回数组只出现一次的数字
func onlyOneNum(param []int) []int {
	ele := make(map[int]int)
	res := []int{}
	for i := range param {
		var value = param[i]
		if len(ele) == 0 {
			ele[value] = 1
		} else {
			for mk := range ele {
				valueExist, exist := ele[value]
				if !exist {
					ele[value] = 1
					break
				} else if exist && mk == value {
					ele[value] = valueExist + 1
				}
			}
		}
	}
	var t int = 0
	for k, v := range ele {
		if v == 1 {
			res = append(res, k)
		}
		t++
	}
	return res
}

// 判断是否为回文数
func isPalindrome(x int) bool {
	var str string = strconv.Itoa(x)
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return str == string(r)
}

// 判断是否为有效括号
func isValid(s string) bool {
	var bl bool = false
	var prepare []rune = []rune(s)
	if len(prepare)%2 != 0 {
		return bl
	}
	if len(prepare) == 0 {
		bl = true
		return bl
	}
	var leftC []rune = []rune{}
	for s := range prepare {
		tar := string(prepare[s])
		if tar == "{" || tar == "[" || tar == "(" {
			leftC = append(leftC, prepare[s])
		}
		if tar == "}" && len(leftC) != 0 {
			if string(leftC[len(leftC)-1]) != "{" {
				return bl
			}
			leftC = append(leftC[0 : len(leftC)-1])
		}
		if tar == "]" && len(leftC) != 0 {
			if string(leftC[len(leftC)-1]) != "[" {
				return bl
			}
			leftC = append(leftC[0 : len(leftC)-1])
		}
		if tar == ")" && len(leftC) != 0 {
			if string(leftC[len(leftC)-1]) != "(" {
				return bl
			}
			leftC = append(leftC[0 : len(leftC)-1])
		}
	}
	if len(leftC) != 0 {
		return bl
	}
	bl = true
	return bl
}

// 返回最长公共前缀
func longestCommonPrefix(strs []string) string {
	var pre []rune = []rune{}
	var shortestWord string
	var wordLen int
	var index int
	for i := range strs {
		temp := strs[i]
		if i == 0 || wordLen > len(temp) {
			wordLen = len(temp)
			shortestWord = strs[i]
			index = i
		}
	}

	strs = append(strs[:index], strs[index+1:]...)

	for out := range strs {
		tem := strs[out]
		var tempRu []rune = []rune{}
		for i := 0; i < wordLen; i++ {
			if shortestWord[i] != tem[i] {
				pre = tempRu
				return string(pre)
			}
			tempC := shortestWord[i]
			var tempr []rune = []rune(string(tempC))
			tempRu = append(tempRu, tempr...)
		}
		if len(tempRu) < len(pre) || out == 0 {
			pre = tempRu
		}
	}
	return string(pre)
}

func plusOne(digits []int) []int {
	var res []int = []int{}
	num := 0
	for i := 0; i < len(digits); i++ {
		roll := len(digits) - 1 - i
		temp := digits[i]
		if len(digits)-1-i != 0 {
			for j := 0; j < roll; j++ {
				temp = temp * 10
			}
		}
		num = num + temp
	}
	num++
	numstr := strconv.Itoa(num)

	for i := range numstr {
		temp := string(numstr[i])
		numInt, error := strconv.ParseInt(temp, 10, 10)
		if error == nil {
			res = append(res, int(numInt))
		}
	}

	return res
}

// 切片去重复元素(使用了额外切片空间)
func removeDuplicates(nums []int) (int, []int) {
	var res []int = []int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := range nums {
		tar := nums[i]
		if len(res) == 0 {
			res = append(res, tar)
		} else {
			var flag bool = false
			for inner := range res {
				innerTar := res[inner]
				if tar == innerTar {
					flag = true
				}
			}
			if !flag {
				res = append(res, tar)
			}
		}

	}
	return len(res), res
}

// 合并区间
func merge(intervals [][]int) [][]int {
	for i := range intervals {
		if i+1 < len(intervals) {
			ele := intervals[i]
			ele2 := intervals[i+1]
			if ele[len(ele)-1] > ele2[0] {
				ele = append(ele[i:len(ele)-1], ele2[len(ele2)-1])
				intervals = append(intervals[i:i+1], intervals[i+2:]...)
				//递归
				intervals = merge(intervals)
			}
		} else {
			break
		}
	}
	return intervals
}

// 两数之和
func twoSum(nums []int, target int) []int {
	res := []int{}
	for i := range nums {
		if i+1 < len(nums) {
			if nums[i] >= target {
				continue
			}
			temp := target - nums[i]
			for j := range nums {
				if nums[j] >= target {
					continue
				}
				if temp == nums[j] {
					res = append(res, i)
					res = append(res, j)
				}
			}
			if len(res) != 0 {
				break
			}
		} else {
			break
		}
	}
	return res
}

func main() {
	//numInt := []int{1, 2, 3, 4, 3, 4, 2, 5, 6}
	//var res []int = onlyOneNum(numInt)
	//for i := range res {
	//fmt.Println("只出现过一次的数有", res[i])
	//}
	//--------------------------------------------------
	//f := isPalindrome(2442)
	//if f {
	//fmt.Print("是")
	//} else {
	//fmt.Print("不是")
	//}
	//--------------------------------------------------
	//res := isValid("(]")
	//if res {
	//fmt.Print("有效")
	//} else {
	//fmt.Print("无效")
	//}
	//--------------------------------------------------
	//var str = []string{"dante", "dan", "dantexn"}
	//fmt.Print(longestCommonPrefix(str))
	//--------------------------------------------------
	//var param []int = []int{4, 9, 9, 9, 9}
	//var fin []int = plusOne(param)
	//for i := range fin {
	//	fmt.Println(fin[i])
	//}
	//--------------------------------------------------
	//var param []int = []int{10, 5, 3, 0, 1, 1, 2, 3, 4, 4, 7, 8, 8, 8, 9}
	//length, resNum := removeDuplicates(param)
	//fmt.Println("长度", length)
	//for i := range resNum {
	//	fmt.Println("元素有：", resNum[i])
	//}
	//--------------------------------------------------
	//var param = [][]int{{1, 3}, {2, 6}, {8, 10}, {11, 18}}
	//res := merge(param)
	//for i := range res {
	//	fmt.Println(res[i])
	//}
	//-----------------------------------------------------
	//var param = []int{2, 7, 11, 15}
	//res := twoSum(param, 9)
	//for i := range res {
	//	fmt.Println(res[i])
	//}
}
