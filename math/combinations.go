package math

var (
	combinationIndecies map[int][][]int = map[int][][]int{}
)

func makeCombinationIndecies(n int) [][]int {
	res, ok := combinationIndecies[n]
	if ok {
		return res
	}
	seed := make([]int, n)
	for i := range seed {
		seed[i] = i
	}
	res = makeCombination([]int{}, seed)
	combinationIndecies[n] = res
	return res
}

func makeCombination(prev []int, seed []int) [][]int {
	res := [][]int{}
	complete := true
	for i := range seed {
		if seed[i] >= 0 {
			complete = false
			newPrev := copyArr(prev)
			newPrev = append(newPrev, i)
			newSeed := copyArr(seed)
			newSeed[i] = -1
			res = append(res, makeCombination(newPrev, newSeed)...)
		}
	}
	if complete {
		res = append(res, prev)
	}
	return res
}
