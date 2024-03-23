package main

import (
	"demo/persional/read_file"
	"fmt"
)

func main1() {
	var files []string
	files, _ = read_file.GetAllFile("/***/THUCNews/")
	fmt.Println("目录下的所有文件如下")
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}

func main2() {
	pathName := "/***/testdata/weibo_data_10w.txt"
	read_file.ReadByLine(pathName)
}

func main() {
	str := []string{"1", "2", "3"}
	fmt.Println(len(str))
	fmt.Println("str", str)

	str = str[1:]
	fmt.Println(len(str))
	fmt.Println("str", str)
}
