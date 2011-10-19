package main

import (
	"strings"
)

var (
	gMasks []Mask = []Mask{
		Mask{1 << 0, "public"},
		Mask{1 << 1, "private"},
		Mask{1 << 2, "protected"}}
)

type Mask struct {
	mask int
	name string
}

func (m Mask) Value() int {
	return m.mask
}

func (m Mask) String() (s string) {
	s = m.name
	return
}

type Maskable interface {
	Has(m Mask) bool
	Set(m Mask, on bool)
	String() string
}

type base_mask struct {
	*int
}

func NewBaseMask(list string) (b base_mask) {
	b.int = new(int)
	for _, mask := range gMasks {
		if strings.Contains(list, mask.String()) {
			b.Set(mask, true)
		}
	}
	return
}

func (b base_mask) Has(m Mask) bool {
	return (m.mask & *b.int) != 0
}

func (b base_mask) Set(m Mask, on bool) {
	if (on && !b.Has(m)) || (!on && b.Has(m)) {
		*b.int = *b.int ^ m.mask
	}
}

func (b base_mask) String() (s string) {
	for _, mask := range gMasks {
		if b.Has(mask) {
			s = mask.String()
		}
	}
	return
}
