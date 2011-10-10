package main

type Mask struct {
	mask int
	name string
}

func (m Mask) String(s string) {
	s = m.name
	return
}

type Maskable interface {
	Has(m Mask) bool
	String() string
}
