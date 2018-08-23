package main


func toLowerCase(str string) string {
	b := []rune(str)
	for i, v := range str {
		if v < 91 && v > 64 {
			b[i] += 32
		}
	}
	return string(b)
}
