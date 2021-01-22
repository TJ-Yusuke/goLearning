package iteration

//引数に書かれた文字列をrepeatCount回繰り返した文字列を返す
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
