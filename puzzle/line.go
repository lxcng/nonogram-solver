package puzzle

import (
	"nonogram-solver/math"
	"sync"
)

type Line struct {
	nums        []int
	spaces      [][]int
	numSpaces   int
	totalSpaces int
	complete    bool
}

func NewLine(nums []int, ln int) *Line {
	numSpaces := len(nums) + 1
	totalNum := 0
	for _, s := range nums {
		totalNum += s
	}
	res := &Line{
		nums:        nums,
		numSpaces:   numSpaces,
		totalSpaces: ln - totalNum,
		complete:    false,
	}
	return res
}

func (l *Line) Spaces() {
	l.spaces = math.Spaces(l.totalSpaces, l.numSpaces)
}

func (l *Line) Traverse(dots []*Dot, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	nVars := 0
	newSpaces := [][]int{}
	for _, spaces := range l.spaces {
		if spaces == nil {
			continue
		}
		if l.traverse(dots, spaces) {
			nVars++
			newSpaces = append(newSpaces, spaces)
		}
	}
	l.spaces = newSpaces
	l.complete = true
	for _, d := range dots {
		if d.n == nVars {
			d.status = Filled
		}
		if d.n == -nVars {
			d.status = Empty
		}
		if d.status == Unknown {
			l.complete = false
		}
		d.n = 0
	}
}

func (l *Line) traverse(dots []*Dot, spaces []int) bool {
	spaceInd := 0
	numInd := 0
	mod := Empty
	for i := 0; i < len(dots); i++ {
		cap := 0
		switch mod {
		case Empty:
			cap = i + spaces[spaceInd]
		case Filled:
			cap = i + l.nums[numInd]
		}
		for j := i; j < cap; j++ {
			switch mod {
			case Empty:
				if dots[j].status == Filled {
					return false
				}
			case Filled:
				if dots[j].status == Empty {
					return false
				}
			}
		}
		switch mod {
		case Empty:
			spaceInd++
			mod = Filled
		case Filled:
			numInd++
			mod = Empty
		}
		i = cap - 1
	}
	spaceInd = 0
	numInd = 0
	mod = Empty
	for i := 0; i < len(dots); i++ {
		cap := 0
		switch mod {
		case Empty:
			cap = i + spaces[spaceInd]
		case Filled:
			cap = i + l.nums[numInd]
		}
		for j := i; j < cap; j++ {
			switch mod {
			case Empty:
				dots[j].n--
			case Filled:
				dots[j].n++
			}
		}
		switch mod {
		case Empty:
			spaceInd++
			mod = Filled
		case Filled:
			numInd++
			mod = Empty
		}
		i = cap - 1
	}
	return true
}

func (l *Line) totalSpace(ln int) int {
	f := 0
	for _, n := range l.nums {
		f += n
	}
	return ln - f
}
