package math

import (
	"fmt"
	"sort"
)

func Spaces(total, numsCount int) [][]int {
	sums := toParts(total, numsCount)
	indecies := makeCombinationIndecies(numsCount)
	res := [][]int{}
	for _, summ := range sums {
		res = append(res, applyCombinations(summ, indecies)...)
	}
	return sortArrs(res)
}

func toParts(total, numsCount int) [][]int {
	parts := make([]int, numsCount)
	for i := 2; i < numsCount; i++ {
		parts[i] = 1
		total--
	}
	parts[len(parts)-1] += total
	store := map[string][]int{arrToStr(parts): parts}
	branch(parts, store)
	res := mapToArr(store)
	return res
}

func branch(parts []int, store map[string][]int) {
	for i := len(parts) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if i-j > 1 && parts[i-1] <= parts[i]-1 && parts[j]+1 <= parts[j+1] ||
				i-j == 1 && parts[j] <= parts[i]-2 {
				newParts := copyArr(parts)
				newParts[i]--
				newParts[j]++
				key := arrToStr(newParts)
				_, ok := store[key]
				if !ok {
					store[key] = newParts
					branch(newParts, store)
				}
			}
		}
	}
}

func copyArr(temp []int) []int {
	newT := make([]int, len(temp))
	copy(newT, temp)
	return newT
}

func mapToArr(store map[string][]int) [][]int {
	res := make([][]int, 0, len(store))
	for _, ns := range store {
		res = append(res, ns)
	}
	return res
}

func sortArrs(res [][]int) [][]int {
	sort.Slice(res,
		func(i, j int) bool {
			for ind := range res[i] {
				if res[i][ind] < res[j][ind] {
					return true
				}
				if res[i][ind] > res[j][ind] {
					return false
				}
			}
			return false
		},
	)
	return res
}

func applyCombinations(arr []int, indecies [][]int) [][]int {
	store := map[string][]int{}
	for _, inds := range indecies {
		combination := make([]int, len(inds))
		flag := true
		for i, in := range inds {
			combination[i] = arr[in]
			// check for non-border zero spaces
			if i > 0 && i < len(inds)-1 && arr[in] == 0 {
				flag = false
				break
			}
		}
		if flag {
			store[arrToStr(combination)] = combination
		}
	}
	return mapToArr(store)
}

func arrToStr(arr []int) string {
	res := ""
	for _, n := range arr {
		res += fmt.Sprintf("%d_", n)
	}
	return res
}
