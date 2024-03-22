package main

import (
	"demo/persional/ldflags/test_import"
	"fmt"
	"os"
)

var buildStamp = ""
var testLog = ""

func main() {
	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		// 该字段在make阶段由外部动态赋值
		fmt.Printf("UTC Build Time : %s\n", buildStamp)
		// 该字段在make阶段由外部静态赋值
		fmt.Printf("testLog : %s\n", testLog)
		// 该字段在make阶段由其他文件中的变量赋值
		fmt.Printf("test_version from other import: %s\n", test_import.TestVersion)
		return
	}
}

// 测试方法: make后, 执行 ./build-version -v 可打印外部注入的变量
