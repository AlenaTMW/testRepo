package thelongestString

func strlen(str string) int {
	length := 0
	for range str {
		length++
	}
	return length
}

func Max(s []string) string {
	thelongestString := ""
	for i := range s {
		if strlen(thelongestString) < strlen(s[i]) {
			thelongestString = s[i]
		}
	}
	return thelongestString
}
