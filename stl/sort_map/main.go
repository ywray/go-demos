package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

type kvPair struct {
	Key   string
	Value uint64
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type kvPairList []kvPair

func (p kvPairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

// A function to turn a map into a kvPairList, then sort and return it.
func sortMapByValue(m map[string]string) kvPairList {
	p := make(kvPairList, len(m))
	i := 0
	for k, v := range m {
		t, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			continue
		}
		p[i] = kvPair{k, t}
		i++
	}
	sort.Slice(p, p.Less)
	return p
}

func main() {
	data := make(map[string]string)
	nowTimeStamp := time.Now().Unix()
	data["a"] = strconv.FormatInt(nowTimeStamp-1, 10)
	data["b"] = strconv.FormatInt(nowTimeStamp-2, 10)
	data["c"] = strconv.FormatInt(nowTimeStamp-3, 10)
	r := sortMapByValue(data)
	for k, v := range r {
		fmt.Println("k:", k, "v:", v)
	}
}
