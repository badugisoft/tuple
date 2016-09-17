// ********************************************************* //
// The content of this file is auto-generated, DO NOT MODIFY //
// ********************************************************* //

package tuple

type StringTriplet struct {
	First  string
	Second string
	Third  string
}

func NewStringTriplet(f string, s string, t string) StringTriplet {
	return StringTriplet{First: f, Second: s, Third: t}
}

func NewpStringTriplet(f string, s string, t string) *StringTriplet {
	return &StringTriplet{First: f, Second: s, Third: t}
}

type IntTriplet struct {
	First  int
	Second int
	Third  int
}

func NewIntTriplet(f int, s int, t int) IntTriplet {
	return IntTriplet{First: f, Second: s, Third: t}
}

func NewpIntTriplet(f int, s int, t int) *IntTriplet {
	return &IntTriplet{First: f, Second: s, Third: t}
}

type BoolTriplet struct {
	First  bool
	Second bool
	Third  bool
}

func NewBoolTriplet(f bool, s bool, t bool) BoolTriplet {
	return BoolTriplet{First: f, Second: s, Third: t}
}

func NewpBoolTriplet(f bool, s bool, t bool) *BoolTriplet {
	return &BoolTriplet{First: f, Second: s, Third: t}
}
