package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func invalidReason(input string) string {
	// 包含非法字符 （A-Z/其他符号）
	regStr := "^[a-z0-9-]+$"
	re := regexp.MustCompile(regStr)
	match := re.MatchString(input)
	//fmt.Println("input", input, "result", match)
	if !match {
		return "包含非法字符（仅支持小写字母、数字及中横线-）"
	}

	// 只能小写字母开头（数字/- 开头）
	if !unicode.IsLetter(rune(input[0])) {
		return "只能以小写字母开头"
	}

	// 只能小写字母/数字结尾（-结尾）
	if input[0] == '-' {
		return "只能以小写字母或数字结尾"

	}
	return ""
}

func main() {
	testData := []string{"0", "-", "a", "0", "0000", "000a", "a?", "A", "?", "a-"}
	for _, v := range testData {
		fmt.Println("input", v, "output", invalidReason(v))
	}
}
