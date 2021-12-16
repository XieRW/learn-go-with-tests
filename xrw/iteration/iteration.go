package iteration

func Repeat(s string) (repeated string) {
	for i := 0; i < 5; i++ {
		repeated += s
	}
	return repeated
}
