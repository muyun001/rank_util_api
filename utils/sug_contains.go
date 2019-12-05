package utils

import "strings"

// IsSugContains 是否下拉包含
// 逐字按顺序匹配，sugStr中允许添加字符，但是最终要全部匹配checkWord中的字
// sugChars循环完，checkChars没循环完，则不匹配
func IsSugContains(checkWord, sugStr string) bool {
	checkChars := strings.Split(checkWord, "")
	sugChars := strings.Split(sugStr, "")
	s := 0
	slen := len(sugChars)
	for i := range checkChars {
		found := false
		for j := s; j < slen; j++ {
			if checkChars[i] == sugChars[j] {
				found = true
				s = j + 1
				break
			}
		}

		if found == false {
			return false
		}
	}

	return true
}
