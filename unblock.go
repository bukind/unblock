// Binary unblock is a solver for Unblock game on Android.
package main

import (
	"fmt"
	"log"
	"strings"
)

const (
	COLS = 6
	ROWS = 6
)

var (
	// The initial state of the desk.
	start = Desk{
		/*
		0x44, 0x46, 0x00,
		0x06, 0x06, 0x66,
		0x06, 0x44, 0x66,
		0x44, 0x60, 0x66,
		0x00, 0x64, 0x40,
		0x44, 0x40, 0x00,
		*/
		0x66, 0x44, 0x46,
		0x66, 0x06, 0x66,
		0x64, 0x46, 0x60,
		0x44, 0x60, 0x60,
		0x00, 0x64, 0x40,
		0x04, 0x45, 0x50,
	}
)

// Hence each cell uses only 3 lowest bits:
// * 000 empty
// * 11x part of a vertical plank #x
// * 10x part of a horizontal plank #x
// Here x shows the ordinal number of a plank in its line
// of movement.  At most there could be two planks
// in the line of movement, hence x can be either 0 or 1.
type Cell byte

// IsEmpty returns true if the cell is not occupied.
func (v Cell) IsEmpty() bool {
	return v == 0
}

// IsVertical returns true if the cell is part of a vertical plank.
func (v Cell) IsVertical() bool {
	return (v & 2) != 0
}

func (v Cell) Ordinal() byte {
	return byte(v) & 1
}

// String returns string representation of the cell.
func (v Cell) String() string {
	if v.IsEmpty() {
		return " . "
	}
	if v.IsVertical() {
		return fmt.Sprintf("|%d|", v&1)
	}
	return fmt.Sprintf("=%d=", v&1)
}

// Pos is a position of a cell.
type Pos struct {
	Col int
	Row int
}

func (p Pos) String() string {
	return fmt.Sprintf("(%d,%d)", p.Col, p.Row)
}

// Desk is a state of the playing desk, with 6x6 cells.
// The cells in two adjacent cells are packed into a byte.
// Thus we have 3 bytes per row.
type Desk [COLS * ROWS / 2]byte

func (d *Desk) At(p Pos) Cell {
	u := d[p.Row*COLS/2+p.Col/2]
	// The cell in left column is packed in higher bits.
	// This is because it would be easier to code it
	// in the hex representation.
	if p.Col&1 == 0 {
		u = u >> 4
	}
	return Cell(u & 0x7)
}

func (d *Desk) Set(p Pos, c Cell) {
	i := p.Row*COLS/2 + p.Col/2
	if p.Col&1 == 0 {
		d[i] = d[i]&0x7 + byte(c)<<4
	} else {
		d[i] = d[i]&0x70 + byte(c)
	}
}

// Empties returns the list of empty positions in the desk.
func (d *Desk) Empties() []Pos {
	result := make([]Pos, 0, 16)
	for row := 0; row < ROWS; row++ {
		for col := 0; col < COLS; col++ {
			pos := Pos{col, row}
			if d.At(pos) == 0 {
				result = append(result, pos)
			}
		}
	}
	return result
}

func (d *Desk) IsSolved() bool {
	const solvingRow = 2
	for col := COLS - 1; col > 0; col-- {
		cell := d.At(Pos{col, solvingRow})
		if cell.IsEmpty() {
			continue
		}
		return !cell.IsVertical()
	}
	// Should never reach here.
	return true
}

