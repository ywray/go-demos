package main

import (
	"fmt"
)

func main() {
	offset := 0 //mean start from 0 (this is important for me)
	text := "123456783}}56"
	text1 := text[offset:]
	text2 := text[offset+5:]
	fmt.Println("text:", text, "1:", text1, "2:", text2)

	// golang底层用byte数组实现string，中文字符在unicode下占2字节，在utf-8下占3字节， golang默认编码为utf-8
	strIn := "12345汉字一二三四五abcde"
	str1 := strIn[3:6]
	fmt.Println("str1:", str1, "str len:", len(strIn))
	runeStr := []rune(strIn)
	runeStrLen := len(runeStr)
	rune2 := string(runeStr[3:6])
	fmt.Println("rune:", rune2, "rune2 len:", runeStrLen)

}
