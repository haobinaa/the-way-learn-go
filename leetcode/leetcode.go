package leetcode

func ReplaceSpace(s string) string {
	var result []byte
	for i, _ := range s {
		if s[i] == ' ' {
			result = append(result, '%', '2', '0')
		} else {
			result = append(result, s[i])
		}
	}
	return string(result)
}