// This tries to move a horizontal plank to the position pos.
// Obviously the plank should have the same row and be on the left.
func (d *Desk) moveHorizontalRight(pos Pos) (Desk, bool) {
	// Moving backward.
	// Note that we don't check the last position in the line,
	// because the planks are two space at least.
	col1 := 0
	var move Cell
	for col := pos.Col - 1; col > 0; col-- {
		cell := d.At(Pos{col, pos.Row})
		switch {
		case cell.IsEmpty():
			continue
		case cell.IsVertical():
			return *d, false
		}
		// We've found the plank's end.
		col1 = col
		move = cell
		break
	}
	if col1 == 0 {
		return *d, false
	}
	// Search for the second end.
	// We can skip one space since planks are 2 spaces at least.
	col2 := col1 - 1
	for col := col2 - 1; col >= 0; col-- {
		cell := d.At(Pos{col, pos.Row})
		if cell.IsEmpty() || cell.IsVertical() || cell.Ordinal() != move.Ordinal() {
			break
		}
		col2 = col
	}
	res := *d
	// Here we move the plank [col2,col1] into pos.col.
	for col := col1 + 1; col <= pos.Col; col++ {
		res.Set(Pos{col, pos.Row}, move)
	}
	for col := col2; col < col2+pos.Col-col1; col++ {
		res.Set(Pos{col, pos.Row}, Cell(0))
	}
	return res, true
}

// This tries to move a horizontal plank to the position pos.
// Obviously the plank should have the same row and be on the left.
func (d *Desk) moveHorizontalLeft(pos Pos) (Desk, bool) {
	// Moving backward.
	// Note that we don't check the last position in the line,
	// because the planks are two space at least.
	col1 := 0
	var move Cell
	for col := pos.Col + 1; col < COLS-1; col++ {
		cell := d.At(Pos{col, pos.Row})
		switch {
		case cell.IsEmpty():
			continue
		case cell.IsVertical():
			return *d, false
		}
		// We've found the plank's end.
		col1 = col
		move = cell
		break
	}
	if col1 == 0 {
		return *d, false
	}
	// Search for the second end.
	// We can skip one space since planks are 2 spaces at least.
	col2 := col1 + 1
	for col := col2 + 1; col <= COLS-1; col++ {
		cell := d.At(Pos{col, pos.Row})
		if cell.IsEmpty() || cell.IsVertical() || cell.Ordinal() != move.Ordinal() {
			break
		}
		col2 = col
	}
	res := *d
	// Here we move the plank [col1,col2] into pos.col.
	for col := col1 - 1; col >= pos.Col; col-- {
		res.Set(Pos{col, pos.Row}, move)
	}
	for col := col2; col > col2-(col1-pos.Col); col-- {
		res.Set(Pos{col, pos.Row}, Cell(0))
	}
	return res, true
}

// This tries to move a vertical plank to the position pos.
// Obviously the plank should have the same col and be above the pos.
func (d *Desk) moveVerticalDown(pos Pos) (Desk, bool) {
	// Moving backward.
	// Note that we don't check the last position in the line,
	// because the planks are two space at least.
	row1 := 0
	var move Cell
	for row := pos.Row - 1; row > 0; row-- {
		cell := d.At(Pos{pos.Col, row})
		switch {
		case cell.IsEmpty():
			continue
		case !cell.IsVertical():
			return *d, false
		}
		// We've found the plank's end.
		row1 = row
		move = cell
		break
	}
	if row1 == 0 {
		return *d, false
	}
	// Search for the second end.
	// We can skip one space since planks are 2 spaces at least.
	row2 := row1 - 1
	for row := row2 - 1; row >= 0; row-- {
		cell := d.At(Pos{pos.Col, row})
		if cell.IsEmpty() || !cell.IsVertical() || cell.Ordinal() != move.Ordinal() {
			break
		}
		row2 = row
	}
	res := *d
	// Here we move the plank [row2,row1] into pos.Row.
	for row := row1 + 1; row <= pos.Row; row++ {
		res.Set(Pos{pos.Col, row}, move)
	}
	for row := row2; row < row2+pos.Row-row1; row++ {
		res.Set(Pos{pos.Col, row}, Cell(0))
	}
	return res, true
}

