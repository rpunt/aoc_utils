package utils

type Intrange struct {
	Min int
	Max int
}

func (i Intrange) Contains(j Intrange) bool {
	return j.Min >= i.Min && j.Max <= i.Max
}

func (i Intrange) Equals(j Intrange) bool {
	return i.Min == j.Min && i.Max == j.Max
}

func (i Intrange) Overlaps(j Intrange) bool {
	return (i.Min <= j.Min && i.Max >= j.Min) || (i.Min <= j.Max && i.Max >= j.Max) || i.Contains(j) || j.Contains(i)
}
