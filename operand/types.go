package operand

import (
	"fmt"

	"github.com/mmcloughlin/avo/reg"
)

type Mem struct {
	Disp  int
	Base  reg.Register
	Index reg.Register
	Scale uint8
}

func (m Mem) Asm() string {
	a := ""
	if m.Disp != 0 {
		a += fmt.Sprintf("%d", m.Disp)
	}
	if m.Base != nil {
		a += fmt.Sprintf("(%s)", m.Base.Asm())
	}
	if m.Index != nil && m.Scale != 0 {
		a += fmt.Sprintf("(%s*%d)", m.Index.Asm(), m.Scale)
	}
	return a
}

type Imm uint64

func (i Imm) Asm() string {
	return fmt.Sprintf("%#x", i)
}

// Rel is an offset relative to the instruction pointer.
type Rel int32

func (r Rel) Asm() string {
	return fmt.Sprintf(".%+d", r)
}

// LabelRef is a reference to a label.
type LabelRef string

func (l LabelRef) Asm() string {
	return string(l)
}