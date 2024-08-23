package iteration

func Repeat(character string, iter_factor int) string {
	var repeated string
	for i := 0; i < iter_factor; i++ {
		repeated += character
	}
	return repeated
}
