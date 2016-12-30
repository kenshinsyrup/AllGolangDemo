package second

func Second() {
	s := make([]int, 0, 10)
	count := 10000
	for i := 0; i < count; i++ {
		s = append(s, i)
	}
}
