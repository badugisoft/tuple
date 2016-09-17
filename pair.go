package tuple

type Pair struct {
	First  interface{}
	Second interface{}
}

func NewPair(f interface{}, s interface{}) Pair {
	return Pair{First: f, Second: s}
}

func NewpPair(f interface{}, s interface{}) *Pair {
	return &Pair{First: f, Second: s}
}
