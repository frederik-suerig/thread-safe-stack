package threadsafestack

type StringStack struct {
	Value []string
}

func (s *StringStack) push(value string) {
	s.Value = append(s.Value, value)
}
