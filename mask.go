package main

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
