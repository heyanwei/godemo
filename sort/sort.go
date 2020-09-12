package sort

// Person ..
type Person struct {
	name string
	age  int
}

// PersonSlice ..
type PersonSlice []Person

// Len ..
func (p PersonSlice) Len() int {
	return len(p)
}

// Less ..
func (p PersonSlice) Less(i, j int) bool {
	return p[i].age < p[j].age
}

// Swap ..
func (p PersonSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
