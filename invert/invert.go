package invert

func Invert(s string) string {
	b := []byte(s)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}

	return string(b)
}
