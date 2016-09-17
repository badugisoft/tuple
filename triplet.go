package tuple

type Triplet struct {
	First  interface{}
	Second interface{}
	Third  interface{}
}

func NewTriplet(f interface{}, s interface{}, t interface{}) Triplet {
	return Triplet{First: f, Second: s, Third: t}
}

func NewpTriplet(f interface{}, s interface{}, t interface{}) *Triplet {
	return &Triplet{First: f, Second: s, Third: t}
}
