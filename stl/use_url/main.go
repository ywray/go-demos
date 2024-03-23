package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main1() {
	urlStr := "https://galaxyfs-dev.dev.ihuman.com/nas/zyw-ci-0630/ci_pipeline_logs/6953175578786660352/6953175672227364864.log?aa=1"
	u, err := url.Parse(urlStr)
	if err != nil {

	}
	fmt.Println(u)
}

func extractPath(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	// 提取路径部分
	path := parsedURL.Path

	// 去除 "/fe/" 部分
	path = strings.Replace(path, "/fe/", "/", 1)

	//// 如果路径以斜杠开头，去掉开头的斜杠
	//if len(path) > 0 && path[0] == '/' {
	//	path = path[1:]
	//}
	params := parsedURL.Query()
	fmt.Println(params)

	return path, nil
}

func main() {
	url1 := "https://www.456.com/fe/nas/resource?insName=apk-storage&envName=core-wangjing&path=/ihuman-code-app/ihuman-code-app-abroad/60/"
	url2 := "https://www.123.com/fe/artifact/repo/artifact?envId=core-wangjing&version=master-2311231424&repositoryId=data-platform-etl-scala&repositoryName=data-platform-etl-scala"

	path1, err1 := extractPath(url1)
	if err1 != nil {
		fmt.Println("Error extracting path from URL 1:", err1)
	} else {
		fmt.Println("Extracted path from URL 1:", path1)
	}

	path2, err2 := extractPath(url2)
	if err2 != nil {
		fmt.Println("Error extracting path from URL 2:", err2)
	} else {
		fmt.Println("Extracted path from URL 2:", path2)
	}

	urlStr := "https://galaxy-dev.ihuman.pwrdgp.com/ci/start-build?cid=6993416375683252224&name=jump-android"
	u, _ := url.Parse(urlStr)
	params, _ := url.ParseQuery(u.RawQuery)
	if cid, ok := params["cid"]; ok {
		ciddd := cid[0]
		fmt.Println("ciddd", ciddd)
	}

}
