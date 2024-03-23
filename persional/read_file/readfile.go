package read_file

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
	//"sync"
)

var L sync.Mutex
var L2 sync.Mutex
var byteNum = 4096
var rOff, wOff = -byteNum, -byteNum
var retResev, retSend = false, false

var wg = sync.WaitGroup{}
var ch = make(chan []byte, 8)

func quick() {
	filePath := "/***/riot/testdata/weibo_data.txt"
	start := time.Now()
	newfile, err := os.OpenFile(filePath+"123", os.O_CREATE|os.O_WRONLY, 0666)
	checkErr(err)
	defer newfile.Close()

	file, e := os.Open(filePath)
	checkErr(e)
	defer file.Close()

	for i := 0; i < 4; i++ {
		//接收数据
		wg.Add(1)
		go resevCh(newfile)

		wg.Add(1)
		//发送数据
		go SendCh(file)
	}
	wg.Wait()
	end := time.Now()

	fmt.Println(end.Sub(start))
	fmt.Println("主程序退出")
}

func resevCh(newfile *os.File) {
	L2.Lock() //必须是与上面不同的锁
	defer func() {
		//fmt.Println("写入子程序退出....")
		L2.Unlock()
		wg.Done()
	}()
	//限制口，防止在其中一个g程写入文件完成后然而其他g程还没结束
	if retResev {
		return
	}
	for !retResev {
		ls, ok := <-ch
		if ok {
			wOff += byteNum
			//var n1 int
			_, err := newfile.WriteAt(ls, int64(wOff))
			//fmt.Println("len(ls)：", len(ls))
			//fmt.Println("线程写入resevCh的字节数为：", n1)
			checkErr(err)

		} else {
			retResev = true
		}

	}

}

func SendCh(file *os.File) {
	L.Lock()
	defer func() {
		fmt.Println("读取SendCh子程序退出！！！")
		L.Unlock()
		wg.Done()
	}()
	//限制口，防止在其中一个g程读取文件完成后然而其他g程还没结束
	if retSend {
		return
	}
	for !retSend {
		rOff += byteNum
		ls := make([]byte, byteNum) //每个g程都维护一个这样的结构

		n, err := file.ReadAt(ls, int64(rOff))
		//fmt.Println("某某线程读取了的字节数为：", n)
		checkErr(err)
		ch <- ls[:n] //通信
		if err == io.EOF {
			retSend = true
			close(ch) //务必要关闭,而且只能由一个g程关闭，而不是所有的g程,
			//所以不能放到defer func中去
			fmt.Println("关闭ch")
			return
		}

	}

}

func checkErr(err error) {

	if err != nil {
		fmt.Println(err)
	}
}

func ReadByLine(fileName string) {
	// open file
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	count := 0
	tmp := make([]string, 0, 10)
	word := ""
	runeWord := make([]rune, 0)
	// cutIndex 控制截取
	cutIndex := 64
	maxLen := 0
	maxLenIndex := 0
	for scanner.Scan() {
		// do something with a line
		tmp = strings.Split(scanner.Text(), "||||")
		if len(tmp) < 10 {
			continue
		}
		word = strings.Split(tmp[9], "http")[0]
		word = strings.ReplaceAll(word, " ", "")
		runeWord = []rune(word)
		if utf8.RuneCountInString(word) > cutIndex {
			word = string(runeWord[:cutIndex])
		}
		fmt.Printf("[%d], len1:%d -->  %s\n", count, len(runeWord), word)
		if len(runeWord) > maxLen {
			maxLen = len(runeWord)
			maxLenIndex = count
		}
		count++
	}

	fmt.Println("maxLen:", maxLen, "maxLenIndex:", maxLenIndex)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
