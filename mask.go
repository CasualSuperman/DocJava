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

type FullMask struct {
	base_mask
	extra_masks []Mask
}

func (f FullMask) String() (s string) {
	s = f.base_mask.String()
	for _, mask := range f.extra_masks {
		if f.Has(mask) {
			if s != "" {
				s += " "
			}
			s += mask.String()
		}
	}
	return
}

func BitMasker(extras []Mask) (f FullMask) {
	f.base_mask = NewBaseMask("")
	f.extra_masks = append([]Mask{}, extras...)
	return
}

func (f FullMask) Apply(list string) (r FullMask) {
	r.base_mask = NewBaseMask(list)
	r.extra_masks = f.extra_masks
	for _, mask := range f.extra_masks {
		if strings.Contains(list, mask.String()) {
			r.Set(mask, true)
		}
	}
	return
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