// This tries to move a vertical plank to the position pos.
// Obviously the plank should have the same col and be below the pos.
func (d *Desk) moveVerticalUp(pos Pos) (Desk, bool) {
	// Moving backward.
	// Note that we don't check the last position in the line,
	// because the planks are two space at least.
	row1 := 0
	var move Cell
	for row := pos.Row + 1; row < ROWS-1; row++ {
		cell := d.At(Pos{pos.Col, row})
		switch {
		case cell.IsEmpty():
			continue
		case !cell.IsVertical():
			return *d, false
		}
		// We've found the plank's end.
		row1 = row
		move = cell
		break
	}
	if row1 == 0 {
		return *d, false
	}
	// Search for the second end.
	// We can skip one space since planks are 2 spaces at least.
	row2 := row1 + 1
	for row := row2 + 1; row <= ROWS-1; row++ {
		cell := d.At(Pos{pos.Col, row})
		if cell.IsEmpty() || !cell.IsVertical() || cell.Ordinal() != move.Ordinal() {
			break
		}
		row2 = row
	}
	// log.Printf("moving to %s plank @ [%d,%d]", pos, row1, row2)
	res := *d
	// Here we move the plank [row1,row2] into pos.Row.
	for row := row1 - 1; row >= pos.Row; row-- {
		res.Set(Pos{pos.Col, row}, move)
	}
	for row := row2; row > row2-(row1-pos.Row); row-- {
		res.Set(Pos{pos.Col, row}, Cell(0))
	}
	return res, true
}

// String returns string representation of the Desk.
func (d Desk) String() string {
	var sb strings.Builder
	for row := 0; row < ROWS; row++ {
		for col := 0; col < COLS; col++ {
			fmt.Fprintf(&sb, "%s", d.At(Pos{col, row}))
		}
		fmt.Fprintf(&sb, "\n")
	}
	return sb.String()
}

type Replay struct {
	steps int
	prev  Desk
}

func (r Replay) Prev() (Desk, bool) {
	if r.steps > 0 {
		return r.prev, true
	}
	return Desk{}, false
}

type Solver struct {
	states map[Desk]Replay
}

func NewSolver(start Desk) *Solver {
	return &Solver{
		states: map[Desk]Replay{
			start: Replay{steps: 0},
		},
	}
}

func (s *Solver) PrintChain(fin Desk) {
	if prev, ok := s.states[fin].Prev(); ok {
		s.PrintChain(prev)
	}
	fmt.Println(fin)
}

// Step makes one step in solving the puzzle.
// It takes all the states from the previous step and mutate them
// to produce states for the next step.  While doing so, it filters out
// all states that were already seen.
//
// It also checks if any of the produced steps is a solution to the puzzle.
// Returns:
//   - if the solution is not found it returns all the next states and false;
//   - if the solution is found, it returns just this state and true.
func (s *Solver) Step(step int, desks ...Desk) ([]Desk, bool) {
	log.Printf("step #%d input desks %d", step, len(desks))
	nextstep := step + 1
	var result []Desk
	for _, d := range desks {
		empties := d.Empties()
		for _, pos := range empties {
			for _, fn := range []func(Pos) (Desk, bool){
				d.moveHorizontalRight,
				d.moveHorizontalLeft,
				d.moveVerticalUp,
				d.moveVerticalDown,
			} {
				next, ok := fn(pos)
				if !ok {
					continue
				}
				if _, ok := s.states[next]; ok {
					// The step was already there.
					continue
				}
				s.states[next] = Replay{
					steps: nextstep,
					prev:  d,
				}
				if next.IsSolved() {
					log.Printf("step #%d has found the solution after %d attempts (%d total)", step, len(result), len(s.states))
					return []Desk{next}, true
				}
				result = append(result, next)
			}
		}
	}
	log.Printf("step #%d has finished with %d states (%d total)", step, len(result), len(s.states))
	return result, false
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("The starting desk follows:")
	fmt.Println(start)
	ds := []Desk{start}
	s := NewSolver(start)
	maxsteps := 40
	for step := 0; step < maxsteps; step++ {
		var ok bool
		ds, ok = s.Step(step, ds...)
		if ok {
			fmt.Printf("============================================\nThe solutions is found in %d steps!!!\n", step+1)
			s.PrintChain(ds[0])
			return
		}
	}
	fmt.Printf("The solution is not found in %d steps.\n", maxsteps)
}
