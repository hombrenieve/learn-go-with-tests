package iteration

func Repeat(character string, iterate int) string {
	var repeated string
	for i := 0; i < iterate; i++ {
		repeated += character
	}
	return repeated
}
