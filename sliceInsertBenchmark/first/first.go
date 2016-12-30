package first

func First() {
	s := make([]int, 10)
	count := 100000
	for i := 0; i < count; i++ {
		s = append(s, i)
	}
}
