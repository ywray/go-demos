package main

import (
	"fmt"
	"time"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

func runFuzzy() {
	fuzzy.Match("twl", "cartwheel")  // true
	fuzzy.Match("cart", "cartwheel") // true
	fuzzy.Match("cw", "cartwheel")   // true
	fuzzy.Match("ee", "cartwheel")   // true
	fuzzy.Match("art", "cartwheel")  // true
	fuzzy.Match("eeel", "cartwheel") // false
	fuzzy.Match("dog", "cartwheel")  // false
	fuzzy.Match("kitten", "sitting") // false

	fuzzy.RankMatch("kitten", "sitting") // -1
	fuzzy.RankMatch("cart", "cartwheel") // 5

	words := []string{"cartwheel", "foobar", "wheel", "baz"}
	fuzzy.Find("whl", words) // [cartwheel wheel]

	r := fuzzy.RankFind("whl", words) // [{whl cartwheel 6 0} {whl wheel 2 2}]
	fmt.Println("r:", r)
	// Unicode normalized matching.
	fuzzy.MatchNormalized("cartwheel", "cartwhéél") // true

	// Case insensitive matching.
	fuzzy.MatchFold("ArTeeL", "cartwheel") // true
}

func testRandomString(n int) []string {
	tmp := make([]string, 0, n)
	st := time.Now()
	for i := 0; i < n; i++ {
		tmp = append(tmp, RandStringBytesMaskImprSrc(5))
	}
	fmt.Println("time cost:", time.Now().Sub(st))
	return tmp
}
func main() {
	n := 200000
	stringSlice := testRandomString(n)
	st := time.Now()
	for i := 0; i < n; i++ {
		fmt.Printf("[%d]: %s\n", i, stringSlice[i])
	}
	fmt.Println("time cost2:", time.Now().Sub(st))
}
