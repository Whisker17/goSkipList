package src

//Get key from Element
func (e *Element) Key() interface{} {
	return e.key
}

//Get next element
func (e *Element) Next() *Element {
	return e.next[0]
}

//Get next element at a specific level
func (e *Element) NextLevel(level int) *Element {
	if level >= len(e.next) || level < 0 {
		panic("invalid NextLevel!")
	}

	return e.next[level]
}