package puzzle

import (
	"image"
	"image/color"
	"image/png"
	"nonogram-solver/parser"
	"os"
	"strings"
	"sync"
)

type Puzzle struct {
	name          string
	rows          [][]*Dot
	cols          [][]*Dot
	rowNumbers    []*Line
	columnNumbers []*Line
	w, h          int
}

type stat struct {
}

func NewPuzzleFromConfig(path string) (*Puzzle, error) {
	c, err := parser.Parse(path)
	if err != nil {
		return nil, err
	}
	return NewPuzzle(path, c.W, c.H, c.RowNumbers, c.ColumnNumbers), nil
}

func NewPuzzle(name string, w, h int, rowNumbers, columnNumbers [][]int) *Puzzle {
	rows := make([][]*Dot, h, h)
	for i := range rows {
		rows[i] = make([]*Dot, w, w)
	}
	cols := make([][]*Dot, w, w)
	for i := range cols {
		cols[i] = make([]*Dot, h, h)
	}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			d := NewDot()
			rows[y][x] = d
			cols[x][y] = d
		}
	}

	return &Puzzle{
		name:          trimName(name),
		rows:          rows,
		cols:          cols,
		rowNumbers:    convertLines(rowNumbers, h),
		columnNumbers: convertLines(columnNumbers, w),
		w:             w,
		h:             h,
	}
}

func (p *Puzzle) Solve() {
	i := 0
	for {
		r := p.traverseRows()
		c := p.traverseCols()
		i++
		if r && c {
			p.saveBmp(p.name)
			break
		}
	}
}

func (p *Puzzle) traverseRows() bool {
	var wg sync.WaitGroup
	complete := true
	for i := range p.rows {
		if !p.rowNumbers[i].complete {
			wg.Add(1)
			complete = false
			go p.rowNumbers[i].Traverse(p.rows[i], &wg)
		}
	}
	wg.Wait()
	return complete
}

func (p *Puzzle) traverseCols() bool {
	var wg sync.WaitGroup
	complete := true
	for i := range p.cols {
		if !p.columnNumbers[i].complete {
			wg.Add(1)
			complete = false
			go p.columnNumbers[i].Traverse(p.cols[i], &wg)
		}
	}
	wg.Wait()
	return complete
}

func convertLines(lines [][]int, ln int) []*Line {
	res := make([]*Line, 0, len(lines))
	for _, l := range lines {
		line := NewLine(l, ln)
		line.Spaces()
		res = append(res, line)
	}
	return res
}

func (p *Puzzle) saveBmp(name string) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{p.w, p.h}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for x := 0; x < p.w; x++ {
		for y := 0; y < p.h; y++ {
			switch p.cols[x][y].status {
			case Filled:
				img.Set(x, y, color.Black)
			case Empty:
				img.Set(x, y, color.White)
			case Unknown:
				img.Set(x, y, color.Gray16{0x4444})
			}
		}
	}
	f, _ := os.Create(name + ".png")
	png.Encode(f, img)
}

func trimName(name string) string {
	ind := strings.LastIndex(name, ".")
	if ind > 0 {
		name = name[:ind]
	}
	ind = strings.LastIndex(name, "/")
	if ind > 0 && ind+1 < len(name) {
		name = name[ind+1:]
	}
	ind = strings.LastIndex(name, "\\")
	if ind > 0 && ind+1 < len(name) {
		name = name[ind+1:]
	}
	return name
}
