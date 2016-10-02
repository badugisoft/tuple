package tuple

type Pair struct {
	First, Second interface{}
}

func NewPair(f, s interface{}) Pair {
	return Pair{First: f, Second: s}
}

func NewpPair(f, s interface{}) *Pair {
	return &Pair{First: f, Second: s}
}

type Triplet struct {
	First, Second, Third interface{}
}

func NewTriplet(f, s, t interface{}) Triplet {
	return Triplet{First: f, Second: s, Third: t}
}

func NewpTriplet(f, s, t interface{}) *Triplet {
	return &Triplet{First: f, Second: s, Third: t}
}

type Quad struct {
	First, Second, Third, Fourth interface{}
}

func NewQuad(f, s, t, fo interface{}) Quad {
	return Quad{First: f, Second: s, Third: t, Fourth: fo}
}

func NewpQuad(f, s, t, fo interface{}) *Quad {
	return &Quad{First: f, Second: s, Third: t, Fourth: fo}
}
